package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	config := initialize.ParseConfigPath(path...)
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GvaConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GvaConfig); err != nil {
		fmt.Println(err)
	}
	return v
}
