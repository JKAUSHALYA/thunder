/*
 * Copyright (c) 2025, WSO2 LLC. (http://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Package model defines the data structures for the application module.
package model

import "github.com/asgardeo/thunder/internal/application/constants"

// ApplicationDTO represents the data transfer object for application service operations.
type ApplicationDTO struct {
	ID                        string
	Name                      string
	Description               string
	AuthFlowGraphID           string
	RegistrationFlowGraphID   string
	IsRegistrationFlowEnabled bool

	URL     string
	LogoURL string

	Certificate       *Certificate
	InboundAuthConfig []InboundAuthConfig
}

// BasicApplicationDTO represents a simplified data transfer object for application service operations.
type BasicApplicationDTO struct {
	ID                        string
	Name                      string
	Description               string
	AuthFlowGraphID           string
	RegistrationFlowGraphID   string
	IsRegistrationFlowEnabled bool
	ClientID                  string
}

// ApplicationProcessedDTO represents the processed data transfer object for application service operations.
type ApplicationProcessedDTO struct {
	ID                        string
	Name                      string
	Description               string
	AuthFlowGraphID           string
	RegistrationFlowGraphID   string
	IsRegistrationFlowEnabled bool

	URL     string
	LogoURL string

	Certificate       *Certificate
	InboundAuthConfig []InboundAuthConfigProcessed
}

// Certificate represents a certificate used in the application.
// TODO: Move Certificate to a separate package.
type Certificate struct {
	ID    string                    `json:"-"`
	Type  constants.CertificateType `json:"type"`
	Value string                    `json:"value"`
}

// InboundAuthConfig represents the inbound authentication configuration for an application.
// TODO: Need to refactor when supporting other/multiple inbound auth types.
type InboundAuthConfig struct {
	Type           constants.InboundAuthType
	OAuthAppConfig *OAuthAppConfig
}

// InboundAuthConfigProcessed represents the processed inbound authentication configuration for an application.
type InboundAuthConfigProcessed struct {
	Type           constants.InboundAuthType
	OAuthAppConfig *OAuthAppConfigProcessed
}

// ApplicationRequest represents the request structure for creating or updating an application.
type ApplicationRequest struct {
	Name                      string       `json:"name"`
	Description               string       `json:"description"`
	ClientID                  string       `json:"client_id"`
	ClientSecret              string       `json:"client_secret"`
	RedirectURIs              []string     `json:"redirect_uris"`
	GrantTypes                []string     `json:"grant_types"`
	AuthFlowGraphID           string       `json:"auth_flow_graph_id,omitempty"`
	RegistrationFlowGraphID   string       `json:"registration_flow_graph_id,omitempty"`
	IsRegistrationFlowEnabled bool         `json:"is_registration_flow_enabled,omitempty"`
	URL                       string       `json:"url,omitempty"`
	LogoURL                   string       `json:"logo_url,omitempty"`
	Certificate               *Certificate `json:"certificate,omitempty"`
}

// ApplicationCompleteResponse represents the complete response structure for an application.
type ApplicationCompleteResponse struct {
	ID                        string       `json:"id,omitempty"`
	Name                      string       `json:"name"`
	Description               string       `json:"description,omitempty"`
	ClientID                  string       `json:"client_id"`
	ClientSecret              string       `json:"client_secret"`
	RedirectURIs              []string     `json:"redirect_uris"`
	GrantTypes                []string     `json:"grant_types"`
	AuthFlowGraphID           string       `json:"auth_flow_graph_id,omitempty"`
	RegistrationFlowGraphID   string       `json:"registration_flow_graph_id,omitempty"`
	IsRegistrationFlowEnabled bool         `json:"is_registration_flow_enabled,omitempty"`
	URL                       string       `json:"url,omitempty"`
	LogoURL                   string       `json:"logo_url,omitempty"`
	Certificate               *Certificate `json:"certificate,omitempty"`
}

// ApplicationGetResponse represents the response structure for getting an application.
type ApplicationGetResponse struct {
	ID                        string       `json:"id,omitempty"`
	Name                      string       `json:"name"`
	Description               string       `json:"description,omitempty"`
	ClientID                  string       `json:"client_id"`
	RedirectURIs              []string     `json:"redirect_uris"`
	GrantTypes                []string     `json:"grant_types"`
	AuthFlowGraphID           string       `json:"auth_flow_graph_id,omitempty"`
	RegistrationFlowGraphID   string       `json:"registration_flow_graph_id,omitempty"`
	IsRegistrationFlowEnabled bool         `json:"is_registration_flow_enabled,omitempty"`
	URL                       string       `json:"url,omitempty"`
	LogoURL                   string       `json:"logo_url,omitempty"`
	Certificate               *Certificate `json:"certificate,omitempty"`
}

// BasicApplicationResponse represents a simplified response structure for an application.
type BasicApplicationResponse struct {
	ID                        string `json:"id,omitempty"`
	Name                      string `json:"name"`
	Description               string `json:"description,omitempty"`
	ClientID                  string `json:"client_id"`
	AuthFlowGraphID           string `json:"auth_flow_graph_id,omitempty"`
	RegistrationFlowGraphID   string `json:"registration_flow_graph_id,omitempty"`
	IsRegistrationFlowEnabled bool   `json:"is_registration_flow_enabled,omitempty"`
}
