package geoip

import (
	"fmt"
	"testing"
	"time"
)

var cases = map[string]string{
	"1.208.0.0":       "KR",
	"114.114.114.114": "CN",
	"101.226.103.106": "CN",
	"14.17.32.211":    "CN",
	"8.8.8.8":         "US",
	"183.79.227.111":  "JP",
	"255.255.255.255": "",
	"192.168.0.1":     "",
	"224.0.0.1":       "",
}

func TestGeoip(t *testing.T) {

	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		QueryCountryByString("8.8.8.8")
	}
	t2 := time.Now()
	fmt.Println("time=", t2.Sub(t1))

	t3 := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println(IsDirectDomainRecursive("xxxxxx.xxx.wangsu.com"))
	}
	t4 := time.Now()
	fmt.Println("time=", t4.Sub(t3))
}
