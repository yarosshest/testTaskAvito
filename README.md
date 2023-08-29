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

