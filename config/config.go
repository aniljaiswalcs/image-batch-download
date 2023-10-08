package cfg

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ImageUrlPath string
	BatchSize    uint
	BaseRootPath string
}

func NewConfig(filepath string) *Config {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.SetConfigFile(filepath)

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	setDefaults(v)

	return &Config{
		ImageUrlPath: v.GetString("imageurlpath"),
		BatchSize:    v.GetUint("batchsize"),
		BaseRootPath: v.GetString("storagerootpath"),
	}

}

func setDefaults(v *viper.Viper) {
	v.SetDefault("batchsize", 25)
	v.SetDefault("storagerootpath", "/mnt/d/go-workspace/src/all_repo/getsafe/image-batch-download")
}
