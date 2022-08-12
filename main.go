/*
Based on work of Copyright 2021 Pelotech.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/containerd/containerd/runtime/v2/shim"

	readonlyrootfsShim "readonlyrootfs-runc-shim/pkg/shim"
)

func main() {
	// init and execute the shim
	shim.Run("io.containerd.test-shim.v1", readonlyrootfsShim.New)
}
