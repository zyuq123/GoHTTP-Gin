package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger(path string) *zap.Logger {

	hook := lumberjack.Logger{
		Filename:   path, // 日志文件路径
		MaxSize:    128,  // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,   // 日志文件最多保存多少个备份
		MaxAge:     7,    // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别:低于级别的不打印
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	Logger = zap.New(core, caller, development, filed)
	return Logger

}

func Warn(msg string, fileds ...zap.Field) {
	Logger.Warn(msg, fileds...)
}

func Info(msg string, fileds ...zap.Field) {
	Logger.Info(msg, fileds...)
}

func Error(msg string, fileds ...zap.Field) {
	Logger.Error(msg, fileds...)
}

//谨慎使用Fatal，除非是到了迫不得已需要终止的地方，比如数据库连接失败
func Fatal(msg string, fileds ...zap.Field) {
	Logger.Fatal(msg, fileds...)
}
