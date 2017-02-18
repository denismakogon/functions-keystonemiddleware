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

import (
	"github.com/iron-io/functions/api/server"
	"github.com/iron-io/functions/api/models"
	"net/http"
	"strings"
	"encoding/json"
	"errors"
	"bytes"
)


type KeystoneV3TokenAuth struct {

	AuthURL string
	// Project-scoped token that comes as X-Auth-Token
	Token string
	// Project ID that comes as X-Project-ID
	ProjectID string
	// Project Domain ID that comes as X-Domain-ID
	ProjectDomainID string
}


func (tkv3 *KeystoneV3TokenAuth) Serve(ctx server.MiddlewareContext, w http.ResponseWriter, r *http.Request, app *models.App) error {

	tkv3.Token, tkv3.ProjectID, tkv3.ProjectDomainID =
		r.Header.Get("X-Auth-Token"),
		r.Header.Get("X-Project-ID"),
		r.Header.Get("X-Domain-ID")
	token_url := strings.Join([]string{tkv3.AuthURL, "auth/tokens"}, "/")
	bytesData, err := PrepareAuthData(tkv3.Token, tkv3.ProjectID, tkv3.ProjectDomainID)

	if err != nil {
		m := map[string]string{"error": "Invalid Authorization token"}
		json.NewEncoder(w).Encode(m)
		return errors.New("Invalid authorization token.")
	}

	response, err := http.Post(token_url, "Content-Type: application/json", bytes.NewReader(bytesData))
	if response.StatusCode != http.StatusCreated {
		m := map[string]string{"error": "Invalid Authorization token"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(m)
		return errors.New("Invalid authorization token.")
	}

	return nil
}
