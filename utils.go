package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"time"
)

const ua = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

var (
	api     = resty.New()
	apiVc   = resty.New()
	apiLive = resty.New()
	client  = resty.New()
	login   = resty.New()
)

// 读取配置
func readConf(filename string) (*Config, error) {
	var (
		conf *Config
		err  error
	)

	if _, err = toml.DecodeFile(filename, &conf); err != nil {
		log.Fatal(err)
	}

	return conf, err
}

// 写入配置
func writeConf(filename string) error {
	f, err := os.Create(filename)

	if err != nil {
		// failed to create/open the file
		log.Fatal(err)
	}

	if err = toml.NewEncoder(f).Encode(config); err != nil {
		// failed to encode
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		// failed to close the file
		log.Fatal(err)
	}

	return err
}

// 创建数据库
func createDB() {
	createTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS '%v' "+
		"('ID'	INTEGER NOT NULL, "+
		"'Mid'	INTEGER NOT NULL UNIQUE, "+
		"'Forward'	INTEGER NOT NULL DEFAULT 0, "+
		"'Comment'	INTEGER NOT NULL DEFAULT 0, "+
		"'Like'	INTEGER NOT NULL DEFAULT 0, "+
		"'Relation'	INTEGER NOT NULL DEFAULT 0, "+
		"'VIP'	INTEGER NOT NULL DEFAULT 0, "+
		"'Weights'	INTEGER DEFAULT 0, "+
		"PRIMARY KEY('ID' AUTOINCREMENT));", tableName)

	_, err := db.Exec(createTable)
	checkErr(err)
}

// 搜索数据
func searchDB(key string, mid int) int {
	var result int

	searchTable := fmt.Sprintf("SELECT %v FROM %v WHERE Mid = %v", key, tableName, mid)

	rows, err := db.Query(searchTable)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&result)
		checkErr(err)
	}

	return result
}

// 写入数据
func insertDB(key string, value int) {
	insertTable := fmt.Sprintf("INSERT INTO %v(%v) values(?)", tableName, key)

	i, err := db.Prepare(insertTable)
	checkErr(err)

	_, err = i.Exec(value)
	checkErr(err)
}

// 更新数据
func updateDB(key string, value, mid int) {
	updateTable := fmt.Sprintf("UPDATE %v SET %v = ? WHERE MID = ?", tableName, key)

	u, err := db.Prepare(updateTable)
	checkErr(err)

	_, err = u.Exec(value, mid)
	checkErr(err)
}

func initHTTP() {
	headers := map[string]string{
		"dnt":        "1",
		"user-agent": ua,
		"refer":      fmt.Sprintf("https://t.bilibili.com/%v", config.Dynamic.ID),
	}

	api.SetHeaders(headers)
	apiVc.SetHeaders(headers)
	apiLive.SetHeaders(headers)

	if config.ProxyURL.API != "" {
		api.SetBaseURL(config.ProxyURL.API + "/x")
	} else {
		api.SetBaseURL("https://api.bilibili.com/x")
	}

	if config.ProxyURL.APIVc != "" {
		apiVc.SetBaseURL(config.ProxyURL.APIVc)
	} else {
		apiVc.SetBaseURL("https://api.vc.bilibili.com")
	}

	if config.ProxyURL.APILive != "" {
		apiLive.SetBaseURL(config.ProxyURL.APILive)
	} else {
		apiLive.SetBaseURL("https://api.live.bilibili.com")
	}

	api.SetTimeout(10 * time.Second)
	apiVc.SetTimeout(10 * time.Second)
	apiLive.SetTimeout(10 * time.Second)

	api.SetRetryCount(3)
	apiVc.SetRetryCount(3)
	apiLive.SetRetryCount(3)
	client.SetRetryCount(3)
	login.SetRetryCount(3)

	client.SetBaseURL("https://api.bilibili.com/x")
	login.SetBaseURL("https://passport.bilibili.com")
}
