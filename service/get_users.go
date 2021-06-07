package service

import (
	"fmt"

	"smith.lambda/swagger-generate/models"
	"smith.lambda/swagger-generate/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetUsers(p operations.GetUsersParams) middleware.Responder {
	ctx := p.HTTPRequest.Context()
	users, err := RepositoryGetUsers(ctx)
	if err != nil {
		return operations.NewGetUsersInternalServerError().WithPayload(&models.Error{
			Message: fmt.Sprintf("scan users error: %v", err),
		})
	}
	var resp models.Users
	for _, u := range users {
		u := u
		resp = append(resp, &models.User{
			ID:   &u.Id,
			Name: &u.Name,
		})
	}
	return operations.NewGetUsersOK().WithPayload(resp)
}
