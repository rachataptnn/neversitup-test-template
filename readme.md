# Project Structure 

## OverAll

&nbsp;&nbsp;&nbsp;&nbsp;My teammate, who defined this structure, was inspired by clean architecture, focusing on the handler first. The flow goes through the service layer afterward, then to the repository.<br>
&nbsp;&nbsp;&nbsp;&nbsp;Most unit tests are written at the service layer, which contains the majority of the logic. Additionally, we conduct external testing using a custom tool that I designed with my QA team member for faster retesting and easier maintenance.

## Image of the structure
```
└── 📁product-master-service
    └── 📁.vscode
        └── launch.json
    └── 📁config
        └── app.go
        └── config.go
        └── google.go
        └── local.yaml
        └── mongo.go
        └── tracing.go
    └── 📁http
        └── 📁handler
            └── category.go
            └── subCategory.go
            └── healthz.go
        └── 📁router
            └── category.go
            └── subCategory.go
            └── healthz.go
    └── 📁model
        └── category.go
        └── subCategory.go
        └── pagination.go
        └── seo.go
    └── 📁repository
        └── category.go
        └── subCategory.go
    └── 📁service
        └── 📁error
            └── errors.go
            └── errors_test.go
        └── 📁string
            └── string.go
            └── string_test.go
        └── category.go
        └── category.go
        └── subCategory.go
    └── main.go
    └── go.mod
    └── go.sum
    └── .gitignore
    └── version
    └── README.md
    └── Makefile
```

## Structure Explaining

### .vscode
    Stores the configuration used for running and debugging the server.

### config 
    Stores constants that change their values depending on the environment, and constants for controlling program workflow.

### http 
- handler
    
        Stores code used for logging request and response details, and mapping errors retrieved from `service` to HTTP status codes.

- router 
        
        Stores routes indicating which function in `http/handler` to forward each route to.

### model
    Stores the structs for each scope.


### repository
    Stores the code used for connecting to or executing queries on databases.

### service
    Stores the code used for business logic, such as data modification or validation, and unit tests are written here.


