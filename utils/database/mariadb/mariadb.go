package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bellatrijuliana/agoratix-app/config" // 設定パッケージをインポートします。
	_ "github.com/go-sql-driver/mysql"               // MariaDBドライバーです。
)

// InitDBはデータベースへの接続を初期化し、返します。
func InitDB(cfg *config.AppConfig) *sql.DB {
	// 設定から接続文字列を作成します。
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	// データベース接続を開きます。
	db, err := sql.Open("mysql", dbURI)
	if err != nil { // エラーが発生した場合
		log.Fatal("Failed to open database connection:", err) // データベース接続のオープンに失敗しましたという致命的なエラーをログに出力します
	}

	// パフォーマンス向上のため、接続プールを設定します。
	// 毎回新しいタクシーを呼ぶのではなく、数台のタクシーを待機させておくようなものです。
	db.SetMaxIdleConns(10)                  // アイドル状態の接続を保持する数です。
	db.SetMaxOpenConns(100)                 // 開くことができる最大接続数です。
	db.SetConnMaxIdleTime(5 * time.Minute)  // 接続がアイドル状態を維持できる最長期間です。
	db.SetConnMaxLifetime(60 * time.Minute) // 接続が全体で存続できる最長期間です。

	// 接続が機能していることを確認するためにデータベースにpingを送信します。
	err = db.Ping()
	if err != nil { // エラーが発生した場合
		log.Fatal("Failed to connect to the database:", err) // データベースへの接続に失敗しましたという致命的なエラーをログに出力します
	}

	fmt.Println("Successfully connected to the database!") // データベースへの接続に成功しました！
	return db                                              // データベース接続を返します
}
