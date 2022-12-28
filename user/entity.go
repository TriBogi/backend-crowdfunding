package user

import "time"

type User struct {
	ID               int       //`json:"id"`
	Name             string    //`json:"name"`
	Ocupation        string    //`json:"occupation"`
	Email            string    //`json:"email"`
	Password_Hash    string    //`json:"password"`
	Avatar_File_Name string    //`json:"avatarFileName"`
	Roleuser         string    //`json:"roleuser"`
	Created_At       time.Time //`json:"createdAt"`
	Updated_At       time.Time //`json:"updatedAt"`
}
