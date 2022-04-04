// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import "github.com/AssetMantle/modules/schema/types"

type height struct {
	Value int64 `json:"height"`
}

var _ types.Height = (*height)(nil)

func (height height) Get() int64 { return height.Value }
func (height height) Compare(compareHeight types.Height) int {
	if height.Get() > compareHeight.Get() {
		return 1
	} else if height.Get() < compareHeight.Get() {
		return -1
	}

	return 0
}
func NewHeight(value int64) types.Height {
	return height{Value: value}
}
