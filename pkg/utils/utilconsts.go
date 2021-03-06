// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import "k8s.io/klog"

const (
	infoLevel  = klog.Level(5)
	debugLevel = klog.Level(10)
	// HelmCRKind is kind of the Helm CR
	HelmCRKind = "HelmRelease"
	// HelmCRAPIVersion is APIVersion of the Helm CR
	HelmCRAPIVersion = "app.ibm.com/v1alpha1"
	// HelmCRChartName is spec.ChartName of the Helm CR
	HelmCRChartName = "chartName"
	// HelmCRReleaseName is spec.ReleaseName of the Helm CR
	HelmCRReleaseName = "releaseName"
	// HelmCRVersion is spec.Version of the Helm CR
	HelmCRVersion = "version"
	// HelmCRSource is spec.Source of the Helm CR
	HelmCRSource = "source"
	// HelmCRSourceType is spec.Source.Type of the Helm CR
	HelmCRSourceType = "type"
	// HelmCRSourceHelm is spec.Source.Helmrepo of the Helm CR
	HelmCRSourceHelm = "helmrepo"
	// HelmCRSourceGit is spec.Source.Github of the Helm CR
	HelmCRSourceGit = "github"
	// HelmCRRepoURL is spec.Source.Github.Urls or spec.Source.Helmrepo.Urls of the Helm CR
	HelmCRRepoURL = "urls"
	// HelmCRGitRepoChartPath is spec.Source.Github.ChartPath of the Helm CR
	HelmCRGitRepoChartPath = "chartPath"
)
