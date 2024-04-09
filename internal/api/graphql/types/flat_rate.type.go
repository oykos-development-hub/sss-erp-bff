package types

import "github.com/graphql-go/graphql"

var FlatRateOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRateOverview",
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
			Type: graphql.NewList(FlatRateType),
		},
	},
})

var FlatRateInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRateInsert",
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
			Type: FlatRateType,
		},
	},
})

var FlatRateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRateType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"flat_rate_type": &graphql.Field{
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
		"flat_rate_details": &graphql.Field{
			Type: FlatRateDetailsType,
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

var FlatRateDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRateDetails",
	Fields: graphql.Fields{
		"all_payments_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_grace_period": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_grace_period_due_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"amount_grace_period_available": &graphql.Field{
			Type: graphql.Boolean,
		},
		"left_to_pay_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"court_costs_paid": &graphql.Field{
			Type: graphql.Float,
		},
		"court_costs_left_to_pay_amount": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var FlatRateDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRateDelete",
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
