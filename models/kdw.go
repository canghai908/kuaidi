package models

import (
	"errors"
	//"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"strconv"
	"time"
	"wxapi/util"
)

type User struct {
	Id       int64
	Username string `xorm:"unique not null"`
	Password string
	Count    int64
	Key      string
	Version  int `xorm:"version"`
}

var x *xorm.Engine

func Init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "info.db")
	if err != nil {
		log.Fatal("Fail to create engine:%v\n", err)
	}
	if err = x.Sync(new(User)); err != nil {
		log.Fatal("Fail to sync database:%v\n", err)
	}
}

func AddUser(username string, password string) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := r.Intn(1000000)
	p := strconv.Itoa(t) + "kuaidi"
	dpasswd := util.Encrypter(p)
	_, err := x.Insert(&User{Username: username, Password: password, Key: dpasswd})
	return err
}

//fun Getinfo(cemskind,cno,key string)
func GetUser(username string) (*User, error) {
	a := &User{}
	has, err := x.Where("username=?", username).Get(a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("User does no exist")
	}
	return a, err
}
