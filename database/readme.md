1 - Build docker image

docker build . -t go_db

2 - Run Docker image

docker run -p 54321:5432 go_db