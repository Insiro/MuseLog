package apps

import "your_spotify/internal/app/bean"

func initConfig() {
	config := bean.InitConfig()
	bean.InitBean(config)
}
