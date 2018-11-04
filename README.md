# crypto-wallet

This project aims to provide an api a own wallet to follow purchase and sale of crypto currency.

Useful infos:
-------------

Price show full price for an operation.
Ex:
- 100 for quantity 2 means 1 unit is equal to 50
- 100 for quantity 0.5 means 1 unit is equal to 200

Euro price for a crypto is the price for the crypto which is use to buy another
Ex:
- If I buy 2 ether with 1 bitcoin, euro price is the price of the bitcoin

Operation API:
--------------

Create operation:

```
curl localhost:8080/api/operation -X POST -d '{"quantity": 1, "currency": 1, "description": "", 
"buy_order": {"price": 1, "euro_price": 1, "date": "YYY-MM-DD", "currency": 1}, 
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
"buy_order": {"price": 3 , "currency": 3, "euro_price": 0.5, "date": "2018-01-20"}}'
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

Get crypto internal code:

```
curl localhost:8080/api/cryptos
```

Get budgets:

```
curl localhost:8080/api/budgets

curl localhost:8080/api/budget/2
```

Add euro:

```
curl localhost:8080/api/budget/euro/-10 -X POST
```