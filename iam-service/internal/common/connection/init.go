package connection

import "log"

func Init() {
	if err := loadConfig(); err != nil {
		log.Fatalln("Loading config faild. Error: ", err)
		return
	}
}
