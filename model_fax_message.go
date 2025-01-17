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

// Base model for Fax Messages
type FaxMessage struct {
	// Your method of sending e.g. 'wordpress', 'php', 'c#'.
	Source string `json:"source,omitempty"`
	// Recipient fax number in E.164 format.
	To string `json:"to"`
	// Your list ID if sending to a whole list. Can be used instead of 'to'.
	ListId int32 `json:"list_id,omitempty"`
	// Your sender id. Must be a valid fax number.
	From string `json:"from,omitempty"`
	// Leave blank for immediate delivery. Your schedule time in unix format http://help.clicksend.com/what-is-a-unix-timestamp
	Schedule int32 `json:"schedule,omitempty"`
	// Your reference. Will be passed back with all replies and delivery reports.
	CustomString string `json:"custom_string,omitempty"`
	// Recipient country.
	Country string `json:"country,omitempty"`
	// An email address where the reply should be emailed to.
	FromEmail string `json:"from_email,omitempty"`
}
