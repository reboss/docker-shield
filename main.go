package main

import (
    "log"
    "fmt"
    "github.com/docker/go-plugins-helpers/authorization"
    "github.com/reboss/docker-shield/plugin"
)

func main() {
    dockerShield := plugin.NewDockerShield()
    h := authorization.NewHandler(dockerShield)

    fmt.Println("docker-shield initializing")

    if err := h.ServeUnix("docker-shield", 0); err != nil {
        log.Fatal(err)
    }
}
