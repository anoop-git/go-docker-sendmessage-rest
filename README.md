## BUILD

This applcation calls the McKesson producer in the dockerized-kafka_kafka network and sends the message to the rest service exposed to port 8081
```
go build
docker-compose build
docker-compose up
```

