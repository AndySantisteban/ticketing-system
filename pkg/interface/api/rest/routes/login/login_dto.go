package login

type LoginRouteDTO struct {
	Username string `query:"username"`
	Password string `query:"password"`
}
