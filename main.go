package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	lastRefreshTimeStamp int64
	rc                   *RedisCache
	username             = `missdeer`
	project              = `coredns-custom-build`
	token                = ``
)

func handler(c *gin.Context) {
	baseName := filepath.Base(c.Param("baseName"))

	targetLink, err := rc.GetString(baseName)
	if err != nil || targetLink == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	acceptLang := c.GetHeader("Accept-Language")
	if strings.Contains(acceptLang, "zh") {
		c.HTML(http.StatusOK, "index.zh.tmpl", gin.H{
			"targetLink": targetLink,
		})
		return
	}
	c.HTML(http.StatusOK, "index.en.tmpl", gin.H{
		"targetLink": targetLink,
	})
}

func main() {
	redis := os.Getenv("REDIS")
	if redis == "" {
		redis = "127.0.0.1:6379"
	}

	if rc = RedisInit(redis); rc == nil {
		return
	}

	updateLinkMap := func() {
		now := time.Now().Unix()
		if atomic.LoadInt64(&lastRefreshTimeStamp)+3600 > now {
			return
		}
		atomic.StoreInt64(&lastRefreshTimeStamp, now)
		a := &Appveyor{}
		a.UpdateLinkMap()
	}

	go updateLinkMap()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://github.com/missdeer/coredns_custom_build")
	})
	r.GET("/dl/*baseName", handler)
	r.GET("/refresh", func(c *gin.Context) {
		updateLinkMap()
		c.JSON(200, gin.H{"result": "OK"})
	})
	r.POST("/refresh", func(c *gin.Context) {
		updateLinkMap()
		c.JSON(200, gin.H{"result": "OK"})
	})

	bind := os.Getenv("BIND")
	if bind == "" {
		bind = ":8765"
	}
	log.Fatal(r.Run(bind))
}
