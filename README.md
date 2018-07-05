# medicalcard

## Usecasec

1. Как пользователь, я хочу иметь возможность просмотра моей медицинской карты.
1. Как пользователь, я хочу иметь возможность создавать медицинскую карту.
1. Как пользователь, я хочу иметь возможность редактирования личные данные карты.(ФИО, адрес, телефон. etc.)
1. Как пользователь, я хочу иметь возможность удаления своей медицинской карты.

## REST API

### GET /api/medicalcard/contacts

Ответ: 200 OK
```json
    [{
        "name": "Имя",
        "surname": "Фамилия",
        "secondname": "Отчество",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "adress": "Обнинск, Ленина 69, 442"
    }]
```

### POST /api/medicalcard/contacts

Тело запроса:

```json
    {
        "name": "Имя",
        "surname": "Фамилия",
        "secondname": "Отчество",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "adress": "Обнинск, Ленина 69 "
    }
```

    Ответ: 201 Created
    Location: /api/medicalcard/mycard/1

### PUT /api/medicalcard/contacts

Тело запроса:

```json
    {
        "name": "Имя",
        "surname": "Фамилия",
        "secondname": "Отчество",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "adress": "Обнинск, Ленина 69"
    }
```

    Ответ: 202 Accepted
    Location: /api/medicalcard/mycard/1

### DELETE /api/medicalcard/contacts


    Ответ: 204 No Content
   