// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package school

import (
	"encoding/json"
)

// DefaultSessionServicePath is the default URI for SessionService collections.
const DefaultSessionServicePath = "/redfish/v1/SessionService"

// Session describes a single connection (session) between a client and a
// Redfish service instance.
type Session struct {
	Entity
	Description string
	Modified    string
	UserName    string
}

// GetSession will get a Session instance from the Redfish service.
func GetSession(c Client, uri string) (*Session, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var t Session
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ListReferencedSessions gets the collection of Sessions
func ListReferencedSessions(c Client, link string) ([]*Session, error) {
	var result []*Session
	links, err := GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, sLink := range links.ItemLinks {
		s, err := GetSession(c, sLink)
		if err != nil {
			return result, err
		}
		result = append(result, s)
	}

	return result, nil
}

// ListSessions gets all Session in the system
func ListSessions(c Client) ([]*Session, error) {
	return ListReferencedSessions(c, DefaultSessionServicePath)
}