package main

import (
    "github.com/gin-gonic/gin"
    _ "gin/connections"
    "strconv"
    "gin/models/mioji-label"
    "gin/models/base-data"
    "gin/api"
)

var DB = make(map[string]string)

func SetupRouter() *gin.Engine {
    // Disable Console Color
    // gin.DisableConsoleColor()
    r := gin.Default()

    r.NoRoute(handle404)

    r.GET("/", func(c *gin.Context) {
        apiName := c.Query("type")

        var result interface{}
        switch apiName {
        case "apitest":
            result = api.ApiTest(c)
        case "api00001":
            result = api.Api00001(c)
        default:
            result = gin.H{"c": 0, "m": "不存在该接口", "d": gin.H{}}
            return
        }

        c.JSON(200, result)
    })

    r.GET("/test", func(c *gin.Context) {
        return
    })

    // Ping test
    r.GET("/users", func(c *gin.Context) {
        users, err := mioji_label.GetAllUsers()
        if err != nil {
            c.JSON(200, gin.H{"A": "B"})
            return
        }

        c.JSON(200, users)
    })

    r.GET("/user/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Params.ByName("id"))

        user, err := mioji_label.GetUser(id)
        if err != nil {
            c.JSON(200, gin.H{"A": "B"})
            return
        }

        c.JSON(200, user)
    })

    r.GET("/countries", func(c *gin.Context) {
        countries, err := base_data.GetAllCountries()
        if err != nil {
            c.JSON(200, gin.H{"A": "B"})
            return
        }

        c.JSON(200, countries)
    })

    // Authorized group (uses gin.BasicAuth() middleware)
    // Same than:
    // authorized := r.Group("/")
    // authorized.Use(gin.BasicAuth(gin.Credentials{
    //	  "foo":  "bar",
    //	  "manu": "123",
    //}))
    authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
        "foo":  "bar", // user:foo password:bar
        "manu": "123", // user:manu password:123
    }))

    authorized.POST("admin", func(c *gin.Context) {
        user := c.MustGet(gin.AuthUserKey).(string)

        // Parse JSON
        var json struct {
            Value string `json:"value" binding:"required"`
        }

        if c.Bind(&json) == nil {
            DB[user] = json.Value
            c.JSON(200, gin.H{"status": "ok"})
        }
    })

    return r
}

func handle404(c *gin.Context) {
    c.JSON(404, gin.H{"c": 404, "m": "路由错误", "d": gin.H{}})
}

func main() {
    r := SetupRouter()
    // Listen and Server in 0.0.0.0:8080
    r.Run(":8080")
}
