package mioji_label

import (
    "gin/connections"
    "errors"
)

type User struct {
    Id           int    `json:"id"`
    Name         string `json:"name"`
    Account      string `json:"account"`
    Password     string `json:"-"`
    RoleId       int    `json:"roleId"`
    LabelTarget  int    `json:"label_target"`
    ReviewTarget int    `json:"review_target"`
    Ctime        int    `json:"ctime"`
    Utime        string `json:"utime"`
    //Disable      int    `json:"disable"`
    //DisableTs    int    `json:"disable_ts"`
}

func (t *User) TableName() string {
    return "user"
}

//获取所有用户数据
func GetAllUsers() (users []*User, err error) {
    var t *User
    noRecord := connections.MySQLClient["mioji_label"].Table(t.TableName()).Find(&users).RecordNotFound()

    if noRecord {
        return nil, errors.New("Users not exists")
    }

    return users, nil
}

//根据用户id查询单条数据
func GetUser(id int) (user *User, err error) {
    var u User
    noRecord := connections.MySQLClient["mioji_label"].Table(u.TableName()).Where("id = ?", id).First(&u).RecordNotFound()

    if noRecord {
        return nil, errors.New("User not exists")
    }

    return &u, nil
}
