package bubble

// Copyright [2020] [freshpassport@gmail.com]
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

import "sort"

func Sort(dataItems sort.Interface) {
	totalLen := dataItems.Len()
	for i := 0; i < totalLen; i++ {
		isChanged := false
		for j := 0; j < totalLen-i-1; j++ {
			if dataItems.Less(j, j+1) {
				dataItems.Swap(j, j+1)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}
