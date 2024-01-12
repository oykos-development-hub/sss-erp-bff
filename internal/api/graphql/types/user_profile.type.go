package types

import "github.com/graphql-go/graphql"

var UserProfilesOverviewItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfilesOverviewItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_judge": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_judge_president": &graphql.Field{
			Type: graphql.Boolean,
		},
		"role": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"job_position": &graphql.Field{
			Type: DropdownItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfileBasicItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileBasicItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"middle_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"birth_last_name": &graphql.Field{
			Type: graphql.String,
		},
		"country_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"city_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"nationality": &graphql.Field{
			Type: graphql.String,
		},
		"national_minority": &graphql.Field{
			Type: graphql.String,
		},
		"citizenship": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"father_name": &graphql.Field{
			Type: graphql.String,
		},
		"mother_name": &graphql.Field{
			Type: graphql.String,
		},
		"mother_birth_last_name": &graphql.Field{
			Type: graphql.String,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"bank_name": &graphql.Field{
			Type: graphql.String,
		},
		"personal_id": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_id": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_document_number": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_document_issuer": &graphql.Field{
			Type: graphql.String,
		},
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"single_parent": &graphql.Field{
			Type: graphql.Boolean,
		},
		"housing_done": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_president": &graphql.Field{
			Type: graphql.Boolean,
		},
		"housing_description": &graphql.Field{
			Type: graphql.String,
		},
		"marital_status": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_taking_oath": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_becoming_judge": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"job_position": &graphql.Field{
			Type: DropdownItemType,
		},
		"contract": &graphql.Field{
			Type: ContractItemType,
		},
		"position_in_organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"secondary_mail": &graphql.Field{
			Type: graphql.String,
		},
		"pin": &graphql.Field{
			Type: graphql.String,
		},
		"is_judge": &graphql.Field{
			Type: graphql.Boolean,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfileItemUpdateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileUpdateItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"middle_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"birth_last_name": &graphql.Field{
			Type: graphql.String,
		},
		"country_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"city_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"nationality": &graphql.Field{
			Type: graphql.String,
		},
		"citizenship": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"father_name": &graphql.Field{
			Type: graphql.String,
		},
		"mother_name": &graphql.Field{
			Type: graphql.String,
		},
		"mother_birth_last_name": &graphql.Field{
			Type: graphql.String,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"bank_name": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_id": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_document_number": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_document_issuer": &graphql.Field{
			Type: graphql.String,
		},
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"single_parent": &graphql.Field{
			Type: graphql.Boolean,
		},
		"housing_done": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_president": &graphql.Field{
			Type: graphql.Boolean,
		},
		"housing_description": &graphql.Field{
			Type: graphql.String,
		},
		"marital_status": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_taking_oath": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_becoming_judge": &graphql.Field{
			Type: graphql.String,
		},
		"position_in_organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"is_judge": &graphql.Field{
			Type: graphql.Boolean,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfilesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfilesOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(UserProfilesOverviewItemType),
		},
	},
})

var UserProfileContractsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileContracts",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(ContractItemType),
		},
	},
})

var UserProfileBasicType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileBasic",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: UserProfileBasicItemType,
		},
	},
})

var UserProfileBasicInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileBasicInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: UserProfileBasicItemType,
		},
	},
})

var UserProfileContractInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileContractsInsertType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: ContractItemType,
		},
	},
})

var UserProfileContractDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileContractDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfileUpdateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileBasicUpdate",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: UserProfileItemUpdateType,
		},
	},
})

var UserProfileEducationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileEducation",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(UserProfileEducationItemType),
		},
	},
})

var UserProfileEducationItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileEducationItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: DropdownItemType,
		},
		"date_of_certification": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"academic_title": &graphql.Field{
			Type: graphql.String,
		},
		"expertise_level": &graphql.Field{
			Type: graphql.String,
		},
		"score": &graphql.Field{
			Type: graphql.String,
		},
		"certificate_issuer": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
	},
})

var UserProfileEducationInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileEducationInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: UserProfileEducationItemType,
		},
	},
})

var UserProfileEducationDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EducationDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfileExperienceItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileExperienceItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"relevant": &graphql.Field{
			Type: graphql.Boolean,
		},
		"organization_unit": &graphql.Field{
			Type: graphql.String,
		},
		"amount_of_experience": &graphql.Field{
			Type: graphql.Int,
		},
		"amount_of_insured_experience": &graphql.Field{
			Type: graphql.Int,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"reference_file_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var UserProfileExperienceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileExperience",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(UserProfileExperienceItemType),
		},
	},
})

var UserProfileExperienceInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileExperienceInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: UserProfileExperienceItemType,
		},
	},
})

var UserProfileExperienceDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExperienceDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfileFamilyItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileFamilyItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"birth_last_name": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"country_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"city_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"nationality": &graphql.Field{
			Type: graphql.String,
		},
		"national_minority": &graphql.Field{
			Type: graphql.String,
		},
		"citizenship": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"father_name": &graphql.Field{
			Type: graphql.String,
		},
		"mother_name": &graphql.Field{
			Type: graphql.String,
		},
		"mother_birth_last_name": &graphql.Field{
			Type: graphql.String,
		},
		"personal_id": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_id": &graphql.Field{
			Type: graphql.String,
		},
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"handicapped_person": &graphql.Field{
			Type: graphql.Boolean,
		},
		"insurance_coverage": &graphql.Field{
			Type: graphql.String,
		},
		"employee_relationship": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfileFamilyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileFamily",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(UserProfileFamilyItemType),
		},
	},
})

var UserProfileFamilyInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileFamilyInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: UserProfileFamilyItemType,
		},
	},
})

var UserProfileFamilyDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FamilyDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
