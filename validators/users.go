package usersvalidator

type PostUserBody struct {
	ID        string `json:"id,omitempty" binding:"omitempty,uuid"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required,email"`
	Password  string `json:"password" binding:"required,gte=10"`
	Type      string `json:"type,omitempty" binding:"isdefault"`
}

type PostUserParams struct {
	Name string `uri:"name" binding:"required,gte=3"`
}
