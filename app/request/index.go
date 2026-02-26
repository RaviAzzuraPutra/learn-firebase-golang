package request

type Member_Request struct {
	Name  *string `form:"name"`
	Email *string `form:"email"`
	Phone *string `form:"phone"`
}
