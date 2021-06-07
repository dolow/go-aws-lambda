package service

import (
	"context"
	"errors"

  "smith.lambda/db"

	"github.com/guregu/dynamo"
)

func RepositoryGetUsers(ctx context.Context) ([]User, error) {
	var resp []User
  // TODO
	table := db.Gdb.Table(db.UsersTable)
	if err := table.Scan().AllWithContext(ctx, &resp); err != nil {
		// 0件の場合も正常とします
		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, nil
		}
		return resp, err
	}
	return resp, nil
}
