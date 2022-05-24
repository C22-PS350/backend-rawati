# Backend
### App & tools:
  * Go version 1.17 or newer
  * Docker
  * Make (apparently only available on Unix family or via WSL on Windows)

### Database:
```bash
  make db-up (start)
  make db-stop (stop)
```

### Run the app:
```bash
  go mod tidy (first time only)
  go run github.com/C22-PS350/backend-rawati/cmd/rawati
```

#### Access the app:
  ``localhost:8080``
