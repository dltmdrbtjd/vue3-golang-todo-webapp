package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Convenience-Tools/convenience-server/pkg/db"
	"github.com/Convenience-Tools/convenience-server/pkg/middleware"
	"github.com/Convenience-Tools/convenience-server/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		db.Close(ctx)
		cancel()
	}()

	db.Initialize(ctx)

	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Use(middleware.CORSMiddleware())
	engine.Use(middleware.JSONMiddleware())
	routes.Register(engine)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("ListenAndServe has been failed. Error %s", err.Error())
		panic(err)
	}
}
