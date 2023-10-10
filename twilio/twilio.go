package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	twilioAccountSid = "ACdba33859d1797a271d64188ed200dae1"
	twilioAuthToken  = "4b49710ae7effdf0f0d2ad168e015e1b"
	twilioFromNumber = "+12314000751"
)

func SendSMS(to, body string) error {
	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", twilioAccountSid)
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", twilioFromNumber)
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(twilioAccountSid, twilioAuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("failed to send SMS: %s", resp.Status)
}