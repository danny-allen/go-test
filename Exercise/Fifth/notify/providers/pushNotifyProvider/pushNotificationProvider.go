package pushNotifyProvider

import (
	"fmt"
	"test/demo/Exercise/Fifth/notify"
)

type Provider struct {
	notify.ProviderInterface
	Name 	string
}

func (p *Provider) Queue(uid int64, message string) bool {
	fmt.Println("Push Notify: " + message)
	return true
}

func (p *Provider) Send() bool {
	fmt.Println("Push Notify: Sending!!!")
	return true
}



