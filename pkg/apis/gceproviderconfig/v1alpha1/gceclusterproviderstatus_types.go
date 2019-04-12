package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GCEClusterProviderStatus is the Schema for the gceclusterproviderstatus API
// +k8s:openapi-gen=true
type GCEClusterProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GCEClusterProviderStatusList contains a list of GCEClusterProviderStatus
type GCEClusterProviderStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCEClusterProviderSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GCEClusterProviderStatus{}, &GCEClusterProviderStatusList{})
}
