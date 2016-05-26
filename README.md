# Golang web application with Gorilla toolkit

* Gorilla toolkit (not a framework)
  * Benefits:
    * Massive concurrency
    * Session management
    * Security
    * Routing

* Demo application (auto parts)
  * go run src/main.go (replace the import paths to: "github.com/golang-webapp-gorilla/src/bw/ctrl" if it's not under GOPATH)
  * structure ("Overview of the demo code"):
    * main.go
    * controller: ctrl folder
    * model
    * viewmodel: vm folder, pass to views
    * template: html template
