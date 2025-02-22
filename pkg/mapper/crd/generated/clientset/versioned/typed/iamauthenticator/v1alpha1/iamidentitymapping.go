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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	iamauthenticatorv1alpha1 "sigs.k8s.io/aws-iam-authenticator/pkg/mapper/crd/apis/iamauthenticator/v1alpha1"
	scheme "sigs.k8s.io/aws-iam-authenticator/pkg/mapper/crd/generated/clientset/versioned/scheme"
)

// IAMIdentityMappingsGetter has a method to return a IAMIdentityMappingInterface.
// A group's client should implement this interface.
type IAMIdentityMappingsGetter interface {
	IAMIdentityMappings() IAMIdentityMappingInterface
}

// IAMIdentityMappingInterface has methods to work with IAMIdentityMapping resources.
type IAMIdentityMappingInterface interface {
	Create(ctx context.Context, iAMIdentityMapping *iamauthenticatorv1alpha1.IAMIdentityMapping, opts v1.CreateOptions) (*iamauthenticatorv1alpha1.IAMIdentityMapping, error)
	Update(ctx context.Context, iAMIdentityMapping *iamauthenticatorv1alpha1.IAMIdentityMapping, opts v1.UpdateOptions) (*iamauthenticatorv1alpha1.IAMIdentityMapping, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, iAMIdentityMapping *iamauthenticatorv1alpha1.IAMIdentityMapping, opts v1.UpdateOptions) (*iamauthenticatorv1alpha1.IAMIdentityMapping, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*iamauthenticatorv1alpha1.IAMIdentityMapping, error)
	List(ctx context.Context, opts v1.ListOptions) (*iamauthenticatorv1alpha1.IAMIdentityMappingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *iamauthenticatorv1alpha1.IAMIdentityMapping, err error)
	IAMIdentityMappingExpansion
}

// iAMIdentityMappings implements IAMIdentityMappingInterface
type iAMIdentityMappings struct {
	*gentype.ClientWithList[*iamauthenticatorv1alpha1.IAMIdentityMapping, *iamauthenticatorv1alpha1.IAMIdentityMappingList]
}

// newIAMIdentityMappings returns a IAMIdentityMappings
func newIAMIdentityMappings(c *IamauthenticatorV1alpha1Client) *iAMIdentityMappings {
	return &iAMIdentityMappings{
		gentype.NewClientWithList[*iamauthenticatorv1alpha1.IAMIdentityMapping, *iamauthenticatorv1alpha1.IAMIdentityMappingList](
			"iamidentitymappings",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *iamauthenticatorv1alpha1.IAMIdentityMapping {
				return &iamauthenticatorv1alpha1.IAMIdentityMapping{}
			},
			func() *iamauthenticatorv1alpha1.IAMIdentityMappingList {
				return &iamauthenticatorv1alpha1.IAMIdentityMappingList{}
			},
		),
	}
}
