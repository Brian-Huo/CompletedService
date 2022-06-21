package orderqueue

import (
	"cleaningservice/common/variables"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

var orderawaitqueue map[int64]int
var ordertransferqueue map[int64]string

func Init() {
	orderawaitqueue = make(map[int64]int)
	ordertransferqueue = make(map[int64]string)
}

func Insert(orderid int64) {
	orderawaitqueue[orderid] = 0
}

func Delete(orderid int64) {
	delete(orderawaitqueue, orderid)
	delete(ordertransferqueue, orderid)
}

func GetAwaitQueue() map[int64]int {
	return orderawaitqueue
}

func GetTransferQueue() map[int64]string {
	return ordertransferqueue
}

func CountOrder() {
	for k := range orderawaitqueue {
		orderawaitqueue[k]++
	}
}

func AwaitQueueToMsg() string {
	msg := ""
	for k, v := range orderawaitqueue {
		msg += fmt.Sprintf("<b>order</b> %d is remaining in the system for %d days.<br>", k, v)
	}
	return msg + "</br> Please inform your manager and make sure the order(s) is fully reviewed.</br>Regards.</br>QME Technology Team."
}

func TransferQueueToMsg() string {
	msg := ""
	for k, v := range ordertransferqueue {
		msg += fmt.Sprintf("<b>order</b> %d is requiring immediately transfer with contact details: %s.<br>", k, v)
	}
	return msg + "</br> Please inform your manager and negotiate with our customers ASAP.</br>Regards.</br>QME Technology Team."
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
		if len(orderawaitqueue) > 0 {
			SendQueue(AwaitQueueToMsg())
			logx.Info("Send order queue email to QME Reception.")
		}
		if len(ordertransferqueue) > 0 {
			SendQueue(TransferQueueToMsg())
			logx.Info("Send order transfer email to QME Reception.")
		}
	}
}

func OrderTransferStart(orderId int64, contact string) {
	ordertransferqueue[orderId] = contact
	SendQueue(TransferQueueToMsg())
	logx.Info("Order %d transfer email send to QME Reception.", orderId)
}
