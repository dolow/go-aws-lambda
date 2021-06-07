package service

import (
	"context"
  "encoding/json"
  "fmt"
	"net/http/httptest"
	"testing"

  "smith.lambda/db"
	"smith.lambda/swagger-generate/restapi/operations"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
  "github.com/go-openapi/runtime"
	"github.com/guregu/dynamo"

  "github.com/stretchr/testify/assert"
)

func init() {
	dbEndpoint := "http://localhost:4566"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "dummy",
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint:   aws.String(dbEndpoint),
			DisableSSL: aws.Bool(true),
		},
	}))
	db.Gdb = dynamo.New(sess)
}

func TestGetUsers(t *testing.T) {
  err := db.Gdb.CreateTable(db.UsersTable, User{}).Provision(1, 1).RunWithContext(context.TODO())
  if err != nil {
    assert.Fail(t, fmt.Sprintf("dynamo create table %s: %v", db.UsersTable, err))
  }

  t.Cleanup(func() {
    if err := db.Gdb.Table(db.UsersTable).DeleteTable().RunWithContext(context.TODO()); err != nil {
      t.Fatalf("dynamo delete table %s: %v", db.UsersTable, err)
    }
  })

  subject := func() *httptest.ResponseRecorder {
    p := operations.NewGetUsersParams()
    p.HTTPRequest = httptest.NewRequest("GET", "/v1/users", nil)
    resp := GetUsers(p)

    w := httptest.NewRecorder()
    resp.WriteResponse(w, runtime.JSONProducer())
    return w
  }

  t.Run("When record exists", func(t *testing.T) {
    expect := []User{{Id: "001", Name: "gopher"}, {Id: "002", Name: "rubyist"}}
    for _, u := range expect {
      if err := db.Gdb.Table(db.UsersTable).Put(u).RunWithContext(context.TODO()); err != nil {
        assert.Fail(t, fmt.Sprintf("dynamo input user %v: %v", u, err))
      }
    }

    t.Cleanup(func() {
      for _, user := range expect {
        if err := db.Gdb.Table(db.UsersTable).Delete("id", user.Id).RunWithContext(context.TODO()); err != nil {
          t.Fatalf("dynamo delete table %s: %v", db.UsersTable, err)
        }
      }
    })

    t.Run("Should return all users", func(t *testing.T) {
      w := subject()

      actual := []User{}
      err := json.Unmarshal(w.Body.Bytes(), &actual)
      if err != nil {
        assert.Fail(t, fmt.Sprintf("%s", err))
      }

      assert.Equal(t, w.Result().StatusCode, 200)

      for i, u := range actual {
        assert.Equal(t, u.Id, expect[i].Id)
        assert.Equal(t, u.Name, expect[i].Name)
      }
    })
  })

  t.Run("When record does not exist", func(t *testing.T) {
    w := subject()

    actual := []User{}
    err := json.Unmarshal(w.Body.Bytes(), &actual)
    if err != nil {
      assert.Fail(t, fmt.Sprintf("%s", err))
    }

    assert.Equal(t, w.Result().StatusCode, 200)
    assert.Equal(t, len(actual), 0)
  })
}
