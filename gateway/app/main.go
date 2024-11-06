package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetServiceURL(service string) string {
	url := "http://" + service + ":9000/"
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
