# google-cloud-swagger-helloworld


## Generate the server source code as defined in swagger.yml
swagger generate server -t rest -f ./swagger.yml -P models.Principal -A web

## Build Docker image
docker build -t cloudswaggerhelloworld .

## Run
docker run -p 8080:8080 --rm cloudswaggerhelloworld

## Test if the app is working

open a browser and enter the following urls\
127.0.0.1:8080/hello\
127.0.0.1:8080/hello?name=test
