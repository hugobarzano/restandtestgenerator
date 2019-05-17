# restandtestgenerator

1º- Define your desired service inside input folder as file.json

```

[{
  "name": "Miauuu as a service",
  "service": "/cats",
  "url": "http://localhost",
  "port":":8080",
  "body": {
    "name":"Wurst",
    "alias": "Mi bolita gatita bonita",
    "reina": true,
    "age": 1.5,
    "color": "Negro como mi alma"
  }
}]

```

2º- Execute restandtestgenerator main program

MacOS

```
 ./restandtestgenerator.darwin
 ```
 Linux
 ```
  ./restandtestgenerator.linux
 
 ```
 
 
 3º- Go to output and check your /cats service code  and CRUD test
    
    ├── goservice_cats
    │   ├── controller
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go
    │   ├── models
    │   └── mongo
    ├── nodeservice_cats
    │   ├── README.md
    │   ├── api
    │   ├── package.json
    │   └── server.js
    └── test_cats.sh
