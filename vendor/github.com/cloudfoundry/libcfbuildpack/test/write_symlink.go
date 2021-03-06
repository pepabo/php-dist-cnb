/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// WriteSymlink creates newName as a symbolic link to oldName.  Before writing, it creates all required parent
// directories for the newName.
func WriteSymlink(t *testing.T, oldName string, newName string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(newName), 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.Symlink(oldName, newName); err != nil {
		t.Fatal(fmt.Errorf("error while creating '%s' as symlink to '%s': %v", newName, oldName, err))
	}
}
