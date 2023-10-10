package config

type DomainConfig = map[string]DomainS3Infos

type DomainS3Infos struct {
	AccessKey string `mapstructure:"access-key"`
	SecretKey string `mapstructure:"secret-key"`
	Endpoint  string
	Region    string
	Bucket    string
}
