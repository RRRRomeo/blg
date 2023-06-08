package cnf

import (
	"blg/tools/common"
	"fmt"
	"log"
	"path"

	"github.com/spf13/viper"
)

type ServerCnf struct {
	Addr string
	Port string
}

type LogCnf struct {
	Level string
}

type DbCnf struct {
	Host string
	Port string
	User string
	Name string
	Pass string
}

type CNF struct {
	Server ServerCnf
	Log    LogCnf
	Db     DbCnf
}

var GlobalCnf CNF

func readCnf() {
	cnfPath := path.Join(common.GetCurPath(), "cnf")
	fmt.Printf("cnfPath:%s\n", cnfPath)
	// 初始化 Viper
	viper.SetConfigName("blg")   // 配置文件的文件名（不包含扩展名）
	viper.SetConfigType("yaml")  // 配置文件的类型，这里使用 YAML 格式
	viper.AddConfigPath(cnfPath) // 配置文件所在的路径，这里使用当前目录

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	// // 读取配置项的值
	GlobalCnf.Server.Addr = viper.GetString("blg.server.addr")
	GlobalCnf.Server.Port = viper.GetString("blg.server.port")
	GlobalCnf.Db.Host = viper.GetString("blg.db.host")
	GlobalCnf.Db.Port = viper.GetString("blg.db.port")
	GlobalCnf.Db.User = viper.GetString("blg.db.user")
	GlobalCnf.Db.Pass = viper.GetString("blg.db.pass")
	GlobalCnf.Db.Name = viper.GetString("blg.db.name")
	GlobalCnf.Log.Level = viper.GetString("blg.log.level")

	// // 打印配置项的值
	log.Printf("GlobalCnf:%+v\n", GlobalCnf)
}

// func VipperReadSeverCnf() {

// }
