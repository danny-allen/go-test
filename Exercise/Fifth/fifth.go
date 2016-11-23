package fifth

import (
	"test/demo/Exercise/Fifth/notify"
	"test/demo/Exercise/Fifth/notify/providers/emailNotifyProvider"
	"test/demo/Exercise/Fifth/notify/providers/pushNotifyProvider"
	"test/demo/Exercise/Fifth/notify/providers/smsNotifyProvider"
)


// Our main function.
func Run() {

	// Register email provider.
	notify.RegisterProvider(&notify.NotifyProvider{
		Provider: &emailNotify.Provider{},
	})

	// Register push notification provider.
	notify.RegisterProvider(&notify.NotifyProvider{
		Provider: &pushNotifyProvider.Provider{},
	})

	// Register push notification provider.
	notify.RegisterProvider(&notify.NotifyProvider{
		Provider: &smsNotify.Provider{},
	})

	// Jump to what happens later - use the services.
	Later();
}


// Make use of the notify feature, which should know what providers to use.
func Later() {

	// Queue a message.
	notify.Queue(int64(4), "New Message")

	// Send the message.
	notify.Send()
}