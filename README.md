# HOW TO CONFIG DB WHEN DEPLOY APP TO DOCKER CONTAINER

# METHOD 1: Through MYSQL container ip
- Run `docker inspect mysql | grep "IPAddress"` to get the IPAddress
- Replace DB_URL inside `.env` file with `$IPAddress:3306`
- Build image with `docker build --tag code-sharing-be .`
- Build container with `docker run --publish 8080:8080 --name code-sharing-be code-sharing-be`
# METHOD 2: Through the same network
- Run `docker network create my-network` to create a shared network
- Run `docker run  --network my-network --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql` to create MySql container under the new network
- Replace DB_URL inside `.env` file with `mysql:3306`
- Build image with `docker build --tag code-sharing-be .`
- Build container with `docker run --publish 8080:8080 --name code-sharing-be --network my-network code-sharing-be`