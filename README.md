
# Withdrawal System

Withdrawal System Project is an API that allows users to request a withdrawal amount. The system validates the input to ensure that the amount is not negative, for example. It then calculates the smallest possible number of banknotes to meet the requested amount, returning the final result.


## Techs

* [Golang](https://go.dev) - Programming language
* [Fiber](https://gofiber.io) - Web framework for Go
* [Swagger](https://swagger.io) - API documentation tool
## Endpoints

#### Transaction
- **POST** `localhost:3000/transaction`: Make a withdrawal
- **GET** `localhost:3000/swagger`: Swagger documentation
## Environment Setting

1. **Clone the repository**

```bash
 git clone https://github.com/jotaGGod/withdrawal-system.git
```

2. **Navigate to the project directory**
```bash
 cd withdrawal-system
```

3. **Install dependencies**
```bash
 go mod tidy
```

4. **Run the application**
```bash
 go run main.go
```
5. **Run Tests (Optional)**
```bash
 go test ./... -v
```