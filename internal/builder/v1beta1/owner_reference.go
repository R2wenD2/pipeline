/*
Copyright 2019 The Tekton Authors

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

package builder

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OwnerReferenceOp is an operation which modifies an OwnerReference struct.
type OwnerReferenceOp func(*metav1.OwnerReference)

// OwnerReferenceAPIVersion sets the APIVersion to the OwnerReference.
func OwnerReferenceAPIVersion(version string) OwnerReferenceOp {
	return func(o *metav1.OwnerReference) {
		o.APIVersion = version
	}
}

// Controller sets the Controller to the OwnerReference.y
func Controller(o *metav1.OwnerReference) {
	o.Controller = &trueB
}

// BlockOwnerDeletion sets the BlockOwnerDeletion to the OwnerReference.
func BlockOwnerDeletion(o *metav1.OwnerReference) {
	o.BlockOwnerDeletion = &trueB
}
