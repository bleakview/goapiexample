# GOApiExample

Hi ! This is an example written in golang for a start point on how to write a Gorilla with Gorm.
It supports testing with jest and you can test api with swagger via http://localhost:\<port>/docs/.
If you invoke docker with docker run --name goapiexample -d -p 4000:4000 bleakview/goapiexample:1.0.0 you can reach swagger with http://localhost:4000/docs/
If you directly run docker container it will use sqlite.
You can use docker compose it will start app with mysql.
You can change default port with PORT environment variable.
You can change default database with mysql with the following syntax.
\<username>:\<password>@tcp(\<ip>:\<port>)/\<database>
the environment variable is DB_CONNECTION_URL.

On very change code test will be run and a new docker image will bu pushed to
[https://hub.docker.com/r/bleakview/goapiexample](https://hub.docker.com/r/bleakview/goapiexample).
