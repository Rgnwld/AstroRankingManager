# Astro Ranking Manager (ARM)

### This ranking is still in early development and doesnt have any kind of security
This is only for study and early tests in Astro development.
Please dont use it in production.
- NOTE: Working on ["JWT"](https://jwt.io/) system


### How to start up this manager:
- Need Docker installed in your computer. 
    - If you dont have Docker installed, you can get it here [Docker](https://www.docker.com/products/docker-desktop/)
- ARM uses mysql to persist ranking data. On this repo you can use a DockerFile to create a container with base configurations to use it.
    - To start the Dockerfile, run in this folder ```docker build -t astromanager . ; docker run -dp 127.0.0.1:3306:3306```
    - Make sure that port 3306 is free to use!
- To run the project, you still need the [go compiler](https://go.dev/learn/)
- Now that everything is setup, just run ```go run ./```


### Patch Notes
- 16/10: 
    - Started to implement JWT system. Sign In already working
    - Added authentication for some routes (ex: Ranking)