package main

import (
	"errors"
	"fmt"
	"time"
)

type Operator interface {
	Balance() int64
	Withdraw(amount int64) error
	Deposit(amount int64) error
	Transactions() []Tx
}

type ActionKind string

const (
	ActionKindIncr ActionKind = "+"
	ActionKindDecr ActionKind = "-"
	DateFotm                  = "2006-01-02 15:04:05"
)

type Tx struct {
	value     int64      // значение на которое изменилось
	action    ActionKind // действие, прибавляем или отнимаем
	createdAt time.Time
}

var ErrNegativeValue = errors.New("negative value")

var _ error = (*withdrawError)(nil)

type withdrawError struct {
	err        error
	ActionTime time.Time
}

func (b withdrawError) Error() string {

	return b.err.Error()
}
func (b withdrawError) Print() string {
	b.ActionTime = time.Now()
	return fmt.Sprintf("Ошибка транзакции!!! баланс недостаточен для проведения данной операции %s", b.ActionTime.Format(DateFotm))
}

// Нужно вывести данные транзакции в формате сумма: +-value, time: время создания транзакции
func (t Tx) Print() string {
	return fmt.Sprintf("sum: %s%d, time %s", t.action, t.value, t.createdAt.Format(DateFotm))
}

var _ Operator = (*txOperator)(nil)

type txOperator struct {
	balance      int64
	transactions []Tx
}

func (t *txOperator) Balance() int64 {
	// TODO implement me
	return t.balance
}

func (t *txOperator) Withdraw(amount int64) error {
	// TODO implement me
	if amount > t.balance {
		var err *withdrawError
		return err
	}
	t.balance -= amount
	t.transactions = append(t.transactions, Tx{value: amount, action: ActionKindDecr, createdAt: time.Now()})
	return nil
}

func (t *txOperator) Deposit(amount int64) error {
	if amount < 0 {
		return ErrNegativeValue
	}
	t.balance += amount
	t.transactions = append(t.transactions, Tx{value: amount, action: ActionKindIncr, createdAt: time.Now()})
	return nil
}

func (t *txOperator) Transactions() []Tx {
	return t.transactions
}

func main() {
	var op Operator = &txOperator{}
	_ = op
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Deposit(-100); err != nil {
		if errors.Is(err, ErrNegativeValue) {
			fmt.Println(err)
		}
	}
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Withdraw(90); err != nil {
		var e *withdrawError
		if errors.Is(err, e) {
			var e withdrawError
			fmt.Println(e.Print())
		}

	}

	fmt.Println("Текущий баланс", op.Balance())
	fmt.Println("транзакции", op.Transactions())
	for _, t := range op.Transactions() {
		fmt.Println(t.Print())
	}
}
