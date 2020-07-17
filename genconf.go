package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

var (
	corefileTemplate *template.Template
	tlsRegexp        *regexp.Regexp
)

func init() {
	var err error
	const tmpl = `Corefile.tmpl`
	corefileTemplate, err = template.New(tmpl).ParseFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	tlsRegexp = regexp.MustCompile(`^tls:\/\/\[?([^\]]+)\]?:[0-9]+$`)
}

func configurationGenerator(c *gin.Context) {
	acceptLang := c.GetHeader("Accept-Language")
	if strings.Contains(acceptLang, "zh") {
		c.HTML(http.StatusOK, "genconf.zh.tmpl", gin.H{
			"ChinaDNSServerList":  chinaDNSServers,
			"AbroadDNSServerList": abroadDNSServers,
		})
		return
	}
	c.HTML(http.StatusOK, "genconf.en.tmpl", gin.H{
		"ChinaDNSServerList":  chinaDNSServers,
		"AbroadDNSServerList": abroadDNSServers,
	})
}

type TLSDNSServerInfo struct {
	Port          int
	TLSServerName string
	ServerList    string
}

// CorefileData pass data to Corefile template
type CorefileData struct {
	Listen                 string
	AdsEnabled             bool
	DefaultAdsPolicy       bool
	AdsBlockList           string
	AdsWhiteList           string
	AdsUpdateInterval      string
	AdsCacheEnabled        bool
	HostsEnabled           bool
	AbroadDNSServerList    string
	ChinaDomainNameList    string
	ChinaDNSServerList     string
	BogusEnabled           bool
	BogusIPList            string
	LogEnabled             bool
	TTLCacheEnabled        bool
	RedisEnabled           bool
	RedisServerAddress     string
	HealthCheckEnabled     bool
	HotReloadEnabled       bool
	NetflixIPSet           bool
	AbroadTLSDNSServerList []*TLSDNSServerInfo
}

func generateConfiguration(c *gin.Context) {
	c.MultipartForm()

	customChinaDNSServers := c.PostForm("chinadnsservers")
	chinaDNSServerList := strings.Split(customChinaDNSServers, " ")

	customAbroadDNSServers := c.PostForm("abroaddnsservers")
	abroadDNSServerList := strings.Split(customAbroadDNSServers, " ")

	for key, _ := range c.Request.PostForm {
		if _, ok := chinaDNSServers[key]; ok {
			chinaDNSServerList = append(chinaDNSServerList, key)
		}
		if _, ok := abroadDNSServers[key]; ok {
			abroadDNSServerList = append(abroadDNSServerList, key)
		}
	}

	var abroadNonTLSDNSServerList []string
	var abroadTLSDNSServerList []string
	for _, s := range abroadDNSServerList {
		if !strings.HasPrefix(s, "tls://") {
			abroadNonTLSDNSServerList = append(abroadNonTLSDNSServerList, s)
		} else {
			abroadTLSDNSServerList = append(abroadTLSDNSServerList, s)
		}
	}

	var abroadTLSDNSServerInfoList []*TLSDNSServerInfo
	counter := 0
	for _, s := range abroadTLSDNSServerList {
		ss := tlsRegexp.FindStringSubmatch(s)
		if len(ss) <= 1 {
			log.Println("extracting host failed from", s)
			continue
		}
		tlsServerName, ok := tlsNameMap[ss[1]]
		if !ok {
			log.Println("can't recognize", ss[1])
			continue
		}
		duplicatedServerName := false
		for _, as := range abroadTLSDNSServerInfoList {
			if as.TLSServerName == tlsServerName {
				as.ServerList = as.ServerList + " " + s
				duplicatedServerName = true
				continue
			}
		}
		if duplicatedServerName {
			continue
		}
		counter++
		port := 53000 + counter
		abroadTLSDNSServerInfoList = append(abroadTLSDNSServerInfoList, &TLSDNSServerInfo{
			Port:          port,
			TLSServerName: tlsServerName,
			ServerList:    s,
		})
		abroadNonTLSDNSServerList = append(abroadNonTLSDNSServerList, fmt.Sprintf("127.0.0.1:%d", port))
	}

	d := CorefileData{
		Listen:                 c.PostForm("listen"),
		AdsEnabled:             c.PostForm("ads") == "on",
		DefaultAdsPolicy:       c.PostForm("defaultadspolicy") == "on",
		AdsBlockList:           c.PostForm("adsblocklist"),
		AdsWhiteList:           c.PostForm("adswhitelist"),
		AdsUpdateInterval:      c.PostForm("adsupdateinterval"),
		AdsCacheEnabled:        c.PostForm("adscache") == "on",
		HostsEnabled:           c.PostForm("hosts") == "on",
		BogusEnabled:           c.PostForm("bogus") == "on",
		LogEnabled:             c.PostForm("log") == "on",
		TTLCacheEnabled:        c.PostForm("ttlcache") == "on",
		RedisEnabled:           c.PostForm("redis") != "",
		RedisServerAddress:     c.PostForm("redis"),
		HealthCheckEnabled:     c.PostForm("healthcheck") == "on",
		HotReloadEnabled:       c.PostForm("hotreload") == "on",
		ChinaDNSServerList:     strings.Replace(strings.Join(chinaDNSServerList, " "), "dns://", "", -1),
		AbroadDNSServerList:    strings.Replace(strings.Join(abroadNonTLSDNSServerList, " "), "dns://", "", -1),
		ChinaDomainNameList:    strings.Join(chinaDomainList, " "),
		BogusIPList:            strings.Join(bogusIPList, " "),
		NetflixIPSet:           false,
		AbroadTLSDNSServerList: abroadTLSDNSServerInfoList,
	}

	var configurations bytes.Buffer
	if err := corefileTemplate.Execute(&configurations, d); err != nil {
		log.Println("executing Corefile template failed", err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename=Corefile")
	c.Data(http.StatusOK, "text/plain;charset=UTF-8", configurations.Bytes())
}
