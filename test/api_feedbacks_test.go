/*
Партнерский API Маркета

Testing FeedbacksAPIService

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

func Test_ya-client_FeedbacksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeedbacksAPIService GetFeedbackAndCommentUpdates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FeedbacksAPI.GetFeedbackAndCommentUpdates(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
