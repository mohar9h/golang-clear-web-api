package logging

import (
	"github.com/mohar9h/golang-clear-web-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	config *config.Config
	logger *zap.SugarLogger
}

func newZapLogger(config *config.Config) *zapLogger {
	logger := &zapLogger{config: config}
	logger.Init()
	return logger
}

func (z *zapLogger) getLogLevel() zapcore.Level {
	level, exists := logLevelMap[z.config.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (z *zapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   z.config.LogFile,
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		z.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller,
		zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

	z.logger = logger
}

func (z *zapLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	prepareLogKeys(extra, category, subCategory, z, message)
}

func (z *zapLogger) Debugf(template string, args ...interface{}) {
	z.logger.Debugf(template, args...)
}

func (z *zapLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	prepareLogKeys(extra, category, subCategory, z, message)
}

func (z *zapLogger) Infof(template string, args ...interface{}) {
	z.logger.Infof(template, args...)
}

func (z *zapLogger) Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	prepareLogKeys(extra, category, subCategory, z, message)
}

func (z *zapLogger) Warnf(template string, args ...interface{}) {
	z.logger.Warnf(template, args...)
}

func (z *zapLogger) Error(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	prepareLogKeys(extra, category, subCategory, z, message)
}

func (z *zapLogger) Errorf(err error, template string, args ...interface{}) {
	z.logger.Errorf(template, args...)
}

func (z *zapLogger) Fatal(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	prepareLogKeys(extra, category, subCategory, z, message)
}

func (z *zapLogger) Fatalf(err error, template string, args ...interface{}) {
	z.logger.Fatalf(template, args...)
}

func prepareLogKeys(extra map[ExtraKey]interface{}, category Category, subCategory SubCategory, z *zapLogger, message string) {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}

	extra["Category"] = category
	extra["SubCategory"] = subCategory

	params := mapToZapParams(extra)

	z.logger.Debugw(message, params...)
}
