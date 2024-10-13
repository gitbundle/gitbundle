// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	container_module "github.com/gitbundle/modules/packages/container"
	packages_model "github.com/gitbundle/server/pkg/packages"

	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

func addContainerRepositoryProperty(x *xorm.Engine) (err error) {
	switch x.Dialect().URI().DBType {
	case schemas.SQLITE:
		_, err = x.Exec(
			"INSERT INTO package_property (ref_type, ref_id, name, value) SELECT ?, p.id, ?, u.lower_name || '/' || p.lower_name FROM package p JOIN `user` u ON p.owner_id = u.id WHERE p.type = ?",
			packages_model.PropertyTypePackage,
			container_module.PropertyRepository,
			packages_model.TypeContainer,
		)
	case schemas.MSSQL:
		_, err = x.Exec(
			"INSERT INTO package_property (ref_type, ref_id, name, value) SELECT ?, p.id, ?, u.lower_name + '/' + p.lower_name FROM package p JOIN `user` u ON p.owner_id = u.id WHERE p.type = ?",
			packages_model.PropertyTypePackage,
			container_module.PropertyRepository,
			packages_model.TypeContainer,
		)
	default:
		_, err = x.Exec(
			"INSERT INTO package_property (ref_type, ref_id, name, value) SELECT ?, p.id, ?, CONCAT(u.lower_name, '/', p.lower_name) FROM package p JOIN `user` u ON p.owner_id = u.id WHERE p.type = ?",
			packages_model.PropertyTypePackage,
			container_module.PropertyRepository,
			packages_model.TypeContainer,
		)
	}
	return err
}
