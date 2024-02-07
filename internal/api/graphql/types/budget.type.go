package types

import "github.com/graphql-go/graphql"

var BudgetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_type": &graphql.Field{
			Type: graphql.Int,
		},
		"limits": &graphql.Field{
			Type: graphql.NewList(BudgetLimitType),
		},
	},
})

var BudgetLimitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetLimitType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"limit": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var FinancialBudgetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetType",
	Fields: graphql.Fields{
		"account_version": &graphql.Field{
			Type: graphql.Int,
		},
		"accounts": &graphql.Field{
			Type: graphql.NewList(GetAccountType()),
		},
	},
})

var BudgetOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetOverview",
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
			Type: graphql.NewList(BudgetType),
		},
	},
})

var FinancialBudgetOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetOverview",
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
		"item": &graphql.Field{
			Type: FinancialBudgetType,
		},
	},
})

var BudgetDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetDelete",
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

var BudgetSendType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetSend",
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

var BudgetInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetInsert",
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
			Type: BudgetType,
		},
	},
})
