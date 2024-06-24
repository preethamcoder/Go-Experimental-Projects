package main

import (
	"errors"
	"fmt"
)

type Account struct {
	income   float64
	tax      float64
	expenses float64
}

func Create_acct(in, tx, exp float64) (*Account, error) {
	if in < 0.0 || tx < 0.0 || exp < 0.0 {
		return &Account{}, errors.New("Bad values. Recheck them.")
	}
	return &Account{
		income:   in,
		tax:      tx,
		expenses: exp,
	}, nil
}

func (act Account) Get_acct() {
	fmt.Println("Income is: ", act.income, "Tax is: ", act.tax, "Expenses are: ", act.expenses)
}

func (act Account) Calc_prof() (float64, error) {
	if act.income < 0 {
		return 0.0, errors.New("Invalid Income")
	}
	profit := act.income - act.expenses
	var rem_profit float64 = profit * (1 - act.tax)
	fmt.Println(rem_profit)
	return rem_profit, nil
}

func main() {
	var inc float64
	var tx float64
	var exp float64
	var temp *Account
	// var e error
	fmt.Println("Enter your income: ")
	_, err := fmt.Scan(&inc)
	if err != nil {
		return
	}
	fmt.Println("Enter your Tax: make sure it is a decimal between 0 and 1")
	_, err = fmt.Scan(&tx)
	if err != nil {
		return
	}
	fmt.Println("Enter you expenses: ")
	_, err = fmt.Scan(&exp)
	if err != nil {
		return
	}
	temp, err = Create_acct(inc, tx, exp)
	if err != nil {
		return
	}
	temp.Get_acct()
	var final float64
	final, _ = temp.Calc_prof()
	fmt.Println("Profit is: ", final)
}
