// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/gitbundle/server/build/codeformat"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Usage: gitea-format-imports [files...]")
	}

	for _, file := range os.Args[1:] {
		if err := codeformat.FormatGoImports(file, true, true); err != nil {
			log.Fatalf("can not format file %s, err=%v", file, err)
		}
	}
}
