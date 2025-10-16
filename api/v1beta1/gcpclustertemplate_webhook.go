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
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var gcpclustertemplatelog = logf.Log.WithName("gcpclustertemplate-resource")

func (r *GCPClusterTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		WithDefaulter(r). // registers webhook.CustomDefaulter
		WithValidator(r). // registers webhook.CustomValidator
		Complete()
}

//+kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-gcpclustertemplate,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=gcpclustertemplates,versions=v1beta1,name=default.gcpclustertemplate.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1
//+kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-gcpclustertemplate,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=gcpclustertemplates,versions=v1beta1,name=validation.gcpclustertemplate.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var _ webhook.CustomDefaulter = &GCPClusterTemplate{}

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (r *GCPClusterTemplate) Default(ctx context.Context, obj runtime.Object) error {
	r, ok := obj.(*GCPClusterTemplate)
	if !ok {
		return fmt.Errorf("expected *GCPClusterTemplate, got %T", obj)
	}
	gcpclustertemplatelog.Info("default", "name", r.Name)
	return nil
}

var _ webhook.CustomValidator = &GCPClusterTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPClusterTemplate) ValidateCreate(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	r, ok := obj.(*GCPClusterTemplate)
	if !ok {
		return nil, fmt.Errorf("expected *GCPClusterTemplate, got %T", obj)
	}
	gcpclustertemplatelog.Info("validate create", "name", r.Name)

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPClusterTemplate) ValidateUpdate(ctx context.Context, oldRaw runtime.Object, newRaw runtime.Object) (warnings admission.Warnings, err error) {
	r, ok := newRaw.(*GCPClusterTemplate)
	if !ok {
		return nil, fmt.Errorf("expected *GCPClusterTemplate, got %T", newRaw)
	}
	old, ok := oldRaw.(*GCPClusterTemplate)
	if !ok {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("expected an GCPClusterTemplate but got a %T", oldRaw))
	}

	if !reflect.DeepEqual(r.Spec, old.Spec) {
		return nil, apierrors.NewBadRequest("GCPClusterTemplate.Spec is immutable")
	}
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPClusterTemplate) ValidateDelete(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	gcpclustertemplatelog.Info("validate delete", "name", r.Name)
	return nil, nil
}
