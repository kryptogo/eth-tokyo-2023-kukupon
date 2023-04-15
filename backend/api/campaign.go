package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryptogo/eth-tokyo-2023-kukuponpon/pkg/service"
)

// GetCampaigns response campaigns
func GetCampaigns(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetCampaigns())
}

type VerifyCampaignReq struct {
	CampaignId string `json:"campaign_id"`
	From       string `json:"from"`
	Signature  string `json:"signature"`
}

// VerifyCampaign verify signature
func VerifyCampaign(c *gin.Context) {
	var req VerifyCampaignReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify signature
	verifyMsg := map[string]interface{}{
		"campaign_id": req.CampaignId,
	}
	verifyMsgStr, err := json.Marshal(verifyMsg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	verified := service.VerifySignature(req.From, req.Signature, string(verifyMsgStr))
	if !verified {
		// do nothing for demo
	}

	// verify campaign condition
	isMembership := service.VerifyMembership(req.CampaignId, req.From)
	if !isMembership {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not a member"})
		return
	}

	c.SetCookie("verified_membership", "cookie_value", 3600, "/", "localhost", false, true)
	c.String(http.StatusOK, "ok")
}

type RetrieveCouponReq struct {
	CampaignId string `json:"campaign_id"`
}

// RetrieveCoupon retrieve coupon, the string can convert to a private key
func RetrieveCoupon(c *gin.Context) {
	// verify cookie
	cookieVakue, err := c.Cookie("verified_membership")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if cookieVakue != "cookie_value" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cookie"})
		return
	}

	var req RetrieveCouponReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// coupons := service.GetCoupons(req.CampaignId)

	// c.JSON(http.StatusOK, gin.H{
	// 	"coupons": coupons,
	// })
}
