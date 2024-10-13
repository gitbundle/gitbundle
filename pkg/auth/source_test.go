// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package auth

import (
	"strings"
	"testing"

	"github.com/gitbundle/modules/json"
	"github.com/gitbundle/server/pkg/db"
	"github.com/gitbundle/server/pkg/unittest"

	"github.com/stretchr/testify/assert"
	"xorm.io/xorm/schemas"
)

type TestSource struct {
	Provider                      string
	ClientID                      string
	ClientSecret                  string
	OpenIDConnectAutoDiscoveryURL string
	IconURL                       string
}

// FromDB fills up a LDAPConfig from serialized format.
func (source *TestSource) FromDB(bs []byte) error {
	return json.Unmarshal(bs, &source)
}

// ToDB exports a LDAPConfig to a serialized format.
func (source *TestSource) ToDB() ([]byte, error) {
	return json.Marshal(source)
}

func TestDumpAuthSource(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	authSourceSchema, err := db.TableInfo(new(Source))
	assert.NoError(t, err)

	RegisterTypeConfig(OAuth2, new(TestSource))

	CreateSource(&Source{
		Type:     OAuth2,
		Name:     "TestSource",
		IsActive: false,
		Cfg: &TestSource{
			Provider: "ConvertibleSourceName",
			ClientID: "42",
		},
	})

	sb := new(strings.Builder)

	db.DumpTables([]*schemas.Table{authSourceSchema}, sb)

	assert.Contains(t, sb.String(), `"Provider":"ConvertibleSourceName"`)
}
