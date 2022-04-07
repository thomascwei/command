package api

import (
	"command/internal"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var (
	Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	sugarLogger *zap.SugaredLogger
)

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "log/route.log",
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
func init() {
	InitLogger()
	defer sugarLogger.Sync()
}

// 檢查若有錯誤將錯誤信息返回給client
func ResponseErrorInfo(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"result": "fail",
		"error":  err.Error(),
	})
}

// request完成返回結果
func FinishReturnResult(c *gin.Context, result interface{}) {
	c.JSON(200, gin.H{
		"result": "ok",
		"data":   result,
	})
	return
}

func GetAllCommandTemplateRoute(c *gin.Context) {
	result, err := internal.GetAllCommandTemplate()
	if err != nil {
		sugarLogger.Errorf(err.Error())
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, result)
}
