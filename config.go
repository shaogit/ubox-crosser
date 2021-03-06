package crosser

type Config struct {
	Password string `json:"password"`
	Method   string `json:"method"` // encryption method
}

type ClientConfig struct {
	Config
	TargetAddress string `json:"target_address"`
}

type ServerConfig struct {
	Config
	NorthAddress string `json:"north_address"`
	SouthAddress string `json:"south_address"`
}
