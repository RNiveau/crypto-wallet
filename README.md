# crypto-wallet

This project aims to provide an api a own wallet to follow purchase and sale of crypto currency.

Operation API:
--------------

Create operation:

```
curl localhost:8080/api/operation -X POST -d '{"quantity": 1, "currency": 1, "description": "", 
"buy_order": {"total_price": 1, "price": 1, "euro_price": 1, "date": "", "currency": 1}, 
"parent_id": "5b82cf8dc15df939fb776bce" }'
```

Get all operations:

```
curl localhost:8080/api/operations
```

Get one operation:

```
curl localhost:8080/api/operation/{id}
```