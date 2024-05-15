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
		"id": &graphql.Field{
			Type: graphql.Int,
		},
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
		"date": &graphql.Field{
			Type: graphql.String,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var PaymentOrdersForAccountingOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentOrdersForAccountingOverviewType",
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
			Type: graphql.NewList(PaymentOrdersForAccountingType),
		},
	},
})

var PaymentOrdersForAccountingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentOrdersForAccountingType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"payment_order_id": &graphql.Field{
			Type: graphql.Int,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"date": &graphql.Field{
			Type: graphql.String,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
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
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"date": &graphql.Field{
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
		"supplier_id": &graphql.Field{
			Type: graphql.Int,
		},
		"invoice": &graphql.Field{
			Type: DropdownItemType,
		},
		"salary": &graphql.Field{
			Type: DropdownItemType,
		},
		"payment_order": &graphql.Field{
			Type: DropdownItemType,
		},
		"enforced_payment": &graphql.Field{
			Type: DropdownItemType,
		},
		"return_enforced_payment": &graphql.Field{
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
			Type: graphql.NewList(AccountingEntryType),
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
		"item": &graphql.Field{
			Type: AccountingEntryType,
		},
	},
})

var AccountingEntryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingEntryType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"id_of_entry": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"credit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"debit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"date_of_booking": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(AccountingEntryItemType),
		},
	},
})

var AccountingEntryItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountingEntryItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"date": &graphql.Field{
			Type: graphql.String,
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
		"payment_order": &graphql.Field{
			Type: DropdownItemType,
		},
		"enforced_payment": &graphql.Field{
			Type: DropdownItemType,
		},
		"return_enforced_payment": &graphql.Field{
			Type: DropdownItemType,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var AnalyticalCardOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AnalyticalCardOverviewType",
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
			Type: graphql.NewList(AnalyticalCardType),
		},
	},
})

var AnalyticalCardType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AnalyticalCardType",
	Fields: graphql.Fields{
		"initial_state": &graphql.Field{
			Type: graphql.Float,
		},
		"sum_debit_amount_in_period": &graphql.Field{
			Type: graphql.Float,
		},
		"sum_credit_amount_in_period": &graphql.Field{
			Type: graphql.Float,
		},
		"sum_debit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"sum_credit_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(AnalyticalCardItemType),
		},
	},
})

var AnalyticalCardItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AnalyticalCardItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"date": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
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
		"balance": &graphql.Field{
			Type: graphql.Float,
		},
		"document_title": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})
