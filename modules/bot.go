package modules

import (
    "github.com/fluffle/goirc/client"
    "crypto/tls"
    "strings"
)

var (
    BOT *client.Conn
)

func InitBot() {
    config := client.NewConfig(
        CONFIG.Section("irc").Key("nickname").String(), // nick
        CONFIG.Section("irc").Key("nickname").String(), // ident
        CONFIG.Section("irc").Key("username").String(), // username
    )
    config.Server = CONFIG.Section("irc").Key("server").String()
    config.SSL = CONFIG.Section("irc").Key("ssl").MustBool(false)
    config.SSLConfig = &tls.Config{ServerName: strings.Split(CONFIG.Section("irc").Key("server").String(), ":")[0]}
    if CONFIG.Section("irc").HasKey("password") {
        config.Pass = CONFIG.Section("irc").Key("password").String()
    }

    BOT = client.Client(config)
}
