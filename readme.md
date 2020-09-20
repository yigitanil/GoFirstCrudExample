# CUSTOMER API

1. Instructions
2. Tech Stack
3. Model


## 1.Instructions
### Run

```
    go run main.go
```

## 2.Tech Stack

Used Libraries and Frameworks:
- Gorm
- Echo

## 3.Model

**Relation: Customer, 0..1 -> N Address**


*Customer- 0..1*


```
- firstName: String
- lastName: String
- phoneNumber: String
- email: String
```

*Address- N*
```
- city: String
- country: String
- address: String
- customerId: long

