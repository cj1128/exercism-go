package account

import "sync"

type Account struct {
	closed  bool
	balance int64
	sync.Mutex
}

func Open(initial int64) *Account {
	if initial < 0 {
		return nil
	}
	return &Account{
		balance: initial,
	}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	newBalance := a.balance + amount
	if newBalance < 0 {
		return 0, false
	}
	a.balance = newBalance
	return newBalance, true
}
