![Screenshot 2024-06-28 121705](https://github.com/Prakash333singh/LinkMinify/assets/110618721/ab134abc-c206-413e-89d2-a5d70d1c3951)

# URL Shortener

This project is a URL shortener service built using Golang, the Gin framework, Redis database, and Docker. The service provides various endpoints to shorten URLs, retrieve them, add tags, edit URLs, and delete them.

## Features

- **URL Shortening**: Shorten a long URL and get a shortened ID.
- **Retrieve URL**: Get the original URL from the shortened ID.
- **Add Tag**: Add a tag to a given URL.
- **Edit URL**: Edit the URL and its expiry.
- **Delete URL**: Delete the URL from the Redis database.

## Endpoints

- **POST** `/api/v1`

  - This API will shorten the URL and return the shortened ID.
  - Example request body:
    ```json
    {
      "url": "https://www.example.com",
      "short": "",
      "expiry": 40
    }
    ```

- **GET** `/api/v1/:shortID`

  - This API will retrieve the original URL from the shortened ID.

- **POST** `/api/v1/addTag`

  - This API will add a tag to the given URL.
  - Example request body:
    ```json
    {
      "url": "https://www.example.com",
      "tag": "sports"
    }
    ```

- **PUT** `/api/v1/:shortID`

  - This API will edit the URL and its expiry.
  - Example request body:
    ```json
    {
      "url": "https://www.newexample.com",
      "expiry": 60
    }
    ```

- **DELETE** `/api/v1/:shortID`
  - This API will delete the URL from the Redis database.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/Prakash333singh/LinkMinify
cd url_shortener

```

![Screenshot 2024-06-28 121642](https://github.com/Prakash333singh/LinkMinify/assets/110618721/d3d47bee-81d7-4895-a021-87d483106a60)
