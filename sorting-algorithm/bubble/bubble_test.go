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

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortInt(t *testing.T) {
	data := sort.IntSlice{10, -1, 15, 7, 3, 11}
	Sort(data)
	fmt.Println(data)
}

type User struct {
	Name    string
	RegTime int64
}

func (u User) String() string {
	return fmt.Sprintf("@(Name: %s RegTime: %d)", u.Name, u.RegTime)
}

type UserSlice []User

func (p UserSlice) Len() int           { return len(p) }
func (p UserSlice) Less(i, j int) bool { return p[i].RegTime < p[j].RegTime }
func (p UserSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestSortUser(t *testing.T) {
	users := []User{
		User{Name: "张三", RegTime: 1592179200},
		User{Name: "李四", RegTime: 1600732800},
		User{Name: "赵武", RegTime: 1605548113},
	}
	Sort(UserSlice(users))
	fmt.Println(users)
}
