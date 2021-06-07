package service

import (
	"smith.lambda/swagger-generate/models"
	"smith.lambda/swagger-generate/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func DeleteUsers(p operations.DeleteUsersParams) middleware.Responder {
	var resp models.Users
	return operations.NewDeleteUsersOK().WithPayload(resp)
}
