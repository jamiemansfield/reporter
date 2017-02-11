package main

import (
    "fmt"
    "log"
    "net/http"
    "gopkg.in/macaron.v1"
    "github.com/jamiemansfield/reporter/modules"
)

func main() {
    // Initialise webserver, configuration, and IRC bot
    m := macaron.Classic()
    modules.InitConfig()
    modules.InitBot()

    // Connect to IRC server
    if err := modules.BOT.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }

    // Run webserver
    log.Fatal(http.ListenAndServe(":" + modules.CONFIG.Section("web").Key("port").String(), m))
}
