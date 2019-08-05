package common

import "strconv"
//应用配置项
type AppCfg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (a *AppCfg) Addr() string  {
	return a.Address + ":" + strconv.Itoa(a.Port)
}
