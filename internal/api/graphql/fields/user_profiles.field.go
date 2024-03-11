package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserProfilesOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.UserProfilesOverviewResolver,
	}
}

func (f *Field) UserProfileContractsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileContractsType,
		Description: "Returns a data of User Profile's contracts",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileContractsResolver,
	}
}

func (f *Field) UserProfileBasicField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileBasicType,
		Description: "Returns a data of User Profile for displaying inside Basic tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileBasicResolver,
	}
}

func (f *Field) UserProfileBasicInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileBasicInsertType,
		Description: "Inserts a data of User Profile for displaying inside Basic tab",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileBasicInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileBasicInsertResolver,
	}
}

func (f *Field) UserProfileUpdateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileBasicInsertType,
		Description: "Updates a data of User Profile",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileUpdateMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileUpdateResolver,
	}
}

func (f *Field) UserProfileContractInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileContractInsertType,
		Description: "Inserts or updates contract of User Profile",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileContractInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileContractInsertResolver,
	}
}

func (f *Field) UserProfileContractDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileContractDeleteType,
		Description: "Deletes existing User Profile's Contract",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileContractDeleteResolver,
	}
}

func (f *Field) UserProfileEducationField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileEducationType,
		Description: "Returns a data of User Profile for displaying inside Education tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"education_type": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.UserProfileEducationResolver,
	}
}

func (f *Field) UserProfileEducationInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileEducationInsertType,
		Description: "Creates new or alter existing User Profile's Education item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileEducationInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileEducationInsertResolver,
	}
}

func (f *Field) UserProfileEducationDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileEducationDeleteType,
		Description: "Deletes existing User Profile's Education",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileEducationDeleteResolver,
	}
}

func (f *Field) UserProfileExperienceField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileExperienceType,
		Description: "Returns a data of User Profile for displaying inside Experience tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileExperienceResolver,
	}
}

func (f *Field) UserProfileExperienceInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileExperienceInsertType,
		Description: "Creates new or alter existing User Profile's Experience item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileExperienceInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileExperienceInsertResolver,
	}
}

func (f *Field) UserProfileExperiencesInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileExperiencesInsertType,
		Description: "Creates new or alter existing User Profile's Experiences item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.UserProfileExperienceInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileExperiencesInsertResolver,
	}
}

func (f *Field) UserProfileExperienceDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileExperienceDeleteType,
		Description: "Deletes existing User Profile's Experience",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileExperienceDeleteResolver,
	}
}

func (f *Field) UserProfileFamilyField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileFamilyType,
		Description: "Returns a data of User Profile for displaying inside Family tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileFamilyResolver,
	}
}

func (f *Field) UserProfileFamilyInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileFamilyInsertType,
		Description: "Creates new or alter existing User Profile's Family item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileFamilyInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileFamilyInsertResolver,
	}
}

func (f *Field) UserProfileFamilyDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileFamilyDeleteType,
		Description: "Deletes existing User Profile's Family",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileFamilyDeleteResolver,
	}
}
