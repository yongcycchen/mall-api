package config

import (
	"log"
	"strings"

	"github.com/yongcycchen/mall-api/internal/config"
)

func MapConfig(section string, v interface{}) {
	if strings.HasPrefix(section, "web-") {
		log.Fatalf("[err] section name can't have web- prefix")
	}
	config.MapConfig(section, v)
}
