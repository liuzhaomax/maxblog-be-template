package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"maxblog-be-template/internal/conf"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/src/pb"
	"maxblog-be-template/src/service"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type options struct {
	ConfigDir  string
	ConfigFile string
}

type Option func(*options)

func SetConfigDir(configDir string) Option {
	return func(opts *options) {
		opts.ConfigDir = configDir
	}
}

func SetConfigFile(configFile string) Option {
	return func(opts *options) {
		opts.ConfigFile = configFile
	}
}

func InitConfig(opts *options) {
	cfg := conf.GetInstanceOfConfig()
	cfg.Load(opts.ConfigDir, opts.ConfigFile)
	logger.WithFields(logger.Fields{
		"path": opts.ConfigDir + "/" + opts.ConfigFile,
	}).Info(core.Config_File_Load_Succeeded)
}

func InitDB() (*gorm.DB, func(), error) {
	cfg := conf.GetInstanceOfConfig()
	logger.Info(core.DB_Connection_Started)
	db, clean, err := cfg.NewDB()
	if err != nil {
		logger.Fatal(core.DB_Connection_Failed, err)
		return nil, clean, err
	}
	err = cfg.AutoMigrate(db)
	if err != nil {
		logger.Fatal(core.DB_Auto_Migration_Failed, err)
		return nil, clean, err
	}
	return db, clean, err
}

func InitServer(ctx context.Context, service *service.BData) func() {
	cfg := conf.GetInstanceOfConfig()
	host := flag.String("host", cfg.Server.Host, "Enter host")
	port := flag.Int("port", cfg.Server.Port, "Enter port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	server := grpc.NewServer()
	pb.RegisterDataServiceServer(server, service)
	go func() {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Fatal(core.Server_Listen_Failed, err)
			panic(err)
		}
		logger.WithContext(ctx).Infof("Server is running at %s", addr)
		err = server.Serve(listen)
		if err != nil {
			logger.Fatal(core.Server_Serve_Failed, err)
			panic(err)
		}
	}()
	return func() {
		_, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.Server.ShutdownTimeout))
		defer cancel()
		server.Stop()
		logger.Info(core.Server_Stoped)
	}
}

func Init(ctx context.Context, opts ...Option) func() {
	// initialising options
	options := options{}
	for _, opt := range opts {
		opt(&options)
	}
	// init config
	InitConfig(&options)
	// init injector and DB
	injector, injectorClean, _ := InitInjector()
	cfg := conf.GetInstanceOfConfig()
	logger.WithFields(logger.Fields{
		"db_type":   cfg.DB.Type,
		"db_name":   cfg.Mysql.DBName,
		"user_name": cfg.Mysql.UserName,
		"host":      cfg.Mysql.Host,
		"port":      cfg.Mysql.Port,
	}).Info(core.DB_Connecetion_Succeeded)
	// init server
	serverClean := InitServer(ctx, injector.Service)
	return func() {
		serverClean()
		injectorClean()
	}
}

func Launch(ctx context.Context, opts ...Option) {
	logger.Info(core.Server_Launch_Start)
	clean := Init(ctx, opts...)
	cfg := conf.GetInstanceOfConfig()
	logger.WithFields(logger.Fields{
		"app_name": cfg.App.AppName,
		"version":  cfg.App.Version,
		"pid":      os.Getpid(),
		"host":     cfg.Server.Host,
		"port":     cfg.Server.Port,
	}).Info(core.Server_Started)
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
LOOP:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("%s [%s]", core.Server_Interrupt_Received, sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break LOOP
		case syscall.SIGHUP:
		default:
			break LOOP
		}
	}
	defer logger.WithContext(ctx).Infof(core.Server_Shutting_Down)
	defer time.Sleep(time.Second)
	defer os.Exit(state)
	defer clean()
}
