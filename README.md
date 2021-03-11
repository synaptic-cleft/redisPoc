# Redis Proof of Concept

[Official Redis website](https://redis.io/clients#go) suggests using one of those three clients:  
- Go-Redis
- Radix
- Redigo

## Run Redis
```
docker pull redis
docker run --name redis-test-instance -p 6379:6379 -d redis

docker stats
docker stop CONTAINER_ID
```

`go run *.go`

## Tutorial references
Go-Redis  
[https://tutorialedge.net/golang/go-redis-tutorial/](https://tutorialedge.net/golang/go-redis-tutorial/)  
[https://blog.logrocket.com/how-to-use-redis-as-a-database-with-go-redis/](https://blog.logrocket.com/how-to-use-redis-as-a-database-with-go-redis/)

Redigo  
[https://www.alexedwards.net/blog/working-with-redis](https://www.alexedwards.net/blog/working-with-redis)