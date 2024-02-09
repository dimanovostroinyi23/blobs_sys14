# Generate microservice skeleton for TokenD system

## Docker way

**Create** directory, where files will be generated **beforehand**.

Then run following command:
```sh
# -v will mount created directory to the `/generated` dir inside container, use absolute path, please
docker run --pull always -it --rm -v $(pwd):/generated registry.gitlab.com/tokend/yo-generator-tokend-module
```
> Note: If you didn't create target directory before running generation, you'll have to run:
>
> `sudo chown -R $(id -u):$(id -g) <PATH_TO_DESIRED_DIR>`

When generation finishes, just go to the `<PATH_TO_DESIRED_DIR>` and run `go mod init`. Now you are ready to deploy your fabulous microservice for TokenD.

## No-silly-containers way

First, install [Yeoman](http://yeoman.io):

```sh
npm install -g yo
```

Then go to the root of this repo and link package:
```sh
# Run in the root of this repo
npm link
```

Then generate your new project:
```sh
cd <PATH_TO_DESIRED_DIR>
yo tokend-module
```

And prepare it to work:
```sh
go mod init
go mod vendor
```
