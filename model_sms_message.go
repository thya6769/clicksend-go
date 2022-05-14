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

// SmsMessage model
type SmsMessage struct {
	// Your sender id - more info: http://help.clicksend.com/SMS/what-is-a-sender-id-or-sender-number.
	From string `json:"from,omitempty"`
	// Your message.
	Body string `json:"body"`
	// Recipient phone number in E.164 format.
	To string `json:"to,omitempty"`
	// Your method of sending e.g. 'wordpress', 'php', 'c#'.
	Source string `json:"source,omitempty"`
	// Leave blank for immediate delivery. Your schedule time in unix format http://help.clicksend.com/what-is-a-unix-timestamp
	Schedule int32 `json:"schedule,omitempty"`
	// Your reference. Will be passed back with all replies and delivery reports.
	CustomString string `json:"custom_string,omitempty"`
	// Your list ID if sending to a whole list. Can be used instead of 'to'.
	ListId int32 `json:"list_id,omitempty"`
	// Recipient country.
	Country string `json:"country,omitempty"`
	// An email address where the reply should be emailed to. If omitted, the reply will be emailed back to the user who sent the outgoing SMS.
	FromEmail string `json:"from_email,omitempty"`
}
