// Copyright 2022-2023 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package accessor

import "github.com/sacloud/iaas-api-go/types"

/************************************************
 switchID
************************************************/

// SwitchID is accessor interface of SwitchID field
type SwitchID interface {
	GetSwitchID() types.ID
	SetSwitchID(id types.ID)
}
