# Alfred Workflow for Cursor Project Manager

## 概要

このワークフローは、CursorのProject Managerに登録されているプロジェクトをAlfredから簡単に検索・起動できるものです。

- `cr <プロジェクト名>` でプロジェクト候補を表示
- 選択するとCursorでそのプロジェクトを開く
- Go製シングルバイナリで動作し、追加インストール不要

## 使い方

## 前提
- Cursorのプロジェクト情報はデフォルトでは `Library/Application Support/Cursor/User/globalStorage/alefragnani.project-manager/projects.json` に保存されていることを前提に動作します
  - 上記以外のパスに `projects.json` が存在する場合、Alfred Workflow の環境変数 `CURSOR_PROJECTS_FILE` にフルパスを指定してください。
- Goがインストールされていればビルド可能です。

---

何か問題があればIssueやPRをお願いします。

## 開発向け

### ビルド方法
以下のコマンドでビルドできます：
```
go build -o cursor-alfred main.go
```