package types

import "github.com/graphql-go/graphql"

var PropBenConfOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfOverview",
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
			Type: graphql.NewList(PropBenConfType),
		},
	},
})

var PropBenConfInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfInsert",
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
			Type: PropBenConfType,
		},
	},
})

var PropBenConfType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"property_benefits_confiscation_type": &graphql.Field{
			Type: graphql.String,
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
			Type: graphql.String,
		},
		"court_costs": &graphql.Field{
			Type: graphql.Float,
		},
		"court_account": &graphql.Field{
			Type: DropdownItemType,
		},
		"property_benefits_confiscation_details": &graphql.Field{
			Type: PropBenConfDetailsType,
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

var PropBenConfDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfDetails",
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

var PropBenConfDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfDelete",
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
