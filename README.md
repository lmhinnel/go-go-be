## RUN

Inside the newly cloned project folder run the following command:

```bash
$ docker compose up -d
```

## API

### GET ALL /api/v1/students

```bash
curl --location 'http://localhost:3000/api/v1/students?limit=10&offset=0'
```

### GET DETAIL /api/v1/students/3

```bash
curl --location 'http://localhost:3000/api/v1/students/3'
```

### POST /api/v1/students

```bash
curl --location 'http://localhost:3000/api/v1/students' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Perry3",
    "last_name": "Platypus",
    "phone": "0322456781",
    "email": "perry_x_sp2@gmail.com",
    "birthday": "2000-01-08T07:00:00+07:00",
    "gender": 1,
    "university": "Tech University",
    "address": "456 Elm Street"
}'
```

### PUT /api/v1/students/3

```bash
curl --location 'http://localhost:3000/api/v1/students/3' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 3,
    "first_name": "Perry4",
    "last_name": "Platypus",
    "phone": "0322456781",
    "email": "perry_x_sp2@gmail.com",
    "birthday": "2000-01-08T07:00:00+07:00",
    "gender": 1,
    "university": "Tech University",
    "address": "456 Elm Street"
}'
```

### DELETE /api/v1/students/3

```bash
curl --location 'http://localhost:3000/api/v1/students/3' \
--request DELETE
```
