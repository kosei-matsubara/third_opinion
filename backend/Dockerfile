# ========================================
# 📦 ビルド用ステージ（本番バイナリ作成）
# ========================================
FROM golang:1.24-alpine AS builder
WORKDIR /app

# 📋 依存関係ファイルのみ先にコピー（Docker layer cache効率化）
# go.mod = Goプロジェクトの設定ファイル（package.jsonのGo版）
# go.sum = 依存関係のバージョン固定ファイル（package-lock.jsonのGo版）
COPY go.mod go.sum ./

# 🌐 依存関係ダウンロード（GitHubなどから必要なパッケージをダウンロード）
# この時点でgin-gonicなどのライブラリがダウンロードされる
RUN go mod download

# 📂 アプリケーションソースコード全体をコピー
COPY . .

# 🔨 本番用バイナリファイル作成
# CGO_ENABLED=0: C言語ライブラリ依存なしでビルド（Alpine環境対応）
# GOOS=linux: Linux OS用にビルド
# -ldflags '-extldflags "-static"': 静的リンク（依存関係を全て含む実行ファイル）
# -o main: 出力ファイル名を"main"に指定
# ./cmd/: メインファイル（main.go）の場所
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main ./cmd/

# ========================================
# 🚀 本番環境用ステージ（軽量化）
# ========================================
FROM alpine:latest AS deploy

# 📡 アプリケーションが使用するポート番号を公開
EXPOSE 8080

# 🔐 HTTPS通信用証明書をインストール（外部API連携で必要）
RUN apk --no-cache add ca-certificates

# 本番用作業ディレクトリ
WORKDIR /app

# 📦 ビルドステージで作成したバイナリファイルをコピー
# --from=builder: 前のステージから取得
COPY --from=builder /app/main .

# 🏃‍♂️ アプリケーション起動コマンド
CMD ["./main"]

# ========================================
# 🛠️ 開発環境用ステージ（開発ツール充実）
# ========================================
FROM golang:1.24-alpine AS dev
EXPOSE 8080
WORKDIR /app

# 🕒 データベース起動待機スクリプト追加
# 問題: Goアプリが先に起動→DB未準備→接続エラー→アプリ落ちる
# 解決: wait-for-itでDB準備完了まで待機してからアプリ起動
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /usr/local/bin/
RUN chmod 755 /usr/local/bin/wait-for-it.sh

# 🏗️ 開発に必要な基本ツールインストール
# git: Gitリポジトリ操作
# build-base: C言語コンパイラ（一部のGoライブラリで必要）
# bash: より高機能なシェル（デバッグ時に便利）
# curl: HTTP通信テスト用
# nodejs, npm, yarn: フロントエンド開発ツール（フルスタック対応）
RUN apk add --no-cache git build-base bash curl nodejs npm yarn && \
    sed -i 's/bin\/ash/bin\/bash/g' /etc/passwd

# 🔄 Go開発ツールのインストール（GitHubから最新版取得）
# 
# 1. Air（ホットリロード）
# GitHub: https://github.com/cosmtrek/air
# 機能: ソースコード変更を検知して自動でアプリ再起動
# 使用例: ファイル保存 → 自動でgo run実行 → ブラウザで即座に確認
RUN go install github.com/air-verse/air@latest && \
    \
    # 2. Migrate（データベースマイグレーション）
    # GitHub: https://github.com/golang-migrate/migrate
    # 機能: SQLファイルを使ってDB構造を自動更新
    # 使用例: migrate up → テーブル作成/更新が自動実行
    # -tags 'mysql': MySQL対応版をビルド
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1 && \
    \
    # 3. golangci-lint（コード品質チェック）
    # GitHub: https://github.com/golangci/golangci-lint
    # 機能: 文法エラー、未使用変数、効率の悪いコード等をチェック
    # 使用例: 保存時に自動チェック → 問題があれば警告表示
    # curl: インストールスクリプトをダウンロード実行
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

# 🧹 ビルドキャッシュクリーンアップ（イメージサイズ削減）
RUN go clean -modcache

# 📂 開発用ソースコードコピー
COPY . .

# 🏃‍♂️ 開発環境起動コマンド
# wait-for-it.sh: DB起動待機スクリプト
# -t 30: 最大30秒待機
# db:3306: 接続先（docker-compose.ymlのdbサービス:3306ポート）
# --: 待機完了後に実行するコマンドの開始
# air: ホットリロード開発サーバー起動
#
# 実行の流れ:
# 1. DBが起動するまで最大30秒待機
# 2. DB接続確認できたらairコマンド実行
# 3. airがmain.goを監視しながらサーバー起動
# 4. ファイル変更時に自動再起動
CMD ["wait-for-it.sh", "-t", "30", "db:3306", "--", "air"]

# ========================================
# 📚 各ツールの詳細説明
# ========================================
#
# 🔄 Air (github.com/cosmtrek/air)
# - ファイル変更監視 → 自動ビルド → 自動再起動
# - 設定: .air.toml ファイルで詳細設定可能
# - 開発効率: 手動再起動が不要になり開発速度向上
#
# 🗃️ Migrate (github.com/golang-migrate/migrate)
# - up: マイグレーション実行（テーブル作成・更新）
# - down: マイグレーション取り消し（テーブル削除・巻き戻し）
# - ファイル命名: 000001_create_users.up.sql
#
# 🔍 golangci-lint
# - 30以上のリンター（チェックツール）を統合
# - 設定: .golangci.yml で有効/無効チェック項目設定
# - CI/CD: 本番デプロイ前の品質ゲートとして使用
#
# ⏳ wait-for-it
# - シェルスクリプト（Bash製）
# - TCP接続テスト → 成功まで待機 → 指定コマンド実行
# - Docker Compose環境での依存関係解決に必須
# ========================================
