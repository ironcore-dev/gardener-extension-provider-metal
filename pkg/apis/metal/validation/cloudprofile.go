// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"fmt"

	gardenercore "github.com/gardener/gardener/pkg/apis/core"
	gardenercorehelper "github.com/gardener/gardener/pkg/apis/core/helper"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/strings/slices"

	apismetal "github.com/ironcore-dev/gardener-extension-provider-metal/pkg/apis/metal"
)

// ValidateCloudProfileConfig validates a CloudProfileConfig object.
func ValidateCloudProfileConfig(cpConfig *apismetal.CloudProfileConfig, machineImages []gardenercore.MachineImage, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	machineImagesPath := fldPath.Child("machineImages")

	for _, image := range machineImages {
		var processed bool
		for i, imageConfig := range cpConfig.MachineImages {
			if image.Name == imageConfig.Name {
				allErrs = append(allErrs, validateVersions(imageConfig.Versions, gardenercorehelper.ToExpirableVersions(image.Versions), machineImagesPath.Index(i).Child("versions"))...)
				processed = true
				break
			}
		}
		if !processed && len(image.Versions) > 0 {
			allErrs = append(allErrs, field.Required(machineImagesPath, fmt.Sprintf("must provide an image mapping for image %q", image.Name)))
		}
	}

	return allErrs
}

func validateVersions(versionsConfig []apismetal.MachineImageVersion, versions []gardenercore.ExpirableVersion, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	for _, version := range versions {
		var processed bool
		for j, versionConfig := range versionsConfig {
			jdxPath := fldPath.Index(j)
			if version.Version == versionConfig.Version {
				if len(versionConfig.Image) == 0 {
					allErrs = append(allErrs, field.Required(jdxPath.Child("image"), "must provide an image"))
				}
				if !slices.Contains(v1beta1constants.ValidArchitectures, *versionConfig.Architecture) {
					allErrs = append(allErrs, field.NotSupported(jdxPath.Child("architecture"), *versionConfig.Architecture, v1beta1constants.ValidArchitectures))
				}
				processed = true
				break
			}
		}
		if !processed {
			allErrs = append(allErrs, field.Required(fldPath, fmt.Sprintf("must provide an image mapping for version %q", version.Version)))
		}
	}

	return allErrs
}
