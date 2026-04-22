package main

import "fmt"

type Observer interface {
    Update(orderID string, status string)
}

type Subject interface {
    Subscribe(observer Observer)
    Unsubscribe(observer Observer)
    Notify()
}

type BaseNotifier struct {
    channelName string
}

func (b BaseNotifier) Name() string {
    return b.channelName
}

type RepairOrder struct {
    orderID   string
    status    string
    observers []Observer
}

func NewRepairOrder(orderID string) *RepairOrder {
    return &RepairOrder{
        orderID:   orderID,
        status:    "Новая заявка",
        observers: make([]Observer, 0),
    }
}

func (r *RepairOrder) Subscribe(observer Observer) {
    r.observers = append(r.observers, observer)
}

func (r *RepairOrder) Unsubscribe(observer Observer) {
    for i, current := range r.observers {
        if current == observer {
            r.observers = append(r.observers[:i], r.observers[i+1:]...)
            break
        }
    }
}

func (r *RepairOrder) Notify() {
    for _, observer := range r.observers {
        observer.Update(r.orderID, r.status)
    }
}

func (r *RepairOrder) SetStatus(status string) {
    fmt.Printf("[Заказ %s] смена статуса: %s -> %s\n", r.orderID, r.status, status)
    r.status = status
    r.Notify()
}

type EmailNotifier struct {
    BaseNotifier
    smtpHost string
    email    string
}

func NewEmailNotifier(smtpHost, email string) *EmailNotifier {
    return &EmailNotifier{
        BaseNotifier: BaseNotifier{channelName: "Email"},
        smtpHost:     smtpHost,
        email:        email,
    }
}

func (e *EmailNotifier) Update(orderID string, status string) {
    fmt.Printf("  [Email -> %s через %s] Заказ %s: %s\n",
        e.email, e.smtpHost, orderID, status)
}

type SmsNotifier struct {
    BaseNotifier
    gateway string
    phone   string
}

func NewSmsNotifier(gateway, phone string) *SmsNotifier {
    return &SmsNotifier{
        BaseNotifier: BaseNotifier{channelName: "SMS"},
        gateway:      gateway,
        phone:        phone,
    }
}

func (s *SmsNotifier) Update(orderID string, status string) {
    fmt.Printf("  [SMS -> %s через %s] Заказ %s: %s\n",
        s.phone, s.gateway, orderID, status)
}

type DashboardNotifier struct {
    BaseNotifier
    screenID string
}

func NewDashboardNotifier(screenID string) *DashboardNotifier {
    return &DashboardNotifier{
        BaseNotifier: BaseNotifier{channelName: "Dashboard"},
        screenID:     screenID,
    }
}

func (d *DashboardNotifier) Update(orderID string, status string) {
    fmt.Printf("  [Панель %s] Заказ %s: %s\n", d.screenID, orderID, status)
}

func main() {
    order := NewRepairOrder("A-1042")

    emailNotifier := NewEmailNotifier("smtp.autoservice.local", "client@example.com")
    smsNotifier := NewSmsNotifier("sms-gw-1", "+79990001122")
    dashboardNotifier := NewDashboardNotifier("operator-screen-3")

    order.Subscribe(emailNotifier)
    order.Subscribe(smsNotifier)
    order.Subscribe(dashboardNotifier)

    order.SetStatus("Автомобиль принят в работу")
    fmt.Println()

    order.SetStatus("Ожидание запчастей")
    fmt.Println()

    order.Unsubscribe(smsNotifier)
    order.SetStatus("Работы завершены")
}
