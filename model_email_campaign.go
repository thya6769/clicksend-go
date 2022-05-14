/*
 * ClickSend v3 API
 *
 *  This is an official SDK for [ClickSend](https://clicksend.com)  Below you will find a current list of the available methods for clicksend.  *NOTE: You will need to create a free account to use the API. You can register [here](https://dashboard.clicksend.com/#/signup/step1/)..*
 *
 * API version: 3.1
 * Contact: support@clicksend.com
 * Generated by: Swagger Codegen (https://github.com/clicksend-api/clicksend-codegen.git)
 */

package clicksend

// Campaign Model for Email
type EmailCampaign struct {
	// Your campaign name.
	Name string `json:"name"`
	// Your campaign subject.
	Subject string `json:"subject"`
	// Your campaign message.
	Body string `json:"body"`
	// The allowed email address id.
	FromEmailAddressId float32 `json:"from_email_address_id"`
	// Your name or business name.
	FromName string `json:"from_name"`
	// Your template id.
	TemplateId float32 `json:"template_id,omitempty"`
	// Your contact list id.
	ListId float32 `json:"list_id"`
	// Your schedule timestamp.
	Schedule int32 `json:"schedule,omitempty"`
}
