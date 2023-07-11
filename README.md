# Qleave
qleave 是請假管理系統，採取前後端分離模式，前端使用vue3，而後端則是gin。

## 技術 (Technology stacks)
| 名稱 | 版本 |
| :-- | :--: |
| Go |  1.19|
| Gin | 1.9.0 |
| JavaScript | es6 |
| Vue |3.2.45 |
| Css |3 |
| tailwind |3.2.7 |

## API 目錄結構
```
├─cmd (主要應用程式)
├─database (資料庫管理相關。例如：migrations 文件。)
├─deploy (部署相關。 例如：docker-compose，docker。)
├─docs (設計和使用者文件)
├─domain (放置關於 domain interface)
├─infrastructure (放置與外部鏈接相關，例如資料庫，外部 api)
├─internal (私有應用程式和函式庫)
├─modal (sqlc 生產的檔案)
└─pkg (可讓外部應用使用的函式庫)		
```

## 執行 API 伺服器
1. 進入到 backstage 目錄下。
```
cd backstage
```
2. 安裝套件
```
go mod download
```
3. 初始化資料庫，內部有已建立資料，為了展現使用
```
go run cmd/app/main.go init
```
4. 初始化資料庫，內部有已建立資料，為了展現使用
```
go run cmd/app/main.go init
```
5. 執行
```
go run cmd/app/main.go run
```

## 執行 Web
1. 進入到 forestage 目錄下。
```
cd forestage
```
2. 安裝套件
```
npm install
```
3. 執行
```
npm run dev
```