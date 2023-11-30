package fields

import (
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var LoginField = &graphql.Field{
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
	Resolve: resolvers.LoginResolver,
}

var LogoutField = &graphql.Field{
	Type:        types.LogoutType,
	Description: "Logout the user",
	Resolve:     resolvers.LogoutResolver,
}

var RefreshField = &graphql.Field{
	Type:        types.RefreshTokenType,
	Description: "Returns a basic data for logged in user",
	Resolve:     resolvers.RefreshTokenResolver,
}

var UserForgotPassword = &graphql.Field{
	Type:        types.ForgotPasswordType,
	Description: "Sends an e-mail",
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.ForgotPasswordResolver,
}

var UserValidateMail = &graphql.Field{
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
	Resolve: resolvers.UserValidateMailResolver,
}

var UserResetPassword = &graphql.Field{
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
	Resolve: resolvers.UserResetPasswordResolver,
}
