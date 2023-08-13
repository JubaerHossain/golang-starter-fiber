package config

import (
	"attendance/utils"
)

func InitDatabase() {
	utils.InitMongoDBConnection()
}
