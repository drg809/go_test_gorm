# go_test_gorm

Test con golang y su ORM gorm basado en el tutorial https://tutorialedge.net/golang/golang-orm-tutorial/

Se usa MySQL cómo base de datos.

URL:
localhost:8081

Para ejecutar:

```
go build main.go
./main
```

Rutas:

| MÉTODO  | ENDPOINT | ACCIÓN |
| ------------- | ------------- | ------------- |
| GET  | /users  | Listado de usuarios.  |
| POST  | /user/{name}/{email}  | Crea un nuevo usuario.  |
| PUT  | /user/{name}/{email}  | Actualiza un usuario existente.  |
| DELETE  | /user/{email}  | Elimina un usuario existente.  |


Mis agradecimientos a los autores del mismo.

No me pertenecen los derechos de autor ni copyright del mismo, el único propósito de esto fue el aprendizaje del lenguaje.
