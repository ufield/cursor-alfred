# Alfred Workflow for Cursor Project Manager

## 概要

このワークフローは、CursorのProject Managerに登録されているプロジェクトをAlfredから簡単に検索・起動できるものです。

- `csr <プロジェクト名>` でプロジェクト候補を表示
- 選択するとCursorでそのプロジェクトを開く
- Go製シングルバイナリで動作し、追加インストール不要

## 使い方

1. `csr.go` をビルドして `csr` バイナリを作成します。
   ```sh
   go build -o csr csr.go
   ```
2. Alfred WorkflowのScript Filterに `csr` バイナリを指定します。
   - キーワード: `csr`
   - Script Filter: `/path/to/csr`
3. アクションには以下のように設定します。
   - `{query}` を引数として受け取り、
   - `/usr/bin/open -a "Cursor" "{query}"`
     でプロジェクトを開きます。

## 前提
- Cursorのプロジェクト情報は `~/.cursor/projects.json` に保存されている必要があります。
- Goがインストールされていればビルド可能です。

## カスタマイズ
- プロジェクトファイルのパスが異なる場合は、`csr.go` 内のパスを修正してください。

---

何か問題があればIssueやPRをお願いします。

## 開発向け

### ビルド方法
以下のコマンドでビルドできます：
```
go build -o cursor-alfred main.go
```