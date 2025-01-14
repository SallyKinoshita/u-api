# API 実装課題

## 概要

本リポジトリは、ある課題に基づき作成された API の実装です。本プロジェクトでは、Go 言語を使用し、オニオンアーキテクチャを採用しています。

## 使用技術

- **言語**: Go
- **データベース**: MySQL
- **インフラ**: Docker / Docker Compose
- **アーキテクチャ**: オニオンアーキテクチャ

## ディレクトリ構成

```plaintext
.
├── cmd
│   └── api                  # アプリケーションのエントリーポイント
│       └── main.go
├── internal
│   ├── application
│   │   └── usecases         # ビジネスロジック
│   ├── domain
│   │   ├── models           # ドメインモデル
│   │   └── repositories     # リポジトリインターフェース
│   ├── infrastructure
│   │   ├── db               # データベース接続設定
│   │   └── persistence      # リポジトリ実装
│   ├── interfaces
│   │   └── controllers      # コントローラー
│   └── gen
│       └── openapi          # OpenAPI生成コード
├── migrations               # マイグレーションファイル
├── docs                     # OpenAPI仕様書
├── docker　　　　　　　　　　  # Docker構成
├── docker-compose.yml       # Docker構成ファイル
├── .env                     # 環境変数ファイル
├── atlas.hcl                # DBスキーマ管理ファイル
├── go.mod                   # Goモジュール設定
└── Makefile                 # 開発タスクのスクリプト
```

## セットアップ

### 必要条件

以下がインストールされている必要があります：

- Docker
- Docker Compose

### 起動手順

以下の手順でプロジェクトをセットアップし、MySQL と Web サーバーを起動します。

```bash
git clone https://github.com/SallyKinoshita/u-api.git
cd u-api
cp .env.example .env
make up
```

## テスト

### テストの実行

以下のコマンドでテストを実行します。

```bash
make go-test
```

## 補足説明

### アーキテクチャ

- **オニオンアーキテクチャ**:
  - **Domain**: 業務ロジックやエンティティを管理。
  - **Application**: ユースケースや DTO を定義。
  - **Infrastructure**: データベースや永続化に関する実装。
  - **Interfaces**: 外部とのインターフェース（HTTP ハンドラー）。

### マイグレーション

- **Atlas** を使用してデータベースのスキーマを管理。
- `migrations` ディレクトリにマイグレーションファイルを格納。
- 一部 seeder ファイルも `migrations` 配下で管理している。

#### マイグレーションファイル自動生成 (差分反映)

atlas.hcl を更新後、下記を実行

```shell
make db-migrate-diff
```

#### マイグレーション (反映)

※一部 seeder を含む。

```shell
make db-migrate-apply
```

## 残 TODO

時間切れにより下記残 TODO となりました…

- 請求書データの新規作成の動作確認。
  - データ生成処理ロジック一部未実装(CompanyID の取得)
- ユニットテストの一部。テスト実装箇所も失敗しているものがある。
- 全体的にエラーハンドリングが雑。カスタム error パッケージを用意して整理。
- カスタム context パッケージを用意して ctx 周りの実装を整理する。
- middleware 実装。認証・認可の考慮。
