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

// Array of MmsMessage items
type MmsMessageCollection struct {
	// Media file you want to send
	MediaFile string `json:"media_file"`
	// Array of MmsMessage models
	Messages []MmsMessage `json:"messages"`
}
