package main

import (
    "gopkg.in/macaron.v1"
)

func main() {
    // Initialise macaron
    m := macaron.Classic()

    // Run webserver
    m.Run()
}
