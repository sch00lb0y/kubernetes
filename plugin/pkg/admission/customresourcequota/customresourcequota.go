/*
Copyright 2014 The Kubernetes Authors.

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

package customresourcequota

import (
	"errors"
	"io"

	"k8s.io/apiserver/pkg/admission"
)

const PluginName = "CRResourceQuota"

func Register(plugins *admission.Plugins) {
	plugins.Register(PluginName,
		func(config io.Reader) (admission.Interface, error) {
			return NewCustomResourceQuotaAdmission(), nil
		})
}

type CustomResourceQuotaAdmission struct {
}

func NewCustomResourceQuotaAdmission() *CustomResourceQuotaAdmission {
	return &CustomResourceQuotaAdmission{}
}
func (a *CustomResourceQuotaAdmission) Handles(operation admission.Operation) bool {
	return true
}
func (a *CustomResourceQuotaAdmission) Validate(attr admission.Attributes) error {
	return errors.New("balaji error")
}
