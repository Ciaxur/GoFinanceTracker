package DataStructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration file Structure from JSON
type Config struct {
	Properties *ConfigProperties `json:"config"`
}
type ConfigProperties struct {
	Data_Length        int     `json:"data_length"`
	Savings_Percentage float32 `json:"savings_percentage"`
	Liquid_Precentage  float32 `json:"liquid_percentage"`
}

// DEFAULT VALUES (not exposed) -> (80/20)
const savingsP float32 = 80.00
const liquidP float32 = 100.00 - savingsP

/**
 * Loads Configuration Section from JSON File
 * @param fileName - The Name of the file that's being read
 * @return Config Object of data stored
 */
func LoadConfig(fileName string) *Config {
	var config Config

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Load Config: No Config File, loading Default!\n")
		config.Properties = &ConfigProperties{
			Data_Length:        0,
			Savings_Percentage: savingsP,
			Liquid_Precentage:  liquidP,
		}
	} else {
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)

		err = json.Unmarshal(byteValue, &config)

		// Validate Savings/Liquid Ratio
		if config.Properties.Liquid_Precentage+config.Properties.Savings_Percentage != 100.00 {
			config.Properties.Liquid_Precentage = liquidP
			config.Properties.Savings_Percentage = savingsP
		}
	}

	return &config
}
