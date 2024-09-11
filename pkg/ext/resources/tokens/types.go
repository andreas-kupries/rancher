package tokens

import (
	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Token is the new extension Token structure
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the spec of RancherToken
	Spec TokenSpec `json:"spec"`
	// +optional
	Status TokenStatus `json:"status"`
}

// TokenSpec contains the user-specifiable parts of the Token
type TokenSpec struct {
	// UserID is the user id
	UserID string `json:"userID"`
	// Human readable description
	// +optional
	Description string `json:"description, omitempty"`
	// ClusterName is the cluster that the token is scoped to. If empty, the token
	// can be used for all clusters.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`
	// TTL is the time-to-live of the token, in milliseconds
	TTL int64 `json:"ttl"`
	// Enabled indicates an active token
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// ??
	IsDerived bool `json:"isDerived"`
}

// TokenStatus contains the data derived from the specification or otherwise generated.
type TokenStatus struct {
	// TokenValue is the access key. Shown only on token creation. Not saved.
	TokenValue string `json:"tokenValue,omitempty"`
	// TokenHash is the hash of the value. Only thing saved.
	TokenHash string `json:"tokenHash,omitempty"`
	// Time derived data
	// Expired flag, derived from creation time and time-to-live
	Expired bool `json:"expired"`
	// ExpiredAt is creation time + time-to-live
	ExpiredAt string `json:"expiredAt"`
	// User derived data
	// AuthProvider names the auth provider managing the user
	AuthProvider string `json:"authProvider"`
	// UserPrincipal holds the detailed user data
	UserPrincipal apiv3.Principal `json:"userPrincipal"`
	// GroupPrincipals holds detailed group information
	GroupPrincipals []apiv3.Principal `json:"groupsPrincipals"`
	// ProviderInfo provides provider-specific details
	ProviderInfo map[string]string `json:"providerInfo"`
	// Time of last change to the token
	LastUpdateTime string `json:"lastUpdateTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TokenList is the standard structure for a list of Tokens in the kube API
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Token `json:"items"`
}

// Implement the TokenAccessor interface

func (t *Token) GetName() string {
	return t.ObjectMeta.Name
}

func (t *Token) IsEnabled() bool {
	return t.Spec.Enabled
}

func (t *Token) GetUserID() string {
	return t.Spec.UserID
}

func (t *Token) ObjClusterName() string {
	return t.Spec.ClusterName
}

func (t *Token) GetAuthProvider() string {
	return t.Status.AuthProvider
}

func (t *Token) GetUserPrincipal() apiv3.Principal {
	return t.Status.UserPrincipal
}

func (t *Token) GetGroupPrincipals() []apiv3.Principal {
	return t.Status.GroupPrincipals
}

func (t *Token) GetProviderInfo() map[string]string {
	return t.Status.ProviderInfo
}
