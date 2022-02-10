# cache-img

## Работа с базой данных
* Создаем файлы для <a href='https://github.com/golang-migrate/migrate'>миграции</a>
  
    > migrate create -ext sql -dir ./schema -seq init
* Применяет миграции к базе данных
    > migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
  
* Отмена мирации
  > migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' down