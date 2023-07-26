package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var UserProfilesOverviewField = &graphql.Field{
	Type:        types.UserProfilesOverviewType,
	Description: "Returns a data of User Profiles for displaying on Overview screen",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"job_position_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"is_active": &graphql.ArgumentConfig{
			Type:         graphql.Boolean,
			DefaultValue: nil,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.UserProfilesOverviewResolver,
}

var UserProfileContractsField = &graphql.Field{
	Type:        types.UserProfileContractsType,
	Description: "Returns a data of User Profile's contracts",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.UserProfileContractsResolver,
}

var UserProfileBasicField = &graphql.Field{
	Type:        types.UserProfileBasicType,
	Description: "Returns a data of User Profile for displaying inside Basic tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileBasicResolver,
}

var UserProfileBasicInsertField = &graphql.Field{
	Type:        types.UserProfileBasicInsertType,
	Description: "Inserts a data of User Profile for displaying inside Basic tab",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileBasicInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileBasicInsertResolver,
}

var UserProfileUpdateField = &graphql.Field{
	Type:        types.UserProfileBasicInsertType,
	Description: "Updates a data of User Profile",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileUpdateMutation),
		},
	},
	Resolve: resolvers.UserProfileUpdateResolver,
}

var UserProfileContractInsertField = &graphql.Field{
	Type:        types.UserProfileContractInsertType,
	Description: "Inserts or updates contract of User Profile",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileContractInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileContractInsertResolver,
}

var UserProfileContractDeleteField = &graphql.Field{
	Type:        types.UserProfileContractDeleteType,
	Description: "Deletes existing User Profile's Contract",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.UserProfileContractDeleteResolver,
}

var UserProfileEducationField = &graphql.Field{
	Type:        types.UserProfileEducationType,
	Description: "Returns a data of User Profile for displaying inside Education tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileEducationResolver,
}

var UserProfileEducationInsertField = &graphql.Field{
	Type:        types.UserProfileEducationInsertType,
	Description: "Creates new or alter existing User Profile's Education item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileEducationInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileEducationInsertResolver,
}

var UserProfileEducationDeleteField = &graphql.Field{
	Type:        types.UserProfileEducationDeleteType,
	Description: "Deletes existing User Profile's Education",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.UserProfileEducationDeleteResolver,
}

var UserProfileExperienceField = &graphql.Field{
	Type:        types.UserProfileExperienceType,
	Description: "Returns a data of User Profile for displaying inside Experience tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileExperienceResolver,
}

var UserProfileExperienceInsertField = &graphql.Field{
	Type:        types.UserProfileExperienceInsertType,
	Description: "Creates new or alter existing User Profile's Experience item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileExperienceInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileExperienceInsertResolver,
}

var UserProfileExperienceDeleteField = &graphql.Field{
	Type:        types.UserProfileExperienceDeleteType,
	Description: "Deletes existing User Profile's Experience",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileExperienceDeleteResolver,
}

var UserProfileFamilyField = &graphql.Field{
	Type:        types.UserProfileFamilyType,
	Description: "Returns a data of User Profile for displaying inside Family tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileFamilyResolver,
}

var UserProfileFamilyInsertField = &graphql.Field{
	Type:        types.UserProfileFamilyInsertType,
	Description: "Creates new or alter existing User Profile's Family item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileFamilyInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileFamilyInsertResolver,
}

var UserProfileFamilyDeleteField = &graphql.Field{
	Type:        types.UserProfileFamilyDeleteType,
	Description: "Deletes existing User Profile's Family",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileFamilyDeleteResolver,
}
