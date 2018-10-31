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

// ResellerAccount model
type ResellerAccount struct {
	// Account username
	Username string `json:"username"`
	// Account password (unhashed)
	Password string `json:"password"`
	// Account email
	UserEmail string `json:"user_email"`
	// Account phone number
	UserPhone string `json:"user_phone"`
	// Account owner first name
	UserFirstName string `json:"user_first_name"`
	// Account owner last name
	UserLastName string `json:"user_last_name"`
	// Account name (usually company name)
	AccountName string `json:"account_name"`
	// Country of account holder
	Country string `json:"country"`
}