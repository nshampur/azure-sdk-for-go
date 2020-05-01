package digitaltwins

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// IoTHubsClient is the azure Digital Twins Client for managing DigitalTwinsInstance
type IoTHubsClient struct {
	BaseClient
}

// NewIoTHubsClient creates an instance of the IoTHubsClient client.
func NewIoTHubsClient(subscriptionID uuid.UUID) IoTHubsClient {
	return NewIoTHubsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIoTHubsClientWithBaseURI creates an instance of the IoTHubsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIoTHubsClientWithBaseURI(baseURI string, subscriptionID uuid.UUID) IoTHubsClient {
	return IoTHubsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List get DigitalTwinsInstance IoTHubs.
// Parameters:
// resourceGroupName - the name of the resource group that contains the DigitalTwinsInstance.
// resourceName - the name of the DigitalTwinsInstance.
func (client IoTHubsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result IntegrationResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTHubsClient.List")
		defer func() {
			sc := -1
			if result.irlr.Response.Response != nil {
				sc = result.irlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("digitaltwins.IoTHubsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.IoTHubsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.irlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "digitaltwins.IoTHubsClient", "List", resp, "Failure sending request")
		return
	}

	result.irlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.IoTHubsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IoTHubsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/integrationResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IoTHubsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IoTHubsClient) ListResponder(resp *http.Response) (result IntegrationResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IoTHubsClient) listNextResults(ctx context.Context, lastResults IntegrationResourceListResult) (result IntegrationResourceListResult, err error) {
	req, err := lastResults.integrationResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "digitaltwins.IoTHubsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "digitaltwins.IoTHubsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.IoTHubsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IoTHubsClient) ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result IntegrationResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTHubsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, resourceName)
	return
}
