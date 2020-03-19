# SWS
SWS (Simple Web Server) is a simple web server that prints the contents of the `MESSAGE` environment variable.

## Build
```shell script
docker build -t sws .
docker run -p 8000:8000 -e MESSAGE='Hello World' sws
```
