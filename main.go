package twiliogo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Twilio stores the API keys and makes the call to the Twilio API
type Twilio struct {
	AccountSid string
	AuthToken  string
}

// SendSMS sends a sms message to a given number
func (t Twilio) SendSMS(from, to, message string) error {
	// Set initial variables
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + t.AccountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", to)
	v.Set("From", from)
	v.Set("Body", message)
	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}

	req, err := http.NewRequest("POST", urlStr, &rb)
	if err != nil {
		return err
	}
	req.SetBasicAuth(t.AccountSid, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return err
		}
		fmt.Println(data["sid"])
	} else {
		// TODO: create error?
		fmt.Println(resp.Status)
	}
	return nil
}
