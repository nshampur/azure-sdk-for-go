package contentmoderator

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
	"io"
	"net/http"
)

// ListManagementImageClient is the you use the API to scan your content as it is generated. Content Moderator then
// processes your content and sends the results along with relevant information either back to your systems or to the
// built-in review tool. You can use this information to take decisions e.g. take it down, send to human judge, etc.
//
// When using the API, images need to have a minimum of 128 pixels and a maximum file size of 4MB.
// Text can be at most 1024 characters long.
// If the content passed to the text API or the image API exceeds the size limits, the API will return an error code
// that informs about the issue.
type ListManagementImageClient struct {
	BaseClient
}

// NewListManagementImageClient creates an instance of the ListManagementImageClient client.
func NewListManagementImageClient(endpoint string) ListManagementImageClient {
	return ListManagementImageClient{New(endpoint)}
}

// AddImage add an image to the list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
// tag - tag for the image.
// label - the image label.
func (client ListManagementImageClient) AddImage(ctx context.Context, listID string, tag *int32, label string) (result Image, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageClient.AddImage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddImagePreparer(ctx, listID, tag, label)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImage", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddImageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImage", resp, "Failure sending request")
		return
	}

	result, err = client.AddImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImage", resp, "Failure responding to request")
	}

	return
}

// AddImagePreparer prepares the AddImage request.
func (client ListManagementImageClient) AddImagePreparer(ctx context.Context, listID string, tag *int32, label string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	queryParameters := map[string]interface{}{}
	if tag != nil {
		queryParameters["tag"] = autorest.Encode("query", *tag)
	}
	if len(label) > 0 {
		queryParameters["label"] = autorest.Encode("query", label)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}/images", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddImageSender sends the AddImage request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageClient) AddImageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddImageResponder handles the response to the AddImage request. The method always
// closes the http.Response Body.
func (client ListManagementImageClient) AddImageResponder(resp *http.Response) (result Image, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AddImageFileInput add an image to the list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
// imageStream - the image file.
// tag - tag for the image.
// label - the image label.
func (client ListManagementImageClient) AddImageFileInput(ctx context.Context, listID string, imageStream io.ReadCloser, tag *int32, label string) (result Image, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageClient.AddImageFileInput")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddImageFileInputPreparer(ctx, listID, imageStream, tag, label)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImageFileInput", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddImageFileInputSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImageFileInput", resp, "Failure sending request")
		return
	}

	result, err = client.AddImageFileInputResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImageFileInput", resp, "Failure responding to request")
	}

	return
}

// AddImageFileInputPreparer prepares the AddImageFileInput request.
func (client ListManagementImageClient) AddImageFileInputPreparer(ctx context.Context, listID string, imageStream io.ReadCloser, tag *int32, label string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	queryParameters := map[string]interface{}{}
	if tag != nil {
		queryParameters["tag"] = autorest.Encode("query", *tag)
	}
	if len(label) > 0 {
		queryParameters["label"] = autorest.Encode("query", label)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("image/gif"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}/images", pathParameters),
		autorest.WithFile(imageStream),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddImageFileInputSender sends the AddImageFileInput request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageClient) AddImageFileInputSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddImageFileInputResponder handles the response to the AddImageFileInput request. The method always
// closes the http.Response Body.
func (client ListManagementImageClient) AddImageFileInputResponder(resp *http.Response) (result Image, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AddImageURLInput add an image to the list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
// contentType - the content type.
// imageURL - the image url.
// tag - tag for the image.
// label - the image label.
func (client ListManagementImageClient) AddImageURLInput(ctx context.Context, listID string, contentType string, imageURL ImageURL, tag *int32, label string) (result Image, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageClient.AddImageURLInput")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddImageURLInputPreparer(ctx, listID, contentType, imageURL, tag, label)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImageURLInput", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddImageURLInputSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImageURLInput", resp, "Failure sending request")
		return
	}

	result, err = client.AddImageURLInputResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "AddImageURLInput", resp, "Failure responding to request")
	}

	return
}

// AddImageURLInputPreparer prepares the AddImageURLInput request.
func (client ListManagementImageClient) AddImageURLInputPreparer(ctx context.Context, listID string, contentType string, imageURL ImageURL, tag *int32, label string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	queryParameters := map[string]interface{}{}
	if tag != nil {
		queryParameters["tag"] = autorest.Encode("query", *tag)
	}
	if len(label) > 0 {
		queryParameters["label"] = autorest.Encode("query", label)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}/images", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Content-Type", autorest.String(contentType)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddImageURLInputSender sends the AddImageURLInput request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageClient) AddImageURLInputSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddImageURLInputResponder handles the response to the AddImageURLInput request. The method always
// closes the http.Response Body.
func (client ListManagementImageClient) AddImageURLInputResponder(resp *http.Response) (result Image, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteAllImages deletes all images from the list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
func (client ListManagementImageClient) DeleteAllImages(ctx context.Context, listID string) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageClient.DeleteAllImages")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteAllImagesPreparer(ctx, listID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "DeleteAllImages", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAllImagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "DeleteAllImages", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAllImagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "DeleteAllImages", resp, "Failure responding to request")
	}

	return
}

// DeleteAllImagesPreparer prepares the DeleteAllImages request.
func (client ListManagementImageClient) DeleteAllImagesPreparer(ctx context.Context, listID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}/images", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAllImagesSender sends the DeleteAllImages request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageClient) DeleteAllImagesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAllImagesResponder handles the response to the DeleteAllImages request. The method always
// closes the http.Response Body.
func (client ListManagementImageClient) DeleteAllImagesResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteImage deletes an image from the list with list Id and image Id passed.
// Parameters:
// listID - list Id of the image list.
// imageID - id of the image.
func (client ListManagementImageClient) DeleteImage(ctx context.Context, listID string, imageID string) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageClient.DeleteImage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteImagePreparer(ctx, listID, imageID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "DeleteImage", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteImageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "DeleteImage", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "DeleteImage", resp, "Failure responding to request")
	}

	return
}

// DeleteImagePreparer prepares the DeleteImage request.
func (client ListManagementImageClient) DeleteImagePreparer(ctx context.Context, listID string, imageID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"ImageId": autorest.Encode("path", imageID),
		"listId":  autorest.Encode("path", listID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}/images/{ImageId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteImageSender sends the DeleteImage request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageClient) DeleteImageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteImageResponder handles the response to the DeleteImage request. The method always
// closes the http.Response Body.
func (client ListManagementImageClient) DeleteImageResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllImageIds gets all image Ids from the list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
func (client ListManagementImageClient) GetAllImageIds(ctx context.Context, listID string) (result ImageIds, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageClient.GetAllImageIds")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAllImageIdsPreparer(ctx, listID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "GetAllImageIds", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllImageIdsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "GetAllImageIds", resp, "Failure sending request")
		return
	}

	result, err = client.GetAllImageIdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageClient", "GetAllImageIds", resp, "Failure responding to request")
	}

	return
}

// GetAllImageIdsPreparer prepares the GetAllImageIds request.
func (client ListManagementImageClient) GetAllImageIdsPreparer(ctx context.Context, listID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}/images", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllImageIdsSender sends the GetAllImageIds request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageClient) GetAllImageIdsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAllImageIdsResponder handles the response to the GetAllImageIds request. The method always
// closes the http.Response Body.
func (client ListManagementImageClient) GetAllImageIdsResponder(resp *http.Response) (result ImageIds, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
