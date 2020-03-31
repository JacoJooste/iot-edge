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

package things

import (
	"github.com/ForgeRock/iot-edge/internal/amtest"
	"testing"
)

const (
	testAddress = "127.0.0.1:8008"
	testURL     = "http://" + testAddress
)

func TestAMClient_Initialise(t *testing.T) {
	server := amtest.NewSimpleServer().Start(testAddress)
	defer server.Close()
	c, err := NewAMClient(testURL, amtest.SimpleTestRealm).Initialise()
	if err != nil {
		t.Fatal(err)
	}
	// check that the cookName has been set on the struct
	if c.(*AMClient).cookieName != amtest.CookieName {
		t.Error("Cookie name has not been set")
	}
}

func TestAMClient_Authenticate(t *testing.T) {
	server := amtest.NewSimpleServer().Start(testAddress)
	defer server.Close()
	c, err := NewAMClient(testURL, amtest.SimpleTestRealm).Initialise()
	if err != nil {
		t.Fatal(err)
	}
	reply, err := c.Authenticate(amtest.SimpleTestAuthTree, amtest.SimpleAuthPayload)
	if err != nil {
		t.Fatal(err)
	}
	// check that the reply has a token
	if reply.TokenID == "" {
		t.Errorf("Expected an token in reply: %v", reply)
	}
}
