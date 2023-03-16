// Copyright 2023 Undistro Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import "fmt"

type CheckStatus int

const (
	StatusPassed CheckStatus = iota
	StatusSkipped
	StatusFailed
	StatusError
)

func (s CheckStatus) String() string {
	switch s {
	case StatusPassed:
		return "Passed"
	case StatusSkipped:
		return "Skipped"
	case StatusFailed:
		return "Failed"
	case StatusError:
		return "Error"
	default:
		return ""
	}
}

func (s CheckStatus) MarshalJSON() ([]byte, error) {
	v := fmt.Sprintf(`"%s"`, s.String())
	return []byte(v), nil
}