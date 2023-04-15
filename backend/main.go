package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kryptogo/eth-tokyo-2023-kukuponpon/api"
)

func main() {
	appPort := "8088"
	r := gin.New()

	r.GET("/ok", api.GetOk)
	r.GET("/campaigns", api.GetCampaigns)
	r.POST("/verify_campaign", api.VerifyCampaign)
	r.POST("/get_coupon", api.RetrieveCoupon)

	r.Run(":" + appPort)

}
