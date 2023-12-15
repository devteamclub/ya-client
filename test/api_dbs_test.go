/*
Партнерский API Маркета

Testing DbsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ya-client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/devteamclub/ya-client"
)

func Test_ya-client_DbsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DbsAPIService AcceptOrderCancellation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.AcceptOrderCancellation(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService AddHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.AddHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService AddOffersToArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.AddOffersToArchive(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService ConfirmBusinessPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.ConfirmBusinessPrices(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService ConfirmCampaignPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.ConfirmCampaignPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService CreateOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.CreateOutlet(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService DeleteCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.DeleteCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService DeleteHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.DeleteHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService DeleteOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.DeleteOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService DeleteOffersFromArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.DeleteOffersFromArchive(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService DeleteOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var outletId int64

		resp, httpRes, err := apiClient.DbsAPI.DeleteOutlet(context.Background(), campaignId, outletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService DeleteOutletLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.DeleteOutletLicenses(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GenerateGoodsRealizationReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GenerateGoodsRealizationReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GenerateOrderLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var shipmentId int64
		var boxId int64

		resp, httpRes, err := apiClient.DbsAPI.GenerateOrderLabel(context.Background(), campaignId, orderId, shipmentId, boxId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GenerateOrderLabels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.GenerateOrderLabels(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GeneratePricesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GeneratePricesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GenerateShowsSalesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GenerateShowsSalesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GenerateStocksOnWarehousesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GenerateStocksOnWarehousesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GenerateUnitedMarketplaceServicesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GenerateUnitedMarketplaceServicesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GenerateUnitedNettingReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GenerateUnitedNettingReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetAllOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetAllOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetBidsInfoForBusiness", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetBidsInfoForBusiness(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetBidsRecommendations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetBidsRecommendations(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetBusinessQuarantineOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetBusinessQuarantineOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaign", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCampaign(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaignFeedCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCampaignFeedCategories(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaignLogins", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCampaignLogins(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaignQuarantineOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCampaignQuarantineOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaignRegion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCampaignRegion(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaignSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCampaignSettings(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaigns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GetCampaigns(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCampaignsByLogin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var login string

		resp, httpRes, err := apiClient.DbsAPI.GetCampaignsByLogin(context.Background(), login).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetCategoryContentParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.DbsAPI.GetCategoryContentParameters(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetDeliveryServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GetDeliveryServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.DbsAPI.GetFeed(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetFeedCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.DbsAPI.GetFeedCategories(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetFeedIndexLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.DbsAPI.GetFeedIndexLogs(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetFeedbackAndCommentUpdates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetFeedbackAndCommentUpdates(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetFeeds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetFeeds(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetGoodsStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetGoodsStats(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetModel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var modelId int64

		resp, httpRes, err := apiClient.DbsAPI.GetModel(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetModelOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var modelId int64

		resp, httpRes, err := apiClient.DbsAPI.GetModelOffers(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetModels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GetModels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetModelsOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.GetModelsOffers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOfferCardsContentStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOfferCardsContentStatus(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOfferRecommendations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOfferRecommendations(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOrder(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOrderBuyerInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOrderBuyerInfo(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOrders(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOrdersStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOrdersStats(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var outletId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOutlet(context.Background(), campaignId, outletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOutletLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOutletLicenses(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetOutlets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetOutlets(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetPricesByOfferIds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetPricesByOfferIds(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetReportInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reportId string

		resp, httpRes, err := apiClient.DbsAPI.GetReportInfo(context.Background(), reportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetReturn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64

		resp, httpRes, err := apiClient.DbsAPI.GetReturn(context.Background(), campaignId, orderId, returnId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetReturnApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64

		resp, httpRes, err := apiClient.DbsAPI.GetReturnApplication(context.Background(), campaignId, orderId, returnId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetReturnPhoto", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64
		var itemId int64
		var imageHash string

		resp, httpRes, err := apiClient.DbsAPI.GetReturnPhoto(context.Background(), campaignId, orderId, returnId, itemId, imageHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetReturns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.GetReturns(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetSuggestedOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetSuggestedOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService GetWarehouses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.GetWarehouses(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService ProvideOrderDigitalCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.ProvideOrderDigitalCodes(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService ProvideOrderItemIdentifiers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.ProvideOrderItemIdentifiers(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService PutBidsForBusiness", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.PutBidsForBusiness(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService PutBidsForCampaign", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.PutBidsForCampaign(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService RefreshFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.DbsAPI.RefreshFeed(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SearchModels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.SearchModels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SearchRegionChildren", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var regionId int64

		resp, httpRes, err := apiClient.DbsAPI.SearchRegionChildren(context.Background(), regionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SearchRegionsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var regionId int64

		resp, httpRes, err := apiClient.DbsAPI.SearchRegionsById(context.Background(), regionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SearchRegionsByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DbsAPI.SearchRegionsByName(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SetFeedParams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.DbsAPI.SetFeedParams(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SetOrderDeliveryDate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.SetOrderDeliveryDate(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SetOrderDeliveryTrackCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.SetOrderDeliveryTrackCode(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SetOrderShipmentBoxes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.DbsAPI.SetOrderShipmentBoxes(context.Background(), campaignId, orderId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SetReturnDecision", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64

		resp, httpRes, err := apiClient.DbsAPI.SetReturnDecision(context.Background(), campaignId, orderId, returnId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService SubmitReturnDecision", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64

		resp, httpRes, err := apiClient.DbsAPI.SubmitReturnDecision(context.Background(), campaignId, orderId, returnId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateBusinessPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateBusinessPrices(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOfferContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateOfferContent(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOrderItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		httpRes, err := apiClient.DbsAPI.UpdateOrderItems(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOrderStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateOrderStatus(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOrderStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateOrderStatuses(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOrderStorageLimit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateOrderStorageLimit(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var outletId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateOutlet(context.Background(), campaignId, outletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateOutletLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateOutletLicenses(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdatePrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdatePrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DbsAPIService UpdateStocks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.DbsAPI.UpdateStocks(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
