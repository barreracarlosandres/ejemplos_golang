![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/barreracarlosandres/ejemplos_golang)![GitHub Repo stars](https://img.shields.io/github/stars/barreracarlosandres/ejemplos_golang)
![GitHub issues](https://img.shields.io/github/issues/barreracarlosandres/ejemplos_golang)![GitHub watchers](https://img.shields.io/github/watchers/barreracarlosandres/ejemplos_golang)![GitHub followers](https://img.shields.io/github/followers/barreracarlosandres)

## Ejemplos GoLang

En cada una de las carpetas se tiene ejemplos de los temas realaciones y sus test
- Ejemplo de [punteros](./pointers/main.go) - [test](./pointers/main_test.go)
  ```shell
  go run main.go pointers
  ```
  ```shell
  go test -coverpkg=./... pointers/main_test.go
  ```
  - Ejemplo de [api Rest con Gin](./apirestgin/main.go)
  ```shell
  go run main.go apiRestGin
  ```
  - Ejemplo de [structs](./structs/main.go)
  ```shell
  go run main.go structs
  ```

Nota: Estaré ajustando los ejemplos y adicionando test

## Authors

- [Carlos Barrera](https://linkedin.com/in/carlos-andres-barrera-994031080)
