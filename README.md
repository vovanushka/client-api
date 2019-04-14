# Client API


## Installation

Use the docker to install client api.

```bash
docker build -t client-api .
doker run client-api
```
or use docker compose
```bash
docker-compose up
```
## Usage

File uploading
```
curl -X PUT \
  http://localhost:8080/port/upload \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F file=@path-to-file-here/mock.json
```
Port retrieving

```
curl -X GET \
  http://localhost:8080/port/port-id \
  -H 'Postman-Token: 86bf6f7d-597a-4a9c-a0c4-acd01b9f5c09' \
  -H 'cache-control: no-cache'
```
