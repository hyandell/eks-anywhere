// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import (
	"fmt"
	"path/filepath"

	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	"github.com/pkg/errors"
)

// GetCliArtifacts returns the artifacts for eks-a cli
func (r *ReleaseConfig) GetEksACliArtifacts() ([]Artifact, error) {
	osList := []string{"linux", "darwin"}
	arch := "amd64"

	artifacts := []Artifact{}
	for _, os := range osList {
		releaseName := fmt.Sprintf("eks-a-%s-%s-%s.tar.gz", r.ReleaseVersion, os, arch)

		var sourceS3Key string
		var sourceS3Prefix string
		var releaseS3Path string

		if r.DevRelease || r.ReleaseEnvironment == "development" {
			sourceS3Key = fmt.Sprintf("eks-a-%s-%s.tar.gz", os, arch)
			sourceS3Prefix = fmt.Sprintf("eks-a-cli/latest/%s", os)
		} else {
			sourceS3Key = fmt.Sprintf("eks-a-%s-%s-%s.tar.gz", r.ReleaseVersion, os, arch)
			sourceS3Prefix = fmt.Sprintf("releases/eks-a/%d/artifacts/eks-a/%s/%s", r.ReleaseNumber, r.ReleaseVersion, os)
		}

		if r.DevRelease {
			releaseS3Path = fmt.Sprintf("eks-anywhere/%s/eks-a-cli/%s", r.DevReleaseUriVersion, os)
		} else {
			releaseS3Path = fmt.Sprintf("releases/eks-a/%d/artifacts/eks-a/%s/%s", r.ReleaseNumber, r.ReleaseVersion, os)
		}

		cdnURI, err := r.GetURI(filepath.Join(releaseS3Path, releaseName))
		if err != nil {
			return nil, errors.Cause(err)
		}

		archiveArtifact := &ArchiveArtifact{
			SourceS3Key:    sourceS3Key,
			SourceS3Prefix: sourceS3Prefix,
			ArtifactPath:   filepath.Join(r.ArtifactDir, "eks-a", r.CliRepoHead),
			ReleaseName:    releaseName,
			ReleaseS3Path:  releaseS3Path,
			ReleaseCdnURI:  cdnURI,
			OS:             os,
			Arch:           []string{arch},
		}

		artifacts = append(artifacts, Artifact{Archive: archiveArtifact})
	}
	return artifacts, nil
}

func (r *ReleaseConfig) GetEksARelease() (anywherev1alpha1.EksARelease, error) {
	artifacts, err := r.GetEksACliArtifacts()
	if err != nil {
		return anywherev1alpha1.EksARelease{}, errors.Cause(err)
	}

	var bundleManifestUrl string
	bundleArchiveArtifacts := map[string]anywherev1alpha1.Archive{}

	for _, artifact := range artifacts {
		archiveArtifact := artifact.Archive

		tarfile := filepath.Join(archiveArtifact.ArtifactPath, archiveArtifact.ReleaseName)
		sha256, sha512, err := r.readShaSums(tarfile)
		if err != nil {
			return anywherev1alpha1.EksARelease{}, errors.Cause(err)
		}

		if r.DevRelease {
			bundleManifestUrl = fmt.Sprintf("%s/bundle-release.yaml", r.CDN)
		} else {
			bundleManifestUrl = fmt.Sprintf("%s/releases/bundles/%d/manifest.yaml", r.CDN, r.BundleNumber)
		}

		bundleArchiveArtifact := anywherev1alpha1.Archive{
			Name:        fmt.Sprintf("eks-a-%s", archiveArtifact.OS),
			Description: "EKS Anywhere CLI",
			OS:          archiveArtifact.OS,
			Arch:        archiveArtifact.Arch,
			URI:         archiveArtifact.ReleaseCdnURI,
			SHA256:      sha256,
			SHA512:      sha512,
		}

		bundleArchiveArtifacts[fmt.Sprintf("eks-a-%s", archiveArtifact.OS)] = bundleArchiveArtifact
	}

	eksARelease := anywherev1alpha1.EksARelease{
		Date:      r.ReleaseDate.String(),
		Version:   r.ReleaseVersion,
		Number:    r.ReleaseNumber,
		GitCommit: r.CliRepoHead,
		GitTag:    r.ReleaseVersion,
		EksABinary: anywherev1alpha1.BinaryBundle{
			LinuxBinary:  bundleArchiveArtifacts["eks-a-linux"],
			DarwinBinary: bundleArchiveArtifacts["eks-a-darwin"],
		},
		BundleManifestUrl: bundleManifestUrl,
	}

	return eksARelease, nil
}