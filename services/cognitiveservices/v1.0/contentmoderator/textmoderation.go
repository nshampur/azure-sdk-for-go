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

// TextModerationClient is the you use the API to scan your content as it is generated. Content Moderator then
// processes your content and sends the results along with relevant information either back to your systems or to the
// built-in review tool. You can use this information to take decisions e.g. take it down, send to human judge, etc.
//
// When using the API, images need to have a minimum of 128 pixels and a maximum file size of 4MB.
// Text can be at most 1024 characters long.
// If the content passed to the text API or the image API exceeds the size limits, the API will return an error code
// that informs about the issue.
type TextModerationClient struct {
	BaseClient
}

// NewTextModerationClient creates an instance of the TextModerationClient client.
func NewTextModerationClient(endpoint string) TextModerationClient {
	return TextModerationClient{New(endpoint)}
}

// DetectLanguage this operation will detect the language of given input content. Returns the <a
// href="http://www-01.sil.org/iso639-3/codes.asp">ISO 639-3 code</a> for the predominant language comprising the
// submitted text. Over 110 languages supported.
// Parameters:
// textContentType - the content type.
// textContent - content to screen.
func (client TextModerationClient) DetectLanguage(ctx context.Context, textContentType string, textContent io.ReadCloser) (result DetectedLanguage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TextModerationClient.DetectLanguage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetectLanguagePreparer(ctx, textContentType, textContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.TextModerationClient", "DetectLanguage", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectLanguageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.TextModerationClient", "DetectLanguage", resp, "Failure sending request")
		return
	}

	result, err = client.DetectLanguageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.TextModerationClient", "DetectLanguage", resp, "Failure responding to request")
	}

	return
}

// DetectLanguagePreparer prepares the DetectLanguage request.
func (client TextModerationClient) DetectLanguagePreparer(ctx context.Context, textContentType string, textContent io.ReadCloser) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("text/plain"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPath("/contentmoderator/moderate/v1.0/ProcessText/DetectLanguage"),
		autorest.WithFile(textContent),
		autorest.WithHeader("Content-Type", autorest.String(textContentType)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectLanguageSender sends the DetectLanguage request. The method will close the
// http.Response Body if it receives an error.
func (client TextModerationClient) DetectLanguageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectLanguageResponder handles the response to the DetectLanguage request. The method always
// closes the http.Response Body.
func (client TextModerationClient) DetectLanguageResponder(resp *http.Response) (result DetectedLanguage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ScreenText detects profanity in more than 100 languages and match against custom and shared blacklists.
// Parameters:
// textContentType - the content type.
// textContent - content to screen.
// language - language of the text.
// autocorrect - autocorrect text.
// pii - detect personal identifiable information.
// listID - the list Id.
// classify - classify input.
func (client TextModerationClient) ScreenText(ctx context.Context, textContentType string, textContent io.ReadCloser, language string, autocorrect *bool, pii *bool, listID string, classify *bool) (result Screen, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TextModerationClient.ScreenText")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ScreenTextPreparer(ctx, textContentType, textContent, language, autocorrect, pii, listID, classify)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.TextModerationClient", "ScreenText", nil, "Failure preparing request")
		return
	}

	resp, err := client.ScreenTextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.TextModerationClient", "ScreenText", resp, "Failure sending request")
		return
	}

	result, err = client.ScreenTextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.TextModerationClient", "ScreenText", resp, "Failure responding to request")
	}

	return
}

// ScreenTextPreparer prepares the ScreenText request.
func (client TextModerationClient) ScreenTextPreparer(ctx context.Context, textContentType string, textContent io.ReadCloser, language string, autocorrect *bool, pii *bool, listID string, classify *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	queryParameters := map[string]interface{}{}
	if len(language) > 0 {
		queryParameters["language"] = autorest.Encode("query", language)
	}
	if autocorrect != nil {
		queryParameters["autocorrect"] = autorest.Encode("query", *autocorrect)
	} else {
		queryParameters["autocorrect"] = autorest.Encode("query", false)
	}
	if pii != nil {
		queryParameters["PII"] = autorest.Encode("query", *pii)
	} else {
		queryParameters["PII"] = autorest.Encode("query", false)
	}
	if len(listID) > 0 {
		queryParameters["listId"] = autorest.Encode("query", listID)
	}
	if classify != nil {
		queryParameters["classify"] = autorest.Encode("query", *classify)
	} else {
		queryParameters["classify"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("text/plain"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPath("/contentmoderator/moderate/v1.0/ProcessText/Screen/"),
		autorest.WithFile(textContent),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Content-Type", autorest.String(textContentType)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ScreenTextSender sends the ScreenText request. The method will close the
// http.Response Body if it receives an error.
func (client TextModerationClient) ScreenTextSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ScreenTextResponder handles the response to the ScreenText request. The method always
// closes the http.Response Body.
func (client TextModerationClient) ScreenTextResponder(resp *http.Response) (result Screen, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
