package notify


// Declare Providers array.
var Providers []*NotifyProvider


// Notify provider interface.
type ProviderInterface interface {
	Queue(uid int64, message string) (bool)
	Send() (bool)
}

// Abstract Provider.
type NotifyProvider struct {
	Provider ProviderInterface
}

// Register your provider (must use NotifyProvider and implement its methods).
func RegisterProvider(e *NotifyProvider) {

	// Add the providers to the slice.
	Providers = append(Providers, e)
}

// Queue wrapper.
func Queue(uid int64, message string) {

	// Loop the providers.
	for _, v := range Providers {

		// On call the method on the provider.
		v.Provider.Queue(uid, message)
	}
}

// Send wrapper.
func Send() {

	// Loop the providers.
	for _, v := range Providers {

		// On call the method on the provider.
		v.Provider.Send()
	}
}

//
//type AbstractProvider struct {
//	ProviderInterface
//	Name string
//}
//
//// Method to register providers.
//func RegisterProvider(pkg string, st string) {
//
//	providers[pkg] = st
//}
//
//// Notification data structure.
//type Notification struct {
//	Message 	string
//	UserId		int64
//}
//
//
//// Queue the notification
//func (AbstractProvider) Queue(uid int64, message string) {
//
//	// Loop the providers.
//	for i, provider := range providers {
//
//		fmt.Println(providers[i])
//		fmt.Println(provider)
//		// Queue for each of the providers.
//		//provider.Queue
//	}
//}
//




/*
	Notifications triggers
		- Create Event
		- Update Event
		- Create Comment

	Notification Output
		- Timeline
		- Email

	Notifications Storage
		- Database

	Notifications Providers
		- Email
		- Push Notifications
		- SMS

	Methods
		- notification.queue(uid, message)
		- notification.send()
 */