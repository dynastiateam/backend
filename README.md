## Build
`docker-compose build`

## Run
`docker-compose up`

## Endpoints
`/users` -> http://localhost:9001

`/auth` -> http://localhost:9002

## API 

Request:
```
POST /auth/v1/login

{"email":"test@test.com", "password":"qwerty"}
```
Response:
```
#success
HTTP/1.1 200 OK

{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTgsIk5hbWUiOiJKb2huIERvZSIsIlJvbGUiOjQsImF1ZCI6ImR5bmFwcCIsImV4cCI6MTU3NzM2MzU2MiwiaWF0IjoxNTc3Mjc3MTYyLCJpc3MiOiJhdXRoLmR5bmFwcCJ9.asGix3XEgR0CwlRYZYyEYyqPcptPp04OjYZojlYBpyI",
    "refresh_token": "395a6fac-3aee-4891-8ed9-ec5546b8777c"
}

# Any internal error
HTTP/1.1 500 Internal Server Error
{"error":"error text"}
```

Request:
```
POST /auth/v1/refresh

{"token":"395a6fac-3aee-4891-8ed9-ec5546b8777c"}
```
Response:
```
#success
HTTP/1.1 200 OK

{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTgsIk5hbWUiOiJKb2huIERvZSIsIlJvbGUiOjQsImF1ZCI6ImR5bmFwcCIsImV4cCI6MTU3NzM2MzU2MiwiaWF0IjoxNTc3Mjc3MTYyLCJpc3MiOiJhdXRoLmR5bmFwcCJ9.asGix3XEgR0CwlRYZYyEYyqPcptPp04OjYZojlYBpyI",
    "refresh_token": "395a6fac-3aee-4891-8ed9-ec5546b8777c"
}

# Any internal error
HTTP/1.1 500 Internal Server Error
{"error":"error text"}
```


Request:
```
POST /users/v1/register

{"email":"test@test.com","first_name":"John","last_name":"Doe","apartment":1,"phone":"380671234567","password":"qwerty", "building_id": 2}
```
Response:
```
#success
HTTP/1.1 200 OK

{
    "id": 51,
    "email": "test2@test.com"
}

# Any internal error
HTTP/1.1 500 Internal Server Error
{"error":"error text"}
```

Request:
```
GET /users/v1/user/{id}

```
Response:
```
#success
HTTP/1.1 200 OK

{
    "id": 18,
    "first_name": "John",
    "last_name": "Doe",
    "phone": "380671234567",
    "email": "test@test.com"
}

# Any internal error
HTTP/1.1 500 Internal Server Error
{"error":"error text"}
```