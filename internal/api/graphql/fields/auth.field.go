package fields

import (
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (r *Field) LogoutField() *graphql.Field {
	return &graphql.Field{
		Type:        types.LogoutType,
		Description: "Logout the user",
		Resolve:     r.Resolvers.LogoutResolver,
	}
}

func (r *Field) LoginField() *graphql.Field {
	return &graphql.Field{
		Type:        types.LoginType,
		Description: "Returns a basic data for logged in user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: r.Resolvers.LoginResolver,
	}
}

func (f *Field) RefreshField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RefreshTokenType,
		Description: "Returns a basic data for logged in user",
		Resolve:     f.Resolvers.RefreshTokenResolver,
	}
}
func (f *Field) UserForgotPassword() *graphql.Field {
	return &graphql.Field{
		Type:        types.ForgotPasswordType,
		Description: "Sends an e-mail",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.ForgotPasswordResolver,
	}
}
func (f *Field) UserValidateMail() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserValidateMailType,
		Description: "Validate e-mail",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"token": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.UserValidateMailResolver,
	}
}
func (f *Field) UserResetPassword() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserResetPasswordType,
		Description: "Reset password",
		Args: graphql.FieldConfigArgument{
			"encrypted_email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.UserResetPasswordResolver,
	}
}
