package service

type User struct {
	Id   string `dynamo:"id,hash"`
	Name string `dynamo:"name"`
}
