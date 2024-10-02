# Infositel Order Request API

Bienvenido a la API de Infositel Order Request (InfositelOR), una API desarrollada en Golang utilizando el framework Echo. Esta API está diseñada para gestionar solicitudes de órdenes de manera eficiente y rápida.

## Características

- Creación de órdenes: Permite crear nuevas órdenes.
- Consulta de órdenes: Permite consultar el estado y los detalles de las órdenes existentes.
- Actualización de órdenes: Permite actualizar la información de órdenes existentes.
- Eliminación de órdenes: Permite eliminar órdenes.

## Requisitos

- Go 1.22+
- Echo Framework v4

## Instalación

1. Clonar el repositorio:

   git clone https://github.com/AndySantisteban/InfositelOR

2. Ir al directorio del proyecto:

   cd infositelor

3. Instalar las dependencias:

   go mod download

## Uso

1. Iniciar el servidor:

   go run cmd/main.go

## Estructura del Proyecto

- cmd/main.go: Punto de entrada de la aplicación.
- pkg/application: Contiene la implementación de las peticiones con la DB.
- pkg/domain: Contiene las definiciones de los modelos de datos.
- pkg/infrastructure: Contiene las llamadas a la base de datos y conexión con la base de datos.
- pkg/interface: Contiene la aplicación web y rest API con proxy a ["/api/*"]

## Contribuir

Si deseas contribuir a este proyecto, por favor sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una nueva rama (feature/nueva-funcionalidad).
3. Realiza tus cambios.
4. Envía un pull request.

Gracias por utilizar InfositelOR API. Si tienes alguna pregunta o sugerencia, no dudes en abrir un issue o enviar un pull request.
