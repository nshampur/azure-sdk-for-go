package reservationsapi

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/reservations/mgmt/2019-07-19-preview/reservations"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetAppliedReservationList(ctx context.Context, subscriptionID string) (result reservations.AppliedReservations, err error)
	GetCatalog(ctx context.Context, subscriptionID string, reservedResourceType string, location string) (result reservations.ListCatalog, err error)
}

var _ BaseClientAPI = (*reservations.BaseClient)(nil)

// QuotaClientAPI contains the set of methods on the QuotaClient type.
type QuotaClientAPI interface {
	ListStatus(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string) (result reservations.CurrentQuotaLimitBase, err error)
}

var _ QuotaClientAPI = (*reservations.QuotaClient)(nil)

// QuotaRequestClientAPI contains the set of methods on the QuotaRequestClient type.
type QuotaRequestClientAPI interface {
	Create(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest reservations.CurrentQuotaLimitBase, ifMatch string) (result reservations.QuotaRequestCreateFuture, err error)
	Update(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest reservations.CurrentQuotaLimitBase, ifMatch string) (result reservations.QuotaRequestUpdateFuture, err error)
}

var _ QuotaRequestClientAPI = (*reservations.QuotaRequestClient)(nil)

// QuotasClientAPI contains the set of methods on the QuotasClient type.
type QuotasClientAPI interface {
	ListStatus(ctx context.Context, subscriptionID string, providerID string, location string) (result reservations.QuotaLimitsPage, err error)
	ListStatusComplete(ctx context.Context, subscriptionID string, providerID string, location string) (result reservations.QuotaLimitsIterator, err error)
}

var _ QuotasClientAPI = (*reservations.QuotasClient)(nil)

// QuotaRequestsClientAPI contains the set of methods on the QuotaRequestsClient type.
type QuotaRequestsClientAPI interface {
	GetStatus(ctx context.Context, subscriptionID string, providerID string, location string, ID string) (result reservations.QuotaRequestDetails, err error)
	ListStatus(ctx context.Context, subscriptionID string, providerID string, location string, filter string, top *int32, skiptoken string) (result reservations.QuotaRequestDetailsListPage, err error)
	ListStatusComplete(ctx context.Context, subscriptionID string, providerID string, location string, filter string, top *int32, skiptoken string) (result reservations.QuotaRequestDetailsListIterator, err error)
}

var _ QuotaRequestsClientAPI = (*reservations.QuotaRequestsClient)(nil)

// AutoQuotaIncreaseClientAPI contains the set of methods on the AutoQuotaIncreaseClient type.
type AutoQuotaIncreaseClientAPI interface {
	Create(ctx context.Context, subscriptionID string, autoQuotaIncreaseRequest reservations.AutoQuotaIncreaseDetail) (result reservations.AutoQuotaIncreaseDetail, err error)
	GetProperties(ctx context.Context, subscriptionID string) (result reservations.AutoQuotaIncreaseDetail, err error)
}

var _ AutoQuotaIncreaseClientAPI = (*reservations.AutoQuotaIncreaseClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	AvailableScopes(ctx context.Context, reservationOrderID string, reservationID string, body []string) (result reservations.ReservationAvailableScopesFuture, err error)
	Get(ctx context.Context, reservationID string, reservationOrderID string, expand string) (result reservations.Response, err error)
	List(ctx context.Context, reservationOrderID string) (result reservations.ListPage, err error)
	ListComplete(ctx context.Context, reservationOrderID string) (result reservations.ListIterator, err error)
	ListRevisions(ctx context.Context, reservationID string, reservationOrderID string) (result reservations.ListPage, err error)
	ListRevisionsComplete(ctx context.Context, reservationID string, reservationOrderID string) (result reservations.ListIterator, err error)
	Merge(ctx context.Context, reservationOrderID string, body reservations.MergeRequest) (result reservations.ReservationMergeFuture, err error)
	Split(ctx context.Context, reservationOrderID string, body reservations.SplitRequest) (result reservations.SplitFuture, err error)
	Update(ctx context.Context, reservationOrderID string, reservationID string, parameters reservations.Patch) (result reservations.ReservationUpdateFuture, err error)
}

var _ ClientAPI = (*reservations.Client)(nil)

// OrderClientAPI contains the set of methods on the OrderClient type.
type OrderClientAPI interface {
	Calculate(ctx context.Context, body reservations.PurchaseRequest) (result reservations.CalculatePriceResponse, err error)
	Get(ctx context.Context, reservationOrderID string, expand string) (result reservations.OrderResponse, err error)
	List(ctx context.Context) (result reservations.OrderListPage, err error)
	ListComplete(ctx context.Context) (result reservations.OrderListIterator, err error)
	Purchase(ctx context.Context, reservationOrderID string, body reservations.PurchaseRequest) (result reservations.OrderPurchaseFuture, err error)
}

var _ OrderClientAPI = (*reservations.OrderClient)(nil)

// OperationClientAPI contains the set of methods on the OperationClient type.
type OperationClientAPI interface {
	List(ctx context.Context) (result reservations.OperationListPage, err error)
	ListComplete(ctx context.Context) (result reservations.OperationListIterator, err error)
}

var _ OperationClientAPI = (*reservations.OperationClient)(nil)
