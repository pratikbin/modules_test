package split

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryKeeper struct {
	mapper utility.Mapper
}

var _ utility.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest utility.QueryRequest) utility.QueryResponse {
	return newQueryResponse(mapper.NewSplits(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).SplitID))
}

func initializeQueryKeeper(mapper utility.Mapper) utility.QueryKeeper {
	return queryKeeper{mapper: mapper}
}