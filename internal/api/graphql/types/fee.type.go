package types

import "github.com/graphql-go/graphql"

var FeeOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeeOverview",
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
			Type: graphql.NewList(FeeType),
		},
	},
})

var FeeInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeeInsert",
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
			Type: FeeType,
		},
	},
})

var FeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeeType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"fee_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"fee_subcategory": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"decision_number": &graphql.Field{
			Type: graphql.String,
		},
		"decision_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"subject": &graphql.Field{
			Type: graphql.String,
		},
		"jmbg": &graphql.Field{
			Type: graphql.String,
		},
		"residence": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"payment_reference_number": &graphql.Field{
			Type: graphql.String,
		},
		"debit_reference_number": &graphql.Field{
			Type: graphql.String,
		},
		"execution_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"payment_deadline_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: DropdownItemType,
		},
		"court_account": &graphql.Field{
			Type: DropdownItemType,
		},
		"fee_details": &graphql.Field{
			Type: FeeDetailsType,
		},
		"file": &graphql.Field{
			Type: graphql.NewList(FileDropdownItemType),
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var FeeDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeeDetails",
	Fields: graphql.Fields{
		"fee_all_payments_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"fee_amount_grace_period": &graphql.Field{
			Type: graphql.Float,
		},
		"fee_amount_grace_period_due_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"fee_amount_grace_period_available": &graphql.Field{
			Type: graphql.Boolean,
		},
		"fee_left_to_pay_amount": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var FeeDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeeDelete",
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
