package infrastructure

import (
	"log"
	"service-exercise/infrastructure/gorm/config"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// データベース接続の*gorm.DBをfxに保持する
// DBModule:fx.Option
var DBModule = fx.Provide(func() (*gorm.DB, error) {
	// データベース接続情報を取得する
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	// MySQLに接続するコネクタを取得する
	connector := config.NewMySQLConnector(conf)
	// データベース接続した結果の*gorm.DBを
	return connector.Connect()
})

// インフラストラクチャ層の依存関係を構築する
var InfrastructureModule = fx.Options(
	DBModule, // データベース接続の*gorm.DBをfxに保持する
	fx.Provide(),
	fx.Invoke(setupEnd), // 依存関係構築完了メッセージを出力する
)

func setupEnd() {
	log.Println("インフラストラクチャ層の構築が完了しました。")
}
