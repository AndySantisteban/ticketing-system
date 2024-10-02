package public

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/labstack/echo/v4"

	client_features "InfositelOR/pkg/application/features/client/queries"
	comment_features "InfositelOR/pkg/application/features/comment/queries"
	equipment_features "InfositelOR/pkg/application/features/equipement/queries"
	equipment_type_features "InfositelOR/pkg/application/features/equipement_type/queries"
	order_features "InfositelOR/pkg/application/features/order/queries"
	order_status_features "InfositelOR/pkg/application/features/order_status/queries"
	user_features "InfositelOR/pkg/application/features/user/queries"
	"InfositelOR/pkg/domain/entities"
)

type PublicRouter struct {
	getOrderById         *order_features.GetOrderQueryHandler
	getClientById        *client_features.GetClientByIDQueryHandler
	getUserById          *user_features.GetUserByIDQueryHandler
	getCommentsList      *comment_features.ListCommentsByOrderIDQueryHandler
	getOrderStatusById   *order_status_features.GetOrderStatusByIDQueryHandler
	getEquipmentById     *equipment_features.GetEquipmentQueryHandler
	getEquipmentTypeById *equipment_type_features.GetEquipmentTypeQueryHandler
}

func RestPublicRouter(

	getOrderById *order_features.GetOrderQueryHandler,
	getClientById *client_features.GetClientByIDQueryHandler,
	getUserById *user_features.GetUserByIDQueryHandler,
	getCommentsList *comment_features.ListCommentsByOrderIDQueryHandler,
	getOrderStatusById *order_status_features.GetOrderStatusByIDQueryHandler,
	getEquipmentById *equipment_features.GetEquipmentQueryHandler,
	getEquipmentTypeById *equipment_type_features.GetEquipmentTypeQueryHandler,

) *PublicRouter {

	return &PublicRouter{
		getOrderById,
		getClientById,
		getUserById,
		getCommentsList,
		getOrderStatusById,
		getEquipmentById,
		getEquipmentTypeById,
	}

}

func (p *PublicRouter) GetTracking(c echo.Context) error {
	query := new(GetTrackingInfoDTO)
	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	order, err := p.getOrderById.Handler(order_features.GetOrderQuery{
		Id: query.Id,
	})
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}
	client, _ := p.getClientById.Handler(client_features.GetClientByIDQuery{
		Id: order.ClientID.Int32,
	})
	user, _ := p.getUserById.Handler(user_features.GetUserByIDQuery{
		ID: order.AssignedTo.Int32,
	})
	equipement, _ := p.getEquipmentById.Handler(equipment_features.GetEquipmentQuery{
		ID: order.EquipmentID.Int32,
	})
	equipementType, _ := p.getEquipmentTypeById.Handler(equipment_type_features.GetEquipmentTypeQuery{
		ID: equipement.TypeID.Int32,
	})
	comments, _ := p.getCommentsList.Handler(comment_features.ListCommentsByOrderIDQuery{
		Offset:  0,
		Limit:   10000,
		OrderID: order.ID,
	})

	tracking := map[string]interface{}{
		"User": entities.User{
			Name:  user.Name,
			Email: user.Email,
		},
		"Order": order,
		"Client": entities.Client{
			Name:  client.Name,
			Email: client.Email,
		},
		"Comments":  comments,
		"Equipment": equipement,
		"EquipmentType": entities.EquipmentType{
			Name: equipementType.Name,
		},
	}

	return c.JSON(200, tracking)
}

func enviarCorreo(para string, asunto string, mensaje string) error {
	// Configuración de Gmail
	from := os.Getenv("EMAIL_GMAIL")        // Correo del remitente (tu correo de Gmail)
	password := os.Getenv("EMAIL_PASSWORD") // Contraseña del correo o token de aplicación
	fmt.Println(from)
	fmt.Println(password)
	// Servidor SMTP de Gmail
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Autenticación
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Cuerpo del correo
	to := []string{para}
	message := []byte("To: " + para + "\r\n" +
		"Subject: " + asunto + "\r\n" +
		"\r\n" +
		mensaje + "\r\n")

	// Enviar el correo
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}

func (p *PublicRouter) SendSupportMessage(c echo.Context) error {
	solicitud := new(TecSupportDTO)
	if err := c.Bind(solicitud); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
	}

	// Enviar el correo de soporte técnico
	err := enviarCorreo(solicitud.Correo, solicitud.Asunto, solicitud.Mensaje)
	if err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo enviar el correo"})
	}

	return c.JSON(http.StatusOK, echo.Map{"mensaje": "Correo enviado exitosamente"})

}
