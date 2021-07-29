# go-mysql-crud
Sample crud operation using Golang and MySql

## API ENDPOINTS

### All Posts
- Path : `/posts`
- Method: `GET`
- Response: `200`

### Create Post
- Path : `/posts`
- Method: `POST`
- Fields: `title, content`
- Response: `201`

### Details a Post
- Path : `/posts/{id}`
- Method: `GET`
- Response: `200`

### Update Post
- Path : `/posts/{id}`
- Method: `PUT`
- Fields: `title, content`
- Response: `200`

### Delete Post
- Path : `/posts/{id}`
- Method: `DELETE`
- Response: `204`

## Required Packages
- Database
    * [MySql](https://github.com/go-sql-driver/mysql)
- Routing
    * [chi](https://github.com/go-chi/chi)

## Quick Run Project
First clone the repo then go to go-mysql-crud folder. After that build your image and run by docker. Make sure you have docker on your machine. 

```
git clone https://github.com/jamesattensure/go-mysql-crud.git

cd go-mysql-crud

docker-compose build

docker-compose up
```

## Kubectl commands
```
kubectl apply -f mysql-pvc-0.yaml
kubectl apply -f mysql-pvc-1.yaml
kubectl apply -f mysql-deployment.yaml
kubectl apply -f mysql-service.yaml
kubectl apply -f web-deployment.yaml
kubectl apply -f web-service.yaml
```

```
kubectl port-forward deployment/web 8005:8005
```