# Simple HelloWorld

Based on Google Cloud Run example.

## Structure

```text
build.sh   - all you need to build
run.sh     - just starts container
main.go    - all go code
Dockerfile - build file for Docker
```

## Build

```text
$ docker build -t helloworld-go .
```

## Run

```text
$ docker run -it --rm -p 7000:8080 helloworld-go
```

```text
$ curl http://localhost:7000/
```
