package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

type EmailNotification struct{}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending SMS notification")
}
func (EmailNotification) SendNotification() {
	fmt.Println("Sending EMAIL notification")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type SMSNotificationSender struct {
}
type EmailNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNoticationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNoticationFactory("SMS")
	emaiFactory, _ := getNoticationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emaiFactory)

	getMethod(smsFactory)
	getMethod(emaiFactory)
}
