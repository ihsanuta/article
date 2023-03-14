# article

## running DB and Redis Docker
```
docker-compose up -d
```

## running migration
```
make migrate
```

## running test
```
make tests
```

## running application
```
go run .
```

## cURL API
### Create Article
```
curl --location --request POST 'localhost:8880/api/v1/articles' \
--header 'Content-Type: application/json' \
--data-raw '{
    "author":"saya",
    "title":"jejak",
    "body":"test"
}'
```

### Get Articles
```
curl --location --request GET 'localhost:8880/api/v1/articles?author=saya&query=jak'
```