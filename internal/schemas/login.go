package schemas

//TODO: Необходимо настроить валидацию для пользовательских данных

type Login struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}