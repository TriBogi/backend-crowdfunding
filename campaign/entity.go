package campaign

import (
	"bogistartup/user"
	"github.com/leekchan/accounting"
	"time"
)

type Campaign struct {
	ID                int
	User_ID           int
	Name              string
	Short_Description string
	Description       string
	Perks             string
	Backer_Count      int
	Goal_Amount       int
	Current_Amount    int
	Slug              string
	Created_At        time.Time
	Updated_At        time.Time
	CampaignImages    []CampaignImage
	User              user.User
}

func (c Campaign) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.Goal_Amount)
}

func (c Campaign) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.Current_Amount)
}

type CampaignImage struct {
	ID          int
	Campaign_ID int
	File_Name   string
	Is_Primary  int
	Created_At  time.Time
	Updated_At  time.Time
}
