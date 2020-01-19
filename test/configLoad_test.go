package test

import (
	"fmt"
	"testing"

	. "../src/DataStructure"
)

// Returns Individual Configuaration Properties Data
func getConfigVals(confP *ConfigProperties) (int, float32, float32, float32) {
	return confP.DataLength, confP.SavingsPercentage, confP.LiquidPercentage, confP.InvestPercentage
}

// DEBUG: Print out Configuration Values
func printConfigVals(confP *ConfigProperties) {
	fmt.Printf("DataLength: %d\n", confP.DataLength)
	fmt.Printf("Savings Percentage: %.2f\n", confP.SavingsPercentage)
	fmt.Printf("Liquid Percentage: %.2f\n", confP.LiquidPercentage)
	fmt.Printf("Invest Percentage: %.2f\n", confP.InvestPercentage)
}

// Make Sure Loading Configuration Method works Properly
func TestLoadConfig(t *testing.T) {
	// Configuration 1: Without Investage Percentage
	conf := LoadConfig("data/configs/config1.json")
	dLength, sPerc, lPerc, iPerc := getConfigVals(conf.Properties)

	if dLength != 1 || sPerc != 80. || lPerc != 20. || iPerc != 0. {
		t.Errorf("Config1 Error: [%d, %.2f, %.2f, %.2f]", dLength, sPerc, lPerc, iPerc)
	}

	// Configuration 2: Without Any JSON Data (Should be Default Values)
	conf = LoadConfig("data/configs/config2.json")
	dLength, sPerc, lPerc, iPerc = getConfigVals(conf.Properties)

	if dLength != 0 || sPerc != 40. || lPerc != 40. || iPerc != 20. {
		t.Errorf("Config1 Error: [%d, %.2f, %.2f, %.2f]", dLength, sPerc, lPerc, iPerc)
	}

	// Configuration 3: Load all Properties Correctly
	conf = LoadConfig("data/configs/config3.json")
	dLength, sPerc, lPerc, iPerc = getConfigVals(conf.Properties)

	printConfigVals(conf.Properties)
	if dLength != 1 || sPerc != 40. || lPerc != 40. || iPerc != 20. {
		t.Errorf("Config1 Error: [%d, %.2f, %.2f, %.2f]", dLength, sPerc, lPerc, iPerc)
	}
}
