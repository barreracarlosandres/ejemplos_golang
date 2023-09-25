![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/barreracarlosandres/ejemplos_golang)![GitHub Repo stars](https://img.shields.io/github/stars/barreracarlosandres/ejemplos_golang)
![GitHub issues](https://img.shields.io/github/issues/barreracarlosandres/ejemplos_golang)![GitHub watchers](https://img.shields.io/github/watchers/barreracarlosandres/ejemplos_golang)![GitHub followers](https://img.shields.io/github/followers/barreracarlosandres)

## Ejemlos GoLang
En cada una de las carpetas se tiene ejemplos de los temas realaciones

| Example                            | Ready | Test                           |
|------------------------------------|-------|--------------------------------|
| [pointers](./pointers/main.go)     | yes   | [yes](./pointers/main_test.go) |
| channel                            | yes   | -o-                            |
| maps                               | yes   | -o-                            |
| structs                            | yes   | -o-                            |
| [apiRestGin](./apirestgin/main.go) | yes   | -o-                            |

# ¿Cómo ejecuto los ejemplos y test?

- En la raiz del projecto ejecutar todos los ejemplos
  ```
  go run main.go
  ```
- Ejecutar un ejemplo así:
  ```
  go run main.go apiRestGin
  ```
- Correr todos los test y vaidar la covertura
  ```
  go test -coverpkg=./... ./...
  ```
Nota: Estaré ajustando los ejemplos y adicionando test

## Authors
- [Carlos Barrera](https://linkedin.com/in/carlos-andres-barrera-994031080)
