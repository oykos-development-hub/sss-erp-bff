package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/graphql-go/graphql"
)

var PublicProcurementPlanItemDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	var items []*dto.ProcurementItemResponseItem
	var authToken = params.Context.Value(config.TokenKey).(string)
	loggedInAccount, err := getLoggedInUser(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	processItem := func(item *structs.PublicProcurementItem, ch chan<- *dto.ProcurementItemResponseItem, wg *sync.WaitGroup) {
		defer wg.Done()
		resItem, _ := buildProcurementItemResponseItem(item, *loggedInAccount)
		ch <- resItem
	}

	if id != nil && id.(int) > 0 {
		item, err := getProcurementItem(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildProcurementItemResponseItem(item, *loggedInAccount)
		items = append(items, resItem)
	} else {
		procurements, err := getProcurementItemList(nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		// Initialize a WaitGroup based on the number of items to be processed
		var wg sync.WaitGroup
		ch := make(chan *dto.ProcurementItemResponseItem, len(procurements))

		for _, item := range procurements {
			wg.Add(1)
			go processItem(item, ch, &wg)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for resItem := range ch {
			items = append(items, resItem)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   len(items),
	}, nil
}

var PublicProcurementPlanItemInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	response := dto.ResponseSingle{
		Status: "success",
	}
	var authToken = params.Context.Value(config.TokenKey).(string)
	loggedInAccount, err := getLoggedInUser(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	var data structs.PublicProcurementItem

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementItem(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildProcurementItemResponseItem(res, *loggedInAccount)

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		res, err := createProcurementItem(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(res, *loggedInAccount)

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil

}

var PublicProcurementPlanItemDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteProcurementItem(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createProcurementItem(item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ITEMS_ENDPOINT, item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementItem(id int, item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ITEMS_ENDPOINT+"/"+strconv.Itoa(id), item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteProcurementItem(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ITEMS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getProcurementItem(id int) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ITEMS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getProcurementItemList(input *dto.GetProcurementItemListInputMS) ([]*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ITEMS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func buildProcurementItemResponseItem(item *structs.PublicProcurementItem, loggedInAccount structs.UserAccounts) (*dto.ProcurementItemResponseItem, error) {
	plan, _ := getProcurementPlan(item.PlanId)
	planDropdown := dto.DropdownSimple{Id: plan.Id, Title: plan.Title}

	var articles []*dto.ProcurementArticleResponseItem
	articlesRaw, err := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &item.Id})
	if err != nil {
		return nil, err
	}
	for _, article := range articlesRaw {
		articleResItem, err := buildProcurementArticleResponseItem(article)
		if err != nil {
			return nil, err
		}
		articles = append(articles, articleResItem)
	}

	userProfile, _ := getUserProfileByUserAccountID(loggedInAccount.Id)
	organizationUnitID, _ := getOrganizationUnitIdByUserProfile(userProfile.Id)

	planStatus := "U toku"

	if organizationUnitID != nil {
		filledArticles, _ := getOrganizationUnitArticles(plan.Id, *organizationUnitID)
		if len(filledArticles) >= len(articlesRaw) {
			planStatus = "Obrađen"
		}
	}

	res := dto.ProcurementItemResponseItem{
		Id:                item.Id,
		Title:             item.Title,
		BudgetIndent:      dto.DropdownSimple{},
		Plan:              planDropdown,
		IsOpenProcurement: item.IsOpenProcurement,
		ArticleType:       item.ArticleType,
		Status:            &planStatus,
		SerialNumber:      item.SerialNumber,
		DateOfAwarding:    (*string)(item.DateOfAwarding),
		DateOfPublishing:  (*string)(item.DateOfPublishing),
		FileId:            item.FileId,
		Articles:          articles,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	return &res, nil
}
