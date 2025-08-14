package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_NAME     string
	DB_PORT     uint
	SERVER_PORT uint
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {

	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {

		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	// これは、.envファイルがローカル開発マシンで使用され、本番サーバーでは使用されないことを意味します。
	if _, exist := os.LookupEnv("SECRET"); !exist { // "SECRET"環境変数が存在しない場合
		if err := godotenv.Load(".env"); err != nil { // .envファイルを読み込みます
			log.Println("Error loading .env file:", err) // .envファイルの読み込みエラーをログに出力します
		}
	}

	// SERVER_PORTを読み取り、文字列から数値に変換します。
	cnvServerPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil { // エラーが発生した場合
		log.Fatal("Cannot parse Server Port variable") // サーバーポート変数をパースできませんという致命的なエラーをログに出力します
		return nil
	}

	// 環境変数からAppConfig構造体に値を設定します。
	defaultConfig.SERVER_PORT = uint(cnvServerPort)
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")

	return &defaultConfig // デフォルト設定へのポインタを返します
}
