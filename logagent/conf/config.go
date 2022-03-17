package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
	Maxsize int    `ini:"maxsize"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}

// NO USE
type TaillogConf struct {
	Filename string `ini:"filename"`
}
