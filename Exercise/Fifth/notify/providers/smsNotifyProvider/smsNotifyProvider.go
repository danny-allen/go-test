package smsNotify

import (
	"fmt"
	"test/demo/Exercise/Fifth/notify"
)


// Email provider struct.
type Provider struct {
	notify.ProviderInterface
}

// Queue the sms message.
func (p *Provider) Queue(uid int64, message string) bool {
	fmt.Println("SMS Notify: " + message)
	return true
}

// Send the sms message.
func (p *Provider) Send() bool {
	fmt.Println("SMS Notify: Sending!!!")
	return true
}
