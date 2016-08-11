package keyvault

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/uuid"
	"net/http"
)

// KeyPermissions enumerates the values for key permissions.
type KeyPermissions string

const (
	// All specifies the all state for key permissions.
	All KeyPermissions = "all"
	// Backup specifies the backup state for key permissions.
	Backup KeyPermissions = "backup"
	// Create specifies the create state for key permissions.
	Create KeyPermissions = "create"
	// Decrypt specifies the decrypt state for key permissions.
	Decrypt KeyPermissions = "decrypt"
	// Delete specifies the delete state for key permissions.
	Delete KeyPermissions = "delete"
	// Encrypt specifies the encrypt state for key permissions.
	Encrypt KeyPermissions = "encrypt"
	// Get specifies the get state for key permissions.
	Get KeyPermissions = "get"
	// Import specifies the import state for key permissions.
	Import KeyPermissions = "import"
	// List specifies the list state for key permissions.
	List KeyPermissions = "list"
	// Restore specifies the restore state for key permissions.
	Restore KeyPermissions = "restore"
	// Sign specifies the sign state for key permissions.
	Sign KeyPermissions = "sign"
	// Unwrapkey specifies the unwrapkey state for key permissions.
	Unwrapkey KeyPermissions = "unwrapkey"
	// Update specifies the update state for key permissions.
	Update KeyPermissions = "update"
	// Verify specifies the verify state for key permissions.
	Verify KeyPermissions = "verify"
	// Wrapkey specifies the wrapkey state for key permissions.
	Wrapkey KeyPermissions = "wrapkey"
)

// SecretPermissions enumerates the values for secret permissions.
type SecretPermissions string

const (
	// SecretPermissionsAll specifies the secret permissions all state for
	// secret permissions.
	SecretPermissionsAll SecretPermissions = "all"
	// SecretPermissionsDelete specifies the secret permissions delete state
	// for secret permissions.
	SecretPermissionsDelete SecretPermissions = "delete"
	// SecretPermissionsGet specifies the secret permissions get state for
	// secret permissions.
	SecretPermissionsGet SecretPermissions = "get"
	// SecretPermissionsList specifies the secret permissions list state for
	// secret permissions.
	SecretPermissionsList SecretPermissions = "list"
	// SecretPermissionsSet specifies the secret permissions set state for
	// secret permissions.
	SecretPermissionsSet SecretPermissions = "set"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Premium specifies the premium state for sku name.
	Premium SkuName = "premium"
	// Standard specifies the standard state for sku name.
	Standard SkuName = "standard"
)

// AccessPolicyEntry is an array of 0 to 16 identities that have access to the
// key vault. All identities in the array must use the same tenant ID as the
// key vault's tenant ID.
type AccessPolicyEntry struct {
	TenantID      *uuid.UUID   `json:"tenantId,omitempty"`
	ObjectID      *uuid.UUID   `json:"objectId,omitempty"`
	ApplicationID *uuid.UUID   `json:"applicationId,omitempty"`
	Permissions   *Permissions `json:"permissions,omitempty"`
}

// Permissions is permissions the identity has for keys and secrets
type Permissions struct {
	Keys    *[]KeyPermissions    `json:"keys,omitempty"`
	Secrets *[]SecretPermissions `json:"secrets,omitempty"`
}

// Resource is key Vault resource
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Sku is sKU details
type Sku struct {
	Family *string `json:"family,omitempty"`
	Name   SkuName `json:"name,omitempty"`
}

// Vault is resource information with extended details.
type Vault struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Properties        *VaultProperties    `json:"properties,omitempty"`
}

// VaultCreateOrUpdateParameters is parameters for creating or updating a vault
type VaultCreateOrUpdateParameters struct {
	Location   *string             `json:"location,omitempty"`
	Tags       *map[string]*string `json:"tags,omitempty"`
	Properties *VaultProperties    `json:"properties,omitempty"`
}

// VaultListResult is list of vaults
type VaultListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Vault `json:"value,omitempty"`
	NextLink          *string  `json:"nextLink,omitempty"`
}

// VaultListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client VaultListResult) VaultListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// VaultProperties is properties of the vault
type VaultProperties struct {
	VaultURI                     *string              `json:"vaultUri,omitempty"`
	TenantID                     *uuid.UUID           `json:"tenantId,omitempty"`
	Sku                          *Sku                 `json:"sku,omitempty"`
	AccessPolicies               *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`
	EnabledForDeployment         *bool                `json:"enabledForDeployment,omitempty"`
	EnabledForDiskEncryption     *bool                `json:"enabledForDiskEncryption,omitempty"`
	EnabledForTemplateDeployment *bool                `json:"enabledForTemplateDeployment,omitempty"`
}
