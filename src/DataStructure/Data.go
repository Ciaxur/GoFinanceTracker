package DataStructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../Utils"
)

type Data struct {
	Block []*DataBlock `json:"data"`
}

type DataBlock struct {
	Month        int       `json:"month"`
	Date         []string  `json:"date"`
	Payment      []float32 `json:"pay"`
	Savings      []float32 `json:"savings"`
	Liquid       []float32 `json:"liquid"`
	Transactions []float32 `json:"transactions"`
}

/**
 * Prints Float32 Array with given format
 * @param arr - Array to Print
 * @param formatStr - String Format to Use for Printing
 * @param end - String at the End of Printing Entire Array
 */
func printFloatArr(arr []float32, formatStr string, end string) {
	fmt.Print("[ ")
	for _, val := range arr {
		fmt.Printf(formatStr, val)
	}
	fmt.Print("]" + end)
}

/**
 * DataBlock Method
 *  Prints DataBlock
 */
func (d *DataBlock) Print() {
	// Output Header
	Utils.Out.Info.Print("====== ")
	Utils.Out.Error.Print(d.Date)
	Utils.Out.Info.Println(" ======")

	// Output Data
	Utils.Out.Info.Print("- Payments: \t\t")
	printFloatArr(d.Payment, "$%.2f ", "\n")
	Utils.Out.Info.Print("- Savings: \t\t")
	printFloatArr(d.Savings, "$%.2f ", "\n")
	Utils.Out.Info.Print("- Liquid: \t\t")
	printFloatArr(d.Liquid, "$%.2f ", "\n")
	Utils.Out.Info.Print("- Transactions: \t")
	printFloatArr(d.Transactions, "$%.2f ", "\n")

	fmt.Println()
}

/**
 * DataBlock Method
 *  Prints Out Block Summary
 *  	- Data Summation
 *		- Data Usage
 */
func (d *DataBlock) PrintSummary() {
	// CALCULATE RESULTS
	var totalPay, totalSavings, totalLiquids, totalTransactions float32 // Summations
	var liquidUsed, savingsUsed float32                                 // Usage
	var liquidFlag, savingsFlag bool                                    // Over-Usage Flags

	// Sum all Data
	for _, v := range d.Payment {
		totalPay += v
	}
	for _, v := range d.Savings {
		totalSavings += v
	}
	for _, v := range d.Liquid {
		totalLiquids += v
	}
	for _, v := range d.Transactions {
		totalTransactions += v
	}

	// Usage
	liquidUsed = totalLiquids - totalTransactions
	if liquidUsed < 0 { // Used Savings!
		liquidFlag = true                       // Used up all Liquid
		savingsUsed = totalSavings + liquidUsed // Carry to Savings
		liquidUsed = 0                          // Used up ALL Liquid

		if savingsUsed < 0 { // Over-Usage! Used up Savings!
			savingsFlag = true
		}
	}

	// Output Header
	Utils.Out.Info.Print("====== ")
	Utils.Out.Error.Print(d.Date)
	Utils.Out.Info.Println(" ======")

	// Output Sum Results
	Utils.Out.Info.Print("Total Payments: \t$")
	Utils.Out.Important.Println(totalPay)
	Utils.Out.Info.Print("Total Savings: \t\t$")
	Utils.Out.Important.Println(totalSavings)
	Utils.Out.Info.Print("Total Liquid: \t\t$")
	Utils.Out.Important.Println(totalLiquids)
	Utils.Out.Info.Print("Total Transactions: \t$")
	Utils.Out.Important.Println(totalTransactions, "\n")

	// Output Usage Results
	Utils.Out.Info.Print("Usage Liquid: \t\t$")
	if liquidFlag { // Over-Used Liquid
		Utils.Out.Important.Print(liquidUsed)
		Utils.Out.Error.Println(" [OVER-USED]")
	} else {
		Utils.Out.Important.Println(liquidUsed)
	}

	Utils.Out.Info.Print("Usage Savings: \t\t$")
	if savingsFlag { // Over-Used Savings
		Utils.Out.Important.Print(savingsUsed)
		Utils.Out.Error.Println(" [OVER-USED]")
	} else {
		Utils.Out.Important.Println(savingsUsed, "\n")
	}
}

/**
 * Loads Data from JSON File
 * @param fileName - The Name of the file that's being read
 * @return Data Object of data stored
 */
func LoadData(fileName string) *Data {
	var data Data

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Load Data: No Data File, loading Default!\n")
	} else {

		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)

		err = json.Unmarshal(byteValue, &data)
		// // data.Block = make([]data.Block, )
		// fmt.Printf("Data Size: %d\n", len(data.Block))
	}

	return &data
}

/**
 * Saves Data and COnfig into JSON Files
 * @param data - Pointer to data Object
 * @param config - Pointer to config Object
 * @param configFileName - The Name of the config file that's being saved
 * @param dataFileName - The Name of the data file that's being saved
 */
func SaveData(data *Data, config *Config, dataFileName string, configFileName string) (error, error) {
	// Update some Config Data
	config.Properties.Data_Length = len(data.Block)

	// Update Data in JSON File
	file, err := json.MarshalIndent(data, "", "\t")
	file2, err2 := json.MarshalIndent(config, "", "\t")

	if err == nil && err2 == nil {
		_ = ioutil.WriteFile(dataFileName, file, 0644)
		_ = ioutil.WriteFile(configFileName, file2, 0644)
	}

	return err, err2
}

/**
 * Adds Data to Objects
 * @param data - Pointer to data Object
 * @param block - Block to Add
 */
func AddData(data *Data, block *DataBlock) {
	// Add new Block
	data.Block = append(data.Block, block)
}

/**
 * Displays Loaded Data
 * @param data - Pointer to data Object
 * @param numData - The Last 'n' Elements
 */
func ViewData(data *Data, numData int, fn func(b *DataBlock)) {
	// Validate Number of Elements
	if numData > len(data.Block) {
		numData = len(data.Block)
	}

	// Splice Block to View
	block := data.Block[len(data.Block)-numData:]
	for i := range block {
		if fn != nil {
			fn(block[i])
		} else {
			Utils.Out.Info.Print("Data Block [")
			Utils.Out.Important.Print(i)
			Utils.Out.Info.Println("]:")

			block[i].Print()
		}
	}
}

/**
 * Searches for given Month in Data and
 *  returns pointer to data
 * @param data - Pointer to the Data
 * @param month - The month to look for
 */
func FindDataMonth(data *Data, month int) *DataBlock {
	// Search backward (from latest month)
	// TODO: Binary Search through Data
	for i := len(data.Block) - 1; i >= 0; i-- {
		if month == data.Block[i].Month {
			return data.Block[i]
		}
	}

	// NOT FOUND :(
	return nil
}
