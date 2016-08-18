# Twilio API wrapper written in Go

How to use it:

``` Go
t := twiliogo.Twilio{AccountSid: twilioAccountSid, AuthToken: twilioAuthToken}
m := "Your message"
from := "ENTER_FROM_NUMBER"
to := "ENTER_TO_NUMBER"
if err := t.SendSMS(from, to, m); err != nil {
	// handle error
}
```