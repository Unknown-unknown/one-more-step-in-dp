package metrics

import (
	"fmt"
	"log"
	"metrics/conf"
	"time"

	"github.com/robfig/cron/v3"
	"gopkg.in/gomail.v2"
)

var (
	c = cron.New()
)

type receiver struct {
	name    string
	address string
}

type EmailReporter struct {
	storage   Storager
	sender    *EmailSender
	receivers []*receiver
	stopSig   chan interface{}
}

type EmailSender struct{}

func NewEmailReporter(storage Storager, sender *EmailSender) *EmailReporter {
	return &EmailReporter{
		storage:   storage,
		sender:    sender,
		receivers: make([]*receiver, 0),
		stopSig:   make(chan interface{}),
	}
}

func (r *EmailReporter) AddToAddress(name, address string) {
	r.receivers = append(r.receivers, &receiver{
		name:    name,
		address: address,
	})
}

func (r *EmailReporter) StartDailyReport(spec string) {
	fmt.Printf("start at: %v\n", time.Now())
	_, err := c.AddFunc(spec, func() {
		if e := r.sendEmail(); e == nil {
			fmt.Println("Daily report sent! ")
		}
	})
	if err != nil {
		panic(err)
	}
	c.Run()
	fmt.Println("end email report")
}

func (r *EmailReporter) sendEmail() error {
	e := conf.C.Email
	d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	s, err := d.Dial()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	today := time.Now().Format("2006-01-02")
	for _, v := range r.receivers {
		m.SetHeader("From", e.Username)
		m.SetAddressHeader("To", v.address, v.name)
		m.SetHeader("Subject", "Metrics Daily Report@"+today)
		m.SetBody("text/html", fmt.Sprintf("Hello %s!", v.name))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", v.address, err)
			continue
		}
		m.Reset()
	}
	return nil
}
