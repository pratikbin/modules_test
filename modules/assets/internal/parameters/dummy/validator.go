// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/types"
)

func validator(i interface{}) error {
	switch value := i.(type) {
	case types.Parameter:
		data := value.GetData().(data.DecData).Get()
		if value.GetID().Compare(ID) != 0 || data.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	case types.Data:
		data := value.(data.DecData).Get()
		if data.IsNegative() {
			return errors.InvalidParameter
		}

		return nil
	default:
		return errors.IncorrectFormat
	}
}
