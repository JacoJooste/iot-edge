/*
 * Copyright 2020 ForgeRock AS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package anvil

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	jose "gopkg.in/square/go-jose.v2"
)

// GenerateConfirmationKey generates a key for signing requests to AM that is accompanied by a restricted PoP SSO token.
func GenerateConfirmationKey(algorithm jose.SignatureAlgorithm) (public jose.JSONWebKeySet, private crypto.Signer, err error) {
	// create a new key
	switch algorithm {
	case jose.ES256:
		private, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case jose.ES384:
		private, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case jose.ES512:
		private, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	case jose.EdDSA:
		_, private, err = ed25519.GenerateKey(rand.Reader)
	case jose.PS256:
		private, err = rsa.GenerateKey(rand.Reader, 2048)
	case jose.PS384:
		private, err = rsa.GenerateKey(rand.Reader, 3072)
	case jose.PS512:
		private, err = rsa.GenerateKey(rand.Reader, 4096)
	default:
		err = fmt.Errorf("unsupported signing algorithm %s", algorithm)
	}
	if err != nil {
		return public, private, err
	}
	return jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{{KeyID: "pop.cnf", Key: private.Public(), Algorithm: string(algorithm), Use: "sig"}},
	}, private, nil
}
