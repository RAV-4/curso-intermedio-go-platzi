package main

import "fmt"

//Notificaciones por SMS o Email

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Tigo"
}

type EmailNotifcation struct {
}

func (EmailNotifcation) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotifcation) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotifcation{}, nil
	}

	return nil, fmt.Errorf("no notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emalFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emalFactory)

	getMethod(smsFactory)
	getMethod(emalFactory)
}
