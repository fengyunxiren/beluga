package server

import (
	"beluga/server/common/config"
	"beluga/server/common/database"
	"beluga/server/common/logger"
	"beluga/server/middleware"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetUp(configPath string) error {
	fmt.Println("Server SetUP")
	configer, err := setupConfig(configPath)
	if err != nil {
		return err
	}
	var logLevel logrus.Level
	if configer.Default.Debug {
		logLevel = logrus.DebugLevel
	} else {
		logLevel = logrus.InfoLevel
	}
	err = setupLogger(configer.Default.LogPath, logLevel)
	if err != nil {
		return err
	}
	err = setupDatabase(configer.Database)
	if err != nil {
		return err
	}
	return nil
}

func setupConfig(configPath string) (*config.Config, error) {
	fmt.Println("config path is ", configPath)
	configer, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	fmt.Println("configer: ", configer)
	config.InitDefaultConfig(configer)
	return configer, nil
}

func setupLogger(logPath string, logLevel logrus.Level) error {
	log, err := logger.NewLogger(logPath, logLevel)
	if err != nil {
		return err
	}
	logger.InitDefaultLogger(log)
	return nil
}

func setupDatabase(data config.Database) error {
	if data.Driver == "" || data.Dsn == "" {
		return errors.New("Database is not config")
	}
	if data.MaxIdleConn == 0 {
		data.MaxIdleConn = 10
	}
	if data.MaxOpenConns == 0 {
		data.MaxOpenConns = 100
	}
	if data.MaxLifeTime == time.Duration(0) {
		data.MaxLifeTime = time.Hour
	}
	fmt.Println("Database Driver: ", data.Driver)
	fmt.Println("Database MaxIdleConn: ", data.MaxIdleConn)
	fmt.Println("Database MaxOpenConns: ", data.MaxOpenConns)
	fmt.Println("Database maxLifeTime: ", data.MaxLifeTime)
	return database.InitDBPool(data.Driver, data.Dsn, data.MaxIdleConn, data.MaxOpenConns, data.MaxLifeTime)
}

func initMiddleware(router gin.IRouter) {
	router.Use(middleware.RequestId())
	configer := config.GetConfig()
	var logLevel logrus.Level
	if configer.Default.Debug {
		logLevel = logrus.DebugLevel
	} else {
		logLevel = logrus.InfoLevel
	}
	router.Use(middleware.AccessLogger(configer.Default.AccessLogPath, logLevel))
}

func RunServer() error {
	fmt.Println("Welcom to beluga...")
	router := gin.New()
	initMiddleware(router)
	ServerApps.RegisterRouter(router)
	configer := config.GetConfig()
	log := logger.GetLogger()
	router.GET("/ping", func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		log.Info("test for logger")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", configer.Servers.Host, configer.Servers.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Showtdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server showtdown:", err)
	}
	log.Println("Server existing")
	return nil
}
