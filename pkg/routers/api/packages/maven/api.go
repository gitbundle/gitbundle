// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package maven

import (
	"encoding/xml"
	"sort"
	"strings"

	maven_module "github.com/gitbundle/modules/packages/maven"
	packages_model "github.com/gitbundle/server/pkg/packages"
)

// MetadataResponse https://maven.apache.org/ref/3.2.5/maven-repository-metadata/repository-metadata.html
type MetadataResponse struct {
	XMLName    xml.Name `xml:"metadata"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Release    string   `xml:"versioning>release,omitempty"`
	Latest     string   `xml:"versioning>latest"`
	Version    []string `xml:"versioning>versions>version"`
}

func createMetadataResponse(pds []*packages_model.PackageDescriptor) *MetadataResponse {
	sort.Slice(pds, func(i, j int) bool {
		// Maven and Gradle order packages by their creation timestamp and not by their version string
		return pds[i].Version.CreatedUnix < pds[j].Version.CreatedUnix
	})

	var release *packages_model.PackageDescriptor

	versions := make([]string, 0, len(pds))
	for _, pd := range pds {
		if !strings.HasSuffix(pd.Version.Version, "-SNAPSHOT") {
			release = pd
		}
		versions = append(versions, pd.Version.Version)
	}

	latest := pds[len(pds)-1]

	metadata := latest.Metadata.(*maven_module.Metadata)

	resp := &MetadataResponse{
		GroupID:    metadata.GroupID,
		ArtifactID: metadata.ArtifactID,
		Latest:     latest.Version.Version,
		Version:    versions,
	}
	if release != nil {
		resp.Release = release.Version.Version
	}
	return resp
}
