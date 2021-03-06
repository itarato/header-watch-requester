# Header Watch - Requester

End actor for the Header Watch service to collect request information.

Example:

- From one shell: ```go run main.go```
- From another shell: ```curl -s -D - -d '{"locations":[{"url":"http://google.com","id":"LOCATION_STORAGE_ID"}]}' http://localhost:8080```

Endpoints
---------

POST
----

- *Path:* ```/```
- *Post body:* JSON:

```
{
  "locations": [
    {"url": "http://???"}
  ]
}
```

- *Returns:*

```
{
  "locations": [
    {
      "url": "...", // string
      "id": "...", // string, location storage id
      "status_code": "...", // int
      "protocol": "...", // string
      "headers": {
        NAME: [VALUES]
      }
    }
  ]
}
```
