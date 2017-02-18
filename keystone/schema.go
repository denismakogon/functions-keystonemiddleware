// Copyright 2016 Iron.io
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

package keystone

import "encoding/json"


type KeystoneToken struct {
	ID string `json:"id"`
}


type KeystoneIdentity struct {
	Method []string `json:"method"`
	Token KeystoneToken `json:"token"`
}

type KeystoneProjectDomain struct {
	ID string `json:"id"`
}


type KeystoneProject struct {
	ID string `json:"id"`
	Domain KeystoneProjectDomain `json:"domain"`
}


type KeystoneScope struct {
	Project KeystoneProject `json:"project"`
}


type KeystoneAuthConfig struct {
	Identity KeystoneIdentity `json:"identity"`
	Scope KeystoneScope `json:"scope"`
}

type KeystoneV3TokenAuthSchema struct {
	Auth KeystoneAuthConfig `json:"auth"`
}

func PrepareAuthData(token string, projectID string, projectDomainID string) ([]byte, error) {
	identity := KeystoneIdentity{Method: []string{"token",}, Token: KeystoneToken{ID: token}}
	scope := KeystoneScope{Project: KeystoneProject{ID: projectID, Domain: KeystoneProjectDomain{ID: projectDomainID}}}
	auth := KeystoneV3TokenAuthSchema{Auth: KeystoneAuthConfig{Identity: identity, Scope: scope}}

	return json.Marshal(auth)
}
