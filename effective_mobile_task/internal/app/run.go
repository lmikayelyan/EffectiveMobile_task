package app

import (
	"context"
	"effective_mobile_task/docs"
	"effective_mobile_task/internal/config"
	"effective_mobile_task/internal/handler"
	"effective_mobile_task/internal/repository"
	"effective_mobile_task/internal/service"
	pkg2 "effective_mobile_task/pkg"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	logger := pkg2.NewLogger(cfg)
	newLog, err := logger.InitLogger(ctx)
	if err != nil {
		log.Panic(err)
	}

	pgSession := pkg2.NewPgSession(&cfg.Postgres)
	pool, err := pgSession.InitPgSession(ctx)
	if err != nil {
		log.Panic(err)
	}

	carRepo := repository.CarRepo(pool)
	carService := service.CarService(carRepo)
	carHandler := handler.CarHandler(carService, newLog)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	actions := router.Group("/cars")
	{
		actions.POST("/add", carHandler.Create)
		actions.PATCH("/update/:id", carHandler.Update)
		actions.DELETE("/remove/:id", carHandler.Delete)
		actions.GET("/info/:limit/:offset", carHandler.Get)
	}

	if err := router.Run(docs.SwaggerInfo.Host); err != nil {
		log.Panic(err)
	}
}
