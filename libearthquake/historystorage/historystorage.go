// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package historystorage

import (
	"fmt"
	"time"

	. "../equtils"
	"./naive"
)

type HistoryStorage interface {
	CreateStorage()

	Init()
	Name() string

	CreateNewWorkingDir() string
	RecordNewTrace(newTrace *SingleTrace)
	RecordResult(succeed bool, requiredTime time.Duration) error

	NrStoredHistories() int
	GetStoredHistory(id int) (*SingleTrace, error)

	IsSucceed(id int) (bool, error)
}

func New(name, dirPath string) HistoryStorage {
	switch name {
	case "naive":
		return naive.New(dirPath)
	default:
		fmt.Printf("unknown history storage: %s\n", name)
	}

	return nil
}
