package main

import (
	. "github.com/Monibuca/engine/v2"
	_ "github.com/Monibuca/plugin-gateway"
	_ "github.com/Monibuca/plugin-gb28181"

)
func main()  {
	// engine 暴露出来的方法。
	Run("config.toml")
	select {}
}
