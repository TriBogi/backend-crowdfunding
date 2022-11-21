package user

import "time"

type User struct {
	ID               int       //`json:"id"`
	Name             string    //`json:"name"`
	Ocupation        string    //`json:"ocupation"`
	Email            string    //`json:"email"`
	Password_Hash    string    //`json:"password"`
	Avatar_File_Name string    //`json:"avatarFileName"`
	roleuser         string    //`json:"role"`
	Created_At       time.Time //`json:"createdAt"`
	Updated_At       time.Time //`json:"updatedAt"`
}
