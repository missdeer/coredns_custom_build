package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	username = `missdeer`
	project  = `coredns-custom-build`
	token    = ``
	linkMap  = map[string]string{
		"coredns-windows-amd64.zip":            "",
		"coredns-windows-386.zip":              "",
		"coredns-darwin-amd64.zip":             "",
		"coredns-linux-386.zip":                "",
		"coredns-linux-amd64.zip":              "",
		"coredns-linux-armv5.zip":              "",
		"coredns-linux-armv6.zip":              "",
		"coredns-linux-armv7.zip":              "",
		"coredns-linux-arm64.zip":              "",
		"coredns-linux-ppc64.zip":              "",
		"coredns-linux-ppc64le.zip":            "",
		"coredns-linux-mips64-hardfloat.zip":   "",
		"coredns-linux-mips64-softfloat.zip":   "",
		"coredns-linux-mips64le-hardfloat.zip": "",
		"coredns-linux-mips64le-softfloat.zip": "",
		"coredns-linux-mips-hardfloat.zip":     "",
		"coredns-linux-mips-softfloat.zip":     "",
		"coredns-linux-mipsle-hardfloat.zip":   "",
		"coredns-linux-mipsle-softfloat.zip":   "",
		"coredns-linux-s390x.zip":              "",
		"coredns-freebsd-amd64.zip":            "",
		"coredns-freebsd-386.zip":              "",
		"coredns-freebsd-arm.zip":              "",
		"coredns-netbsd-amd64.zip":             "",
		"coredns-netbsd-386.zip":               "",
		"coredns-netbsd-arm.zip":               "",
		"coredns-openbsd-amd64.zip":            "",
		"coredns-openbsd-386.zip":              "",
		"coredns-dragonfly-amd64.zip":          "",
		"coredns-solaris-amd64.zip":            "",
		"coredns-android-amd64.zip":            "",
		"coredns-android-386.zip":              "",
		"coredns-android-arm.zip":              "",
		"coredns-android-aarch64.zip":          "",
	}
)

func handler(c *gin.Context) {
	baseName := filepath.Base(c.Param("baseName"))
	targetLink, ok := linkMap[baseName]
	if ok && targetLink != "" {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"targetLink": targetLink,
		})
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func main() {
	redis := os.Getenv("REDIS")
	if redis == "" {
		redis = "127.0.0.1:6379"
	}
	rc := RedisInit(redis)
	if rc == nil {
		return
	}

	for k := range linkMap {
		l, e := rc.GetString(k)
		if e == nil {
			linkMap[k] = l
		}
	}

	updateLinkMap := func() {
		a := &Appveyor{}
		a.UpdateLinkMap()

		for k, v := range linkMap {
			rc.Put(k, v)
		}
	}

	updateLinkMap()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
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
