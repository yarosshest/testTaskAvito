Проект упакован в docker запускается при помощи команд
```bash
docker compose build
docker compose up
```

api запускается по адресу http://localhost:60122/

swagger располагается по адресу http://localhost:60122/swagger/index.html#

Для предоставления конечного SQL файла с созданием всех необходимых таблиц в БД. Необходимо исполнить команду:
```bash
docker exec -it testtaskavito-postgres-1 pg_dump -U postgres --column-inserts --data-only testTaskAvito > backup_data.sql
```
В результате чего будет получен конченый SQL фалл backup_data.sql

# test task avito API
This is a test task avito server.

## Version: 1.0


### /Segment

#### DELETE
##### Summary:

Delete segment

##### Description:

Delete segment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | query | segment name | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 202 | ok |
| 400 | bad request |
| 404 | segment with this name not found |
| 500 | internal server error |

#### POST
##### Summary:

Add segment

##### Description:

Add segment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | query | segment name | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 | ok |
| 400 | bad request |
| 409 | segment with this name alredy exist |
| 500 | internal server error |

### /User

#### GET
##### Summary:

Get user

##### Description:

Get user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | user id | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | segments | [ string ] |
| 400 | bad request |  |
| 404 | user with this id not found |  |
| 500 | internal server error |  |

#### PUT
##### Summary:

Update user

##### Description:

Update user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query | user id | Yes | integer |
| add | body | Segments to add and del | Yes | [db.QueueUpdateUser](#db.QueueUpdateUser) |

##### Responses

| Code | Description |
| ---- | ----------- |
| 202 | ok |
| 400 | bad request |
| 404 | user with this id not found |
| 500 | internal server error |

### Models


#### db.QueueUpdateUser

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| add | [ string ] |  | No |
| dell | [ string ] |  | No |

