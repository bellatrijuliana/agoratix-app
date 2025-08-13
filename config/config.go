package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// AppConfigは、アプリケーションのすべての設定を保持する構造体です。
// これは、コンフィグレーションの設計図だと考えてください。
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

// GetConfigは、アプリケーションの他の部分が設定を取得するために呼び出す公開関数です。
// 「シングルトンパターン」を使用しており、
// 設定が一度だけ読み込まれることを保証します。
func GetConfig() *AppConfig {
	// ロックは、アプリケーションの複数の部分が
	// まったく同時に設定を要求した場合でも、
	// すべてが一度に読み込もうとしないように保証します。
	// ドアの「一度に一人」のルールのようなものです。
	lock.Lock()         // ロックを取得します
	defer lock.Unlock() // 関数が終了する前にロックを解放します

	// もし設定がまだ読み込まれていない場合（appConfigがnilの場合）...
	if appConfig == nil {
		// ...initConfig()を呼び出して読み込みます。
		appConfig = initConfig()
	}

	// 読み込まれた設定を返します。
	return appConfig
}

// initConfigは、.envファイルから環境変数を実際に読み込むプライベート関数です。
func initConfig() *AppConfig {
	var defaultConfig AppConfig

	// これはよくあるテクニックです：特定の
	// 環境変数（"SECRET"など）が設定されていない場合にのみ.envファイルを読み込みます。
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
