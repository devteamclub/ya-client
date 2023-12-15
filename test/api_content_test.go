/*
Партнерский API Маркета

Testing ContentAPIService

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

func Test_client_ContentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentAPIService GetCategoryContentParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.ContentAPI.GetCategoryContentParameters(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAPIService GetOfferCardsContentStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.ContentAPI.GetOfferCardsContentStatus(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAPIService UpdateOfferContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.ContentAPI.UpdateOfferContent(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
