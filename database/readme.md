- Build docker image

docker build . -t go_db

- Run Docker image

docker run -p 54321:5432 go_db

- Run with Docker Compose

docker-compose up

- Connect PGAdmin4 to local database

    1 - Get Running Containers with docker ps
    2 - Copy Container ID
    3 - Execute the follow command to get private ip address: 
    
    docker ps
    docker inspect [CONTAINERID] | grep IPAddress

