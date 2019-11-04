package config

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/viper"
)

const (
	ProjectEnvironmentPrefix = "CLASSIFIER"
)

type ElasticSearch struct {
	Scheme  string `mapstructure:"scheme"`
	Host    string `mapstructure:"host"`
	Port    string `mapstructure:"port"`
	Timeout uint   `mapstructure:"timeout"`
}

type Configuration struct {
	Redis            Redis                 `mapstructure:"redis"`
	Website          Website               `mapstructure:"website"`
	Server           Server                `mapstructure:"server"`
	BaseDir          string                `mapstructure:"-"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Expiry   int    `mapstructure:"expiry"`
	DB       int    `mapstructure:"db"`
}

func (r Redis) Address() string {
	return net.JoinHostPort(r.Host, r.Port)
}

type Website struct {
	Host     string `mapstructure:"host"`
	Protocol string `mapstructure:"protocol"`
}

func (w Website) DNS() string {
	return fmt.Sprintf("%s://%s", w.Protocol, w.Host)
}

type Server struct {
	Protocol string `mapstructure:"protocol"`
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	TmpPath  string `mapstructure:"tmp_path"`
}

func (s Server) DNS() string {
	return fmt.Sprintf("%s://%s:%s", s.Protocol, s.Host, s.Port)
}

func getCountryCodes() []string {
	return []string{"ae", "eg", "bh", "lb", "ma", "qa", "sa"}
}

func createAppConfigs() Configurations {
	result := Configurations{}
	result.List = make(map[string]Configuration)

	return result
}

type Configurations struct {
	List map[string]Configuration `mapstructure:"list"`
}

func (c Configurations) All() map[string]Configuration {
	return c.List
}

func (c Configurations) Get(countryCode string) (Configuration, error) {
	if result, ok := c.List[countryCode]; ok {
		return result, nil
	}

	return Configuration{}, fmt.Errorf(`invalid countrCode "%s"`, countryCode)
}

func NewConfigs() Configurations {
	configName := "config"
	baseViper := newViper("config", ProjectEnvironmentPrefix, configName, nil)
	generalSettings := baseViper.AllSettings()

	cfgs := createAppConfigs()
	for _, countryCode := range getCountryCodes() {
		cfg := Configuration{}
		configName := countryCode
		envProjectCountryPrefix := strings.ToUpper(fmt.Sprintf("%s_%s", ProjectEnvironmentPrefix, countryCode))
		countryViper := newViper(
			"config/country",
			envProjectCountryPrefix,
			configName,
			generalSettings,
		)
		if e := countryViper.Unmarshal(&cfg); nil != e {
			panic(e)
		}
		cfgs.List[countryCode] = cfg
	}

	return cfgs
}

func newViper(path string, envPrefix string, configName string, generalSettings map[string]interface{}) *viper.Viper {
	v := viper.New()
	configPath := "/go/src/github.com/ustrugany/classifier/" + path
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)

	if err := v.ReadInConfig(); nil != err {
		panic(err)
	}

	v = v.Sub("config")
	if nil != generalSettings {
		// be wary that general settings will take precedence over country specific ones
		if err := v.MergeConfigMap(generalSettings); nil != err {
			panic(err)
		}
	}
	v.AutomaticEnv()
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AddConfigPath("./" + path)
	v.AutomaticEnv()
	return v
}
