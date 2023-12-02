# API для работы с пользователем

- [Главная](../README.md)


## Methods

- [Sign up](#sign-up)
- [Sign in](#sign-in)
- [Sign out](#sign-out)

---

## Sign up

[Наверх][toup]

Регистрация



```plaintext
POST /api/users/signup
```

Параметры тела запроса

```json
{
    "login": "string"
    "password": "string"
}

```

#### Ответ

В случае успеха - статус кода HTTP - 201.
В случае некорректных данных - статуст кода HTTP - 422 с телом ответа:
```json
{
    "code": "integer",
    "message": "string"
}
```


---

## Sign in

[Наверх][toup]

Аутентификация

```plaintext
POST /api/users/signin
```

Параметры тела запроса

```json
{
    "login": "string"
    "password": "string"
}
```

#### Ответ

В случае успеха - статус кода HTTP - 201 c телом ответа:
```json
{
    "token": "string"
}
```

В случае некорректных данных - статуст кода HTTP - 422 с телом ответа:
```json
{
    "code": "integer",
    "message": "string"
}
```



---

## Sign out

[Наверх][toup]

Выход


```plaintext
DELETE /api/users/signout
```

#### Ответ

В случае успеха - статус кода HTTP - 204.
Если пользователь не авторизован - 401.


[toup]: #api-для-работы-с-пользователем
