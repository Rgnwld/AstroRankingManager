# Astro Ranking Manager (ARM)

## This ranking is still in early development

This is only for study and early tests in Astro development.
Please dont use it in production.

### 🚧 Requirements to start:

- Need Docker installed in your computer.
  - If you dont have Docker installed, you can get it here [Docker](https://www.docker.com/products/docker-desktop/)

### 🚧 Start up this manager:

To start this manager, run the `make build` command. 
This command will compile the API and start a database container with all the dependencies necessary for the successful execution of the manager.

For dev proposes, use `make dev` command. It will only build the database container and export its port.

For more commands run `make help`

### 📦 Tools and Features of this project
  - ### API
    - [Go](https://go.dev/)
      - [Gin](https://github.com/gin-gonic/gin) (API)
      - [Colly](https://github.com/gocolly/colly) (WebScrapper) 
    - [JWT](https://jwt.io/)


  - ### Environment
    - ### Database
      - [MySql](https://www.mysql.com/)
      - [Goose](https://github.com/pressly/goose) (Migrations)
    - [Docker](https://www.docker.com/)
      - [Docker-Compose](https://docs.docker.com/compose/)
      - [Docker-File](https://docs.docker.com/engine/reference/builder/)
    - [Make](https://makefiletutorial.com/)