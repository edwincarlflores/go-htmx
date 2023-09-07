# go-htmx

Boilerplate setup for building full stack web applications with Go, HTMX, fiber, templ & tailwindcss

### Pre-requisites

1. Go 1.20+
2. templ CLI (https://templ.guide/quick-start/installation)
3. air (https://github.com/cosmtrek/air)

### Installing Tools

1. Install templ (requires Go 1.20 or greater)
   ```sh
    go install github.com/a-h/templ/cmd/templ@latest
   ```
2. Install air for hot reloading server on file changes
   ```sh
    go install github.com/cosmtrek/air@latest
   ```

### Running

```
air
```

Go to `localhost:8080`

### Built with

- [Go](https://go.dev/)
- [Fiber](https://docs.gofiber.io/)
- [HTMX](https://htmx.org/)
- [TEMPL](https://templ.guide/)
- [tailwindcss](https://tailwindcss.com/)
