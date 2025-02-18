package configs

type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
		Jwt      Jwt      `mapstructure:"jwt"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}

	Jwt struct {
		SecretKey string `mapstructure:"secretKey"`
	}
)
