package model

import (
    "gin/connections"
)

const TableName = "user"

type User struct {
    Id           int
    Name         string
    Account      string
    Password     string
    RoleId       int
    LabelTarget  int
    ReviewTarget int
    Ctime        int
    Utime        string
    Disable      int
    DisableTs    int
}

func GetUsers() []*User {
    var users []*User
    connections.MySQLConn["mioji_label"].Table(TableName).Find(&users)

    return users
}
