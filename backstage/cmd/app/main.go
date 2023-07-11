package main

import (
	"context"
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "leaveManager/internal/autoload"
	"leaveManager/internal/route"
	"leaveManager/internal/sqliteInit"
	"leaveManager/pkg/validator"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var actionParam string = ""
	if len(os.Args) > 1 {
		actionParam = os.Args[1]
	}

	switch actionParam {
	case "init":
		sqliteInit.InitDatabase()
	case "run":
		runServer()
	default:
		info := `
			init : 初始化
			run  : 執行
		`
		log.Println(info)
	}
}

func runServer() {
	post := os.Getenv("POST")

	db, err := sql.Open("sqlite3", "./sqlitedb.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	validator.Init()

	gin.DisableConsoleColor()
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	httpService := gin.Default()

	route.Setup(db, httpService)

	srv := &http.Server{
		Addr:    ":" + post,
		Handler: httpService,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server exiting")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
