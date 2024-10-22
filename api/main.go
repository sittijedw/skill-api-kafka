package main

import (
	"api/database"
	"api/skill"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {
	db := database.ConnectDB()
	database.CreateSkillTable(db)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()

	db := database.ConnectDB()
	skillRepository := skill.NewRepository(db)
	skillHandler := skill.NewHandler(skillRepository)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/skills/:key", skillHandler.GetByKeyHandler)
		v1.GET("/skills", skillHandler.GetAllHandler)
		v1.POST("/skills", skillHandler.CreateHandler)
		v1.PUT("/skills/:key", skillHandler.UpdateByKeyHandler)
		v1.PATCH("/skills/:key/actions/name", skillHandler.UpdateNameByKeyHandler)
		v1.PATCH("/skills/:key/actions/description", skillHandler.UpdateDescriptionByKeyHandler)
		v1.PATCH("/skills/:key/actions/logo", skillHandler.UpdateLogoByKeyHandler)
		v1.PATCH("/skills/:key/actions/tags", skillHandler.UpdateTagsByKeyHandler)
		v1.DELETE("/skills/:key", skillHandler.DeleteByKeyHandler)
	}

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	closedChan := make(chan struct{})

	go func() {
		<-ctx.Done()
		log.Println("Shutting down...")

		ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println(err)
			}
		}
		close(closedChan)
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

	<-closedChan
}
