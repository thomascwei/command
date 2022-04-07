package internal

import (
	"context"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

var (
	Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	Configs           = LoadConfig("./config")
	DBConnection      = fmt.Sprintf("mongodb://%s:%s@%s:%s", Configs.User, Configs.Password, Configs.Host, Configs.Port)
	_, b, _, _        = runtime.Caller(0)
	rootPath          = filepath.Dir(filepath.Dir(b))
	CommandCollection *mongo.Collection
	HeaderCollection  *mongo.Collection
	sugarLogger       *zap.SugaredLogger
	Ctx, CtxCancel    = context.WithTimeout(context.Background(), 30*time.Second)
	Client            *mongo.Client
)

// 讀專案中的config檔
func LoadConfig(MyPath string) (config Config) {
	// 若有同名環境變量則使用環境變量
	viper.AutomaticEnv()
	viper.AddConfigPath(MyPath)
	// 為了讓執行test也能讀到config添加索引路徑
	wd, err := os.Getwd()
	parent := filepath.Dir(wd)
	viper.AddConfigPath(path.Join(parent, MyPath))
	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("can not load config: " + err.Error())
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("can not load config: " + err.Error())
	}
	//Trace.Printf("%+v\n", config)
	return
}
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
		Filename:   "log/internal.log",
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
	Client, err := mongo.NewClient(options.Client().ApplyURI(DBConnection))

	if err != nil {
		log.Fatal(err)
	}
	err = Client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}

	CommandCollection = Client.Database(Configs.DB).Collection(Configs.CommandTable)
	HeaderCollection = Client.Database(Configs.DB).Collection(Configs.HeaderTable)
	sugarLogger.Infof("finish initial")
	Trace.Println("finish initial")
}

func GetAllCommandTemplate() (ResultList []ReadCommandTemplate, err error) {
	cur, err := CommandCollection.Find(Ctx, bson.D{})
	if err != nil {
		Trace.Println(err)
		log.Fatal(err)
	}
	defer cur.Close(Ctx)
	for cur.Next(Ctx) {
		var result ReadCommandTemplate
		err := cur.Decode(&result)
		if err != nil {
			Trace.Println(err)
			log.Fatal(err)
		}
		// do something with result....
		ResultList = append(ResultList, result)
		if err := cur.Err(); err != nil {
			Trace.Println(err)
			log.Fatal(err)
		}
	}
	return
}

func GetAllHeaderTemplate() {}

func GetSingleCommandTemplate() {}

func CreateSingleCommandTemplate() {}

func CreateSingleHeaderTemplate() {}

func UpdateSingleCommandTemplate() {}

func UpdateSingleHeaderTemplate() {}

func DeleteSingleCommandTemplate() {}

func DeleteSingleHeaderTemplate() {}
