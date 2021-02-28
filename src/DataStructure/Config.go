package DataStructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config - Configuration file Structure from JSON
type Config struct {
	Properties *ConfigProperties `json:"config"`
}

// ConfigProperties - JSON Config Struct
type ConfigProperties struct {
	DataLength        int     `json:"data_length"`
	SavingsPercentage float32 `json:"savings_percentage"`
	LiquidPercentage  float32 `json:"liquid_percentage"`
	InvestPercentage  float32 `json:"invest_percentage"`
}

// DEFAULT VALUES (not exposed) -> (40/20/20)
const savingsP float32 = 40.00
const investP float32 = 20.00
const liquidP float32 = 100.00 - (savingsP + investP)

// LoadConfig \
//  Loads Configuration Section from JSON File
//  @param fileName - The Name of the file that's being read
//  @return Config Object of data stored
func LoadConfig(fileName string) *Config {
	var config Config

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Load Config: No Config File, loading Default!\n")
		config.Properties = &ConfigProperties{
			DataLength:        0,
			SavingsPercentage: savingsP,
			LiquidPercentage:  liquidP,
			InvestPercentage:  investP,
		}
	} else {
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)

		err = json.Unmarshal(byteValue, &config)

		// Validate Savings/Liquid Ratio
		if config.Properties.LiquidPercentage+config.Properties.SavingsPercentage != 100.00 {
			config.Properties.LiquidPercentage = liquidP
			config.Properties.SavingsPercentage = savingsP
			config.Properties.InvestPercentage = investP
		}
	}

	return &config
}
