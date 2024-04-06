package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) FixedDepositInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FixedDepositInsertType,
		Description: "Creates new or alter existing fixed deposit",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FixedDepositMutation),
			},
		},
		Resolve: f.Resolvers.FixedDepositInsertResolver,
	}
}

func (f *Field) FixedDepositDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Delete fixed deposit",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FixedDepositDeleteResolver,
	}
}

func (f *Field) FixedDepositOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FixedDepositOverviewType,
		Description: "Returns a data of fixed deposits",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"judge_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FixedDepositOverviewResolver,
	}
}

func (f *Field) FixedDepositItemInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FixedDepositItemInsertType,
		Description: "Creates new or alter existing fixed deposit item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FixedDepositItemMutation),
			},
		},
		Resolve: f.Resolvers.FixedDepositItemInsertResolver,
	}
}

func (f *Field) FixedDepositItemDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Delete fixed deposit item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FixedDepositItemDeleteResolver,
	}
}

func (f *Field) FixedDepositDispatchInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FixedDepositItemInsertType,
		Description: "Creates new or alter existing fixed deposit dispatch",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FixedDepositDispatchMutation),
			},
		},
		Resolve: f.Resolvers.FixedDepositDispatchInsertResolver,
	}
}

func (f *Field) FixedDepositDispatchDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FixedDepositItemInsertType,
		Description: "Delete fixed deposit dispatch",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FixedDepositDispatchDeleteResolver,
	}
}

func (f *Field) FixedDepositJudgeInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineInsertType,
		Description: "Creates new or alter existing fixed deposit judge",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FixedDepositJudgeMutation),
			},
		},
		Resolve: f.Resolvers.FixedDepositJudgeInsertResolver,
	}
}

func (f *Field) FixedDepositJudgeDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FixedDepositItemInsertType,
		Description: "Delete fixed deposit judge",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FixedDepositJudgeDeleteResolver,
	}
}
