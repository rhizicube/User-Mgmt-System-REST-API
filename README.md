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

```
  GET /users
```


#### Get user

```
  GET /user/${id}
```

#### Create user

```
  POST /user
```

#### Update user

```
  PUT /user/${id}
```

#### Delete user

```
  DELETE /user/${id}
```


