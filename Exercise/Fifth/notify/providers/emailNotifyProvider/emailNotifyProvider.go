package emailNotify

import (
	"fmt"
	"test/demo/Exercise/Fifth/notify"
)


// Email provider struct.
type Provider struct {
	notify.ProviderInterface
}

// Queue the email message.
func (p *Provider) Queue(uid int64, message string) bool {
	fmt.Println("Email Notify: " + message)
	return true
}

// Send the email message.
func (p *Provider) Send() bool {
	fmt.Println("Email Notify: Sending!!!")
	return true
}
