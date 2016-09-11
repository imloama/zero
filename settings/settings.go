package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var environments = map[string]string{
	"production": "settings/prod.json",
	"dev":        "settings/dev.json",
}

type Settings struct {
	SecretKey          string //加密key
	JWTExpirationDelta int    //jwt的有效时间
	URL                string //指的系统访问地址
	LogLevel           string //日志级别   debug info fatal error  warn warning all
	LogOutPut          string //输出的日志文件
	Port               int    //系统端口
	Installed          bool   //是否已经安装
	DbURL              string //数据库连接地址
	DbUser             string //数据库用户
	DbPwd              string // 数据库用户密码
}

var settings Settings = Settings{}
var env = "dev"

func Init() {
	env = os.Getenv("DEV_MODEL")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "dev"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	//	fmt.Printf("read file content: %v\n", content)
	log.Printf("read file content: %s\n", content)
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

//SaveSettings 根据当前调整后的变量，保存settings
func SaveSettings() {
	data, err := json.Marshal(settings)
	if err != nil {
		fmt.Println("Error while json settings ", err)
		return
	}
	err = ioutil.WriteFile(environments[env], data, 0666) //os.ModePerm)
	if err != nil {
		fmt.Println("Error while write settings file ", err)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsDevEnvironment() bool {
	return env == "dev"
}
