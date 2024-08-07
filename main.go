package main

import (
    "github.com/Hailemari/library_management/controllers"
    "github.com/Hailemari/library_management/services"
)

func main() {
    library := services.NewLibrary()
    controller := controllers.NewLibraryController(library)
    controller.HandleRequest()
}
