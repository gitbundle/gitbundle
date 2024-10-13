// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package oauth2

// Name returns the provider name of this source
func (source *Source) Name() string {
	return source.Provider
}

// DisplayName returns the display name of this source
func (source *Source) DisplayName() string {
	provider, has := gothProviders[source.Provider]
	if !has {
		return source.Provider
	}
	return provider.DisplayName()
}
