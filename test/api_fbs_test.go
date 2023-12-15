/*
Партнерский API Маркета

Testing FbsAPIService

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

func Test_ya-client_FbsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FbsAPIService AddHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.AddHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService AddOffersToArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.AddOffersToArchive(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService ConfirmBusinessPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.ConfirmBusinessPrices(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService ConfirmCampaignPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.ConfirmCampaignPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService ConfirmShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.ConfirmShipment(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DeleteCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.DeleteCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DeleteHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.DeleteHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DeleteOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.DeleteOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DeleteOffersFromArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.DeleteOffersFromArchive(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DownloadShipmentAct", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.DownloadShipmentAct(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DownloadShipmentDiscrepancyAct", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.DownloadShipmentDiscrepancyAct(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DownloadShipmentInboundAct", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.DownloadShipmentInboundAct(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DownloadShipmentPalletLabels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.DownloadShipmentPalletLabels(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DownloadShipmentReceptionTransferAct", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.DownloadShipmentReceptionTransferAct(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService DownloadShipmentTransportationWaybill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.DownloadShipmentTransportationWaybill(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GenerateGoodsRealizationReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GenerateGoodsRealizationReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GenerateOrderLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var shipmentId int64
		var boxId int64

		resp, httpRes, err := apiClient.FbsAPI.GenerateOrderLabel(context.Background(), campaignId, orderId, shipmentId, boxId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GenerateOrderLabels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbsAPI.GenerateOrderLabels(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GeneratePricesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GeneratePricesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GenerateShowsSalesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GenerateShowsSalesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GenerateStocksOnWarehousesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GenerateStocksOnWarehousesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GenerateUnitedMarketplaceServicesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GenerateUnitedMarketplaceServicesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GenerateUnitedNettingReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GenerateUnitedNettingReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetBidsInfoForBusiness", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetBidsInfoForBusiness(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetBidsRecommendations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetBidsRecommendations(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetBusinessQuarantineOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetBusinessQuarantineOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetCampaign", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetCampaign(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetCampaignLogins", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetCampaignLogins(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetCampaignQuarantineOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetCampaignQuarantineOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetCampaigns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GetCampaigns(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetCampaignsByLogin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var login string

		resp, httpRes, err := apiClient.FbsAPI.GetCampaignsByLogin(context.Background(), login).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetCategoryContentParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.FbsAPI.GetCategoryContentParameters(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetDeliveryServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.GetDeliveryServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetGoodsStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetGoodsStats(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOfferCardsContentStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOfferCardsContentStatus(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOfferMappingEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOfferMappingEntries(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOfferRecommendations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOfferRecommendations(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOrder(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOrderLabelsData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOrderLabelsData(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOrders(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetOrdersStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetOrdersStats(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetPricesByOfferIds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetPricesByOfferIds(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetReportInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reportId string

		resp, httpRes, err := apiClient.FbsAPI.GetReportInfo(context.Background(), reportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetReturn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64

		resp, httpRes, err := apiClient.FbsAPI.GetReturn(context.Background(), campaignId, orderId, returnId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetReturnApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64

		resp, httpRes, err := apiClient.FbsAPI.GetReturnApplication(context.Background(), campaignId, orderId, returnId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetReturnPhoto", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64
		var itemId int64
		var imageHash string

		resp, httpRes, err := apiClient.FbsAPI.GetReturnPhoto(context.Background(), campaignId, orderId, returnId, itemId, imageHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetReturns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetReturns(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.GetShipment(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetShipmentOrdersInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.GetShipmentOrdersInfo(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetStocks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetStocks(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetSuggestedOfferMappingEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetSuggestedOfferMappingEntries(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetSuggestedOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetSuggestedOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetSuggestedPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.GetSuggestedPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService GetWarehouses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.GetWarehouses(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService ProvideOrderItemIdentifiers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbsAPI.ProvideOrderItemIdentifiers(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService PutBidsForBusiness", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.PutBidsForBusiness(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService SearchRegionChildren", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var regionId int64

		resp, httpRes, err := apiClient.FbsAPI.SearchRegionChildren(context.Background(), regionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService SearchRegionsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var regionId int64

		resp, httpRes, err := apiClient.FbsAPI.SearchRegionsById(context.Background(), regionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService SearchRegionsByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbsAPI.SearchRegionsByName(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService SearchShipments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.SearchShipments(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService SetOrderShipmentBoxes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.SetOrderShipmentBoxes(context.Background(), campaignId, orderId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService SetShipmentPalletsCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.SetShipmentPalletsCount(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService TransferOrdersFromShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.FbsAPI.TransferOrdersFromShipment(context.Background(), campaignId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateBusinessPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateBusinessPrices(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateOfferContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateOfferContent(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateOfferMappingEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateOfferMappingEntries(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateOrderItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		httpRes, err := apiClient.FbsAPI.UpdateOrderItems(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateOrderStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateOrderStatus(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateOrderStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateOrderStatuses(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdatePrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdatePrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbsAPIService UpdateStocks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbsAPI.UpdateStocks(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}