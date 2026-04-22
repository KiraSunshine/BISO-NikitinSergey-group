package main

import (
    "fmt"
    "sync"
)

type AppConfig struct {
    smtpHost      string
    smsGateway    string
    telegramToken string
    retryCount    int
}

var (
    instance *AppConfig
    once     sync.Once
)

func GetAppConfig() *AppConfig {
    once.Do(func() {
        fmt.Println("Инициализация AppConfig выполняется один раз")
        instance = &AppConfig{
            smtpHost:      "smtp.autoservice.local",
            smsGateway:    "sms-gw-1",
            telegramToken: "token-987654321",
            retryCount:    3,
        }
    })
    return instance
}

func (a *AppConfig) Summary() string {
    return fmt.Sprintf(
        "smtp=%s, sms=%s, retries=%d",
        a.smtpHost,
        a.smsGateway,
        a.retryCount,
    )
}

type OrderService struct {
    config *AppConfig
}

func NewOrderService() *OrderService {
    return &OrderService{config: GetAppConfig()}
}

func (o *OrderService) PrintConfig() {
    fmt.Println("OrderService использует:", o.config.Summary())
}

func main() {
    serviceA := NewOrderService()
    serviceB := NewOrderService()

    serviceA.PrintConfig()
    serviceB.PrintConfig()

    fmt.Printf("Адрес config в serviceA: %p\n", serviceA.config)
    fmt.Printf("Адрес config в serviceB: %p\n", serviceB.config)
}
