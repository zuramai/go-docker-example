# Golang Docker Example
1. Clone the repository
```bash
$ git clone github.com/zuramai/go-docker-example
```
2. Create docker container
```bash
$ docker build -t go-docker-example .
$ docker container create --name go-docker-example -e PORT=8080 -p 8080:8080 go-docker-example
$ docker container ls -a # check if container created
```
3. Start the container
```bash
$ docker container start go-docker-example # or use the container id instead
$ docker container ls
```

Now check `http://localhost:8080` in your browser.


### Stop the container
```bash
$ docker container stop docker-container-example
$ docker container ls
```

### Delete container
```bash
$ docker container rm docker-container-example
$ docker container ls
```

### Delete image
```bash
$ docker image rm docker-container-example
$ docker images ls
```