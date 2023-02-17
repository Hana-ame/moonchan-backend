package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/meromeromeiro/moonchan-backend/api/v1"
	"github.com/meromeromeiro/moonchan-backend/webfinger"
)

func main() {
	router := gin.Default()
	// webfinger
	router.GET(webfinger.WebFingerPath, webfinger.WebfingerHandler)
	// activitypub
	// router.
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
	_ = api.T
}
