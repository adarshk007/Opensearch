1️⃣ Project Structure (Production-ready)

```
tutorial/
├── config/
│   ├── config.go
│   ├── config.dev.yaml
│   └── config.prod.yaml
├── client/
│   └── opensearch_client.go
├── service/
│   └── search_service.go
├── handler/
│   └── search_handler.go
├── router/
│   └── router.go
├── main.go
├── go.mod
└── go.sum
```

Separation of concerns:
```
config → load environment config
client → singleton OpenSearch client
service → business logic (search, index)
handler → HTTP request handling
router → Gin routes
```
6️⃣ Run everything
From your project root:
docker compose up --build

✅ You should see logs like:
my-gin-api  | Starting server on port 8080
opensearch  | ...

7️⃣ Test it from macOS host

curl "http://localhost:8080/api/search?index=my-index"
You should get an OpenSearch response (even if no docs yet).

8️⃣ Debug tip (on macOS)

If you want to test inside the container:

docker exec -it my-gin-api sh
apk add curl
curl http://opensearch:9200
If that works, your networking is correct.


Option 1: Force rebuild
When you change your code or Dockerfile:
docker compose up -d --build
Or, if you want a clean rebuild from scratch (no cache):
docker compose build --no-cache
docker compose up -d


## DATA

Schema:

```
PUT products
{
  "mappings": {
    "properties": {
      "id":         { "type": "keyword" },
      "name":       { "type": "text" },
      "category":   { "type": "keyword" },
      "price":      { "type": "float" },
      "feature":    { "type": "keyword" },
      "created_at": { "type": "date" }
    }
  }
}
```

Data Dump:

```
POST products/_doc/1
{
  "id": "1",
  "name": "Apple iPhone 15",
  "category": "electronics",
  "price": 1299.99,
  "feature": "premium",
  "created_at": "2025-10-01T10:00:00Z"
}

POST products/_doc/2
{
  "id": "2",
  "name": "OnePlus 12",
  "category": "electronics",
  "price": 899.99,
  "feature": "standard",
  "created_at": "2025-10-02T11:00:00Z"
}

POST products/_doc/3
{
  "id": "3",
  "name": "Google Pixel 9",
  "category": "electronics",
  "price": 999.99,
  "feature": "premium",
  "created_at": "2025-10-03T12:00:00Z"
}

POST products/_doc/4
{
  "id": "4",
  "name": "Sony WH-1000XM5",
  "category": "audio",
  "price": 499.99,
  "feature": "premium",
  "created_at": "2025-10-04T09:00:00Z"
}
```

Search Data:

✅ Option 1 — Search in a specific index
If you want all documents from the products index, just specify the index in the endpoint:
```
GET products/_search
{
  "query": {
    "match_all": {}
  }
}
```

👉 That’s the most common form.
It returns all docs in products.
✅ Option 2 — Search multiple indices
You can also query multiple indices at once:
```
GET products,logs,textsearch/_search
{
  "query": {
    "match_all": {}
  }
}
```

✅ Option 3 — Filter by index using a _index query
If you’re searching across all indices but want to filter only those where _index = products:
```
GET _search
{
  "query": {
    "term": {
      "_index": "products"
    }
  }
}
```
That means:
Search across everything, but return only documents whose index is “products”.
✅ Verify index exists
If it still fails, check the index name:
```
GET _cat/indices?v
```
You should see:
health status index     uuid                   pri rep docs.count store.size
green  open   products  KJHkLkASD23231         1   1   4          12kb

To Test API:
```
curl "http://localhost:8080/api/search?index=products"
```
