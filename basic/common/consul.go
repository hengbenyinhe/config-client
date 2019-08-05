package common

//consul配置
type Consul struct {
	Enabled bool     `json:"enabled"`
	Host    string   `json:"host"`
	Port    int      `json:"port"`
}