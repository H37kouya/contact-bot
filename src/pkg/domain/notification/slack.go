package notification

type SlackNotificaion interface {
	TestNotification() error
}
