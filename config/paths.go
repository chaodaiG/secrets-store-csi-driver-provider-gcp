// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"path/filepath"
)

// JoinPaths is a safer way to append a filename to a base path. This should be
// used when constructing subpaths of targetPath for writing files to avoid path
// traversal conditions.
func JoinPaths(base, filename string) (string, error) {
	if err := ValidateFilename(filename); err != nil {
		return "", err
	}
	return filepath.Join(base, filename), nil
}

// ValidateFilename ensures there are no path separators in the filename.
func ValidateFilename(filename string) error {
	part := filepath.Base(filename)
	if part != filename {
		// filename includes path separator
		return fmt.Errorf("invalid path: %q", filename)
	}
	return nil
}
