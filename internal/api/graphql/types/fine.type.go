package types

import "github.com/graphql-go/graphql"

var FineOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FineOverview",
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
			Type: graphql.NewList(FineType),
		},
	},
})

var FineInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FineInsert",
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
			Type: FineType,
		},
	},
})

var FineType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FineType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"act_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"decision_number": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
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
		"account": &graphql.Field{
			Type: DropdownItemType,
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
		"court_costs": &graphql.Field{
			Type: graphql.Float,
		},
		"court_account": &graphql.Field{
			Type: DropdownItemType,
		},
		"fine_fee_details": &graphql.Field{
			Type: FineFeeDetailsType,
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

var FineFeeDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FineFeeDetails",
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
		"fee_court_costs_paid": &graphql.Field{
			Type: graphql.Float,
		},
		"fee_court_costs_left_to_pay_amount": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var FineDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FineDelete",
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
