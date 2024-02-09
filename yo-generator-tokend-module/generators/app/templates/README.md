# <%= serviceName %>

## Description

<%= serviceDescription %>

## Install

  ```
  git clone <%= packageName %>
  cd <%= serviceName %>
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  <%_ if (useDB) { _%>
  ./main migrate up
  <%_ } _%>
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.

Note: if you are using Gitlab for building project `docs/spec/paths` folder must not be
empty, otherwise only `Build and Publish` job will be passed.  

## Running from docker 
  
Make sure that docker installed.

<%_ if (handleHTTP) { _%>
use `docker run ` with `-p 8080:80` to expose port 80 to 8080
<%_ } _%>

  ```
  docker build -t <%= packageName %> .
  docker run -e KV_VIPER_FILE=/config.yaml <%= packageName %>
  ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
<%_ if (useDB) { _%>
* Launch the service with `migrate up` command to create database schema
<%_ } _%>
* Launch the service with `run service` command


<%_ if (useDB) { _%>
### Database
For services, we do use ***PostgresSQL*** database. 
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).

<%_ } _%>

### Third-party services


## Contact

Responsible <%= contactName %>
The primary contact for this project is  <%= telegramContact %>
