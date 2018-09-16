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

* Buy bitcoin with euro:
```
curl localhost:8080/api/operation -X POST -d '{"quantity": 1, "currency": 1, "description": "",
"buy_order": {"price": 1, "euro_price": 1,  "currency": 2}}'
```

* Buy bitcoin with ether:

```
curl localhost:8080/api/operation -X POST -d '{"quantity": 2, "currency": 1, "description": "",
"buy_order": {"price": 3 , "currency": 3, "euro_price": 0.5}}'
```

* Sell ether:

```
curl localhost:8080/api/operation -X POST -d '{"quantity": 1, "currency": 3, "description": "",
"sell_order": {"price": 2 , "currency": 2}, "parent_id": "5b9e9c80c15df93d3d674b14"}'
```

Get all operations:

```
curl localhost:8080/api/operations
```

Get one operation:

```
curl localhost:8080/api/operation/{id}
```