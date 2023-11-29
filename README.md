# Astro Ranking Manager (ARM)

## 🚧 This ranking is still in early development

This is only for study and early tests in Astro development. Please dont use it in production.<br />
The propose of this project is to provide and handle requisitions of the [Astro Game](https://rgnwld.itch.io/astro).<br />
Such as:
- Handle game times.
- Rank the players by their times in each map.
- Provide informations about the game news.

## 🚧 Requirements to start:

- Need Docker installed in your computer.
  - If you dont have Docker installed, you can get it here [Docker](https://www.docker.com/products/docker-desktop/)

## 🚧 Start up this manager:

To start this manager, run the `make up` command. This command will compile an API and start a database container with all the dependencies necessary for the successful execution of the manager.<br />
For development proposals, use the `make dev` command. It will build and run the containers with the logs attached to the terminal.<br />
For development proposals using only the database, run the `make db` command. It will build the database container and export its port.<br />
For more commands run `make help`.

## 📦 Current Tools and Features of this project
  ### API
  - [Go](https://go.dev/)
    - [Gin](https://github.com/gin-gonic/gin) (API)
    - [Colly](https://github.com/gocolly/colly) (WebScrapper)
  - [JWT](https://jwt.io/)

  ### Database
  - [MySql](https://www.mysql.com/)
  - [Goose](https://github.com/pressly/goose) (Migrations)

  ### Environment
  - [Docker](https://www.docker.com/)
    - [Docker-Compose](https://docs.docker.com/compose/)
    - [Docker-File](https://docs.docker.com/engine/reference/builder/)
  - [Make](https://makefiletutorial.com/)