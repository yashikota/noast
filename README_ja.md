# noast

Win/Mac/WSL/Linux どれからでも使えるトースト通知用CLIアプリケーションです。

**[English README](README.md)**

## インストール

```bash
go install github.com/yashikota/noast@latest
```

## 使い方

```bash
# 基本的な通知
noast -message "Hello, World!"

# タイトルとメッセージを指定
noast -title "重要な通知" -message "作業が完了しました"

# アイコンを指定
noast -title "通知" -message "メッセージ" -icon "/path/to/icon.png"
```
