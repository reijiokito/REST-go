docker run -d --name mysql --privileged=true -e MYSQL_ROOT_PASSWORD="123456" -e MYSQL_USER="food_delivery" -e MYSQL_PASSWORD="123456" -e MYSQL_DATABASE="food_delivery" -p 3306:3306 mysql