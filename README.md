# Golang Todo CLI

一個簡單的 Golang 指令列待辦清單工具，支援新增、查看、完成、刪除任務。
資料使用 SQLite 儲存，讓你每次執行程式都能延續上次的待辦清單。

---

## 功能特色

- 查看所有待辦事項
- 新增待辦事項
- 標記任務為完成
- 刪除任務
- 使用 SQLite 進行本地資料儲存

---

## 使用方式

### 安裝並執行

```bash
git clone https://github.com/samuel361592/golang-todo-cli.git
cd golang-todo-cli
go run main.go
```

## 資料儲存說明

待辦事項會儲存在 data/todos.db

data/ 為執行期資料目錄，已被加入 .gitignore

第一次執行程式時，會自動建立資料夾、資料庫與資料表

Clone 專案後可直接執行，無需額外設定

作者
Samuel