package app

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/project-wind/wind-comment/config"
    "github.com/project-wind/wind-comment/models"
    "github.com/project-wind/wind-comment/routers"
)

// Init initializes the app.
func Init(configPath string) {
    err := config.ReadJSON(configPath)
    if err != nil {
        log.Fatal(err)
    }

    err = models.Init()
    if err != nil {
        log.Fatal(err)
    }

    r := gin.Default()

    r.SetTrustedProxies(nil)
    // middlewares.CORSInit(r)
    routers.Init(r)

    log.Fatal(r.Run(fmt.Sprintf("%v:%v", config.Conf.Address, config.Conf.Port)))
}
