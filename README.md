
# Go Consumer and Producer

Go application that generates orders and save to a queue and another application that read Orders from a queue


## Installation and Run

You can use docker-compose to run the application dependencies:

```bash
  docker-compose up -d
```
There are 3 instances in the docker compose, an instance of RabbitMQ, 
an instance of Prometheus and an instance of Grafana

create a sqlite instance with

```bash
sqlite3 orders.db
```
then create a table with

```bash
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price NOT NULL, PRIMARY KEY (id));
```
and finally, to run the Producer application 
```bash
go run cmd/producer/main.go
```
and for the Consumer application run
```bash
go run cmd/consumer/main.go
```
