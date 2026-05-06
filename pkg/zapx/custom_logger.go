package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CreateLogger() *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	middlePriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.WarnLevel && lvl >= zap.InfoLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zap.DebugLevel
	})

	fileError := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/error.log",
		MaxSize:    1, // 10 megabytes
		MaxBackups: 3,
		MaxAge:     1, // 1 day
		Compress:   true,
		LocalTime:  true,
	})

	fileInfo := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/info.log",
		MaxSize:    1, // 10 megabytes
		MaxBackups: 3,
		MaxAge:     1, // 1 day
		Compress:   true,
		LocalTime:  true,
	})

	fileDebug := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/debug.log",
		MaxSize:    1, // 10 megabytes
		MaxBackups: 3,
		MaxAge:     1, // 1 day
		Compress:   true,
		LocalTime:  true,
	})

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	developmentCfg.CallerKey = "caller"

	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileError, highPriority),
		zapcore.NewCore(fileEncoder, fileInfo, middlePriority),
		zapcore.NewCore(fileEncoder, fileDebug, lowPriority),
	)

	return zap.New(core, zap.AddCaller())
}
