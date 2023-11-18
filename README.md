# my very own scripts for insights capture

Scripts for insights capture to do things like:

-   merge clips from the same day
-   decode backup file
-   upload to youtube

## development environment

```bash
PF=/go/src/insights-capture-scripts docker run --rm -it --name godev -v $(pwd):$PF -e GOCACHE=$PF/.cache -u 1000 -w $PF golang:1.21 /bin/bash

go get
```

## build

### Linux

Inside the docker container:

```bash
OOS=linux GOARCH=amd64 go build -buildvcs=false -o build/insights-capture-scripts insights-capture-scripts
```

### Windows

Inside the docker container:

```bash
GOOS=windows GOARCH=amd64 go build -buildvcs=false -o build/insights-capture-scripts.exe insights-capture-scripts
```
