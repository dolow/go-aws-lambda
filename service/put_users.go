package service

import (
	"smith.lambda/swagger-generate/models"
	"smith.lambda/swagger-generate/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func PutUsers(p operations.PutUsersParams) middleware.Responder {
	var resp models.Users
	return operations.NewPutUsersOK().WithPayload(resp)
}
