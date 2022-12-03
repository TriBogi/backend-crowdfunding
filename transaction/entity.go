package transaction

import (
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
	Created_At  time.Time
	Updated_At  time.Time
}
