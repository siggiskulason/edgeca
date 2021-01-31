/*******************************************************************************
 * Copyright 2021 Darval Solutions Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 *******************************************************************************/

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var configFile string

func GetDefaultTLSPort() int {
	return 50025
}

func GetDefaultTLSHost() string {
	hostName, _ := os.Hostname()
	return hostName
}

func GetDefaultTLSCertDir() string {
	homeDir, _ := os.UserHomeDir()
	configDir := homeDir + "/.edgeca"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		_ = os.Mkdir(configDir, 0755)
	} else {
	}

	defaultTLSCertDir := configDir + "/certs"

	if _, err := os.Stat(defaultTLSCertDir); os.IsNotExist(err) {
		_ = os.Mkdir(defaultTLSCertDir, 0755)
	} else {
	}

	return defaultTLSCertDir
}

func GetConfigDir() string {
	homeDir, _ := os.UserHomeDir()
	configDir := homeDir + "/.edgeca"

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		_ = os.Mkdir(configDir, 0755)
	} else {
	}
	return configDir
}

func InitCLIConfiguration() {

	readConfiguration("configuration.yaml")

}

func readConfiguration(filename string) {
	configDir := GetConfigDir()
	configFile = configDir + "/" + filename
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	//viper.AddConfigPath(".")
	viper.SetDefault("default-csr-values.organization", "")
	viper.SetDefault("default-csr-values.organizationalUnit", "")
	viper.SetDefault("default-csr-values.country", "")
	viper.SetDefault("default-csr-values.province", "")
	viper.SetDefault("default-csr-values.locality", "")

	//	viper.ReadInConfig()
	viper.SafeWriteConfigAs(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Could not read:", err)
	}

}

/*
func SetConfigurationUsingYAML(yaml string) error {
	viper.ReadConfig(bytes.NewBuffer([]byte(yaml)))

	viper.WriteConfig()
	file, err := GetConfigurationFileContents()
	log.Println("Updated configuration:", file)
	return err

}
*/
func SetCSRConfiguration(o string, ou string, c string, p string, l string) error {

	viper.Set("default-csr-values.organization", o)
	viper.Set("default-csr-values.organizationalUnit", ou)
	viper.Set("default-csr-values.country", c)
	viper.Set("default-csr-values.province", p)
	viper.Set("default-csr-values.locality", l)
	/*
		file, err := GetConfigurationFileContents()
		if fn != "" {
			err := ioutil.WriteFile(fn, []byte(file), 0644)
			if err != nil {
			}
		}
	*/
	viper.WriteConfig()
	log.Println("Updated configuration file " + configFile)
	return nil
}

func GetConfigurationFileContents() (string, error) {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		err = fmt.Errorf("unable to marshal config to YAML: %w", err)
	}
	return string(bs), err

}

func WriteConfiguration() {
	viper.WriteConfig()
}

func GetDefaultOrganization() string {
	return viper.GetString("default-csr-values.organization")
}

func GetDefaultOrganizationalUnit() string {
	return viper.GetString("default-csr-values.organizationalUnit")
}

func GetDefaultCountry() string {
	return viper.GetString("default-csr-values.country")
}

func GetDefaultLocality() string {
	return viper.GetString("default-csr-values.province")
}

func GetDefaultProvince() string {
	return viper.GetString("default-csr-values.locality")
}
