// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-gql-sample/app/ent/company"
	"go-gql-sample/app/ent/group"
	"go-gql-sample/app/ent/schema"
	"go-gql-sample/app/ent/team"
	"go-gql-sample/app/ent/todo"
	"go-gql-sample/app/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	companyFields := schema.Company{}.Fields()
	_ = companyFields
	// companyDescName is the schema descriptor for name field.
	companyDescName := companyFields[0].Descriptor()
	// company.DefaultName holds the default value on creation for the name field.
	company.DefaultName = companyDescName.Default.(string)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescName is the schema descriptor for name field.
	teamDescName := teamFields[0].Descriptor()
	// team.DefaultName holds the default value on creation for the name field.
	team.DefaultName = teamDescName.Default.(string)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[1].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = func() func(string) error {
		validators := todoDescText.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(text string) error {
			for _, fn := range fns {
				if err := fn(text); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[3].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoFields[4].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
