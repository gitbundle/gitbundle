// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package unittest

import (
	"log"
	"reflect"
)

func fieldByName(v reflect.Value, field string) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	f := v.FieldByName(field)
	if !f.IsValid() {
		log.Panicf("can not read %s for %v", field, v)
	}
	return f
}

type reflectionValue struct {
	v reflect.Value
}

func reflectionWrap(v interface{}) *reflectionValue {
	return &reflectionValue{v: reflect.ValueOf(v)}
}

func (rv *reflectionValue) int(field string) int {
	return int(fieldByName(rv.v, field).Int())
}

func (rv *reflectionValue) str(field string) string {
	return fieldByName(rv.v, field).String()
}

func (rv *reflectionValue) bool(field string) bool {
	return fieldByName(rv.v, field).Bool()
}
