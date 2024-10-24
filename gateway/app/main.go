package main

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	Method string `redis:"method"`
	Url    string `redis:"url"`
	Query  string `redis:"query"`
}

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
}

func GetServicerInfo(service string) Service {
	db := getClient()
	ctx := context.Background()

	var result Service
	err := db.HGetAll(ctx, service).Scan(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	router := gin.Default()
	router.GET("/:service/*param", func(ctx *gin.Context) {
		// Get service name from url.
		service := ctx.Param("service")
		// Get service info from db (redis)
		if service == "" {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": service + " is not found."})
		}

		// Get and confirm service info
		res := GetServicerInfo(service)
		// for URL.
		remote, err := url.Parse(res.Url)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": service + " is failed."})
		}
		// for Method.
		if !slices.Contains(strings.Split(res.Method, "/"), ctx.Request.Method) {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": service + " is not found."})
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
	})

	router.Run(":9000")
}
