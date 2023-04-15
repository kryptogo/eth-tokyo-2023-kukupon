package service

import (
	"fmt"
	"testing"

	"github.com/kryptogo/eth-tokyo-2023-kukupon/pkg/db"
)

func TestGetMembers(t *testing.T) {
	campaignID := "kygc_holder"
	campaign := db.GetCampaign(campaignID)
	members := getMembers(campaignID, campaign.GraphqlReq.OperationName, campaign.GraphqlReq.Query)
	fmt.Println(len(members))
	// assert.True(t, members["0x3fe5ac82b997255b8226abb6aefd91f405fe2e8e"])
	// assert.True(t, members["0x9bae3a9cbac6e769d980f6ce0fd52397dd45d361"])

}
