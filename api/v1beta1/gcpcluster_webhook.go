/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta1

import (
	"context"
	"fmt"
	"reflect"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// clusterlog is for logging in this package.
var clusterlog = logf.Log.WithName("gcpcluster-resource")

// SetupWebhookWithManager sets up and registers the webhook with the manager.
func (c *GCPCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(c).
		WithDefaulter(c). // registers webhook.CustomDefaulter
		WithValidator(c). // registers webhook.CustomValidator
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-gcpcluster,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=gcpclusters,versions=v1beta1,name=validation.gcpcluster.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-gcpcluster,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=gcpclusters,versions=v1beta1,name=default.gcpcluster.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var (
	_ webhook.CustomValidator = &GCPCluster{}
	_ webhook.CustomDefaulter = &GCPCluster{}
)

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (c *GCPCluster) Default(ctx context.Context, obj runtime.Object) error {
	c, ok := obj.(*GCPCluster)
	if !ok {
		return fmt.Errorf("expected *GCPCluster, got %T", obj)
	}
	clusterlog.Info("default", "name", c.Name)
	return nil
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (c *GCPCluster) ValidateCreate(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	c, ok := obj.(*GCPCluster)
	if !ok {
		return nil, fmt.Errorf("expected *GCPCluster, got %T", obj)
	}
	clusterlog.Info("validate create", "name", c.Name)

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (c *GCPCluster) ValidateUpdate(ctx context.Context, oldRaw runtime.Object, newRaw runtime.Object) (warnings admission.Warnings, err error) {
	c, ok := newRaw.(*GCPCluster)
	if !ok {
		return nil, fmt.Errorf("expected *GCPCluster, got %T", newRaw)
	}
	clusterlog.Info("validate update", "name", c.Name)
	var allErrs field.ErrorList
	old := oldRaw.(*GCPCluster)

	if !reflect.DeepEqual(c.Spec.Project, old.Spec.Project) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Project"),
				c.Spec.Project, "field is immutable"),
		)
	}

	if !reflect.DeepEqual(c.Spec.Region, old.Spec.Region) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Region"),
				c.Spec.Region, "field is immutable"),
		)
	}

	if !reflect.DeepEqual(c.Spec.CredentialsRef, old.Spec.CredentialsRef) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "CredentialsRef"),
				c.Spec.CredentialsRef, "field is immutable"),
		)
	}

	if !reflect.DeepEqual(c.Spec.LoadBalancer, old.Spec.LoadBalancer) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "LoadBalancer"),
				c.Spec.LoadBalancer, "field is immutable"),
		)
	}

	if c.Spec.Network.Mtu < int64(1300) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Network", "Mtu"),
				c.Spec.Network.Mtu, "field cannot be lesser than 1300"),
		)
	}

	if c.Spec.Network.Mtu > int64(8896) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "Network", "Mtu"),
				c.Spec.Network.Mtu, "field cannot be greater than 8896"),
		)
	}

	if len(allErrs) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(GroupVersion.WithKind("GCPCluster").GroupKind(), c.Name, allErrs)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (c *GCPCluster) ValidateDelete(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	clusterlog.Info("validate delete", "name", c.Name)

	return nil, nil
}
