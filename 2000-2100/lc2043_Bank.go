package main

type Bank struct {
	balance []int64
}

func ConstructorBank(balance []int64) Bank {
	return Bank{balance: balance}
}

func (bank *Bank) getAccount(acc int) (*int64, bool) {
	// fmt.Println(bank.balance)
	if acc < 1 || acc > len(bank.balance) {
		return nil, false
	}
	return &bank.balance[acc-1], true
}

func (bank *Bank) Transfer(account1 int, account2 int, money int64) bool {
	var m1, m2 *int64
	var ok bool
	if m1, ok = bank.getAccount(account1); !ok {
		return false
	}
	if m2, ok = bank.getAccount(account2); !ok {
		return false
	}

	if *m1 < money {
		return false
	}
	*m1 -= money
	*m2 += money
	return true
}

func (bank *Bank) Deposit(account int, money int64) bool {
	if m1, ok := bank.getAccount(account); !ok {
		return false
	} else {
		*m1 += money
	}
	return true
}

func (bank *Bank) Withdraw(account int, money int64) bool {
	if m1, ok := bank.getAccount(account); !ok {
		return false
	} else if *m1 < money {
		return false
	} else {
		*m1 -= money
	}
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
