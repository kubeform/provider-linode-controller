/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package provider

type LinodeSpec struct {
	// An HTTP User-Agent Prefix to prepend in API requests.
	// +optional
	ApiVersion *string `json:"apiVersion,omitempty" tf:"api_version"`
	// The rate in milliseconds to poll for events.
	// +optional
	EventPollMs *int64 `json:"eventPollMs,omitempty" tf:"event_poll_ms"`
	// The rate in milliseconds to poll for LKE events.
	// +optional
	LkeEventPollMs *int64 `json:"lkeEventPollMs,omitempty" tf:"lke_event_poll_ms"`
	// The rate in milliseconds to poll for an LKE node to be ready.
	// +optional
	LkeNodeReadyPollMs *int64 `json:"lkeNodeReadyPollMs,omitempty" tf:"lke_node_ready_poll_ms"`
	// Maximum delay in milliseconds before retrying a request.
	// +optional
	MaxRetryDelayMs *int64 `json:"maxRetryDelayMs,omitempty" tf:"max_retry_delay_ms"`
	// Minimum delay in milliseconds before retrying a request.
	// +optional
	MinRetryDelayMs *int64 `json:"minRetryDelayMs,omitempty" tf:"min_retry_delay_ms"`
	// Skip waiting for a linode_instance resource to finish deleting.
	// +optional
	SkipInstanceDeletePoll *bool `json:"skipInstanceDeletePoll,omitempty" tf:"skip_instance_delete_poll"`
	// Skip waiting for a linode_instance resource to be running.
	// +optional
	SkipInstanceReadyPoll *bool `json:"skipInstanceReadyPoll,omitempty" tf:"skip_instance_ready_poll"`
	// The token that allows you access to your Linode account
	Token *string `json:"token" tf:"token"`
	// An HTTP User-Agent Prefix to prepend in API requests.
	// +optional
	UaPrefix *string `json:"uaPrefix,omitempty" tf:"ua_prefix"`
	// The HTTP(S) API address of the Linode API to use.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}
