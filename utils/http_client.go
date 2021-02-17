package utils

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var Client http.Client

//新建爬虫所用的http client
func NewHttpClient(timeout int) {
	//设置cookie
	cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	//使用client是为了能利用cookies和代理
	Client = http.Client{Jar: cookieJar, Timeout: time.Second * time.Duration(timeout)}
}
