package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func getClient() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		Protocol: 2,
	})
}

func GetServiceURL(service string) string {
	db := getClient()
	ctx := context.Background()

	port, err := db.Get(ctx, service).Result()
	if err != nil {
		panic(err)
	}
	url := "http://host.docker.internal:" + port
	fmt.Println(url)
	return url
}

func RunReverseProxy(ctx *gin.Context) {
	// Get service name from url.
	service := ctx.Param("service")
	// Get service info from db (redis)
	if service == "" {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "service is not found."})
		return
	}

	// for URL.
	module := GetServiceURL(service)
	remote, err := url.Parse(module)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	// Make reverce proxy director.
	rp := httputil.NewSingleHostReverseProxy(remote)
	rp.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = ctx.Param("param")
	}

	// go to module.
	rp.ServeHTTP(ctx.Writer, ctx.Request)
}

func MakeRouting(router *gin.Engine) *gin.Engine {
	router.GET("/:service/*param", RunReverseProxy)
	router.POST("/:service/*param", RunReverseProxy)
	router.PUT("/:service/*param", RunReverseProxy)
	router.DELETE("/:service/*param", RunReverseProxy)

	return router
}

func main() {
	router := gin.Default()
	router = MakeRouting(router)
	router.Run(":9000")
}
