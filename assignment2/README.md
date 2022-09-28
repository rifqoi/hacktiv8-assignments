# Assignment 2

## Order

### Get All Orders

- Method : GET
- URL : `/orders`
- Params:
  - limit : get the orders within this limit (default: 5)
- Response body :

```json
{
    "status": 200,
    "message": "SUCCESS_GET_ORDERS",
    "payload": [
        {
            "order_id": number,
            "customer_name": string,
            "ordered_at": "2022-10-02T22:03:04+07:00",
            "items": [
                {
                    "item_id": number,
                    "item_code": string,
                    "description": string,
                    "quantity": number
                },
            ]
        }
    ]
}
```

### Create new order

- Method : POST
- URL : `/orders`
- Request body :

```json
{
	"ordered_at": "2022-10-02T15:03:04.0Z",
	"customer_name": string,
	"items": [
		{
			"item_code": string,
			"description": string,
			"quantity": number
		},
	]
}
```

### Update order

- Method : PUT
- URL : `/orders`
- Request body :

```json
{
	"order_id": number,
	"ordered_at": "2022-11-02T15:03:04.0Z",
	"customer_name": string,
	"items": [
		{
			"item_code": string,
			"description": string,
			"quantity": int
		}
	]
}
```

### Delete order by id

- Method : PUT
- URL : `/orders/{id}`
- Response body :

```json
{
	"status": 200,
	"message": "SUCCESS_DELETE_ORDER",
	"payload": {
		"order_id": number,
		"customer_name": string
	}
}
```
