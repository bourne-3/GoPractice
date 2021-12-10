package bank

import "sync"

//func Run1() {
//	Deposit(100)
//	Withdraw(10)
//	fmt.Println(Balance())
//}
//
var (
	mu sync.Mutex
	balance int
)
//
//func Deposit(amount int) {
//	mu.Lock()
//	defer mu.Unlock()
//	balance = balance + amount
//}
//
//func Balance() int {
//	mu.Lock()
//	defer mu.Unlock()
//	b := balance
//	return b
//}
//
//func Withdraw(amount int) bool {
//	Deposit(-amount)
//	if Balance() < 0{
//		Deposit(amount)
//		return false
//	}
//	return true
//}