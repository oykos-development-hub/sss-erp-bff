package types

import (
	"bff/internal/api/dto"
	"sync"

	"github.com/graphql-go/graphql"
)

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
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: DropdownItemType,
		},
		"limits": &graphql.Field{
			Type: graphql.NewList(BudgetLimitType),
		},
		"number_of_requests": &graphql.Field{
			Type: graphql.Int,
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

var (
	accountWithFilledDataType          *graphql.Object
	onceInitalizeAccountWithFilledData sync.Once
)

func GetAccountWithFilledDataType() *graphql.Object {
	onceInitalizeAccountWithFilledData.Do(func() {
		initAccountWithFilledDataType()
	})
	return accountWithFilledDataType
}

var FilledAccountData = graphql.NewObject(graphql.ObjectConfig{
	Name: "FilledAccountDataType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"account_id": &graphql.Field{
			Type: graphql.Int,
		},
		"current_year": &graphql.Field{
			Type: decimalType,
		},
		"next_year": &graphql.Field{
			Type: decimalType,
		},
		"year_after_next": &graphql.Field{
			Type: decimalType,
		},
		"actual": &graphql.Field{
			Type: decimalType,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func initAccountWithFilledDataType() {
	accountWithFilledDataType = graphql.NewObject(graphql.ObjectConfig{
		Name: "AccountWithFilledDataType",
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"parent_id": &graphql.Field{
					Type: graphql.Int,
				},
				"serial_number": &graphql.Field{
					Type: graphql.String,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"filled_data": &graphql.Field{
					Type: FilledAccountData,
				},
				"children": &graphql.Field{
					Type: graphql.NewList(accountWithFilledDataType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if accountItem, ok := p.Source.(*dto.AccountWithFilledFinanceBudget); ok {
							return accountItem.Children, nil
						}
						return nil, nil
					},
				},
			}
		}),
	})
}

var AccountWithFilledDataOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountWithFilledDataOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"version": &graphql.Field{
			Type: graphql.Int,
		},
		"item": &graphql.Field{
			Type: FinancialBudgetType,
		},
	},
})

var FinancialBudgetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetType",
	Fields: graphql.Fields{
		"account_version": &graphql.Field{
			Type: graphql.Int,
		},
		"status": &graphql.Field{
			Type: DropdownItemType,
		},
		"request_id": &graphql.Field{
			Type: graphql.Int,
		},
		"current_request_id": &graphql.Field{
			Type: graphql.Int,
		},
		"current_status": &graphql.Field{
			Type: DropdownItemType,
		},
		"current_budget_comment": &graphql.Field{
			Type: graphql.String,
		},
		"current_accounts": &graphql.Field{
			Type: graphql.NewList(GetAccountWithFilledDataType()),
		},
		"donation_request_id": &graphql.Field{
			Type: graphql.Int,
		},
		"donation_status": &graphql.Field{
			Type: DropdownItemType,
		},
		"donation_budget_comment": &graphql.Field{
			Type: graphql.String,
		},
		"official_comment": &graphql.Field{
			Type: graphql.String,
		},
		"donation_accounts": &graphql.Field{
			Type: graphql.NewList(GetAccountWithFilledDataType()),
		},
	},
})

var FinancialBudgetDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetDetailsType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: FinancialBudgetDetailsItemType,
		},
	},
})

var FinancialBudgetDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetDetailsItemType",
	Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: graphql.Int,
		},
		"latest_version": &graphql.Field{
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
		"item": &graphql.Field{
			Type: BudgetType,
		},
	},
})

var BudgetSendOnReviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetSendOnReview",
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

var BudgetRequestAcceptType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetRequestAccept",
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

var BudgetRequestRejectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetRequestReject",
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

var FinancialBudgetFillType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetFillType",
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
			Type: graphql.NewList(FinancialBudgetFillItemType),
		},
	},
})

var FinancialBudgetFillItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetFillItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"request_id": &graphql.Field{
			Type: graphql.Int,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"current_year": &graphql.Field{
			Type: decimalType,
		},
		"next_year": &graphql.Field{
			Type: decimalType,
		},
		"year_after_next": &graphql.Field{
			Type: decimalType,
		},
		"actual": &graphql.Field{
			Type: decimalType,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var FinancialBudgetSendOnReviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetSendOnReviewType",
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

var OfficialBudgetRequestOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OfficialBudgetRequestOverview",
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
			Type: graphql.NewList(OfficialBudgetRequestType),
		},
	},
})

var OfficialBudgetRequestType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OfficialBudgetRequestItemType",
	Fields: graphql.Fields{
		"unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"receive_date": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var BudgetRequestsDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetRequestsDetailsType",
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
			Type: BudgetRequestsDetailsItemType,
		},
	},
})

var BudgetRequestsDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetRequestsDetailsItemType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: DropdownItemType,
		},
		"budget": &graphql.Field{
			Type: DropdownItemType,
		},
		"request_id": &graphql.Field{
			Type: graphql.Int,
		},
		"financial": &graphql.Field{
			Type: FinancialBudgetType,
		},
		"non_financial": &graphql.Field{
			Type: NonFinancialBudgetType,
		},
		"limit": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var FinancialBudgetSummaryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FinancialBudgetSummaryType",
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
			Type: FinancialBudgetType,
		},
	},
})
