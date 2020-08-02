# Breaking bad CRUD api

A simple api proxy for [Breaking Bad](https://breakingbadapi.com/documentation) api. The application uses a sqlite database on a localhost and is only used in the get, update operations. 

The structure is as follows. Each subpackage (except under api) represents an interface can contain its multiple implementations (e.g mock, sqlite, mongo, redis etc). 

```
├── db
└── pkg
    ├── api                 // all handlers, middleware and http stuff
    │   ├── characters
    │   ├── episodes
    │   ├── helpers
    │   └── quotes
    ├── clients             // high level package for everything external to app
    │   └── breakingbad
    └── store               // high level package for persistance layer
        ├── characters      // each subpackage repesents an interface and its multiple impls
        ├── episodes
        └── quotes
```

## Endpoints
```
GET     /characters         Get all the characters
GET     /characters/{id}    Get one character
PUT     /characters/{id}    Update a character in local db

GET     /episodes           Get all the episodes
GET     /episodes/{id}      Get one episode
PUT     /episodes/{id}      Update an episode in local db

GET     /quotes             Get all the quotes
GET     /quotes/{id}        Get one quote
PUT     /quotes/{id}        Update a quote in local db
```

## Test
Contains only one test

## Running

__Production__ To run against db boot up the sqlite docker container
- Boot up docker container `docker-compose up -d`
- Run application `go run main.go`


__Mock__ To run against mocks.
- Comment [line 25](https://github.com/PrakharSrivastav/go-api-demo/blob/2b9c57767a1c6e806d7936a529d698b3672becda/pkg/app.go#L25) and uncomment [line 27](https://github.com/PrakharSrivastav/go-api-demo/blob/2b9c57767a1c6e806d7936a529d698b3672becda/pkg/app.go#L27) 
- Run application `go run main.go`
