// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func GenerateRandomID(r *rand.Rand) types.ID {
	return base.NewID(simulation.RandStringOfLength(r, r.Intn(99)))
}

func GenerateRandomIDWithDec(r *rand.Rand) types.ID {
	return base.NewID(sdkTypes.MustNewDecFromStr(strconv.FormatInt(r.Int63(), 10)).String())
}

func GenerateRandomIDWithInt64(r *rand.Rand) types.ID {
	return base.NewID(strconv.FormatInt(r.Int63(), 10))
}
