package helpers

import "github.com/gin-gonic/gin"

//封装返回结果格式
func Response(data interface{}) gin.H {
    return gin.H{
        "code": 0,
        "msg":  "",
        "data": data,
    }
}
