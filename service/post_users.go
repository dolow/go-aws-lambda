package service

import (
	"smith.lambda/swagger-generate/models"
	"smith.lambda/swagger-generate/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func PostUsers(p operations.PostUsersParams) middleware.Responder {
	var resp models.Users
	return operations.NewPostUsersOK().WithPayload(resp)
}
