package product

import (
	"github.com/Albinzr/duke_product_module/product/config"
	"github.com/Albinzr/duke_product_module/product/router"
)

type Config ProductConfig.Config

func (c *Config) Init() {
	routerConfig := (*router.Config)(c)
	routerConfig.Init()
}
