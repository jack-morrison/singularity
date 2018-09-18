// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package libexec

import (
	"fmt"

	library "github.com/singularityware/singularity/src/pkg/client/library"
	"github.com/singularityware/singularity/src/pkg/client/oci"
	shub "github.com/singularityware/singularity/src/pkg/client/shub"
	"github.com/singularityware/singularity/src/pkg/sylog"
)

// PullLibraryImage is the function that is responsible for pulling an image from a Sylabs library.
func PullLibraryImage(image string, libraryRef string, libraryURL string, force bool, authToken string) {
	err := library.DownloadImage(image, libraryRef, libraryURL, force, authToken)
	if err != nil {
		sylog.Fatalf("%v\n", err)
	}
}

// PullShubImage is the function that is responsible for pulling an image from a Singularity Hub.
func PullShubImage(filePath string, shubRef string, force bool) {
	err := shub.DownloadImage(filePath, shubRef, force)
	if err != nil {
		sylog.Fatalf("%v\n", err)
	}
}

// PullOciImage ...
func PullOciImage(uri string) {
	image, err := oci.ImageSHA(uri)
	fmt.Println(image)

	if err != nil {
		sylog.Fatalf("Unable to calculate uri hash: %v", err)
	}
}
