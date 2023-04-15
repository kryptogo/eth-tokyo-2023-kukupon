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
		"from":        req.From,
	}
	verifyMsgStr, err := json.Marshal(verifyMsg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	verified := service.VerifySignature(req.From, req.Signature, string(verifyMsgStr))
	if !verified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		return
	}

	// verify campaign condition
	isMembership := service.VerifyMembership(req.CampaignId, req.From)
	c.JSON(http.StatusOK, gin.H{
		"verified": isMembership,
	})
}

// RetrieveCoupon retrieve coupon
func RetrieveCoupon(c *gin.Context) {
	c.String(http.StatusOK, "coupon")
}
