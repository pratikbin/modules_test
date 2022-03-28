/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/common"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(module.Name,
			dummy.ID.String(),
			func(r *rand.Rand) string {
				bytes, err := common.Codec.MarshalJSON(dummy.Parameter.Mutate(base.NewDecData(sdk.NewDecWithPrec(int64(r.Intn(99)), 2))).GetData())
				if err != nil {
					panic(err)
				}
				return string(bytes)
			}),
	}
}