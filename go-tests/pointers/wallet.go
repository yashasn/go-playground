package main

import (
	"errors"
	"fmt"
)

/*
In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then
it is private outside the package it's defined in
*/

// Go lets you create new types from existing ones. type MyName OriginalType
type Bitcoin int

type Stringer interface {
	String() string
}

// Method on type declration same as struct
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	//since it is lowercase it is private
	balance Bitcoin
}

/* This FAILS if func (w Wallet) is used!! Why?
When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
*/

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

//NOTE : The var keyword allows us to define values global to the package. Here, we can use this error during test assertion

var ErrInsufficientFunds = errors.New("You are broke!!")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	//No need to dereference the pointer w here.
	//These are struct pointers and are automaticalled dereferenced!
	return w.balance
}
