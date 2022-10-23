package logger

import (
	"beluga/global"
	"beluga/utils"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var defaultLogger *logrus.Logger

func InitDefaultLogger(log *logrus.Logger) {
	defaultLogger = log
}

func GetLogger() *logrus.Entry {
	if defaultLogger == nil {
		defaultLogger = logrus.New()
	}
	return logrus.NewEntry(defaultLogger)
}

func GetContextLogger(c *gin.Context) *logrus.Entry {
	contextLogger, exists := c.Get(global.CONTEXT_LOGGER)
	var logger *logrus.Entry
	if !exists {
		logger = GetLogger()
	} else {
		cLogger, ok := contextLogger.(*logrus.Entry)
		if !ok {
			logger = GetLogger()
		} else {
			logger = cLogger
		}
	}
	return logger
}

func NewLogger(logPath string, logLevel logrus.Level) (*logrus.Logger, error) {
	dir, _ := filepath.Split(logPath)
	if err := utils.CheckAndCreatePath(dir); err != nil {
		return logrus.New(), err
	}
	logger := logrus.New()
	logger.SetLevel(logLevel)
	loggerWrite, err := rotatelogs.New(
		strings.TrimSuffix(logPath, ".log")+"_%Y%m%d.log",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return logrus.New(), err
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  loggerWrite,
		logrus.FatalLevel: loggerWrite,
		logrus.DebugLevel: loggerWrite,
		logrus.WarnLevel:  loggerWrite,
		logrus.ErrorLevel: loggerWrite,
		logrus.PanicLevel: loggerWrite,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.0000",
	})
	logger.AddHook(lfHook)
	return logger, nil
}
