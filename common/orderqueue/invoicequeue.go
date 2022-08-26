package orderqueue

var invoicequeue map[int64]bool = make(map[int64]bool)
var iteration int = 0

func PushOne(orderId int64) {
	invoicequeue[orderId] = true
}

func PullAll() map[int64]bool {
	return invoicequeue
}

func DeleteOne(orderId int64) {
	delete(invoicequeue, orderId)
}

func FlushAll() {
	invoicequeue = make(map[int64]bool, 0)
}

func IsEmpty() bool {
	return len(invoicequeue) <= 0
}

func GetIteration() int {
	return iteration
}

func IterationFinish() {
	iteration++
}

func IterationFlushh() {
	iteration = 0
}
