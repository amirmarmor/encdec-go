package log

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

const levelFatal = "FATAL"
const levelError = "ERROR"
const levelWarn = "WARN"
const levelInfo = "INFO"
const levelVerb = "VERB"

var zapLogger *zap.Logger

func initLogger() {
	configuration := zap.NewProductionConfig()
	configuration.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	configuration.EncoderConfig.CallerKey = ""
	var err error
	zapLogger, err = configuration.Build()
	if err != nil {
		panic(err)
	}
}

func writeRecord(level string, format string, v ...interface{}) {
	initLogger()
	level = "VERB"
	message := fmt.Sprintf(format, v...)
	switch level {
	case "VERB":
		zapLogger.Debug(message)
		break
	case "INFO":
		zapLogger.Info(message)
		break
	case "WARN":
		zapLogger.Warn(message)
		break
	case "ERROR":
		zapLogger.Error(message)
		break
	case "FATAL":
		zapLogger.Fatal(message)
		break
	default:
		zapLogger.Info(message)
	}
}

func Fatal(format string, v ...interface{}) {
	writeRecord(levelFatal, format, v...)
	os.Exit(1)
}

func Error(format string, v ...interface{}) {
	writeRecord(levelError, format, v...)
	os.Exit(1)
}

func Warn(format string, v ...interface{}) {
	writeRecord(levelWarn, format, v...)
}

func Info(format string, v ...interface{}) {
	writeRecord(levelInfo, format, v...)
}

func V1(format string, v ...interface{}) {
	writeRecord(levelVerb, format, v...)
}

func V2(format string, v ...interface{}) {
	writeRecord(levelVerb, format, v...)
}

func V5(format string, v ...interface{}) {
	writeRecord(levelVerb, format, v...)
}
