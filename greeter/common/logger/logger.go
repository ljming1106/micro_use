package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var log *zap.SugaredLogger

var logLevel = zap.NewAtomicLevel()

/*
 	文件切割：当前使用了lumberjack库对文件进行切割，建议可以用ELK等文件聚合中间件代替
	对zap info等函数进行包装时，如果需要日志文件和函数时，可以使用zap.AddCallerSkip(1)方法，跳过封装函数的调用
*/

func init() {
	filePath := getFilePath()

	fmt.Println("filePath:", filePath)

	//w := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:  filePath,
	//	MaxSize:   1024, //MB
	//	LocalTime: true,
	//	Compress:  true,
	//})

	w := zapcore.AddSync(os.Stdout)

	config := zap.NewProductionEncoderConfig()
	//config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeTime = formatEncodeTime
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		logLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync() // 将 buffer 中的日志写到文件中
	log = logger.Sugar()
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05"))
}

type Level int8

const (
	DebugLevel Level = iota - 1

	InfoLevel

	WarnLevel

	ErrorLevel

	DPanicLevel

	PanicLevel

	FatalLevel
)

func SetLevel(level Level) {
	logLevel.SetLevel(zapcore.Level(level))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Info(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getFilePath() string {
	logfile := getCurrentDirectory() + "/" + getAppname() + ".log"
	return logfile
}

func getAppname() string {
	full := os.Args[0]
	full = strings.Replace(full, "\\", "/", -1)
	splits := strings.Split(full, "/")
	if len(splits) >= 1 {
		name := splits[len(splits)-1]
		name = strings.TrimSuffix(name, ".exe")
		return name
	}

	return ""
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

// 兼容
func Log(args ...interface{}) {
	Info(args)
}

func Fatal(args ...interface{}) {
	Panic(args)
}
