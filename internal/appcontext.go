/*
Copyright © 2025

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

package runner

import (
	"log"

	sync "github.com/matryer/resync"
)

type AppContext struct {
	vmiName        string
	dataVolumeName string
}

//nolint:gochecknoglobals
var (
	instance *AppContext
	once     sync.Once
)

// NewAppContext creates the AppContext once with the provided values.
// Subsequent calls return the same instance, ignoring new values.
func NewAppContext(vmi, dataVolume string) *AppContext {
	once.Do(func() {
		log.Printf("Registering %s Virtual Machine Instance and %s Data Volume\n", vmi, dataVolume)

		instance = &AppContext{
			vmiName:        vmi,
			dataVolumeName: dataVolume,
		}
	})

	return instance
}

// GetAppContext returns the already initialized AppContext.
// Panics if called before NewAppContext.
func GetAppContext() *AppContext {
	if instance == nil {
		log.Fatal("AppContext not initialized. Call NewAppContext first.")
	}

	return instance
}

// CancelAppContext resets the AppContext to its initial state.
func CancelAppContext() {
	once.Reset()
}

// GetVMIName returns the Virtual Machine Instance Name created for the runner.
func (a *AppContext) GetVMIName() string {
	return a.vmiName
}

// GetDataVolumeName returns the Data Volume Name created for the runner.
func (a *AppContext) GetDataVolumeName() string {
	return a.dataVolumeName
}
