# Slack Poster Cli
コマンドラインからslackの自分のチャンネルに自分のアカウントでメッセージを投稿できる簡単なアプリケーションです。  
勤務中にslackの個人用のチャンネルに何かと投稿することが多く、その際のslackを開く手間を省くこととslackを見ることによって集中力が途切れるのを防ぐために作りました。  
goの勉強も兼ねているので少しずつ改良していきます。

# Requirements
- go 1.14.3 darwin/amd64
- direnv

# Setup
```bash
git clone git@github.com:t-kinomura/slack-poster-cli.git
cd slack-poster-cli
cp .envrc .evnrc.expamle // tokenとチャンネル名を指定する
direnv allow
```

# Usage
```bash
go run main.go // 実行すると入力待ちの状態になる。
// 送信したいメッセージを入力する。マークダウンも一応ちゃんと反映されます。
// CTRL-Dで入力を抜けると入力した内容が送信される。
```
