# Go研修 DDD WebApp

![ogp](https://github.com/user-attachments/assets/1fd81192-e44e-425c-9c8d-8ab1b8d58f1b)

## 概要

このプロジェクトはGoを使用したWebアプリケーションの研修用プロジェクトです。クリーンアーキテクチャの原則に従って設計されています。

## 技術スタック

### 基本環境
- Go 1.23.2

### 主要な依存関係
#### Webフレームワーク
- `gin-gonic/gin` - 高性能なWebフレームワーク
- `gin-contrib/cors` - CORSミドルウェア

#### データベース
- `gorm.io/gorm` - Go用ORMライブラリ
- `gorm.io/driver/mysql` - MySQL用GORMドライバー

#### 依存性注入
- `go.uber.org/fx` - 依存性注入フレームワーク
- `go.uber.org/dig` - 依存性注入コンテナ

#### ロギング
- `go.uber.org/zap` - 高性能な構造化ロギング

#### API ドキュメント
- `swaggo/swag` - Swagger/OpenAPI ドキュメント生成
- `swaggo/gin-swagger` - Gin用Swaggerインテグレーション

#### ユーティリティ
- `google/uuid` - UUID生成
- `go-playground/validator` - 構造体バリデーション

## プロジェクト構造

```
.
├── application/     # アプリケーション層（ユースケース）
├── domain/         # ドメイン層（ビジネスロジック）
├── infrastructure/ # インフラストラクチャ層（外部サービスとの連携）
├── presentation/   # プレゼンテーション層（APIハンドラー）
└── docs/          # プロジェクトドキュメント
```

## セットアップ

1. リポジトリのクローン:
```bash
git clone [https://github.com/manaY-monoX/go-hands-webapp]
cd go-hands-webapp
```

2. 依存関係のインストール:
```bash
go mod download
```

3. アプリケーションの実行:
```bash
go run main.go
```

## 開発ガイドライン

- コードはクリーンアーキテクチャの原則に従って構造化されています
- 各層は明確に分離され、依存関係は内側に向かって流れます
- テストを書く際は、対応するパッケージと同じ階層に `_test.go` ファイルを作成してください

## ライセンス

このプロジェクトは研修用に作成されています。
