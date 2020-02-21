# facade

`facade` предоставляет собой заглушку сервиса обработки платежей для демонстрации паттерна Facade.

Он предоставляет следующую функциональность:
   1. Снять деньги с аккаунта.
   2. Получить баланс аккаунта.
   
```(go)
type PaymentSystem interface {
	Withdraw(id string, amount uint32) (err error)
	Balance(id string) (balance int, err error)
}
```

## Структура проекта
1. `/pkg/wallet`: - кошелек, у которого есть имя пользователя и баланс.

2.  `/pkg/transaction` - система хранения операций.

3.  `/pkg/facade` - сама платежная система.

## Запуск тестов
`go test ./...`
   
