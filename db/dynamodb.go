package db

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var (
	Gdb    *dynamo.DB
	region string

	UsersTable string
)

func init() {
	region = os.Getenv("AWS_REGION")

	UsersTable = os.Getenv("DYNAMO_TABLE_USERS")
	if UsersTable == "" {
		log.Fatal("missing env variable: DYNAMO_TABLE_USERS")
	}

	Gdb = dynamo.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	})))
}
