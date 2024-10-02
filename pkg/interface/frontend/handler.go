package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	typescriptify "github.com/tompston/gut"

	features_activity "InfositelOR/pkg/application/features/activity/commands/CreateActivity"
	get_activity "InfositelOR/pkg/application/features/activity/queries/GetActivityByUid"
	list_activity "InfositelOR/pkg/application/features/activity/queries/ListActivityByOrderID"
	client_commands "InfositelOR/pkg/application/features/client/commands"
	client_features "InfositelOR/pkg/application/features/client/queries"
	comment_commands "InfositelOR/pkg/application/features/comment/commands"
	comment_features "InfositelOR/pkg/application/features/comment/queries"
	equipment_features "InfositelOR/pkg/application/features/equipement/queries"
	equipment_type_features "InfositelOR/pkg/application/features/equipement_type/queries"
	order_commands "InfositelOR/pkg/application/features/order/commands"
	order_features "InfositelOR/pkg/application/features/order/queries"
	order_status_commands "InfositelOR/pkg/application/features/order_status/commands"
	order_status_features "InfositelOR/pkg/application/features/order_status/queries"
	user_features "InfositelOR/pkg/application/features/user/queries"
	"InfositelOR/pkg/interface/api/rest/routes/activity"
	"InfositelOR/pkg/interface/api/rest/routes/client"
	"InfositelOR/pkg/interface/api/rest/routes/equipment"
	"InfositelOR/pkg/interface/api/rest/routes/equipment_type"
	"InfositelOR/pkg/interface/api/rest/routes/order"
	"InfositelOR/pkg/interface/api/rest/routes/user"
)

var (
	//go:embed dist/*
	dist embed.FS
)

func RegisterModels() {
	// concat all of the interfaces together
	interfaces := fmt.Sprintln(
		typescriptify.Convert(features_activity.CreateActivityQuery{}),
		typescriptify.Convert(features_activity.CreateActivityResponse{}),
		typescriptify.Convert(get_activity.GetActivityByUidQuery{}),
		typescriptify.Convert(get_activity.GetActivityByUidResponse{}),
		typescriptify.Convert(list_activity.ListActivityByOrderIDQuery{}),
		typescriptify.Convert(list_activity.ListActivityByOrderIDResponse{}),
		typescriptify.Convert(client_features.GetCliendByIDResponse{}),
		typescriptify.Convert(client_features.GetClientByIDQuery{}),
		typescriptify.Convert(client_features.ListAllClientsQuery{}),
		typescriptify.Convert(client_features.ListAllClientsResponse{}),
		typescriptify.Convert(client_commands.CreateClientQuery{}),
		typescriptify.Convert(client_commands.UpdateClientQuery{}),
		typescriptify.Convert(client_commands.DeleteClientByIDQuery{}),
		typescriptify.Convert(comment_features.GetCommentByIDQuery{}),
		typescriptify.Convert(comment_features.ListCommentsByOrderIDQuery{}),
		typescriptify.Convert(comment_features.GetCommentByIDQueryResponse{}),
		typescriptify.Convert(comment_features.ListCommentsByOrderIDQueryResponse{}),
		typescriptify.Convert(comment_commands.UpdateCommentQuery{}),
		typescriptify.Convert(comment_commands.CreateCommentQuery{}),
		typescriptify.Convert(comment_commands.DeleteCommentQuery{}),
		typescriptify.Convert(order_status_features.GetOrderStatusByIDQuery{}),
		typescriptify.Convert(order_status_features.GetOrderStatusByIDQueryResponse{}),
		typescriptify.Convert(order_status_features.ListOrderStatusQuery{}),
		typescriptify.Convert(order_status_features.ListOrderStatusQueryResponse{}),
		typescriptify.Convert(order_status_commands.CreateOrderStatusQuery{}),
		typescriptify.Convert(order_status_commands.CreateOrderStatusQueryResponse{}),
		typescriptify.Convert(order_status_commands.UpdateOrderStatusQuery{}),
		typescriptify.Convert(order_features.GetOrderQuery{}),
		typescriptify.Convert(order_features.ListOrdersQuery{}),
		typescriptify.Convert(order_features.ListOrdersQueryResponse{}),
		typescriptify.Convert(order_features.GetOrderResponse{}),
		typescriptify.Convert(order_commands.CreateOrderQuery{}),
		typescriptify.Convert(order.CreateRouteDTO{}),
		typescriptify.Convert(order_commands.CreateOrderQueryResponse{}),
		typescriptify.Convert(order_commands.UpdateOrderQuery{}),
		typescriptify.Convert(order.UpdateRouteDTO{}),
		typescriptify.Convert(user_features.GetUserByIDQuery{}),
		typescriptify.Convert(user_features.GetUserByIDQueryResponse{}),
		typescriptify.Convert(user_features.ListUserQuery{}),
		typescriptify.Convert(user_features.ListUserQueryResponse{}),
		typescriptify.Convert(client.CreateClientRouteDTO{}),
		typescriptify.Convert(user.CreateUserDTO{}),
		typescriptify.Convert(activity.CreateActivityQuery{}),
		typescriptify.Convert(activity.ListActivityByOrderIDRouteDTO{}),
		typescriptify.Convert(activity.GetActivityByUidRouteDTO{}),

		typescriptify.Convert(equipment_features.GetEquipmentQueryResponse{}),
		typescriptify.Convert(equipment_features.ListAllEquipmentResponse{}),
		typescriptify.Convert(equipment.GetEquipmentQueryDTO{}),
		typescriptify.Convert(equipment.ListAllEquipmentQueryDTO{}),
		typescriptify.Convert(equipment.CreateEquipmentQueryDTO{}),
		typescriptify.Convert(equipment.DeleteEquipmentTypeQueryDTO{}),

		typescriptify.Convert(equipment_type_features.ListAllEquipmentTypeResponse{}),
		typescriptify.Convert(equipment_type_features.GetEquipmentTypeQueryResponse{}),
		typescriptify.Convert(equipment_type.GetEquipmentTypeQueryDTO{}),
		typescriptify.Convert(equipment_type.ListAllEquipmentTypeQueryDTO{}),
		typescriptify.Convert(equipment_type.CreateEquipmentTypeQueryDTO{}),
		typescriptify.Convert(equipment_type.DeleteEquipmentTypeRouteDTO{}),

		// typescriptify.Convert(&client_commands.CreateClientQueryResponse{}),
	)

	if err := typescriptify.Generate("pkg/interface/frontend/web/src/models/api-models.ts", interfaces); err != nil {
		fmt.Println(err)
	}
}
func RegisterHandlers(e *echo.Echo) {
	if os.Getenv("ENV") == "dev" {
		RegisterModels()
		log.Println("Running in dev mode")
		setupDevProxy(e)
		return
	}
	distDirFS, err := fs.Sub(dist, "dist")
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.GET("/*", func(c echo.Context) error {
		path := c.Request().URL.Path
		if strings.HasPrefix(path, "/api") {
			return echo.NewHTTPError(http.StatusNotFound, "API endpoint not found")
		}

		if _, err := fs.Stat(distDirFS, strings.TrimPrefix(path, "/")); err == nil {
			return echo.WrapHandler(http.StripPrefix("/", http.FileServer(http.FS(distDirFS))))(c)
		}

		return c.Blob(http.StatusOK, "text/html", getFileContents(distDirFS, "index.html"))
	})
}
func getFileContents(fs fs.FS, name string) []byte {
	file, err := fs.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}

	return data
}

func setupDevProxy(e *echo.Echo) {
	url, err := url.Parse("http://localhost:5173")
	if err != nil {
		log.Fatal(err)
	}
	balancer := middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: url,
		},
	})
	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: balancer,
		Skipper: func(c echo.Context) bool {
			// Skip the proxy if the prefix is /api
			return len(c.Path()) >= 4 && c.Path()[:4] == "/api"
		},
	}))
}
