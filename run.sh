docker rmi stori-tx-docker:latest
docker build -t stori-tx-docker:latest .
docker run --rm stori-tx-docker:latest ./main/main