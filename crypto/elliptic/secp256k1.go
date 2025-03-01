// Copyright © 2021 AMIS Technologies
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

package elliptic

import (
	"crypto/elliptic"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

var (
	secp256k1Curve = &secp256k1{
		Curve: btcec.S256(),
	}
)

type secp256k1 struct {
	elliptic.Curve
}

func Secp256k1() *secp256k1 {
	return secp256k1Curve
}

// Warn: does not deal with the original point
func (sep *secp256k1) Neg(x, y *big.Int) (*big.Int, *big.Int) {
	NegY := new(big.Int).Neg(y)
	return new(big.Int).Set(x), NegY.Mod(NegY, sep.Curve.Params().P)
}
