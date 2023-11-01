package resolvers

import (
	"bff/shared"
	"fmt"
)

var STATUSES = map[string]string{
	"not_accessible":       "Nedostupan",
	"admin_in_progress":    "U toku",
	"admin_published":      "Poslat",
	"user_published":       "Obradi",
	"user_requested":       "Na čekanju",
	"user_accepted":        "Odobren",
	"user_rejected":        "Odbijen",
	"pre_budget_closed":    "Zaključen",
	"pre_budget_converted": "Konvertovan",
	"post_budget_closed":   "Objavljen",
}

func PopulateStatus(plan map[string]interface{}, isAdmin bool, organizationUnitId int) string {
	var isPublished = plan["date_of_publishing"] != nil && plan["date_of_publishing"] != ""
	var isClosed = plan["date_of_closing"] != nil && plan["date_of_closing"] != ""
	var isPreBudget = plan["is_pre_budget"]
	var isConverted = false
	var isSentOnRevision = false
	var isRejected = false
	var isAccepted = false

	var conversionTargetPlans = shared.FetchByProperty(
		"public_procurement_plans",
		"PreBudgetId",
		plan["id"],
	)

	if len(conversionTargetPlans) > 0 {
		isConverted = true
	}

	if organizationUnitId > 0 {
		var organizationUnitArticles = GetOrganizationUnitArticles(plan["id"].(int), organizationUnitId)

		if len(organizationUnitArticles) > 0 {
			for _, procurementData := range organizationUnitArticles {
				procurement := procurementData.(map[string]interface{})
				if articles, ok := procurement["articles"].([]interface{}); ok {
					var procurementArticles = articles
					if len(procurementArticles) > 0 {
						for _, procurementArticle := range procurementArticles {
							article := procurementArticle.(map[string]interface{})

							if shared.IsInteger(article["amount"]) && article["amount"].(int) > 0 {
								isSentOnRevision = true
								isRejected = false
								isAccepted = false

								if article["is_rejected"] == true || article["status"] == "rejected" {
									isRejected = true
								} else if article["status"] == "accepted" {
									isAccepted = true
								}
							}
						}
					}
				}
			}
		}
	}

	if isAdmin {
		if isPublished {
			if isClosed {
				if isPreBudget == true {
					if isConverted {
						// Admin converted a closed pre-budget Plan into a new post-budget Plan
						return STATUSES["pre_budget_converted"]
					}
					// Admin closed a pre-budget Plan that can't be edited any further
					return STATUSES["pre_budget_closed"]
				}
				// Admin closed a post-budget Plan that can't be edited any further
				return STATUSES["post_budget_closed"]
			}
			// Admin published a Plan that can be seen by Users in Organization units
			return STATUSES["admin_published"]
		}
		// Draft version of the Plan before it has been published
		return STATUSES["admin_in_progress"]
	} else {
		if isPublished {
			if isClosed {
				if isPreBudget == true {
					if isConverted {
						// Admin converted a closed pre-budget Plan into a new post-budget Plan
						return STATUSES["pre_budget_converted"]
					}
					// Admin closed a pre-budget Plan that can't be edited any further
					return STATUSES["pre_budget_closed"]
				}
				// Admin closed a post-budget Plan that can't be edited any further
				return STATUSES["post_budget_closed"]
			}
			if isSentOnRevision {
				if isRejected {
					return STATUSES["user_rejected"]
				}
				if isAccepted {
					return STATUSES["user_accepted"]
				}

				return STATUSES["user_requested"]
			}
			// Users in Organization units can see Plan and request Articles after it has been published
			return STATUSES["user_published"]
		} else {
			// Not accessible for Users in Organization units before Plan has been published
			return STATUSES["not_accessible"]
		}
	}
}

func GetOrganizationUnitArticles(planId int, unitId int) []interface{} {
	var items []interface{}

	if !shared.IsInteger(planId) || !shared.IsInteger(unitId) {
		return items
	}

	var organizationUnits = shared.FetchByProperty(
		"organization_units",
		"Id",
		unitId,
	)
	var plans = shared.FetchByProperty(
		"public_procurement_plan",
		"Id",
		planId,
	)
	var procurements = shared.FetchByProperty(
		"public_procurement_item",
		"PlanId",
		planId,
	)

	if len(procurements) > 0 && len(plans) > 0 && len(organizationUnits) > 0 {
		var plan = map[string]interface{}{}
		var organizationUnit = map[string]interface{}{}
		var procurementArticles []interface{}

		for _, planData := range plans {
			plan = shared.WriteStructToInterface(planData)
		}

		for _, unitData := range organizationUnits {
			organizationUnit = shared.WriteStructToInterface(unitData)
		}

		for _, procurementData := range procurements {
			var procurement = shared.WriteStructToInterface(procurementData)
			var relatedIndent = shared.FetchByProperty(
				"budget_indent",
				"Id",
				procurement["budget_indent_id"],
			)

			if len(relatedIndent) > 0 {
				for _, indentData := range relatedIndent {
					var indent = shared.WriteStructToInterface(indentData)

					procurement["budget_indent"] = map[string]interface{}{
						"id":    indent["id"],
						"title": indent["title"],
					}
				}
			}

			var relatedArticles = shared.FetchByProperty(
				"public_procurement_article",
				"PublicProcurementId",
				procurement["id"],
			)

			if len(relatedArticles) > 0 {
				for _, articleData := range relatedArticles {
					var article = shared.WriteStructToInterface(articleData)
					var relatedOrganizationUnitArticles = shared.FetchByProperty(
						"public_procurement_organization_unit_article",
						"PublicProcurementArticleId",
						article["id"],
					)

					if len(relatedOrganizationUnitArticles) > 0 {
						for _, articleOrganizationUnitData := range relatedOrganizationUnitArticles {
							var articleOrganizationUnit = shared.WriteStructToInterface(articleOrganizationUnitData)

							if articleOrganizationUnit["organization_unit_id"] == unitId {
								articleOrganizationUnit["organization_unit"] = map[string]interface{}{
									"id":    organizationUnit["id"],
									"title": organizationUnit["title"],
								}
								articleOrganizationUnit["public_procurement_article"] = map[string]interface{}{
									"id":             article["id"],
									"title":          article["title"],
									"net_price":      article["net_price"],
									"vat_percentage": article["vat_percentage"],
									"description":    article["description"],
								}

								procurementArticles = append(procurementArticles, articleOrganizationUnit)
							}
						}
					}
				}

				procurement["articles"] = procurementArticles
			}

			procurement["plan"] = map[string]interface{}{
				"id":    plan["id"],
				"title": plan["title"],
			}

			if len(procurementArticles) > 0 {
				items = append(items, procurement)
			}
		}
	}

	return items
}

func PopulateContractArticleProperties(contractArticles []interface{}, filters ...interface{}) []interface{} {
	var contractId int
	var items []interface{}

	switch len(filters) {
	case 1:
		contractId = filters[0].(int)
	}

	for _, articleData := range contractArticles {
		var article = shared.WriteStructToInterface(articleData)

		if shared.IsInteger(contractId) && contractId > 0 && article["public_procurement_contract_id"] != contractId {
			continue
		}

		var relatedArticles = shared.FetchByProperty(
			"public_procurement_articles",
			"Id",
			article["public_procurement_article_id"],
		)

		if len(relatedArticles) > 0 {
			for _, relatedArticleData := range relatedArticles {
				var relatedArticle = shared.WriteStructToInterface(relatedArticleData)

				article["public_procurement_article"] = map[string]interface{}{
					"id":             relatedArticle["id"],
					"title":          relatedArticle["title"],
					"vat_percentage": relatedArticle["vat_percentage"],
					"description":    relatedArticle["description"],
				}
			}
		}

		var relatedContracts = shared.FetchByProperty(
			"public_procurement_contracts",
			"Id",
			article["public_procurement_contract_id"],
		)

		if len(relatedContracts) > 0 {
			for _, contractData := range relatedContracts {
				var contract = shared.WriteStructToInterface(contractData)

				article["contract"] = map[string]interface{}{
					"id":    contract["id"],
					"title": contract["serial_number"],
				}
			}
		}

		items = append(items, article)
	}

	return items
}

func PopulateContractItemProperties(contracts []interface{}, filters ...interface{}) []interface{} {
	var id, procurementId, supplierId int
	var items []interface{}

	switch len(filters) {
	case 1:
		id = filters[0].(int)
	case 2:
		id = filters[0].(int)
		procurementId = filters[1].(int)
	case 3:
		id = filters[0].(int)
		procurementId = filters[1].(int)
		supplierId = filters[2].(int)
	}

	for _, contractData := range contracts {
		var contract = shared.WriteStructToInterface(contractData)

		if shared.IsInteger(id) && id > 0 && contract["id"] != id {
			continue
		}
		if shared.IsInteger(procurementId) && procurementId > 0 && contract["public_procurement_id"] != procurementId {
			continue
		}
		if shared.IsInteger(supplierId) && supplierId > 0 && contract["supplier_id"] != supplierId {
			continue
		}

		var relatedSuppliers = shared.FetchByProperty(
			"suppliers",
			"Id",
			contract["supplier_id"],
		)

		if len(relatedSuppliers) > 0 {
			for _, supplierData := range relatedSuppliers {
				var supplier = shared.WriteStructToInterface(supplierData)

				contract["supplier"] = map[string]interface{}{
					"id":    supplier["id"],
					"title": supplier["title"],
				}
			}
		}

		var relatedProcurement = shared.FetchByProperty(
			"public_procurement_item",
			"Id",
			contract["public_procurement_id"],
		)

		if len(relatedProcurement) > 0 {
			for _, procurementData := range relatedProcurement {
				var procurement = shared.WriteStructToInterface(procurementData)

				contract["public_procurement"] = map[string]interface{}{
					"id":    procurement["id"],
					"title": procurement["title"],
				}
			}
		}

		items = append(items, contract)
	}

	return items
}

func PopulateProcurementLimitProperties(limits []interface{}, filters ...interface{}) []interface{} {
	var id int
	var items []interface{}

	switch len(filters) {
	case 1:
		id = filters[0].(int)
	}

	for _, limitData := range limits {
		var limit = shared.WriteStructToInterface(limitData)

		if shared.IsInteger(id) && id > 0 && limit["id"] != id {
			continue
		}

		var relatedOrganizationUnit = shared.FetchByProperty(
			"organization_units",
			"Id",
			limit["organization_unit_id"],
		)

		if len(relatedOrganizationUnit) > 0 {
			for _, unitData := range relatedOrganizationUnit {
				var unit = shared.WriteStructToInterface(unitData)

				limit["organization_unit"] = map[string]interface{}{
					"id":    unit["id"],
					"title": unit["title"],
				}
			}
		}

		var relatedProcurement = shared.FetchByProperty(
			"public_procurement_item",
			"Id",
			limit["public_procurement_id"],
		)

		if len(relatedProcurement) > 0 {
			for _, procurementData := range relatedProcurement {
				var procurement = shared.WriteStructToInterface(procurementData)

				limit["public_procurement"] = map[string]interface{}{
					"id":    procurement["id"],
					"title": procurement["title"],
				}
			}
		}

		items = append(items, limit)
	}

	return items
}

func PopulateProcurementArticleProperties(articles []interface{}, filters ...interface{}) []interface{} {
	var id int
	var items []interface{}

	switch len(filters) {
	case 1:
		id = filters[0].(int)
	}

	for _, articleData := range articles {
		var article = shared.WriteStructToInterface(articleData)

		if shared.IsInteger(id) && id > 0 && article["id"] != id {
			continue
		}

		var relatedIndent = shared.FetchByProperty(
			"budget_indent",
			"Id",
			article["budget_indent_id"],
		)

		if len(relatedIndent) > 0 {
			for _, indentData := range relatedIndent {
				var indent = shared.WriteStructToInterface(indentData)

				article["budget_indent"] = map[string]interface{}{
					"id":    indent["id"],
					"title": indent["title"],
				}
			}
		}

		var relatedProcurement = shared.FetchByProperty(
			"public_procurement_item",
			"Id",
			article["public_procurement_id"],
		)

		if len(relatedProcurement) > 0 {
			for _, procurementData := range relatedProcurement {
				var procurement = shared.WriteStructToInterface(procurementData)

				article["public_procurement"] = map[string]interface{}{
					"id":    procurement["id"],
					"title": procurement["title"],
				}
			}
		}

		items = append(items, article)
	}

	return items
}

func PopulateProcurementItemProperties(procurements []interface{}, filters ...interface{}) []interface{} {
	var id int
	var planStatus string
	var items []interface{}

	switch len(filters) {
	case 1:
		id = filters[0].(int)
	case 2:
		id = filters[0].(int)
		planStatus = filters[1].(string)
	}

	for _, procurementData := range procurements {
		var procurement = shared.WriteStructToInterface(procurementData)

		if shared.IsInteger(id) && id > 0 && procurement["id"] != id {
			continue
		}

		var relatedIndent = shared.FetchByProperty(
			"budget_indent",
			"Id",
			procurement["budget_indent_id"],
		)

		if len(relatedIndent) > 0 {
			for _, indentData := range relatedIndent {
				var indent = shared.WriteStructToInterface(indentData)

				procurement["budget_indent"] = map[string]interface{}{
					"id":    indent["id"],
					"title": indent["title"],
				}
			}
		}

		var relatedPlan = shared.FetchByProperty(
			"public_procurement_plan",
			"Id",
			procurement["plan_id"],
		)

		if len(relatedPlan) > 0 {
			for _, planData := range relatedPlan {
				var plan = shared.WriteStructToInterface(planData)

				procurement["plan"] = map[string]interface{}{
					"id":    plan["id"],
					"title": plan["title"],
				}
			}
		}

		procurement["articles"] = PopulateProcurementArticleProperties(shared.FetchByProperty(
			"public_procurement_articles",
			"PublicProcurementId",
			procurement["id"],
		))

		if planStatus == STATUSES["admin_published"] ||
			planStatus == STATUSES["user_published"] ||
			planStatus == STATUSES["user_requested"] ||
			planStatus == STATUSES["user_accepted"] ||
			planStatus == STATUSES["user_rejected"] {
			procurement["status"] = planStatus
		} else {
			procurement["status"] = STATUSES["admin_in_progress"]
		}

		items = append(items, procurement)
	}

	return items
}

func PopulatePlanItemProperties(plans []interface{}, filters ...interface{}) []interface{} {
	var isPreBudget interface{}
	var authToken interface{}
	var year, status string
	var items []interface{}

	switch len(filters) {
	case 1:
		isPreBudget = filters[0]
	case 2:
		isPreBudget = filters[0]
		year = filters[1].(string)
	case 3:
		isPreBudget = filters[0]
		year = filters[1].(string)
		status = filters[2].(string)
	case 4:
		isPreBudget = filters[0]
		year = filters[1].(string)
		status = filters[2].(string)
		authToken = filters[3]
		if authToken == nil {
			authToken = ""
		}
	}

	//@TODO check if token belongs to Admin user
	var isAdmin = authToken == "sss"
	var organizationUnitId = 0

	//@TODO fetch OrganizationUnitId from the token if User is not Admin and belongs to Organization Unit
	if !isAdmin {
		organizationUnitId = 2
	}

	for _, planData := range plans {
		var plan = shared.WriteStructToInterface(planData)

		plan["status"] = PopulateStatus(plan, isAdmin, organizationUnitId)

		if plan["status"] == STATUSES["not_accessible"] {
			fmt.Printf("\n not accesible plan %s\n", plan)
			continue
		}
		if shared.IsString(status) && len(status) > 0 {
			continue
		}
		if isPreBudget != nil && plan["is_pre_budget"] != isPreBudget {
			continue
		}
		if shared.IsString(year) && len(year) > 0 && plan["year"] != year {
			continue
		}
		if shared.IsInteger(plan["pre_budget_id"]) && plan["pre_budget_id"].(int) > 0 {
			var relatedPreBudgetPlan = shared.FetchByProperty(
				"public_procurement_plan",
				"Id",
				plan["pre_budget_id"],
			)

			if len(relatedPreBudgetPlan) > 0 {
				for _, preBudgetPlanData := range relatedPreBudgetPlan {
					var preBudgetPlan = shared.WriteStructToInterface(preBudgetPlanData)

					plan["pre_budget_plan"] = map[string]interface{}{
						"id":    preBudgetPlan["id"],
						"title": preBudgetPlan["title"],
					}
				}
			}
		}

		plan["items"] = PopulateProcurementItemProperties(
			shared.FetchByProperty(
				"public_procurement_items",
				"PlanId",
				plan["id"],
			),
			0,
			plan["status"],
		)

		items = append(items, plan)
	}

	return items
}
