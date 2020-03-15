// +build !ignore_autogenerated

/*
Copyright The Flagger Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1alpha3 "github.com/weaveworks/flagger/pkg/apis/istio/v1alpha3"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertProvider) DeepCopyInto(out *AlertProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertProvider.
func (in *AlertProvider) DeepCopy() *AlertProvider {
	if in == nil {
		return nil
	}
	out := new(AlertProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertProviderCondition) DeepCopyInto(out *AlertProviderCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertProviderCondition.
func (in *AlertProviderCondition) DeepCopy() *AlertProviderCondition {
	if in == nil {
		return nil
	}
	out := new(AlertProviderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertProviderList) DeepCopyInto(out *AlertProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertProviderList.
func (in *AlertProviderList) DeepCopy() *AlertProviderList {
	if in == nil {
		return nil
	}
	out := new(AlertProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertProviderSpec) DeepCopyInto(out *AlertProviderSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertProviderSpec.
func (in *AlertProviderSpec) DeepCopy() *AlertProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AlertProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertProviderStatus) DeepCopyInto(out *AlertProviderStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AlertProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertProviderStatus.
func (in *AlertProviderStatus) DeepCopy() *AlertProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AlertProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Canary) DeepCopyInto(out *Canary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Canary.
func (in *Canary) DeepCopy() *Canary {
	if in == nil {
		return nil
	}
	out := new(Canary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Canary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryAlert) DeepCopyInto(out *CanaryAlert) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryAlert.
func (in *CanaryAlert) DeepCopy() *CanaryAlert {
	if in == nil {
		return nil
	}
	out := new(CanaryAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryAnalysis) DeepCopyInto(out *CanaryAnalysis) {
	*out = *in
	if in.Alerts != nil {
		in, out := &in.Alerts, &out.Alerts
		*out = make([]CanaryAlert, len(*in))
		copy(*out, *in)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]CanaryMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]CanaryWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]v1alpha3.HTTPMatchRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryAnalysis.
func (in *CanaryAnalysis) DeepCopy() *CanaryAnalysis {
	if in == nil {
		return nil
	}
	out := new(CanaryAnalysis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryCondition) DeepCopyInto(out *CanaryCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryCondition.
func (in *CanaryCondition) DeepCopy() *CanaryCondition {
	if in == nil {
		return nil
	}
	out := new(CanaryCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryList) DeepCopyInto(out *CanaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Canary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryList.
func (in *CanaryList) DeepCopy() *CanaryList {
	if in == nil {
		return nil
	}
	out := new(CanaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryMetric) DeepCopyInto(out *CanaryMetric) {
	*out = *in
	if in.ThresholdRange != nil {
		in, out := &in.ThresholdRange, &out.ThresholdRange
		*out = new(CanaryThresholdRange)
		(*in).DeepCopyInto(*out)
	}
	if in.TemplateRef != nil {
		in, out := &in.TemplateRef, &out.TemplateRef
		*out = new(CrossNamespaceObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryMetric.
func (in *CanaryMetric) DeepCopy() *CanaryMetric {
	if in == nil {
		return nil
	}
	out := new(CanaryMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryService) DeepCopyInto(out *CanaryService) {
	*out = *in
	out.TargetPort = in.TargetPort
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TrafficPolicy != nil {
		in, out := &in.TrafficPolicy, &out.TrafficPolicy
		*out = new(v1alpha3.TrafficPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]v1alpha3.HTTPMatchRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rewrite != nil {
		in, out := &in.Rewrite, &out.Rewrite
		*out = new(v1alpha3.HTTPRewrite)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(v1alpha3.HTTPRetry)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(v1alpha3.Headers)
		(*in).DeepCopyInto(*out)
	}
	if in.CorsPolicy != nil {
		in, out := &in.CorsPolicy, &out.CorsPolicy
		*out = new(v1alpha3.CorsPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryService.
func (in *CanaryService) DeepCopy() *CanaryService {
	if in == nil {
		return nil
	}
	out := new(CanaryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpec) DeepCopyInto(out *CanarySpec) {
	*out = *in
	out.TargetRef = in.TargetRef
	if in.AutoscalerRef != nil {
		in, out := &in.AutoscalerRef, &out.AutoscalerRef
		*out = new(CrossNamespaceObjectReference)
		**out = **in
	}
	if in.IngressRef != nil {
		in, out := &in.IngressRef, &out.IngressRef
		*out = new(CrossNamespaceObjectReference)
		**out = **in
	}
	in.Service.DeepCopyInto(&out.Service)
	if in.Analysis != nil {
		in, out := &in.Analysis, &out.Analysis
		*out = new(CanaryAnalysis)
		(*in).DeepCopyInto(*out)
	}
	if in.CanaryAnalysis != nil {
		in, out := &in.CanaryAnalysis, &out.CanaryAnalysis
		*out = new(CanaryAnalysis)
		(*in).DeepCopyInto(*out)
	}
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpec.
func (in *CanarySpec) DeepCopy() *CanarySpec {
	if in == nil {
		return nil
	}
	out := new(CanarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStatus) DeepCopyInto(out *CanaryStatus) {
	*out = *in
	if in.TrackedConfigs != nil {
		in, out := &in.TrackedConfigs, &out.TrackedConfigs
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CanaryCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStatus.
func (in *CanaryStatus) DeepCopy() *CanaryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryThresholdRange) DeepCopyInto(out *CanaryThresholdRange) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryThresholdRange.
func (in *CanaryThresholdRange) DeepCopy() *CanaryThresholdRange {
	if in == nil {
		return nil
	}
	out := new(CanaryThresholdRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryWebhook) DeepCopyInto(out *CanaryWebhook) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryWebhook.
func (in *CanaryWebhook) DeepCopy() *CanaryWebhook {
	if in == nil {
		return nil
	}
	out := new(CanaryWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryWebhookPayload) DeepCopyInto(out *CanaryWebhookPayload) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryWebhookPayload.
func (in *CanaryWebhookPayload) DeepCopy() *CanaryWebhookPayload {
	if in == nil {
		return nil
	}
	out := new(CanaryWebhookPayload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossNamespaceObjectReference) DeepCopyInto(out *CrossNamespaceObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossNamespaceObjectReference.
func (in *CrossNamespaceObjectReference) DeepCopy() *CrossNamespaceObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossNamespaceObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTemplate) DeepCopyInto(out *MetricTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTemplate.
func (in *MetricTemplate) DeepCopy() *MetricTemplate {
	if in == nil {
		return nil
	}
	out := new(MetricTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTemplateCondition) DeepCopyInto(out *MetricTemplateCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTemplateCondition.
func (in *MetricTemplateCondition) DeepCopy() *MetricTemplateCondition {
	if in == nil {
		return nil
	}
	out := new(MetricTemplateCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTemplateList) DeepCopyInto(out *MetricTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTemplateList.
func (in *MetricTemplateList) DeepCopy() *MetricTemplateList {
	if in == nil {
		return nil
	}
	out := new(MetricTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTemplateModel) DeepCopyInto(out *MetricTemplateModel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTemplateModel.
func (in *MetricTemplateModel) DeepCopy() *MetricTemplateModel {
	if in == nil {
		return nil
	}
	out := new(MetricTemplateModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTemplateProvider) DeepCopyInto(out *MetricTemplateProvider) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTemplateProvider.
func (in *MetricTemplateProvider) DeepCopy() *MetricTemplateProvider {
	if in == nil {
		return nil
	}
	out := new(MetricTemplateProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTemplateSpec) DeepCopyInto(out *MetricTemplateSpec) {
	*out = *in
	in.Provider.DeepCopyInto(&out.Provider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTemplateSpec.
func (in *MetricTemplateSpec) DeepCopy() *MetricTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(MetricTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTemplateStatus) DeepCopyInto(out *MetricTemplateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MetricTemplateCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTemplateStatus.
func (in *MetricTemplateStatus) DeepCopy() *MetricTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(MetricTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
