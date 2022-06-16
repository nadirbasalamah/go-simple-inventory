# pull the mariadb image from the Docker Hub
docker pull mariadb
# create a mariadb container
docker run -itd --name some-mariadb -v /your/volume/mariadb-volume:/var/lib/mysql -e MARIADB_ROOT_PASSWORD=password -p 3306:3306 mariadb:latest
# access the container to create the database
docker exec -it some-mariadb mysql -u root -p
