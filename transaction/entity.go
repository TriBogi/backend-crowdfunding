package transaction

import (
	"bogistartup/campaign"
	"bogistartup/user"
	"time"
)

type Transaction struct {
	ID          int
	Campaign_Id int
	User_Id     int
	Amount      int
	Status      string
	Code        string
	User        user.User
	Campaign    campaign.Campaign
	Created_At  time.Time
	Updated_At  time.Time
}
