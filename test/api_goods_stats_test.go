/*
Партнерский API Маркета

Testing GoodsStatsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/devteamclub/ya-client"
)

func Test_client_GoodsStatsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GoodsStatsAPIService GetGoodsStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.GoodsStatsAPI.GetGoodsStats(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
