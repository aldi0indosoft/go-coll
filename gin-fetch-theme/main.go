package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gojektech/heimdall"
)


func defRoute(c *gin.Context) {
	hystrixConfig := heimdall.NewHystrixConfig("fetch_hugo_themes", heimdall.HystrixCommandConfig{
		ErrorPercentThreshold: 20,
		MaxConcurrentRequests: 30,
		Timeout:               1000,
	})
	timeout := 10 * time.Second
	hc := heimdall.NewHystrixHTTPClient(timeout, hystrixConfig)
	// urls := []string{
	// 	"https://themes.gohugo.io/tags/single-page/",
	// 	"https://themes.gohugo.io/tags/white/",
	// }
	// for _, url := range urls {
	// 	resp, err := hc.Get(url, nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("body: ", string(body))

	c.JSON(200, gin.H{
		"message": "render",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", defRoute)
	r.Run(":8155")
}
