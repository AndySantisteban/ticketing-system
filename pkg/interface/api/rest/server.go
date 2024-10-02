package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/swaggo/echo-swagger/example/docs"

	"InfositelOR/pkg/interface/api/rest/routes/activity"
	"InfositelOR/pkg/interface/api/rest/routes/client"
	"InfositelOR/pkg/interface/api/rest/routes/comment"
	"InfositelOR/pkg/interface/api/rest/routes/equipment"
	"InfositelOR/pkg/interface/api/rest/routes/equipment_type"
	"InfositelOR/pkg/interface/api/rest/routes/login"
	"InfositelOR/pkg/interface/api/rest/routes/order"
	orderstatus "InfositelOR/pkg/interface/api/rest/routes/order_status"
	"InfositelOR/pkg/interface/api/rest/routes/public"
	"InfositelOR/pkg/interface/api/rest/routes/user"
	"InfositelOR/pkg/interface/frontend"
)

type Server struct {
	HttpClient    *echo.Echo
	activity      *activity.ActivityRouter
	client        *client.ClientRouter
	comment       *comment.CommentRouter
	order         *order.OrderRouter
	order_status  *orderstatus.OrderStatusRouter
	user          *user.UserRouter
	login         *login.LoginRouter
	equipment     *equipment.EquipmentRouter
	equipmentType *equipment_type.EquipmentTypeRouter
	public        *public.PublicRouter
}

func NewRestServer(
	e *echo.Echo,
	activity *activity.ActivityRouter,
	client *client.ClientRouter,
	comment *comment.CommentRouter,
	order *order.OrderRouter,
	order_status *orderstatus.OrderStatusRouter,
	user *user.UserRouter,
	login *login.LoginRouter,
	equipment *equipment.EquipmentRouter,
	equipmentType *equipment_type.EquipmentTypeRouter,
	public *public.PublicRouter,
) *Server {
	server := Server{
		HttpClient:    e,
		activity:      activity,
		client:        client,
		comment:       comment,
		order:         order,
		order_status:  order_status,
		user:          user,
		login:         login,
		equipment:     equipment,
		equipmentType: equipmentType,
		public:        public,
	}
	server.HttpClient.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}, \n",
	}), middleware.Recover())
	go server.RegisterRoutes()

	return &server
}

func (h *Server) RegisterRoutes() {
	// region ActivityRoutes
	activity := h.HttpClient.Group("/api/Activity")
	activity.Use(login.ProtectedHandler)
	activity.GET("/List", h.activity.ListActivityByOrderIDRoute)
	activity.GET("/Get", h.activity.GetActivityByUid)
	activity.POST("/Create", h.activity.CreateActivity)
	// endregion

	// region ClientRoutes
	client := h.HttpClient.Group("/api/Client")
	client.Use(login.ProtectedHandler)
	client.GET("/List", h.client.ListClientRoute)
	client.GET("/Get", h.client.GetClientByUid)
	client.POST("/Create", h.client.CreateClient)
	client.PUT("/Update", h.client.UpdateClient)
	client.DELETE("/Delete", h.client.DeleteClient)
	// endregion

	// region CommentRoutes
	comment := h.HttpClient.Group("/api/Comment")
	comment.Use(login.ProtectedHandler)
	comment.GET("/List", h.comment.ListCommentRoute)
	comment.GET("/Get", h.comment.GetCommentRouteByUid)
	comment.POST("/Create", h.comment.CreateComment)
	comment.PUT("/Update", h.comment.UpdateComment)
	comment.DELETE("/Delete", h.comment.DeleteComment)
	// endregion

	// region OrderRoutes
	order := h.HttpClient.Group("/api/Order")
	order.Use(login.ProtectedHandler)
	order.GET("/List", h.order.ListOrdersRoute)
	order.GET("/Get", h.order.GetOrderRouteByUid)
	order.PUT("/Update", h.order.UpdateRouteByUid)
	order.POST("/Create", h.order.CreateRouteByUid)

	public := h.HttpClient.Group("/api/Public")
	public.GET("/Tracking", h.public.GetTracking)
	public.POST("/SendSupportMessage", h.public.SendSupportMessage)

	// order.DELETE("/DeleteComment", h.comment.DeleteComment)

	// region OrderStatusRoutes
	orderStatus := h.HttpClient.Group("/api/OrderStatus")
	orderStatus.Use(login.ProtectedHandler)
	orderStatus.GET("/Get/:ID", h.order_status.GetOrderStatusRouteByUid)
	orderStatus.GET("/List", h.order_status.ListOrdersStatusRoute)
	orderStatus.PUT("/Update", h.order_status.UpdateOrderStatusRouteByUid)
	orderStatus.POST("/Create", h.order_status.CreateOrderStatusRouteByUid)
	// order.DELETE("/DeleteComment", h.comment.DeleteComment)

	// region OrderStatusRoutes
	user := h.HttpClient.Group("/api/User")
	user.Use(login.ProtectedHandler)
	user.GET("/List", h.user.ListUsersRoute)
	user.GET("/Get", h.user.GetUserRouteByUid)
	user.POST("/Create", h.user.CreateUserRoute)

	loginRoute := h.HttpClient.Group("/api/login")
	loginRoute.POST("/auth", h.login.AuthLogin)
	loginRoute.GET("/userType", h.login.UserType)
	loginRoute.GET("/userName", h.login.UserName)
	// user.PUT("/CreateOrderStatusRouteByUid", h.order_status.CreateOrderStatusRouteByUid)
	// order.DELETE("/DeleteComment", h.comment.DeleteComment)

	// region EquipmentRoutes
	equipment := h.HttpClient.Group("/api/Equipment")
	equipment.Use(login.ProtectedHandler)
	equipment.GET("/Get", h.equipment.GetByID)
	equipment.GET("/List", h.equipment.List)
	// equipment.PUT("/Update", h.order_status.UpdateOrderStatusRouteByUid)
	equipment.POST("/Create", h.equipment.Create)
	equipment.DELETE("/Delete", h.equipment.Delete)

	// region EquipmentRoutes
	equipmentType := h.HttpClient.Group("/api/EquipmentType")
	equipmentType.Use(login.ProtectedHandler)
	equipmentType.GET("/Get", h.equipmentType.GetByID)
	equipmentType.GET("/List", h.equipmentType.List)
	// equipment.PUT("/Update", h.order_status.UpdateOrderStatusRouteByUid)
	equipmentType.POST("/Create", h.equipmentType.Create)
	equipmentType.DELETE("/Delete", h.equipmentType.Delete)

	frontend.RegisterHandlers(h.HttpClient)

	h.HttpClient.Start(":8080")
}
