package main

import "fmt"

type Notifier interface {
    Send(orderID string, status string)
}

type NotifierCreator interface {
    CreateNotifier() Notifier
}

type BaseCreator struct{}

func (b BaseCreator) NotifyClient(c NotifierCreator, orderID string, status string) {
    notifier := c.CreateNotifier()
    notifier.Send(orderID, status)
}

type EmailNotifier struct {
    smtpHost string
}

func (e *EmailNotifier) Send(orderID string, status string) {
    fmt.Printf("Email: заказ %s, статус '%s' через %s\n", orderID, status, e.smtpHost)
}

type SmsNotifier struct {
    gateway string
}

func (s *SmsNotifier) Send(orderID string, status string) {
    fmt.Printf("SMS: заказ %s, статус '%s' через %s\n", orderID, status, s.gateway)
}

type TelegramNotifier struct {
    botToken string
}

func (t *TelegramNotifier) Send(orderID string, status string) {
    preview := t.botToken
    if len(preview) > 6 {
        preview = preview[:6] + "..."
    }
    fmt.Printf("Telegram: заказ %s, статус '%s' token=%s\n", orderID, status, preview)
}

type EmailCreator struct {
    BaseCreator
    smtpHost string
}

func (e *EmailCreator) CreateNotifier() Notifier {
    return &EmailNotifier{smtpHost: e.smtpHost}
}

type SmsCreator struct {
    BaseCreator
    gateway string
}

func (s *SmsCreator) CreateNotifier() Notifier {
    return &SmsNotifier{gateway: s.gateway}
}

type TelegramCreator struct {
    BaseCreator
    botToken string
}

func (t *TelegramCreator) CreateNotifier() Notifier {
    return &TelegramNotifier{botToken: t.botToken}
}

func main() {
    orderID := "A-1042"
    status := "Готов к выдаче"

    emailCreator := &EmailCreator{smtpHost: "smtp.autoservice.local"}
    smsCreator := &SmsCreator{gateway: "sms-gw-1"}
    telegramCreator := &TelegramCreator{botToken: "bot-token-987654"}

    emailCreator.NotifyClient(emailCreator, orderID, status)
    smsCreator.NotifyClient(smsCreator, orderID, status)
    telegramCreator.NotifyClient(telegramCreator, orderID, status)
}
