package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kryptogo/eth-tokyo-2023-kukupon/api"
	"github.com/kryptogo/eth-tokyo-2023-kukupon/api/middleware"
)

func main() {
	appPort := "8080"
	r := gin.New()
	r.Use(middleware.CorsMiddleware)
	r.GET("/ok", api.GetOk)
	r.GET("/campaigns", api.GetCampaigns)
	r.POST("/verify_campaign", api.VerifyCampaign)
	r.POST("/get_coupon", api.RetrieveCoupon)

	r.Run(":" + appPort)

}
