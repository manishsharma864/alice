// Copyright © 2022 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package refresh

import (
	"github.com/manishsharma864/alice/internal/message"
)

type peer struct {
	*message.Peer
	round2 *round2Data
	round3 *round3Data
}

func newPeer(id string) *peer {
	return &peer{
		Peer: message.NewPeer(id),
	}
}
