package types

import "github.com/graphql-go/graphql"

var ObligationsForAccountingOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ObligationsForAccountingOverviewType",
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
			Type: graphql.NewList(ObligationsForAccountingType),
		},
	},
})

var ObligationsForAccountingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ObligationsForAccountingType",
	Fields: graphql.Fields{
		"invoice_id": &graphql.Field{
			Type: graphql.Int,
		},
		"salary_id": &graphql.Field{
			Type: graphql.Int,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var AccountingOrderForObligationsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingOrderForObligationsOverviewType",
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
			Type: AccountingOrderForObligationsType,
		},
	},
})

var AccountingOrderForObligationsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingOrderForObligationsType",
	Fields: graphql.Fields{
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"date_of_booking": &graphql.Field{
			Type: graphql.String,
		},
		"credit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"debit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(AccountingOrderItemForObligationsType),
		},
	},
})

var AccountingOrderItemForObligationsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingOrderItemForObligationsType",
	Fields: graphql.Fields{
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"credit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"debit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"invoice": &graphql.Field{
			Type: DropdownItemType,
		},
		"salary": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var AccountingEntryOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingEntryOverviewType",
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
			Type: AccountingEntryType,
		},
	},
})

var AccountingEntryInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingEntryInsertType",
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
			Type: AccountingEntryType,
		},
	},
})

var AccountingEntryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingEntryType",
	Fields: graphql.Fields{
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"credit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"debit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"invoice": &graphql.Field{
			Type: DropdownItemType,
		},
		"salary": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var AccountingEntryItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingEntryItemType",
	Fields: graphql.Fields{
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"credit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"debit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"invoice": &graphql.Field{
			Type: DropdownItemType,
		},
		"salary": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})
