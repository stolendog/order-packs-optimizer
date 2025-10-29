## Order packs optimizer

This project provides a web app (frontend + backend) to calculate optimal pack sizes for orders according to the following rules:

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.
(Please note, rule #2 takes precedence over rule #3)

## Project Structure

```
cmd/                    - Entry points
internal/
  ├── api/              - HTTP transport layer
  ├── app/              - Application services
  ├── domain/           - Core business logic
  ├── calculator/       - Specialized packing algorithms
  └── infra/            - Infrastructure implementations
web/                    - Frontend assets
```

### Prerequisites

- Docker

- Golang

#### Development Commands

`make help` - Display all available make commands with descriptions

`make run` - Run the application locally on port 9999. The web interface will be available at http://localhost:9999

`make test` - Run all unit tests with verbose output

`make deps` - Download GO module dependencies

#### Docker Commands

`make docker-build` - Build the Docker image for the application (tagged as `order-pack-optimizer`)

`make docker-run` - Run the application in a Docker container on port 9999. Creates an interactive container that will be removed when stopped


