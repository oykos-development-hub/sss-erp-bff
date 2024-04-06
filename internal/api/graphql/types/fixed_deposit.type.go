package types

import "github.com/graphql-go/graphql"

var FixedDepositOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FixedDepositOverviewType",
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
			Type: graphql.NewList(FixedDepositType),
		},
	},
})

var FixedDepositInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FixedDepositInsertType",
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
			Type: FixedDepositType,
		},
	},
})

var FixedDepositItemInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FixedDepositItemInsertType",
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

var FixedDepositType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FixedDepositType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"subject": &graphql.Field{
			Type: graphql.String,
		},
		"judge": &graphql.Field{
			Type: DropdownItemType,
		},
		"case_number": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_recipiet": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_case": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_finality": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_enforceability": &graphql.Field{
			Type: graphql.String,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(FixedDepositItemType),
		},
		"dispatches": &graphql.Field{
			Type: graphql.NewList(FixedDepositDispatchType),
		},
		"judges": &graphql.Field{
			Type: graphql.NewList(FixedDepositJudgeType),
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var FixedDepositItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FixedDepositItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"deposit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"category": &graphql.Field{
			Type: DropdownItemType,
		},
		"judge": &graphql.Field{
			Type: DropdownItemType,
		},
		"type": &graphql.Field{
			Type: DropdownItemType,
		},
		"unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"currency": &graphql.Field{
			Type: DropdownItemType,
		},
		"case_number": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"date_of_confiscation": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_case": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var FixedDepositDispatchType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FixedDepositDispatchType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"deposit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"category": &graphql.Field{
			Type: DropdownItemType,
		},
		"judge": &graphql.Field{
			Type: DropdownItemType,
		},
		"type": &graphql.Field{
			Type: DropdownItemType,
		},
		"unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"currency": &graphql.Field{
			Type: DropdownItemType,
		},
		"case_number": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"date_of_action": &graphql.Field{
			Type: graphql.String,
		},
		"subject": &graphql.Field{
			Type: graphql.String,
		},
		"action": &graphql.Field{
			Type: DropdownItemType,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var FixedDepositJudgeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FixedDepositJudgeType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"deposit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"will_id": &graphql.Field{
			Type: graphql.Int,
		},
		"judge": &graphql.Field{
			Type: DropdownItemType,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
