package main

import (
	"fmt"
	"github.com/JeremyLoy/config"
	"github.com/gin-gonic/gin"
	"log"
	"rhea/configuration"
	"rhea/proxy"
	responsestore "rhea/response-store"
)

func main() {
	var c configuration.Configuration
	if err := config.From(".env").FromEnv().To(&c); err != nil {
		log.Fatal("error loading configuration")
	}
	s := responsestore.NewResponseStore()
	p := proxy.Proxy{TargetRawUrl: c.TargetRawUrl, Store: s}
	r := gin.Default()
	r.GET("/*path", p.Handler)

	if err := r.Run(fmt.Sprintf(":%d", c.Port)); err != nil {
		log.Printf("Error: %v", err)
	}
}
