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

package sign

import (
	"math/big"

	"github.com/manishsharma864/alice/crypto/birkhoffinterpolation"
	pt "github.com/manishsharma864/alice/crypto/ecpointgrouplaw"
	"github.com/manishsharma864/alice/crypto/homo/paillier"
	"github.com/manishsharma864/alice/crypto/tss"
	paillierzkproof "github.com/manishsharma864/alice/crypto/zkproof/paillier"
	"github.com/manishsharma864/alice/internal/message"
	"github.com/manishsharma864/alice/internal/message/types"
	"github.com/manishsharma864/sirius/log"
)

type Sign struct {
	ph *round1Handler
	*message.MsgMain

	msgMainer types.MessageMain
}

func NewSign(threshold uint32, ssid []byte, share *big.Int, pubKey *pt.ECPoint, partialPubKey, allY map[string]*pt.ECPoint, paillierKey *paillier.Paillier, ped map[string]*paillierzkproof.PederssenOpenParameter, bks map[string]*birkhoffinterpolation.BkParameter, msg []byte, peerManager types.PeerManager, listener types.StateChangedListener) (*Sign, error) {
	peerNum := peerManager.NumPeers()
	ph, err := newRound1Handler(threshold, ssid, share, pubKey, partialPubKey, allY, paillierKey, ped, bks, msg, peerManager)
	if err != nil {
		return nil, err
	}
	ms := message.NewMsgMain(peerManager.SelfID(), peerNum, listener, ph, types.MessageType(Type_Round1), types.MessageType(Type_Round2), types.MessageType(Type_Round3), types.MessageType(Type_Round4))
	msgMainer := message.NewEchoMsgMain(ms, peerManager, types.MessageType(Type_Round4))
	return &Sign{
		ph:        ph,
		MsgMain:   ms,
		msgMainer: msgMainer,
	}, nil
}

// GetResult returns the final result: public key, share, bks (including self bk)
func (d *Sign) GetResult() (*Result, error) {
	if d.GetState() != types.StateDone {
		return nil, tss.ErrNotReady
	}

	h := d.GetHandler()
	rh, ok := h.(*round4Handler)
	if !ok {
		log.Error("We cannot convert to result handler in done state")
		return nil, tss.ErrNotReady
	}

	return rh.result, nil
}

func (d *Sign) AddMessage(msg types.Message) error {
	return d.msgMainer.AddMessage(msg)
}

func (d *Sign) Start() {
	d.MsgMain.Start()

	d.ph.sendRound1Messages()
}
