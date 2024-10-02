package main

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	features_activity_create "InfositelOR/pkg/application/features/activity/commands/CreateActivity"
	features_activity_get "InfositelOR/pkg/application/features/activity/queries/GetActivityByUid"
	features_activity_list "InfositelOR/pkg/application/features/activity/queries/ListActivityByOrderID"
	client_commands "InfositelOR/pkg/application/features/client/commands"
	client_features "InfositelOR/pkg/application/features/client/queries"
	comment_commands "InfositelOR/pkg/application/features/comment/commands"
	comment_features "InfositelOR/pkg/application/features/comment/queries"
	equipment_commands "InfositelOR/pkg/application/features/equipement/commands"
	equipment_features "InfositelOR/pkg/application/features/equipement/queries"
	equipement_type_commands "InfositelOR/pkg/application/features/equipement_type/commands"
	equipment_type_features "InfositelOR/pkg/application/features/equipement_type/queries"
	order_commands "InfositelOR/pkg/application/features/order/commands"
	order_features "InfositelOR/pkg/application/features/order/queries"
	order_status_commands "InfositelOR/pkg/application/features/order_status/commands"
	order_status_features "InfositelOR/pkg/application/features/order_status/queries"
	user_commands "InfositelOR/pkg/application/features/user/commands"
	user_features "InfositelOR/pkg/application/features/user/queries"
	"InfositelOR/pkg/infrastructure/db/mysql/connection"
	"InfositelOR/pkg/infrastructure/repositories"
	"InfositelOR/pkg/interface/api/rest"
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
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			echo.New,
			connection.NewConnection,
			repositories.NewActivityRepository,
			repositories.NewClientRepository,
			repositories.NewCommentRepository,
			repositories.NewOrderRepository,
			repositories.NewOrderStatusRepository,
			repositories.NewUserRepository,
			repositories.NewEquipmentRepository,
			repositories.NewEquipmentTypeRepository,
			features_activity_list.NewListActivityByOrderIDQueryHandler,
			client_features.NewGetClientByIDQueryHandler,
			client_features.NewListClientsQueryHandler,
			client_commands.NewCreateClientQueryHandler,
			client_commands.NewUpdateClientQueryHandler,
			client_commands.NewDeleteClientByIDQueryHandler,
			features_activity_get.NewLGetActivityByUidQueryHandler,
			features_activity_create.NewCreateActivityQueryHandler,
			comment_features.NewGetCommentByIDQueryHandler,
			comment_features.NewListCommentsByOrderIDQueryHandler,
			comment_commands.NewCreateCommentQueryHandler,
			comment_commands.NewUpdateCommentQueryHandler,
			comment_commands.NewDeleteCommentQueryHandler,
			order_features.NewGetOrderQueryHandler,
			order_features.NewListOrdersQueryHandler,
			order_commands.NewCreateOrderQueryHandler,
			order_commands.NewUpdateOrderQQueryHandler,
			order_status_commands.NewCreateOrderStatusQueryHandler,
			order_status_commands.NewUpdateOrderStatusQueryHandler,
			order_status_features.NewGetOrderStatusByIDQueryHandler,
			order_status_features.NewListOrderStatusQueryHandler,
			user_features.NewGetUserByIDQueryHandler,
			user_features.NewListUserQueryHanlder,
			user_commands.NewCreateUserQueryHandler,
			equipment_features.NewGetEquipmentQueryHandler,
			equipment_features.NewListEquipmentQueryHandler,
			equipment_commands.NewCreateEquipmentHandler,
			equipment_commands.NewDeleteEquipmentQueryHandler,
			equipment_type_features.NewGetEquipmentQueryHandler,
			equipment_type_features.NewListEquipmentTypeQueryHandler,
			equipement_type_commands.NewCreateEquipmentTypeQueryHandler,
			equipement_type_commands.NewDeleteEquipementTypeQueryHandler,
			activity.RestActivityRouter,
			client.RestClientRouter,
			comment.RestCommentRouter,
			order.RestOrderRouter,
			orderstatus.RestOrderStatusRouter,
			user.RestUserRouter,
			login.NewLoginRouter,
			equipment.RestEquipmentRouter,
			equipment_type.RestEquipmentRouter,
			public.RestPublicRouter,
		),
		fx.Invoke(
			rest.NewRestServer,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, sql *sql.DB) {
				lc.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							return nil
						},
						OnStop: func(ctx context.Context) error {
							defer sql.Close()
							return nil
						},
					},
				)
			},
		),
	)

	app.Run()
}
