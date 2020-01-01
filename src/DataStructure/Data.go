package DataStructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Printf("==== Month: %d ====\n", d.Month)
	// fmt.Printf("- Date: %v\n", d.Date)
	fmt.Print("- Payments: ")
	printFloatArr(d.Payment, "$%.2f ", "\n")
	fmt.Print("- Savings: ")
	printFloatArr(d.Savings, "$%.2f ", "\n")
	fmt.Print("- Liquid: ")
	printFloatArr(d.Liquid, "$%.2f ", "\n")
	fmt.Print("- Transactions: ")
	printFloatArr(d.Transactions, "$%.2f ", "\n")

	fmt.Println()
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
func ViewData(data *Data, numData int) {
	// Validate Number of Elements
	if numData > len(data.Block) {
		numData = len(data.Block)
	}

	// Splice Block to View
	block := data.Block[len(data.Block)-numData:]
	for i, _ := range block {
		fmt.Printf("Data Block [%v]:\n", i)
		block[i].Print()
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
