package helpers

import (
    "encoding/hex"
    "crypto/md5"
    "time"
)

// 过滤字段串为 "NULL" 的
func FilterNullString(str string) string {
    if str == "NULL" {
        str = ""
    }

    return str
}

// 计算字符串的md5值
func Md5(str string) string {
    md5h := md5.New()
    md5h.Write([]byte(str))

    return hex.EncodeToString(md5h.Sum(nil))
}

// 格式化时间
func DateFormat(date time.Time, layout string) string {
    return date.Format(layout)
}
