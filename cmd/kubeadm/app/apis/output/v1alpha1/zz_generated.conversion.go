// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	output "k8s.io/kubernetes/cmd/kubeadm/app/apis/output"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BootstrapToken)(nil), (*output.BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BootstrapToken_To_output_BootstrapToken(a.(*BootstrapToken), b.(*output.BootstrapToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*output.BootstrapToken)(nil), (*BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_output_BootstrapToken_To_v1alpha1_BootstrapToken(a.(*output.BootstrapToken), b.(*BootstrapToken), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_BootstrapToken_To_output_BootstrapToken(in *BootstrapToken, out *output.BootstrapToken, s conversion.Scope) error {
	out.BootstrapToken = in.BootstrapToken
	return nil
}

// Convert_v1alpha1_BootstrapToken_To_output_BootstrapToken is an autogenerated conversion function.
func Convert_v1alpha1_BootstrapToken_To_output_BootstrapToken(in *BootstrapToken, out *output.BootstrapToken, s conversion.Scope) error {
	return autoConvert_v1alpha1_BootstrapToken_To_output_BootstrapToken(in, out, s)
}

func autoConvert_output_BootstrapToken_To_v1alpha1_BootstrapToken(in *output.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	out.BootstrapToken = in.BootstrapToken
	return nil
}

// Convert_output_BootstrapToken_To_v1alpha1_BootstrapToken is an autogenerated conversion function.
func Convert_output_BootstrapToken_To_v1alpha1_BootstrapToken(in *output.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	return autoConvert_output_BootstrapToken_To_v1alpha1_BootstrapToken(in, out, s)
}
