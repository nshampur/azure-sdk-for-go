package adhybridhealthservice

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
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// AddsServiceMembersClient is the REST APIs for Azure Active Directory Connect Health
type AddsServiceMembersClient struct {
	BaseClient
}

// NewAddsServiceMembersClient creates an instance of the AddsServiceMembersClient client.
func NewAddsServiceMembersClient() AddsServiceMembersClient {
	return NewAddsServiceMembersClientWithBaseURI(DefaultBaseURI)
}

// NewAddsServiceMembersClientWithBaseURI creates an instance of the AddsServiceMembersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAddsServiceMembersClientWithBaseURI(baseURI string) AddsServiceMembersClient {
	return AddsServiceMembersClient{NewWithBaseURI(baseURI)}
}

// Delete deletes a Active Directory Domain Controller server that has been onboarded to Azure Active Directory Connect
// Health Service.
// Parameters:
// serviceName - the name of the service.
// serviceMemberID - the server Id.
// confirm - indicates if the server will be permanently deleted or disabled. True indicates that the server
// will be permanently deleted and False indicates that the server will be marked disabled and then deleted
// after 30 days, if it is not re-registered.
func (client AddsServiceMembersClient) Delete(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, confirm *bool) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServiceMembersClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, serviceName, serviceMemberID, confirm)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AddsServiceMembersClient) DeletePreparer(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, confirm *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceMemberId": autorest.Encode("path", serviceMemberID),
		"serviceName":     autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if confirm != nil {
		queryParameters["confirm"] = autorest.Encode("query", *confirm)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of a server, for a given Active Directory Domain Controller service, that are onboarded to
// Azure Active Directory Connect Health Service.
// Parameters:
// serviceName - the name of the service.
// serviceMemberID - the server Id.
func (client AddsServiceMembersClient) Get(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (result ServiceMember, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServiceMembersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, serviceName, serviceMemberID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AddsServiceMembersClient) GetPreparer(ctx context.Context, serviceName string, serviceMemberID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceMemberId": autorest.Encode("path", serviceMemberID),
		"serviceName":     autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) GetResponder(resp *http.Response) (result ServiceMember, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the details of the Active Directory Domain servers, for a given Active Directory Domain Service, that are
// onboarded to Azure Active Directory Connect Health.
// Parameters:
// serviceName - the name of the service.
// filter - the server property filter to apply.
func (client AddsServiceMembersClient) List(ctx context.Context, serviceName string, filter string) (result AddsServiceMembersPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServiceMembersClient.List")
		defer func() {
			sc := -1
			if result.asm.Response.Response != nil {
				sc = result.asm.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, serviceName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.asm.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List", resp, "Failure sending request")
		return
	}

	result.asm, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AddsServiceMembersClient) ListPreparer(ctx context.Context, serviceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addsservicemembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) ListResponder(resp *http.Response) (result AddsServiceMembers, err error) {
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
func (client AddsServiceMembersClient) listNextResults(ctx context.Context, lastResults AddsServiceMembers) (result AddsServiceMembers, err error) {
	req, err := lastResults.addsServiceMembersPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AddsServiceMembersClient) ListComplete(ctx context.Context, serviceName string, filter string) (result AddsServiceMembersIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServiceMembersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, serviceName, filter)
	return
}

// ListCredentials gets the credentials of the server which is needed by the agent to connect to Azure Active Directory
// Connect Health Service.
// Parameters:
// serviceName - the name of the service.
// serviceMemberID - the server Id.
// filter - the property filter to apply.
func (client AddsServiceMembersClient) ListCredentials(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, filter string) (result Credentials, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServiceMembersClient.ListCredentials")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListCredentialsPreparer(ctx, serviceName, serviceMemberID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "ListCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "ListCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServiceMembersClient", "ListCredentials", resp, "Failure responding to request")
	}

	return
}

// ListCredentialsPreparer prepares the ListCredentials request.
func (client AddsServiceMembersClient) ListCredentialsPreparer(ctx context.Context, serviceName string, serviceMemberID uuid.UUID, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceMemberId": autorest.Encode("path", serviceMemberID),
		"serviceName":     autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/credentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCredentialsSender sends the ListCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServiceMembersClient) ListCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListCredentialsResponder handles the response to the ListCredentials request. The method always
// closes the http.Response Body.
func (client AddsServiceMembersClient) ListCredentialsResponder(resp *http.Response) (result Credentials, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
