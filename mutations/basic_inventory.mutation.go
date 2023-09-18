package mutations

import "github.com/graphql-go/graphql"

var BasicInventoryInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BasicInventoryInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"article_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"class_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"depreciation_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"real_estate": &graphql.InputObjectFieldConfig{
			Type: RealEstateInsertMutation,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"inventory_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"internal_ownership": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"office_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"location": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"target_user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"unit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"net_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"gross_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_purchase": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"source": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"source_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"donor_title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"invoice_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"price_of_assessment": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_assessment": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lifetime_of_assessment_in_months": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"deactivation_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"updated_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"invoice_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var RealEstateInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RealEstateInsert",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"square_area": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"land_serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"estate_serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"ownership_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"ownership_scope": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"ownership_investment_scope": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"limitations_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"property_document": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"limitation_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"document": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
