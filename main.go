package main

/*
参考链接：

动态信息: https://api.bilibili.com/x/polymer/web-dynamic/v1/detail?timezone_offset=-480&id=666364927640862760
动态评论列表: https://api.bilibili.com/x/v2/reply/main?jsonp=jsonp&next=0&type=11&oid=196002511&mode=2&plat=1&ps=30
动态点赞列表: https://api.vc.bilibili.com/dynamic_like/v1/dynamic_like/spec_item_likes?dynamic_id=675746321584357384&pn=1&ps=50
动态转发列表: https://api.live.bilibili.com/dynamic_repost/v1/dynamic_repost/view_repost?dynamic_id=666364927640862760&offset=550
粉丝列表: https://api.bilibili.com/x/relation/followers?vmid=476296860&pn=1&ps=50&order=desc&order_type=attention&jsonp=jsonp

*/

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	config              *Config
	db                  *sql.DB
	filePath, tableName string
)

func init() {
	var (
		fileName, dbName, dbPath string
		err                      error
	)

	flag.StringVar(&fileName, "c", "./config.toml", "Path to config file.")
	flag.StringVar(&dbName, "d", "./bili.db", "Path to db.")

	// 获取配置文件的绝对路径
	filePath, err = filepath.Abs(fileName)
	checkErr(err)

	// 获取数据库绝对路径
	dbPath, err = filepath.Abs(dbName)
	checkErr(err)

	// 读取配置文件
	config, err = readConf(filePath)
	checkErr(err)

	// 连接数据库
	db, err = sql.Open("sqlite3", dbPath)

	// 表名
	tableName = fmt.Sprintf("D_%v", config.Dynamic.ID)

	// 数据库初始化
	createDB()

	// 初始化 resty
	initHTTP()
}

func main() {
	// 初始化log
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	checkErr(err)
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime)

	// 登录
	loginStart()

	// 执行任务
	taskStart()
}

func checkErr(err error) {
	if err != err {
		log.Fatalln(err)
	}
}
