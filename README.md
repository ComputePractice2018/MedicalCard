# medicalcard

## Usecasec

1. Как пользователь, я хочу иметь возможность просмотра моей медицинской карты.
1. Как пользователь, я хочу иметь возможность создания новых записей о приемах.
1. Как пользователь, я хочу иметь возможность редактирования записей о приемах.(ФИО врача/медсестры, жалобы, дата приема, etc.)
1. Как пользователь, я хочу иметь возможность удаления записей о приемах.

## REST API

### GET /api/medicalcard/appointments

Ответ: 200 OK
```json
   [{
        "date": "01.07.2018", 
        "doc_name": "Смирнов Василий Данилович", 
        "muse_name": "Прокофьева Вероника Максимовна", 
        "complaint": "Чихает находясь в помещении с цветами", 
        "checkup": "Покраснения слизистой", 
        "diagnosis": "Стоматит", 
        "treatment": "Изолироватть пациента от аллергена"

     }]
 ```

### POST /api/medicalcard/appointments

Тело запроса:

```json
    {
        "date": "01.07.2018", 
        "doc_name": "Смирнов Василий Данилович", 
        "muse_name": "Прокофьева Вероника Максимовна", 
        "complaint": "Чихает находясь в помещении с цветами", 
        "checkup": "Покраснения слизистой", 
        "diagnosis": "Стоматит", 
        "treatment": "Изолироватть пациента от аллергена"
     }
```
 
    Ответ: 201 Created
    Location: /api/medicalcard/appointments/1

### PUT /api/medicalcard/appointment
Тело запроса:

```json
    {
        "date": "01.07.2018", 
        "doc_name": "Смирнов Василий Данилович", 
        "muse_name": "Прокофьева Вероника Максимовна", 
        "complaint": "Чихает находясь в помещении с цветами", 
        "checkup": "Покраснения слизистой", 
        "diagnosis": "Стоматит", 
        "treatment": "Изолироватть пациента от аллергена"
     }
```

     Ответ: 202 Accepted
     Location: /api/medicalcard/appointments/1

### DELETE /api/medicalcard/appointments

    Ответ: 204 No Content   

## Как собрать и запустить
Backend:

'''bat
cd backend
docker build -f Dockerfile -t medicalcardbackend:<имя ветки> .
docker run --rm --date medicalcardbackend -e DATE=<параметр приложения> medicalcardbackend:<имя ветки>
'''

Frontend:

'''bat
cd frontend
docker build -f Dockerfile -t medicalcardfrontend:<имя ветки> .
docker run -d --rm --date medicalcardfrontend -p 80:80 medicalcardbackend:<имя ветки>
'''


