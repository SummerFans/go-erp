package presenter

import (
	"context"
	"flag"
	"fmt"
	"go-erp/application/handler"
	"go-erp/application/middleware"
	"go-erp/domain/throwable"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	router *gin.Engine
}

func (s *httpServer) Run() {
	// 1. 超时处理
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "gracefully waiting for live connections")
	flag.Parse()

	// 获取服务端口
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8080"
	}

	// service 构建服务
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.router,
	}

	// runing 运行服务
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()
	<-ctx.Done()

	log.Println("Shutting down")
	os.Exit(0)
}

// loadRouter 载入路由
func loadRouter(
	categoriesHandler handler.CategoriesHandler,
	goodsHandler handler.GoodsHandler,
) *gin.Engine {
	gin.ForceConsoleColor()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery()) // 捕获异常恢复，防止服务奔溃
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// 登录认证
	security := middleware.NewSecurityMiddleware()
	router.Use(security.Auth())

	// 基础中间件
	baseMiddleware := middleware.NewBaseMiddleware()
	router.Use(baseMiddleware.Locale())

	goodsRouter(router, goodsHandler)
	categoryRouter(router, categoriesHandler)

	router.NoRoute(func(c *gin.Context) {
		err := throwable.NewNotFound("Not Found")
		baseMiddleware.Error(c.Writer, err)
		c.AbortWithError(http.StatusNotFound, *err)
	})
	return router
}

func NewHttpServer(
	categories handler.CategoriesHandler,
	goodsHandler handler.GoodsHandler) Server {
	router := loadRouter(
		categories,
		goodsHandler,
	)

	return &httpServer{router: router}
}
