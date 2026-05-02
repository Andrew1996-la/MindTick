package users_transport_http

type UserService interface {

}

type UserHttpHandler struct {
	userService UserService
}

func NewUserHttpHandler(userService UserService) *UserHttpHandler {
	return &UserHttpHandler{
		userService: userService,
	}
}