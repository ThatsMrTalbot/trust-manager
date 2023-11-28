/*
Copyright 2023 The cert-manager Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// BundleSourceApplyConfiguration represents an declarative configuration of the BundleSource type for use
// with apply.
type BundleSourceApplyConfiguration struct {
	ConfigMap     *SourceObjectKeySelectorApplyConfiguration `json:"configMap,omitempty"`
	Secret        *SourceObjectKeySelectorApplyConfiguration `json:"secret,omitempty"`
	InLine        *string                                    `json:"inLine,omitempty"`
	UseDefaultCAs *bool                                      `json:"useDefaultCAs,omitempty"`
}

// BundleSourceApplyConfiguration constructs an declarative configuration of the BundleSource type for use with
// apply.
func BundleSource() *BundleSourceApplyConfiguration {
	return &BundleSourceApplyConfiguration{}
}

// WithConfigMap sets the ConfigMap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMap field is set to the value of the last call.
func (b *BundleSourceApplyConfiguration) WithConfigMap(value *SourceObjectKeySelectorApplyConfiguration) *BundleSourceApplyConfiguration {
	b.ConfigMap = value
	return b
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *BundleSourceApplyConfiguration) WithSecret(value *SourceObjectKeySelectorApplyConfiguration) *BundleSourceApplyConfiguration {
	b.Secret = value
	return b
}

// WithInLine sets the InLine field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InLine field is set to the value of the last call.
func (b *BundleSourceApplyConfiguration) WithInLine(value string) *BundleSourceApplyConfiguration {
	b.InLine = &value
	return b
}

// WithUseDefaultCAs sets the UseDefaultCAs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UseDefaultCAs field is set to the value of the last call.
func (b *BundleSourceApplyConfiguration) WithUseDefaultCAs(value bool) *BundleSourceApplyConfiguration {
	b.UseDefaultCAs = &value
	return b
}