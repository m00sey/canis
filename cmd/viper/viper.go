package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	//NewFileConfig constructs the agent config from filesystem
	//noinspection GoUnusedExportedFunction
	pflag.Parse()
	viper.SetConfigName("canis")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/canis/")
	//maybe
	viper.AddConfigPath("./config/proc/")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Fatal error config file", err)
	}

	ds := viper.GetStringMap("mongo")
	st := viper.Sub("steward")

	st.Set("mongo", ds)

	out := map[string]interface{}{}
	err = st.Unmarshal(&out)
	if err != nil {
		log.Fatalln(err)
	}

	d, _ := json.MarshalIndent(out, " ", " ")
	fmt.Println(string(d))

}
