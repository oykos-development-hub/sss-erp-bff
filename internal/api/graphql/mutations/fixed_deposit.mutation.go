package mutations

import "github.com/graphql-go/graphql"

var FixedDepositMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FixedDepositMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"judge_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"subject": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_recipiet": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_case": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_finality": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_enforceability": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var FixedDepositItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FixedDepositItemMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"deposit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"category_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"unit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"currency": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_confiscation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"judge_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var FixedDepositDispatchMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FixedDepositDispatchMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"deposit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"category_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"unit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"currency": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_action": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"subject": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"action": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"judge_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var FixedDepositJudgeMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FixedDepositJudgeMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"deposit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"will_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"judge_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var FixedDepositWillMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FixedDepositWillMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"subject": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number_si": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number_rs": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_receipt_si": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_receipt_rs": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"father_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"jmbg": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var FixedDepositWillDispatchMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FixedDepositWillDispatchMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"will_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"dispatch_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"judge_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_dispatch": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
