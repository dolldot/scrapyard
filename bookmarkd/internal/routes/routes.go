package routes

import (
	"net/http"
	"time"

	"github.com/dolldot/scrapyard/bookmarkd/config"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent"
	"github.com/dolldot/scrapyard/bookmarkd/internal/handlers"
	"github.com/dolldot/scrapyard/bookmarkd/internal/logger"
	"github.com/dolldot/scrapyard/bookmarkd/internal/repos"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *ent.Client) *http.Server {
	router := gin.New()

	router.Use(ginzap.GinzapWithConfig(logger.Logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        false,
		SkipPaths:  []string{"/"},
	}))

	router.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:5173",
	}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	router.Static("/assets", "./assets")

	bookmarkRepo := repos.NewBookmarkRepo(db)
	accountRepo := repos.NewAccountRepo(db)
	bookmarkHandler := handlers.NewBookmarkHandler(*bookmarkRepo)
	accountHandler := handlers.NewAccountHandler(*accountRepo)

	router.GET("/", handlers.Landing)

	account := router.Group("/account")
	{
		account.POST("/", accountHandler.CreateAccount)
		account.GET("/login", accountHandler.RenderLogin)
		account.POST("/login", accountHandler.Login)
	}

	bookmark := router.Group("/:number/bookmark/")
	{
		bookmark.GET("/", bookmarkHandler.GetBookmarks)
		bookmark.GET("/add", bookmarkHandler.AddBookmark)
		bookmark.GET("/:uuid/edit", bookmarkHandler.EditBookmark)
		bookmark.POST("/", bookmarkHandler.CreateBookmark)
		bookmark.PUT("/:uuid", bookmarkHandler.UpdateBookmark)
		bookmark.DELETE("/:uuid", bookmarkHandler.DeleteBookmark)
	}

	router.GET("/ping", handlers.HealthCheck)

	server := http.Server{
		Addr:    config.Load.Server.Port,
		Handler: router,
	}

	return &server
}
