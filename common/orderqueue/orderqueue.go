package orderqueue

import (
	"cleaningservice/common/variables"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

var orderqueue map[int64]int

func Init() {
	orderqueue = make(map[int64]int)
}

func Insert(orderid int64) {
	orderqueue[orderid] = 0
}

func Delete(orderid int64) {
	delete(orderqueue, orderid)
}

func GetQueue() map[int64]int {
	return orderqueue
}

func CountOrder() {
	for k := range orderqueue {
		orderqueue[k]++
	}
}

func QueueToMsg() string {
	msg := ""
	for k, v := range orderqueue {
		msg += fmt.Sprintf("<b>order</b> %d is remaining in the system for %d days.</br>", k, v)
	}
	return msg + "</br> Please inform your manager and make sure the order(s) is fully reviewed.</br>Regards.</br>QME Technology Team."
}

func SendQueue(msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", variables.QME_email)
	m.SetHeader("To", variables.Reception_email)
	m.SetHeader("Subject", "Pending Orders")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer("smtp.gmail.com", 587, variables.QME_email, variables.QME_password)

	// Send the email to QME Reception.
	if err := d.DialAndSend(m); err != nil {
		logx.Info("Send email failed: %v", err)
	}
}

func OrderQueueStart() {
	Init()

	for {
		time.Sleep(time.Second * time.Duration(variables.Check_time))
		CountOrder()
		msg := QueueToMsg()
		if len(orderqueue) > 0 {
			SendQueue(msg)
		}
	}
}
