package dataloader

import (
	"context"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/ent/user"
	"strconv"

	"github.com/graph-gophers/dataloader"
)

type UserReader struct {
	Client *ent.Client
}

// GetUsers implements a batch function that can retrieve many users by ID,
// for use in a dataloader
func (u *UserReader) GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// read all requested users in a single query
	userIDs := make([]int, len(keys))
	for ix, key := range keys {
		userIDs[ix], _ = strconv.Atoi(key.String())
	}
	userById ,_ := u.Client.User.Query().Select(user.FieldID, user.FieldName).Where(user.IDIn(userIDs...)).All(ctx)
	// return users in the same order requested
	output := make([]*dataloader.Result, len(keys))
	for index, _ := range keys {
		user := userById[index]
		output[index] = &dataloader.Result{Data: user, Error: nil}
		// if ok {
		// 	output[index] = &dataloader.Result{Data: user, Error: nil}
		// } else {
		// 	err := fmt.Errorf("user not found %s", userKey.String())
		// 	output[index] = &dataloader.Result{Data: nil, Error: err}
		// }
	}

	return output
}

func LoadUser(ctx context.Context, userID string) ([]*ent.User, error) {
	loaders := For(ctx)

	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()

	if err != nil {
		return nil, err
	}
	users, _ := result.(*ent.User)

	return []*ent.User{users}, nil
}