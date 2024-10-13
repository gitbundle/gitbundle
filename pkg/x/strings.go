// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package x

import "strings"

func TrimHTTPPrefix(url string) string {
	if strings.HasPrefix(url, "http://") {
		url = url[len("http://"):]
	} else {
		url = strings.TrimPrefix(url, "https://")
	}
	return strings.TrimSpace(url)
}
