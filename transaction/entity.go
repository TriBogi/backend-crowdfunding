package transaction

import (
	"bogistartup/campaign"
	"bogistartup/user"
	"github.com/leekchan/accounting"
	"time"
)

type Transaction struct {
	ID          int
	Campaign_Id int
	User_Id     int
	Amount      int
	Status      string
	Code        string
	PaymentURL  string
	User        user.User
	Campaign    campaign.Campaign
	Created_At  time.Time
	Updated_At  time.Time
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
