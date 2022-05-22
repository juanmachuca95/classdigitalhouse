# Cache en Golang

Implementaremos una biblioteca de libros, en la cual podremos consultar libros utilizando un sistema de cache para agilizar el proceso y evitar consultas repetitivas.

1. Conexión a la base de datos en este caso, usaremos el motor de base de datos mysql. 

Utilidades: 
 - Driver MySQL GoDriver
 - https://github.com/go-sql-driver/mysql  
  ```go install github.com/go-sql-driver/mysql``` 

 - Examples: http://go-database-sql.org/index.html
 - Faker: para crear datos ficticios de prueba  
  ```go get -u github.com/bxcodec/faker/v3```
 - Gorilla Mux: 
 ```go get -u github.com/gorilla/mux```
 - Libreria para la variables de entorno:
 ```go get github.com/joho/godotenv```

    

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


Una vez programada la api tendriamos podemos implementar via get o post la petición al servidor:

#### Metodo GET 
  ```curl http://localhost:8080/getbook/Rayuela``` 

#### o POST
  ```curl -X POST http://localhost:8080/getbook -H 'Content-Type: application/json' -d '{"book":"Rayuela"}'```

## Cache Implementación

```go
type Cache interface {
	GetBook(string) (models.Book, error)
	AddBook(models.Book) bool
}

type CacheMemory struct {
	cache map[string]models.Book
	l     sync.Mutex
}

func NewCacheMemory() Cache {
	return &CacheMemory{
		cache: make(map[string]models.Book),
	}
}

func (c *CacheMemory) GetBook(bookName string) (models.Book, error) {
	results, exists := c.cache[bookName]
	if !exists {
		return models.Book{}, errors.New("Este libro o existe en cache")
	}
	return results, nil
}

func (c *CacheMemory) AddBook(book models.Book) bool {
	if len(c.cache) == 0 {
		c.l.Lock()
		c.cache = make(map[string]models.Book)
		c.cache[book.Book] = book
		c.l.Unlock()
		return true
	}

	c.l.Lock()
	c.cache[book.Book] = book
	c.l.Unlock()
	return true
}

```

Cuando el dato no se encuentra en cache: 

```curl
curl -X POST http://localhost:8080/getbook -H 'Content-Type: application/json' -d '{"book":"Don Quijote"}'

Datos desde base de datos 👎
Libro: Don Quijote
Autor:Miguel de Cervantes
Tiempo de busqueda ⏰  796.397µs
```

Cuando el dato está en cache:
```curl
curl -X POST http://localhost:8080/getbook -H 'Content-Type: application/json' -d '{"book":"Don Quijote"}'

Datos desde cache 👍
Libro: Don Quijote
Autor:Miguel de Cervantes
Tiempo de busqueda ⏰ 83.774µs%  
```