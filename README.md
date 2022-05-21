# Cache en Golang

Implementaremos una biblioteca de libros, en la cual podremos consultar libros utilizando un sistema de cache para agilizar el proceso y evitar consultas repetitivas.

1. Conexión a la base de datos en este caso, usaremos el motor de base de datos mysql. 

Utilidades: 
 - Driver MySQL GoDriver: https://github.com/go-sql-driver/mysql

 - Examples: http://go-database-sql.org/index.html
 - Faker: para crear datos ficticios de prueba  
  ```go get -u github.com/bxcodec/faker/v3```

    

Requisitos: 
- Previamente deberemos tener instalado mysql en nuestro entorno local
- Creamos una base de datos - Biblioteca

```sql
mysql -u root -p
// ingresa tu password
```

```sql
CREATE DATABASE biblioteca;
```

```sql
use biblioteca;
```
  - Creamos una tabla ```libros``` en la base de datos biblioteca.

```sql
CREATE TABLE libros(id int NOT NULL primary key AUTO_INCREMENT, book varchar(255) NOT NULL, author varchar(255) NOT NULL, created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL);
```