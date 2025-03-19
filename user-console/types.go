package userconsole

type User struct{
	Uid string `json:"uid"`
	Email string `json:"email"`
	Name string `json:"name"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
	TotalCount int `json:"totalCount"`
}