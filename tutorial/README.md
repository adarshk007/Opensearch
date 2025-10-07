# 🚀 OpenSearch Portal (Incoming) — Go + React + Docker

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white)
![OpenSearch](https://img.shields.io/badge/OpenSearch-2.x-005EB8?logo=opensearch&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)


 It's a Running Opensearch Go tutorial - Production Level Code [Helm Chart will be Included Soon]

---

## ✨ Features [In Progress 🧭]

> 🧭 A lightweight **OpenSearch Management Portal** built with **Go (backend)** and **React (frontend)**.  
> Supports **Product Search**, **Log Analytics**, and **Text Search** — with pagination, filters, and dashboard-ready APIs.

- 🔍 Full-text, keyword & structured search via OpenSearch  
- ⚙️ API support for `from`, `size`, sorting, and filters  
- 🧩 Modular Go services: Product / Logs / TextSearch  
- 💅 React + Tailwind UI (dashboard)  
- 🐳 Docker Compose setup with OpenSearch + Dashboards  
- 🔐 Secure TLS-ready configuration  
- 🚀 Easy local dev setup with `docker-compose up`

---

## 🏷️ Tags

<!-- Insert generated badges here -->
<!-- Example output from generate_tags.go -->

[![OpenSearch](https://img.shields.io/badge/-OpenSearch-blue?style=flat&logo=opensearch)](https://github.com/opensearch-project)
[![Helm](https://img.shields.io/badge/-Helm-0f0f0f?style=flat&logo=helm&logoColor=white)](https://helm.sh)
[![Golang](https://img.shields.io/badge/-Golang-00ADD8?style=flat&logo=go)](https://golang.org)
[![FulltextSearch](https://img.shields.io/badge/-FulltextSearch-lightgrey?style=flat)](https://en.wikipedia.org/wiki/Full-text_search)
[![SearchEngine](https://img.shields.io/badge/-SearchEngine-lightgrey?style=flat)](https://en.wikipedia.org/wiki/Search_engine)
[![Docker](https://img.shields.io/badge/-Docker-2496ED?style=flat&logo=docker)](https://www.docker.com)
[![React](https://img.shields.io/badge/-React-61DAFB?style=flat&logo=react)](https://reactjs.org)
[![Tailwind](https://img.shields.io/badge/-TailwindCSS-38B2AC?style=flat&logo=tailwindcss)](https://tailwindcss.com)
[![Analytics](https://img.shields.io/badge/-Analytics-lightgrey?style=flat)](https://en.wikipedia.org/wiki/Analytics)
[![Microservices](https://img.shields.io/badge/-Microservices-lightgrey?style=flat)](https://en.wikipedia.org/wiki/Microservices)
[![Observability](https://img.shields.io/badge/-Observability-lightgrey?style=flat)](https://en.wikipedia.org/wiki/Observability)
[![Elastic](https://img.shields.io/badge/-Elastic-lightgrey?style=flat&logo=elastic)](https://www.elastic.co)
[![DevTools](https://img.shields.io/badge/-DevTools-lightgrey?style=flat)](https://en.wikipedia.org/wiki/Developer_tools)
[![OpenSource](https://img.shields.io/badge/-OpenSource-lightgrey?style=flat)](https://opensource.org)


---

## ⚡ Getting Started

### Prerequisites
- Go ≥ 1.23  
- Docker + Docker Compose  
- Node.js (if using frontend)


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

![alt text](/tutorial/resources/static/result.png)

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

⚙️ Safe practice for teams / CI / Docker
In your .dockerignore and .gitignore, do not exclude these:
```
!go.mod
!go.sum
```
And for clean reproducible builds:
```
go mod tidy
go mod verify
go build
```
🧩 TL;DR
✅ Commit and push both:
```
go.mod
go.sum
```


💡 The Three Cluster Health States
![alt text](/tutorial/resources/static/image.png)

Color	Meaning	What it indicates

🟢 Green	All primary and replica shards are active.	Perfect state; full redundancy and fault tolerance.

🟡 Yellow	All primary shards are active, but one or more replica shards are missing.	Data is safe (no data loss), but not fully redundant — if a node fails, some data may become unavailable.

🔴 Red	One or more primary shards are missing or unassigned.	Some data is unavailable — serious issue.

🔍 Example
If you run:
```
GET _cluster/health
You might see:
{
  "cluster_name": "docker-cluster",
  "status": "yellow",
  "number_of_nodes": 1,
  "active_primary_shards": 5,
  "active_shards": 5,
  "unassigned_shards": 5
}
```

👉 This means:

You have 1 node.

Every index has 1 primary shard + 1 replica (by default).

The replica shards can’t be allocated — because there’s no second node.

⚙️ Why It Happens
In dev setups (like your Docker example), you typically have only one OpenSearch node:
replica: 1
But there’s nowhere to put the replica. So OpenSearch says “yellow”.

🧭 How to Fix It
Option 1: In single-node dev setup
Tell OpenSearch to not expect replicas:
PUT _settings
{
  "index": {
    "number_of_replicas": 0
  }
}
or in your Docker config:
environment:
  - discovery.type=single-node
✅ That will make your cluster green in single-node mode.

Option 2: For production
Add more nodes (so replicas can be assigned):
GET _cat/nodes
Make sure at least 2 nodes exist so replicas can be distributed.

🧩 TL;DR
🟢 Green: All good
🟡 Yellow: Data OK, replicas missing
🔴 Red: Data missing