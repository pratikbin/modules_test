package asset

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryResponse struct {
	Assets schema.InterNFTs
}

var _ utility.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() utility.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(assets schema.InterNFTs) utility.QueryResponse {
	return queryResponse{Assets: assets}
}