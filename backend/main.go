package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kryptogo/eth-tokyo-2023-kukuponpon/api"
	"github.com/kryptogo/eth-tokyo-2023-kukuponpon/api/middleware"
)

func main() {
	appPort := "8081"
	r := gin.New()
	r.Use(middleware.CorsMiddleware)
	r.GET("/ok", api.GetOk)
	r.GET("/campaigns", api.GetCampaigns)
	r.POST("/verify_campaign", api.VerifyCampaign)
	r.POST("/get_coupon", api.RetrieveCoupon)

	r.Run(":" + appPort)

}
