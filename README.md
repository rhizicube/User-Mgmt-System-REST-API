# User-Mgmt-System-REST-API
Performed the CRUD operations for user management system 

Steps to follow :


 - Download the main.go file
 - Create a new folder. Place the main.go to that folder
 - Open terminal in the current directory
 - In terminal type "go mod init [directory_name]" to generate go module file.
 - Type "go get ." to install libraries
 - Now run "go run ." or "go run main.go"


## API Reference

#### Get all users

```http
  GET /users
```


#### Get user

```http
  GET /user/${id}
```

#### Create user

```http
  POST /user
```

#### Update user

```http
  PUT /user/${id}
```

#### Delete user

```http
  DELETE /user/${id}
```


