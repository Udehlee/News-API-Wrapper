## News-API-Wrapper



## Overview
News-API Wrapper offers a straightforward and efficient method for accessing and caching News infomation.Built in Go, it  leverages on Redis to cache News data, significantly improving response times and reducing the load on the external API.

## Features
- Efficient  Retrieval of news information from the News API.
-  Uses Redis to cache API responses, minimizing repeated API calls and speeding up data retrieval.
-  Handles both cache hits and misses effectively, ensuring up-to-date News information 

## Technologies Used
- Go
- Redis driver(github.com/redis/go-redis/v9)
- docker 
- News API (https://newsapi.org/)

- API Endpoint
```sh
GET /api/news

```

## API Endpoint validation
- You can use postman to test endpoints.
 ```sh
 http://localhost:8080/api/news?q=apple

 ```
 ## Example response
 ```json
 {
    "status": "ok",
    "totalResults": 46078,
    "articles": [
        {
            "author": null,
            "publishedAt": "1970-01-01T00:00:00Z"
        },
        {
            "author": "Sarah Fielding",
            "publishedAt": "2024-07-11T12:30:26Z"
        },
    ]
 }
 ```

## Prerequisites
- Ensure you have Go installed
- Ensure you use Docker to run a Redis instance
- Ensure you have your api key details configured in your environment.

## Installation
- Clone this repository to your local machine:
```sh
git clone https://github.com/Udehlee/News-Api-Wrapper.git
```
- Navigate to the project directory:

- cd News-Api-Wrapper

- Run the server:

```sh
go run main.go
```

- The server will be listening on port 8080.

## Check cache key
```sh
docker exec -it redis-container redis-cli
```
```sh
KEYS *
```

