[![Go CI](https://github.com/bleakview/goapiexample/actions/workflows/go_ci.yml/badge.svg)](https://github.com/bleakview/goapiexample/actions/workflows/go_ci.yml)   [![publish to docker registry](https://github.com/bleakview/goapiexample/actions/workflows/push_to_docker_hub.yml/badge.svg)](https://github.com/bleakview/goapiexample/actions/workflows/push_to_docker_hub.yml)

# Golang Api Example
  
<img src="https://bleakview.github.io/git/goapiexample/images/goapiexample.jpg" alt="swgger example image" width="800"/>

  
This is an example written in golang for a start point on how to write a Gorilla with Gorm.  
It supports golang testing and you can test api with swagger via 
```
http://localhost:4000/docs/
```
If you would like to test with docker comment was given below.
```
docker run --name goapiexample -d -p 4000:4000 bleakview/goapiexample:1.0.0 
```
  
If you directly run docker container it will use sqlite. The docker compose file in this repository will invoke the image with mysql with the command given below.  
```
docker-compose up
```
There are two environment variables that you can use with this image.  
**DB_CONNECTION_URL** MySql DSN string for Gorm  
**PORT** Port number  
    
On very change code test will be run and a new docker image will bu pushed to  
[https://hub.docker.com/r/bleakview/goapiexample](https://hub.docker.com/r/bleakview/goapiexample).


