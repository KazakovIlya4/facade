# facade

`facade` предоставляет собой заглушку сервиса обработки платежей для демонстрации паттерна Facade.

Он предоставляет следующую функциональность:
   1. Снять деньги с аккаунта.
   2. Получить баланс аккаунта.
   
```(go)
type PaymentSystem interface {
	GetMoney(id string, amount uint32, securityCode int, transactionID int) error
	GetBalance(id string, sercurityCode int, transactionID int) (int, error)
}
```

## Структура проекта
1. `/pkg/wallet`: - кошелек, у которого есть имя пользователя и баланс.

2.  `/pkg/security` - система подтверждения операции по защитному коду и id операции.

3.  `/pkg/facade` - сама платежная система.

## Запуск тестов
`go test ./pkg/facade/`
   
