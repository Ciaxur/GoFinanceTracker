package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"

	. "./DataStructure"
	"./Utils"

	"github.com/manifoldco/promptui"
)

const (
	dataPath = "data"
)

// Handles PromptTUI Error
func handlePromptErr(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
}

/**
 * Display Files available to load
 *  for the user to choose, or default
 *  file
 *
 * @param dirpath - The directory path that the data is stored in
 * @returns File Name to Load
 */
func displayDataFiles(dirpath string) string {
	// Open Directory
	dataDir, err := os.Open(dirpath)
	defer dataDir.Close()

	// File to Load
	var fileToLoad string

	// If Directory not Found, create one
	if err != nil {
		Utils.Out.Info.Println("Data Directory not found, creating one")
		os.Mkdir("data", 0755)
	} else {
		// Display which Data to load
		list, err := dataDir.Readdirnames(-1)
		sort.Sort(sort.Reverse(sort.StringSlice(list))) // Sort in Reverse (Latest Year first)

		if err != nil {
			println(err.Error())
		} else {
			if len(list) == 0 { // No Files, load defaul so return nothing
				return ""
			} else if len(list) == 1 { // If there is only ONE file, return that
				return list[0]
			}

			// Display Prompt to get which Data to Load
			prompt := promptui.Select{
				Label: "Load Data File",
				Items: list,
			}
			_, fileToLoad, err = prompt.Run()
			handlePromptErr(err)
		}
	}
	// Return the Data File to Load
	return fileToLoad
}

func main() {
	// Load File Data and Variables
	data := LoadData(dataPath + "/" + displayDataFiles(dataPath))
	config := LoadConfig("config.json")
	reader := bufio.NewReader(os.Stdin)

	// Create Choices
	mainChoices := []string{"1.Add Data", "2.View Data", "3.Summary", "4.Save Data", "5.Settings", "0.Exit"}
	viewChoices := []string{"1.Latest Month", "2.Specific Month", "3.View Last 3 Months", "0.Go Back"}
	addDataChoices := []string{"1.Add Payment", "2.Add Transaction", "0.Go Back"}
	settingChoices := []string{"1.View", "2.Change Savings %", "3.Change Liquid %", "0.Go Back"}

	// MAIN MENU
	for {
		// Create the Prompt
		prompt := promptui.Select{
			Label: "Main Menu",
			Items: mainChoices,
		}
		_, result, err := prompt.Run()

		// Handle Error
		handlePromptErr(err)

		// Go through Main Menu Choices
		switch result {
		case mainChoices[0]: // ADD DATA
			// MENU LOOP
			exitMenu := false

			// GET TODYA'S DATE
			todayDate := time.Now()
			year, month, day := todayDate.Date()
			strDate := fmt.Sprintf("%d/%d/%d", month, day, year)

			// CHECK IF DATA PREV FOUND
			d := FindDataMonth(data, int(month))
			if d != nil { // BLOCK FOUND
				Utils.Out.Info.Println("\nMonth Found!")

				d.Print()
			} else { // NEW DATA BLOCK
				Utils.Out.Info.Printf("Creating new Block for '%s'\n", strDate)

				// Create new Data Block
				d = &DataBlock{}
				AddData(data, d)

				// Allocate Data
				d.Date = make([]string, 0, 2)
				d.Payment = make([]float32, 0, 2)
				d.Savings = make([]float32, 0, 2)
				d.Liquid = make([]float32, 0, 2)
				d.Transactions = make([]float32, 0, 5)

				// Add Date & Month
				d.Date = append(d.Date, strDate)
				d.Month = int(month)
			}

			// GET OPTION
			for exitMenu == false {
				prompt := promptui.Select{
					Label: "Add Data",
					Items: addDataChoices,
				}
				_, result, err := prompt.Run()
				handlePromptErr(err)

				switch result {
				case addDataChoices[0]: // ADD PAYMENT
					var payment float32

					Utils.Out.Prompt.Print("Enter Payment: $")

					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &payment)

					if err != nil {
						Utils.Out.Error.Printf("Add Payment Error: %v\n", err.Error())
					} else if payment > 0 {
						// CALCULATE Savings/Liquid
						savings := payment * (config.Properties.SavingsPercentage / 100.00)
						liquid := payment * (config.Properties.LiquidPercentage / 100.00)

						// ADD PAYMENT INFO
						d.Payment = append(d.Payment, payment)
						d.Savings = append(d.Savings, savings)
						d.Liquid = append(d.Liquid, liquid)

						// OUTPUT INFO
						Utils.Out.Info.Printf("Payment '$%.2f' Added!\n", payment)
						Utils.Out.Info.Printf("Savings/Liquid '($%.2f / $%.2f)' Calculated!\n", savings, liquid)

					} else {
						Utils.Out.Warning.Printf("Payment '$%.2f' is Redundant, not added!\n", payment)
					}

				case addDataChoices[1]: // ADD TRANSACTION
					var transaction float32

					Utils.Out.Prompt.Print("Enter Transaction: $")
					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &transaction)

					if err != nil {
						Utils.Out.Error.Printf("Add Transaction Error: %v\n", err.Error())
					} else if transaction > 0 {
						Utils.Out.Info.Printf("Transaction '$%.2f' Added!\n", transaction)
						d.Transactions = append(d.Transactions, transaction)
					} else {
						Utils.Out.Warning.Printf("Transaction '$%.2f' is Redundant, not added!\n", transaction)
					}

				case addDataChoices[2]: // GO BACK
					fmt.Println("Going back to Main Menu...")
					exitMenu = true

				default:
					Utils.Out.Error.Printf("Add Data Menu: Whoops, something went wrong!\n")
					return
				}

				// SHOW DATA BLOCK
				if exitMenu != true {
					Utils.Out.Important.Println("\n===== Block =====")
					d.Print()
				}
			}

		case mainChoices[1], mainChoices[2]: // VIEW DATA	/ SUMMARY
			// SUMMARY MODE
			var summaryMode bool
			promptLabel := "View Menu"
			if result == mainChoices[2] {
				summaryMode = true
				promptLabel = "Summary Menu"
			}

			// LAUNCH MENU
			prompt = promptui.Select{
				Label: promptLabel,
				Items: viewChoices,
			}
			_, result, err = prompt.Run()
			handlePromptErr(err)

			// ACTION BASED ON SELECTION
			switch result {
			case viewChoices[0]: // DISPLAY LATEST MONTH
				if len(data.Block) > 0 { // Verify there is data

					if summaryMode {
						// Obtain Block
						block := data.Block[len(data.Block)-1]
						block.PrintSummary()
					} else {
						ViewData(data, 1, nil)
					}

				}

			case viewChoices[1]: // ASK FOR SPECIFIC MONTH + DISPLAY
				// Get Month from User
				Utils.Out.Prompt.Printf("Enter Specific Month: ")
				var month int
				text, _ := reader.ReadString('\n')
				_, err := fmt.Sscanf(text, "%d", &month)

				if err != nil {
					Utils.Out.Error.Printf("Specific Month Error: %v\n", err.Error())
				} else {
					// Find Month & Display
					d := FindDataMonth(data, month)
					if d != nil {
						if summaryMode {
							d.PrintSummary()
						} else {
							d.Print()
						}
					} else {
						Utils.Out.Warning.Println("No Data Found for Month!")
					}
				}

			case viewChoices[2]: // VIEW LAST 3 MONTHS
				if summaryMode {
					ViewData(data, 3, func(b *DataBlock) {
						b.PrintSummary()
						fmt.Println()
					})
				} else {
					ViewData(data, 3, nil)
				}

			case viewChoices[3]: // RETURN TO PREVIOUS MENU
				fmt.Println("Returning to Main Menu...")
			default:
				Utils.Out.Error.Printf("View Menu: Whoops, something went wrong!\n")
				return

			}

		case mainChoices[3]: // SAVE DATA
			// GENERATE SAVE FILE NAME BASED ON MONTH DATE
			todayDate := time.Now()
			year, _, _ := todayDate.Date()
			fileName := fmt.Sprintf("%d.json", year)

			// Strip out Previous Year to save into single Year File
			var currentData Data
			currentData.Block = make([]*DataBlock, 0, 12)

			// Add Current Year Data into Block to Save
			for _, block := range data.Block {
				if len(block.Date) > 0 {
					// Obtain each Block's Year
					re := regexp.MustCompile(`[0-9]*$`)
					bYear, err := strconv.Atoi(string(re.Find([]byte(block.Date[0]))))

					// Handle Error, if any
					if err != nil {
						Utils.Out.Error.Println("Save Data Error!")
						panic(err)
					}

					// If Same Year, Append to Data that'll be saved
					if bYear == year {
						currentData.Block = append(currentData.Block, block)
					}
				}
			}

			// SAVE THE FILE
			err1, err2 := SaveData(&currentData, config, dataPath+"/"+fileName, "config.json")
			if err1 != nil {
				Utils.Out.Error.Printf("Data Save Failed! %v\n", err1)
			} else if err2 != nil {
				Utils.Out.Error.Printf("Config Save Failed! %v\n", err2)
			} else {
				Utils.Out.Info.Printf("Data and Config Saved!\n")
			}

		case mainChoices[4]: // SETTINGS
			// VARIABLES USED
			exitMenu := false

			// LAUNCH MENU
			for exitMenu != true {
				prompt := promptui.Select{
					Label: "Settings",
					Items: settingChoices,
				}
				_, result, err = prompt.Run()
				handlePromptErr(err)

				// ACTION BASED ON SELECTION
				switch result {
				case settingChoices[0]: // VIEW SETTINGS
					Utils.Out.Info.Println("==== Configuration Settings ====")
					fmt.Printf("Data Length: %d\n", config.Properties.DataLength)
					fmt.Printf("Savings: %.2f%%\n", config.Properties.SavingsPercentage)
					fmt.Printf("Liquid: %.2f%%\n", config.Properties.LiquidPercentage)
					fmt.Printf("Invest: %.2f%%\n\n", config.Properties.InvestPercentage)

				case settingChoices[1]: // CHANGE SAVINGS %
					// OBTAIN NEW SAVINGS PERCENTAGE
					var savingsP float32

					Utils.Out.Prompt.Print("New Savings %")
					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &savingsP)

					if err != nil {
						Utils.Out.Error.Printf("Settings Savings Error: %v\n", err)
					} else {
						Utils.Out.Info.Printf("Savings Set to '%.2f%%'\n\n", savingsP)
						config.Properties.SavingsPercentage = savingsP

						// UPDATE LIQUID TO MATCH 100%
						config.Properties.LiquidPercentage = 100.00 - savingsP
					}

				case settingChoices[2]: // CHANGE LIQUID %
					// OBTAIN NEW LIQUID PERCENTAGE
					var liquidP float32

					Utils.Out.Prompt.Print("New Liquid %")
					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &liquidP)

					if err != nil {
						Utils.Out.Error.Printf("Settings Liquid Error: %v\n", err)
					} else {
						Utils.Out.Info.Printf("Liquid Set to '%.2f%%'\n\n", liquidP)
						config.Properties.LiquidPercentage = liquidP

						// UPDATE LIQUID TO MATCH 100%
						config.Properties.SavingsPercentage = 100.00 - liquidP
					}

				case settingChoices[3]: // GO BACK
					fmt.Println("Returning to Main Menu...")
					exitMenu = true
				default:
					Utils.Out.Error.Printf("Settings Menu: Whoops, something went wrong!\n")
					return
				}
			}

		case mainChoices[5]: // EXIT
			Utils.Out.Info.Print("Exiting...\n")
			return

		default:
			Utils.Out.Error.Printf("Main Menu: Whoops, something went wrong!\n")
			return
		}

		fmt.Println()
	}
}
