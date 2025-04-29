package main

import "fmt"

/*
In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then
it is private outside the package it's defined in
*/

// Go lets you create new types from existing ones. type MyName OriginalType
type Bitcoin int

type Wallet struct {
	//since it is lowercase it is private
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// Method on type declration same as struct
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/* This FAILS if func (w Wallet) is used!! Why?
When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
*/

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	//No need to dereference the pointer w here.
	//These are struct pointers and are automaticalled dereferenced!
	return w.balance
}
