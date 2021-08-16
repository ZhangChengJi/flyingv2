package logs

import (
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level
var L *zap.Logger

type logger struct {
	level         string
	format        string
	prefix        string
	director      string
	linkName      string
	showLine      bool
	encodeLevel   string
	stacktraceKey string
	logInConsole  bool
}

func NewLog() {
	log := &logger{
		level:         "info",
		format:        "console",
		prefix:        "[flying]",
		director:      "log",
		linkName:      "latest.log",
		showLine:      true,
		encodeLevel:   "LowercaseColorLevelEncoder",
		stacktraceKey: "stacktrace",
		logInConsole:  true,
	}
	L = log.Zap()
}

func (log *logger) Zap() (logger *zap.Logger) {
	if ok, _ := PathExists(log.director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", log.director)
		_ = os.Mkdir(log.director, os.ModePerm)
	}

	switch log.level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(log.getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(log.getEncoderCore())
	}
	if log.showLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func (log *logger) getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  log.stacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     log.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case log.encodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case log.encodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case log.encodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case log.encodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func (log *logger) getEncoder() zapcore.Encoder {
	if log.format == "json" {
		return zapcore.NewJSONEncoder(log.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(log.getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func (log *logger) getEncoderCore() (core zapcore.Core) {
	writer, err := log.GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(log.getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func (log *logger) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(log.prefix + "2006/01/02 - 15:04:05.000"))
}
func (log *logger) GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(log.director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(log.linkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if log.logInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
