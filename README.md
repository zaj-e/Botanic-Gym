## Things our REST API need to do
- Support HTTP/2
- Automatically install TLS certificates (echo)
-

## Tools that help building a REST API
------

### Web Server
#### Requirements
HTTP server should have a default timeout value.
#### Examples
- [net/http lib](https://golang.org/pkg/net/http/) has a http server implementation
- [fasthttp](https://github.com/valyala/fasthttp) lib has a blazing fast http server
- [NGINX](https://www.nginx.com/) is a production ready web server
------

### Migrations
#### Examples
- [Golang-Migrate](https://github.com/golang-migrate/migrate)
- [SQl-Migrate](https://github.com/rubenv/sql-migrate)
- [Goose](https://github.com/pressly/goose)

------
### HTTP ROUTERS
#### Context
[httprouter is fast. gorilla mux is slower. But then again we are talking about routing, and unless all your microservice does is routing then it won't be the bottleneck. Database access or whatever other processing is likely to take longer. So it is also a matter of what extra features you need when choosing a plain router vs a framework vs the stdlib.](https://www.reddit.com/r/golang/comments/a3qcid/httprouter_chi_gin_gorillamux/eb8ars5?utm_source=share&utm_medium=web2x)

#### Examples
Some old benchmarls: https://github.com/julienschmidt/go-http-routing-benchmark#conclusions
- [Mux](https://github.com/gorilla/mux)
- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [net/http lib has a default mux](https://golang.org/pkg/net/http/)

------
### Database Drivers
#### Examples
- [pq (postgres)](https://github.com/lib/pq)
- [pgx (postgres)](https://github.com/jackc/pgx)

------

### Database Libraries
#### Examples
### SQLX
Features:
- f1: Simple mapping to structs(https://www.reddit.com/r/golang/comments/fzs8p1/what_orm_are_you_use/fn5zry1?utm_source=share&utm_medium=web2x). 

Usage:
- Database connection (https://www.reddit.com/r/golang/comments/fzs8p1/what_orm_are_you_use/fn6cwz5?utm_source=share&utm_medium=web2x)

##### (database/sql)[https://golang.org/pkg/database/sql/]

------

### ORM
#### Features
In contrast to a sql library, an ORM gives you a unified syntax, which it's loosely coupled from a particular database. It would be a pain if you are selling a product to enterprises and you had to support oracle and Sql server at the same time. 
Some ORM implementation select all row and and do filter internally. For example, if you do select on table A where Id = 1, a bad ORM implementation will select all row in table A (which cause large memory usage on your application) and then filter the result internally. That is just select operation. If you use a lot of join, it would get worse.

#### Examples
##### GORM
Pros: Saves you the verbosity of SQL + manually mapping fields, helps with managing migrations (https://www.reddit.com/r/golang/comments/a6yo8k/do_you_use_database_in_your_go_project/ebz0s5f?utm_source=share&utm_medium=web2x)

Cons:
Complicated queries are better handled by sql directly. (https://www.reddit.com/r/golang/comments/fzs8p1/what_orm_are_you_use/fn65ugr?utm_source=share&utm_medium=web2x)


##### XORM
##### [go-pg](https://github.com/go-pg/pg)
#### [SQLBOILER](https://github.com/volatiletech/sqlboiler)
It's sort of an ehancement of stdlib and a superset of sqlx. It does:
- Query building
- Struct binding
- Relationship management/querying
Selling post: https://www.reddit.com/r/golang/comments/aq7a59/how_is_sqlboiler_better_than_stdlib_sql_or_sqlx/emwkd59?utm_source=share&utm_medium=web2x

------

### Complete web Frameworks
- Beego
- Buffalo

------

### Micro web Frameworks
#### Examples
##### Echo
Pros:
Cons:
- Middleware for compression implementation is broken. For each request, echo server will create new gzip writer (https://github.com/labstack/echo/blob/f867058e3ba4fa4504e3ca976b31012db3ce1555/middleware/compress.go#L72)(https://blog.klauspost.com/gzip-performance-for-go-webservers)

- Default server has no timeout values https://github.com/labstack/echo/issues/590


