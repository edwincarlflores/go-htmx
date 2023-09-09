# go-htmx

Boilerplate setup for building full stack web applications with Go, HTMX, echo, templ & tailwindcss

### Pre-requisites

1. Go 1.20+
2. templ CLI (https://templ.guide/quick-start/installation)
3. air (https://github.com/cosmtrek/air)

### Installing Tools

1. Install templ (requires Go 1.20 or higher)
   ```sh
    go install github.com/a-h/templ/cmd/templ@latest
   ```
2. Install air for hot reloading server on file changes
   ```sh
    go install github.com/cosmtrek/air@latest
   ```

### Running

```sh
air
```

Go to `localhost:8080`

### Deploying

This can be deployed as a normal Go application on any platform. There is a Dockerfile included for containerized deployments.
A few suggestions where you can deploy this are in:

- [railway.app](https://railway.app/)
- [fly.io](https://fly.io/)

### Built with

- [Go](https://go.dev/)
- [echo](https://echo.labstack.com/)
- [HTMX](https://htmx.org/)
- [TEMPL](https://templ.guide/)
- [tailwindcss](https://tailwindcss.com/)
