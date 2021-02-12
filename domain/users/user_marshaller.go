package users

type PublicUser struct {
	Id       int64  `json:"user_id"`
	CreateAt string `json:"create_at"`
	Status   string `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreateAt  string `json:"create_at"`
	Status    string `json:"status"`
}

func (u *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:       u.Id,
			CreateAt: u.CreateAt,
			Status:   u.Status,
		}
	}
	return PrivateUser{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreateAt:  u.CreateAt,
		Status:    u.Status,
	}
}

func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}
