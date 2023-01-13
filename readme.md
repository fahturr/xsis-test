# XSIS Backend Developer Test
Project to fulfil Backend Developer position at **PT.Xsis Mitra Utama** 

## How To Run

- To run this project, first you need to `clone` this github repository
  or using command `git clone git@github.com:fahturr/xsis-test.git`


- After success cloning the project, the next step is duplicate `cmd/.env.example`
  file. And name the duplicate to `.env`


- Fill the empty variable at `.env` with your preferences

```dotenv
APP_PORT=9999       #fill with your desired number

SQL_DB=xsis-db      #fill with your desired database name
SQL_PORT=3306       #fill with your desired number
SQL_USER=xsis-user  #fill with your desired username
SQL_PASS=xsis-pass  #fill with your desired password
```

- And then you just need to run

```bash
docker compose -f deploy/docker-compose.yml --env-file=./cmd/.env up --build
```

- That's all :wink:

## API SPEC

### GET /v1/movies/list

Get List of movie

Response Body :

```json
{
  "status_code": 200,
  "message": "success fetch list movie",
  "payload": [
    {
      "id": 1,
      "title": "Siti nur baya",
      "description": "Pelem Siti nur baya",
      "rating": 5.1,
      "image": "Siti nur baya.png",
      "created_at": "2023-01-13 14:52:17",
      "updated_at": "2023-01-13 14:52:17"
    }
  ]
}
```

### GET /v1/movies/detail/:id

Get detail of movie by passing movies `id`

Response Body :

```json
{
  "status_code": 200,
  "message": "success fetch detail movie",
  "payload": {
    "id": 1,
    "title": "Harry Potter Askaban update",
    "description": "Pelem herri potter update",
    "rating": 7.5,
    "image": "gambar_heri update.png",
    "created_at": "2023-01-13 14:52:17",
    "updated_at": "2023-01-13 14:52:29"
  }
}
```

### POST /v1/movies/create

Create new movie

Request Body :

```json
{
  "title": "Siti nur baya",
  "description": "Pelem herri Siti nur baya",
  "image": "Siti nur baya.png",
  "rating": 5.1
}
```

Response Body :

```json
{
  "status_code": 201,
  "message": "success create movie"
}
```

### PUT /v1/movies/update/:id

Update movie by `id`

Request Body :

```json
{
  "title": "Siti nur baya",
  "description": "Pelem herri Siti nur baya",
  "image": "Siti nur baya.png",
  "rating": 5.1
}
```

Response Body :

```json
{
  "status_code": 200,
  "message": "success update movie"
}
```

### DELETE /v1/movies/delete/:id

Delete movie by `id`

Response Body :

```json
{
  "status_code": 200,
  "message": "success delete movie"
}
```