package notificationhubs

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen specifies the listen state for access rights.
	Listen AccessRights = "Listen"
	// Manage specifies the manage state for access rights.
	Manage AccessRights = "Manage"
	// Send specifies the send state for access rights.
	Send AccessRights = "Send"
)

// NamespaceType enumerates the values for namespace type.
type NamespaceType string

const (
	// Messaging specifies the messaging state for namespace type.
	Messaging NamespaceType = "Messaging"
	// NotificationHub specifies the notification hub state for namespace type.
	NotificationHub NamespaceType = "NotificationHub"
)

// AdmCredential is description of a NotificationHub AdmCredential.
type AdmCredential struct {
	Properties *AdmCredentialProperties `json:"properties,omitempty"`
}

// AdmCredentialProperties is description of a NotificationHub AdmCredential.
type AdmCredentialProperties struct {
	ClientID     *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	AuthTokenURL *string `json:"authTokenUrl,omitempty"`
}

// ApnsCredential is description of a NotificationHub ApnsCredential.
type ApnsCredential struct {
	Properties *ApnsCredentialProperties `json:"properties,omitempty"`
}

// ApnsCredentialProperties is description of a NotificationHub ApnsCredential.
type ApnsCredentialProperties struct {
	ApnsCertificate *string `json:"apnsCertificate,omitempty"`
	CertificateKey  *string `json:"certificateKey,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty"`
	Thumbprint      *string `json:"thumbprint,omitempty"`
}

// BaiduCredential is description of a NotificationHub BaiduCredential.
type BaiduCredential struct {
	Properties *BaiduCredentialProperties `json:"properties,omitempty"`
}

// BaiduCredentialProperties is description of a NotificationHub
// BaiduCredential.
type BaiduCredentialProperties struct {
	BaiduAPIKey    *string `json:"baiduApiKey,omitempty"`
	BaiduEndPoint  *string `json:"baiduEndPoint,omitempty"`
	BaiduSecretKey *string `json:"baiduSecretKey,omitempty"`
}

// CheckAvailabilityParameters is parameters supplied to the Check Name
// Availability for Namespace and NotificationHubs.
type CheckAvailabilityParameters struct {
	ID           *string             `json:"id,omitempty"`
	Name         *string             `json:"name,omitempty"`
	Type         *string             `json:"type,omitempty"`
	Location     *string             `json:"location,omitempty"`
	Tags         *map[string]*string `json:"tags,omitempty"`
	IsAvailiable *bool               `json:"isAvailiable,omitempty"`
}

// CheckAvailabilityResult is description of a CheckAvailibility resource.
type CheckAvailabilityResult struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	IsAvailiable      *bool               `json:"isAvailiable,omitempty"`
}

// GcmCredential is description of a NotificationHub GcmCredential.
type GcmCredential struct {
	Properties *GcmCredentialProperties `json:"properties,omitempty"`
}

// GcmCredentialProperties is description of a NotificationHub GcmCredential.
type GcmCredentialProperties struct {
	GcmEndpoint  *string `json:"gcmEndpoint,omitempty"`
	GoogleAPIKey *string `json:"googleApiKey,omitempty"`
}

// MpnsCredential is description of a NotificationHub MpnsCredential.
type MpnsCredential struct {
	Properties *MpnsCredentialProperties `json:"properties,omitempty"`
}

// MpnsCredentialProperties is description of a NotificationHub MpnsCredential.
type MpnsCredentialProperties struct {
	MpnsCertificate *string `json:"mpnsCertificate,omitempty"`
	CertificateKey  *string `json:"certificateKey,omitempty"`
	Thumbprint      *string `json:"thumbprint,omitempty"`
}

// NamespaceCreateOrUpdateParameters is parameters supplied to the
// CreateOrUpdate Namespace operation.
type NamespaceCreateOrUpdateParameters struct {
	ID         *string              `json:"id,omitempty"`
	Name       *string              `json:"name,omitempty"`
	Type       *string              `json:"type,omitempty"`
	Location   *string              `json:"location,omitempty"`
	Tags       *map[string]*string  `json:"tags,omitempty"`
	Properties *NamespaceProperties `json:"properties,omitempty"`
}

// NamespaceListResult is the response of the List Namespace operation.
type NamespaceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]NamespaceResource `json:"value,omitempty"`
	NextLink          *string              `json:"nextLink,omitempty"`
}

// NamespaceListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client NamespaceListResult) NamespaceListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// NamespaceProperties is namespace properties.
type NamespaceProperties struct {
	Name               *string       `json:"name,omitempty"`
	ProvisioningState  *string       `json:"provisioningState,omitempty"`
	Region             *string       `json:"region,omitempty"`
	Status             *string       `json:"status,omitempty"`
	CreatedAt          *date.Time    `json:"createdAt,omitempty"`
	ServiceBusEndpoint *string       `json:"serviceBusEndpoint,omitempty"`
	SubscriptionID     *string       `json:"subscriptionId,omitempty"`
	ScaleUnit          *string       `json:"scaleUnit,omitempty"`
	Enabled            *bool         `json:"enabled,omitempty"`
	Critical           *bool         `json:"critical,omitempty"`
	NamespaceType      NamespaceType `json:"namespaceType,omitempty"`
}

// NamespaceResource is description of a Namespace resource.
type NamespaceResource struct {
	autorest.Response `json:"-"`
	ID                *string              `json:"id,omitempty"`
	Name              *string              `json:"name,omitempty"`
	Type              *string              `json:"type,omitempty"`
	Location          *string              `json:"location,omitempty"`
	Tags              *map[string]*string  `json:"tags,omitempty"`
	Properties        *NamespaceProperties `json:"properties,omitempty"`
}

// NotificationHubCreateOrUpdateParameters is parameters supplied to the
// CreateOrUpdate NotificationHub operation.
type NotificationHubCreateOrUpdateParameters struct {
	ID         *string                    `json:"id,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Type       *string                    `json:"type,omitempty"`
	Location   *string                    `json:"location,omitempty"`
	Tags       *map[string]*string        `json:"tags,omitempty"`
	Properties *NotificationHubProperties `json:"properties,omitempty"`
}

// NotificationHubListResult is the response of the List NotificationHub
// operation.
type NotificationHubListResult struct {
	autorest.Response `json:"-"`
	Value             *[]NotificationHubResource `json:"value,omitempty"`
	NextLink          *string                    `json:"nextLink,omitempty"`
}

// NotificationHubListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client NotificationHubListResult) NotificationHubListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// NotificationHubProperties is notificationHub properties.
type NotificationHubProperties struct {
	Name               *string                                    `json:"name,omitempty"`
	RegistrationTTL    *string                                    `json:"registrationTtl,omitempty"`
	AuthorizationRules *[]SharedAccessAuthorizationRuleProperties `json:"authorizationRules,omitempty"`
	ApnsCredential     *ApnsCredential                            `json:"apnsCredential,omitempty"`
	WnsCredential      *WnsCredential                             `json:"wnsCredential,omitempty"`
	GcmCredential      *GcmCredential                             `json:"gcmCredential,omitempty"`
	MpnsCredential     *MpnsCredential                            `json:"mpnsCredential,omitempty"`
	AdmCredential      *AdmCredential                             `json:"admCredential,omitempty"`
	BaiduCredential    *BaiduCredential                           `json:"baiduCredential,omitempty"`
}

// NotificationHubResource is description of a NotificationHub Resource.
type NotificationHubResource struct {
	autorest.Response `json:"-"`
	ID                *string                    `json:"id,omitempty"`
	Name              *string                    `json:"name,omitempty"`
	Type              *string                    `json:"type,omitempty"`
	Location          *string                    `json:"location,omitempty"`
	Tags              *map[string]*string        `json:"tags,omitempty"`
	Properties        *NotificationHubProperties `json:"properties,omitempty"`
}

// PnsCredentialsProperties is description of a NotificationHub PNS
// Credentials.
type PnsCredentialsProperties struct {
	ApnsCredential  *ApnsCredential  `json:"apnsCredential,omitempty"`
	WnsCredential   *WnsCredential   `json:"wnsCredential,omitempty"`
	GcmCredential   *GcmCredential   `json:"gcmCredential,omitempty"`
	MpnsCredential  *MpnsCredential  `json:"mpnsCredential,omitempty"`
	AdmCredential   *AdmCredential   `json:"admCredential,omitempty"`
	BaiduCredential *BaiduCredential `json:"baiduCredential,omitempty"`
}

// PnsCredentialsResource is description of a NotificationHub PNS Credentials.
type PnsCredentialsResource struct {
	autorest.Response `json:"-"`
	ID                *string                   `json:"id,omitempty"`
	Name              *string                   `json:"name,omitempty"`
	Type              *string                   `json:"type,omitempty"`
	Location          *string                   `json:"location,omitempty"`
	Tags              *map[string]*string       `json:"tags,omitempty"`
	Properties        *PnsCredentialsProperties `json:"properties,omitempty"`
}

// PolicykeyResource is namespace/NotificationHub Regenerate Keys
type PolicykeyResource struct {
	PolicyKey *string `json:"policyKey,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResourceListKeys is namespace/NotificationHub Connection String
type ResourceListKeys struct {
	autorest.Response         `json:"-"`
	PrimaryConnectionString   *string `json:"primaryConnectionString,omitempty"`
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`
	PrimaryKey                *string `json:"primaryKey,omitempty"`
	SecondaryKey              *string `json:"secondaryKey,omitempty"`
	KeyName                   *string `json:"keyName,omitempty"`
}

// SharedAccessAuthorizationRuleCreateOrUpdateParameters is parameters
// supplied to the CreateOrUpdate Namespace AuthorizationRules.
type SharedAccessAuthorizationRuleCreateOrUpdateParameters struct {
	ID         *string                                  `json:"id,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	Type       *string                                  `json:"type,omitempty"`
	Location   *string                                  `json:"location,omitempty"`
	Tags       *map[string]*string                      `json:"tags,omitempty"`
	Properties *SharedAccessAuthorizationRuleProperties `json:"properties,omitempty"`
}

// SharedAccessAuthorizationRuleListResult is the response of the List
// Namespace operation.
type SharedAccessAuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SharedAccessAuthorizationRuleResource `json:"value,omitempty"`
	NextLink          *string                                  `json:"nextLink,omitempty"`
}

// SharedAccessAuthorizationRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SharedAccessAuthorizationRuleListResult) SharedAccessAuthorizationRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SharedAccessAuthorizationRuleProperties is sharedAccessAuthorizationRule
// properties.
type SharedAccessAuthorizationRuleProperties struct {
	Rights *[]AccessRights `json:"rights,omitempty"`
}

// SharedAccessAuthorizationRuleResource is description of a Namespace
// AuthorizationRules.
type SharedAccessAuthorizationRuleResource struct {
	autorest.Response `json:"-"`
	ID                *string                                  `json:"id,omitempty"`
	Name              *string                                  `json:"name,omitempty"`
	Type              *string                                  `json:"type,omitempty"`
	Location          *string                                  `json:"location,omitempty"`
	Tags              *map[string]*string                      `json:"tags,omitempty"`
	Properties        *SharedAccessAuthorizationRuleProperties `json:"properties,omitempty"`
}

// SubResource is
type SubResource struct {
	ID *string `json:"id,omitempty"`
}

// WnsCredential is description of a NotificationHub WnsCredential.
type WnsCredential struct {
	Properties *WnsCredentialProperties `json:"properties,omitempty"`
}

// WnsCredentialProperties is description of a NotificationHub WnsCredential.
type WnsCredentialProperties struct {
	PackageSid          *string `json:"packageSid,omitempty"`
	SecretKey           *string `json:"secretKey,omitempty"`
	WindowsLiveEndpoint *string `json:"windowsLiveEndpoint,omitempty"`
}
