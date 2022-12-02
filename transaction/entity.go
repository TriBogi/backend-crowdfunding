package transaction

import "time"

type Transaction struct {
	ID          int
	Campaign_Id int
	User_Id     int
	Amount      int
	Status      string
	Code        string
	Created_At  time.Time
	Updated_At  time.Time
}
