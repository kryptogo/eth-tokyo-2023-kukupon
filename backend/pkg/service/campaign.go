package service

import "github.com/kryptogo/eth-tokyo-2023-kukupon/pkg/db"

func GetCampaigns() []db.Campaign {
	return db.GetCampaigns()
}
