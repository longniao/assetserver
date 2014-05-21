package main

import (
	"github.com/msbranco/goconfig"
	"net/http"
)

func main() {
	// 配置
	c, err := goconfig.ReadConfigFile("config.ini")
	if err != nil {
		print(err)
	}

	webroot, err := c.GetString("server", "webroot")

    http.Handle("/", http.FileServer(http.Dir(webroot)))
    http.ListenAndServe(":8080", nil)

}