/*
 * ClickSend v3 REST API
 *
 *  This is the official [ClickSend](https://clicksend.com) SDK.  *You'll need to create a free account to use the API. You can register [here](https://www.clicksend.com/signup).*  You can use our API documentation along with the SDK. Our API docs can be found [here](https://developers.clicksend.com). 
 *
 * API version: 3.1.0
 * Contact: support@clicksend.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Send Email
type Email struct {
	// Array of To Recipient items.
	To []EmailRecipient `json:"to"`
	// Array of Cc Recipient items.
	Cc []EmailRecipient `json:"cc,omitempty"`
	// Array of Bcc Recipient items.
	Bcc []EmailRecipient `json:"bcc,omitempty"`
	// From Email object.
	From []EmailFrom `json:"from"`
	// Body of the email.
	Body string `json:"body"`
	// Array of Attachment items.
	Attachments []Attachment `json:"attachments,omitempty"`
	// Schedule.
	Schedule float32 `json:"schedule,omitempty"`
}