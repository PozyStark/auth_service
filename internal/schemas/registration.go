package schemas

//TODO: Необходимо настроить валидацию для пользовательских данных

type RegistrationSchema struct {
	Username    string `form:"username" json:"username" binding:"required"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Email       string `form:"email" json:"email"`
	FirstName   string `form:"first_name" json:"first_name"`
	MiddleName  string `form:"middle_name" json:"middle_name"`
	LastName    string `form:"last_name" json:"last_name"`
	Password    string `form:"password" json:"password"`
	Age         int    `form:"age" json:"age" `
}

type UpdateUserBoolean struct {
	ID          string `json:"id"`
	IsSuperuser bool   `json:"is_superuser"`
	IsActive    bool   `json:"is_active"`
}

type UpdateUserSchema struct {
	ID          string `json:"id"`
	Username    string `form:"username" json:"username"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Email       string `form:"email" json:"email"`
	FirstName   string `form:"first_name" json:"first_name"`
	MiddleName  string `form:"middle_name" json:"middle_name"`
	LastName    string `form:"last_name" json:"last_name"`
	Password    string `form:"password" json:"password"`
	Age         int    `form:"age" json:"age"`
}
