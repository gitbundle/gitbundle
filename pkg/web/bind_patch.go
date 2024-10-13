// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package web

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"

	"gitea.com/go-chi/binding"
)

func PatchSliceStruct(form url.Values, formStruct interface{}) (errors binding.Errors) {
	formStructV := reflect.ValueOf(formStruct)
	if formStructV.Kind() == reflect.Ptr {
		formStructV = formStructV.Elem()
	}

	var (
		re = regexp.MustCompile("(.*)\\[([0-9]+)\\]\\[(.*)\\]")
	)

	for key, values := range form {
		matches := re.FindStringSubmatch(key)

		if len(matches) == 4 {
			index, err := strconv.Atoi(matches[2])
			if err != nil {
				errors = append(errors, binding.Error{Message: fmt.Sprintf("strconv.Atoi (%v), err: %v", matches[2], err)})
				continue
			}
			errors = mapFromKV(formStructV, index, matches[1], matches[3], values, errors)
		}
	}

	return
}

// Takes values from the form data and puts them into a struct
func mapFromKV(formStruct reflect.Value, index int, pk, sk string, sv []string, errors binding.Errors) binding.Errors {
	if formStruct.Kind() == reflect.Ptr {
		formStruct = formStruct.Elem()
	}
	typ := formStruct.Type()

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := formStruct.Field(i)

		inputFieldName := parseFormName(typeField.Name, typeField.Tag.Get("form"))
		if len(inputFieldName) == 0 || !structField.CanSet() || inputFieldName != pk || structField.Kind() != reflect.Slice {
			continue
		}

		if structField.IsNil() {
			slice := reflect.MakeSlice(structField.Type(), 1, 8)
			structField.Set(slice)
		}

		for index >= structField.Len() {
			element := structField.Index(0).Type()
			if element.Kind() == reflect.Ptr {
				element = element.Elem()
			}
			obj := reflect.New(element)
			if obj.Kind() == reflect.Ptr {
				obj = obj.Elem()
			}
			structField.Set(reflect.Append(structField, obj))
		}

		indexStruct := structField.Index(index)
		indexStructType := indexStruct.Type()
		for j := 0; j < indexStructType.NumField(); j++ {
			typeField := indexStructType.Field(j)
			structField := indexStruct.Field(j)

			inputFieldName := parseFormName(typeField.Name, typeField.Tag.Get("form"))
			if len(inputFieldName) == 0 || !structField.CanSet() || inputFieldName != sk {
				continue
			}

			if numElems := len(sv); structField.Kind() == reflect.Slice && numElems > 0 {
				sliceOf := structField.Type().Elem().Kind()
				slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
				for i := 0; i < numElems; i++ {
					errors = setWithProperType(sliceOf, sv[i], slice.Index(i), inputFieldName, errors)
				}
				indexStruct.Field(j).Set(slice)
			} else {
				errors = setWithProperType(typeField.Type.Kind(), sv[0], structField, inputFieldName, errors)
			}
			break
		}
		break
	}
	return errors
}

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value, nameInTag string, errors binding.Errors) binding.Errors {
	switch valueKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val == "" {
			val = "0"
		}
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			errors.Add([]string{nameInTag}, binding.ERR_INTERGER_TYPE, "Value could not be parsed as integer")
		} else {
			structField.SetInt(intVal)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val == "" {
			val = "0"
		}
		uintVal, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			errors.Add([]string{nameInTag}, binding.ERR_INTERGER_TYPE, "Value could not be parsed as unsigned integer")
		} else {
			structField.SetUint(uintVal)
		}
	case reflect.Bool:
		if val == "on" {
			structField.SetBool(true)
			break
		}

		if val == "" {
			val = "false"
		}
		boolVal, err := strconv.ParseBool(val)
		if err != nil {
			errors.Add([]string{nameInTag}, binding.ERR_BOOLEAN_TYPE, "Value could not be parsed as boolean")
		} else if boolVal {
			structField.SetBool(true)
		}
	case reflect.Float32:
		if val == "" {
			val = "0.0"
		}
		floatVal, err := strconv.ParseFloat(val, 32)
		if err != nil {
			errors.Add([]string{nameInTag}, binding.ERR_FLOAT_TYPE, "Value could not be parsed as 32-bit float")
		} else {
			structField.SetFloat(floatVal)
		}
	case reflect.Float64:
		if val == "" {
			val = "0.0"
		}
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			errors.Add([]string{nameInTag}, binding.ERR_FLOAT_TYPE, "Value could not be parsed as 64-bit float")
		} else {
			structField.SetFloat(floatVal)
		}
	case reflect.String:
		structField.SetString(val)
	}
	return errors
}

func nameMapper(field string) string {
	newstr := make([]rune, 0, len(field))
	for i, chr := range field {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if i > 0 {
				newstr = append(newstr, '_')
			}
			chr -= 'A' - 'a'
		}
		newstr = append(newstr, chr)
	}
	return string(newstr)
}

func parseFormName(raw, actual string) string {
	if len(actual) > 0 {
		return actual
	}
	return nameMapper(raw)
}
