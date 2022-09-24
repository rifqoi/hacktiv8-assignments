# Latihan Clean Architecture

## User

### Get All Users

- Method : GET
- URL : `/users`
- Response body :

```json
{
	"status": 201,
	"message": "succesfully get all the users!",
	"payload": [
		{
			"id": 1,
			"username": ""
		},
		{
			"id": 2,
			"username": ""
		}
	]
}
```

### Get User by ID

- Method : GET
- URL : `/users/{id}`
- Response body :

```json
{
	"status": 201,
	"message": "succesfully get the user!",
	"payload": {
		"id": 1,
		"username": ""
	}
}
```

### Create new user

- Method : POST
- URL : `/users`
- Request body :

```json
{
	"username": "",
	"password": ""
}
```
