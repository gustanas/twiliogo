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

Twitter: [@gusta_nas](https://twitter.com/gusta_nas)

Website: [gustanas.co](http://www.gustanas.co/)
