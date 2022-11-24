package campaign

import "time"

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
}
