# Backend
### App & tools:
  * Go version 1.17 or newer
  * Docker
  * Make (apparently only available on Unix family or Windows via WSL)

### Database:
```bash
make db-up (start)
make db-stop (stop)
```

### See API docs:
 https://rawatidocs.herokuapp.com/

### Run the app:
```bash
go mod tidy (only first run)
go run github.com/C22-PS350/backend-rawati/cmd/rawati
```

### Access the app:
```bash
localhost:8080
```
