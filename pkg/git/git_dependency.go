// Copyright (c) 2019-2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package git

import (
	"log"
	"regexp"
)

func IsValidGitDependency(gitUrl string) bool {
	valid, err := regexp.MatchString(`((git|ssh|http(s)?)|(git@[\w\.]+))(:)([\w\.@\:/\-~]+)(/)?`, gitUrl)
	if err != nil {
		log.Printf("error when matching regex to validate git dependency: %s", err)
	}
	return valid
}
