package login

import (
	user_commands "InfositelOR/pkg/application/features/user/commands"
	user_features "InfositelOR/pkg/application/features/user/queries"
	"InfositelOR/pkg/domain/entities"
	"database/sql"
	"net/http"

	"github.com/ahmetb/go-linq/v3"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type MyClaims struct {
	jwt.Claims
	User string `json:"user"`
}

var secretKey = []byte("secret-key")

func createToken(username entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": username,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserInfoFromToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		userClaim := (*claims)["user"].(map[string]interface{})
		return userClaim, nil
	}

	return nil, fmt.Errorf("UserID not found")
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ProtectedHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(echo.ErrUnauthorized.Code, "header Authorizathion not found")
		}
		tokenString = tokenString[len("Bearer "):]
		err := verifyToken(tokenString)
		if err != nil {
			return c.JSON(echo.ErrUnauthorized.Code, "Invalid token")
		}
		return next(c)
	}
}

type LoginRouter struct {
	getID *user_features.GetUserByIDQueryHandler
	list  *user_features.ListUserQueryHandler
}

func NewLoginRouter(
	getID *user_features.GetUserByIDQueryHandler,
	list *user_features.ListUserQueryHandler,
	create *user_commands.CreateUserQueryHandler,
) *LoginRouter {

	return &LoginRouter{
		getID: getID,
		list:  list,
	}
}

func (ar *LoginRouter) UserType(c echo.Context) error {

	tokenString := c.Request().Header.Get("Authorization")

	tokenString = tokenString[len("Bearer "):]
	item, _ := GetUserInfoFromToken(tokenString)

	return c.String(200, fmt.Sprintf("%v", item["PermissionType"]))
}

func (ar *LoginRouter) UserName(c echo.Context) error {

	tokenString := c.Request().Header.Get("Authorization")

	tokenString = tokenString[len("Bearer "):]
	item, _ := GetUserInfoFromToken(tokenString)

	return c.String(200, fmt.Sprintf("%v", item["Name"]))
}

func (ar *LoginRouter) AuthLogin(c echo.Context) error {
	query := new(LoginRouteDTO)

	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	response, err := ar.list.Handler(user_features.ListUserQuery{
		Limit:  1000,
		Offset: 0,
	})

	if err != nil {
		return c.String(http.StatusBadRequest, "Users List Failed search")
	}

	users := response
	var foundUser *entities.User
	linq.From(users).
		Where(func(i interface{}) bool {
			user := i.(entities.User)
			return user.Name == query.Username && user.Password == sql.NullString{
				Valid:  true,
				String: query.Password,
			}
		}).
		FirstWith(func(i interface{}) bool {
			user := i.(entities.User)
			foundUser = &user
			return true
		})

	if foundUser == nil {
		return c.JSON(echo.ErrUnauthorized.Code, echo.ErrUnauthorized.Message)
	}

	if len(users) == 0 {
		return c.String(http.StatusBadRequest, "User not found")
	}

	tokenString, err := createToken(*foundUser)
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, tokenString)
}
