package user

import "time"

type ListUsersRouteDTO struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

type GetUserByUidRouteDTO struct {
	Id int32 `query:"id"`
}

type CreateUserDTO struct {
	ID             int32      `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	PermissionType string     `json:"permission_type"`
	CreationDate   *time.Time `json:"creation_date"`
	InactiveStatus *string    `json:"inactive_status"`
	Password       string     `json:"password"`
}

// type UpdateOrderStatusRouteDTO struct {
// 	Id   int32  `json:"id"`
// 	Name string `json:"name"`
// }
