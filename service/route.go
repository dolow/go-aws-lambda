package service

import (
  "smith.lambda/swagger-generate/restapi/operations"
)

func Route(api *operations.LambdaAppAPI) {
  api.GetUsersHandler = operations.GetUsersHandlerFunc(GetUsers)
  api.PostUsersHandler = operations.PostUsersHandlerFunc(PostUsers)
  api.PutUsersHandler = operations.PutUsersHandlerFunc(PutUsers)
  api.DeleteUsersHandler = operations.DeleteUsersHandlerFunc(DeleteUsers)
}
