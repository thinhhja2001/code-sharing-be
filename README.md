# HOW TO CONFIG DB WHEN DEPLOY APP TO DOCKER CONTAINER

# METHOD 1: Through MYSQL container ip
Get the IPAddress
```
docker inspect mysql | grep "IPAddress"
```

Replace DB_URL inside `.env` file with `$IPAddress:3306`

Build image with 
```
docker build --tag code-sharing-be .
```

Run app with
```
docker run --publish 8080:8080 --name code-sharing-be code-sharing-be
```
# METHOD 2: Through the same network
Create a shared network 
```
docker network create my-network
``` 

Create MySql container under the new network
```
docker run  --network my-network --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql
```
Replace DB_URL inside `.env` file with `mysql:3306`

Build image app image 
```
docker build --tag code-sharing-be .
```
Run app 
```
docker run --publish 8080:8080 --name code-sharing-be --network my-network code-sharing-be
```