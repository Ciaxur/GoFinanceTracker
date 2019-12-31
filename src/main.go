package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	. "./DataStructure"

	"github.com/manifoldco/promptui"
)

// Handles PromptTUI Error
func handlePromptErr(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
}

func main() {
	// Load File Data and Variables
	data := LoadData("data.json")
	config := LoadConfig("config.json")
	reader := bufio.NewReader(os.Stdin)

	// Create Choices
	mainChoices := []string{"1.Add Data", "2.View Data", "3.Save Data", "4.Settings", "0.Exit"}
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
				fmt.Println("Month Found!")
				d.Print()
			} else { // NEW DATA BLOCK
				fmt.Printf("Creating new Block for '%s'\n", strDate)

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

					fmt.Print("Enter Payment: $")
					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &payment)

					if err != nil {
						fmt.Printf("Add Payment Error: %v\n", err.Error())
					} else if payment > 0 {
						// CALCULATE Savings/Liquid
						savings := payment * (config.Properties.Savings_Percentage / 100.00)
						liquid := payment * (config.Properties.Liquid_Precentage / 100.00)

						// ADD PAYMENT INFO
						d.Payment = append(d.Payment, payment)
						d.Savings = append(d.Savings, savings)
						d.Liquid = append(d.Liquid, liquid)

						// OUTPUT INFO
						fmt.Printf("Payment '$%.2f' Added!\n", payment)
						fmt.Printf("Savings/Liquid '($%.2f / $%.2f)' Calculated!\n", savings, liquid)

					} else {
						fmt.Printf("Payment '$%.2f' is Redundant, not added!\n", payment)
					}

				case addDataChoices[1]: // ADD TRANSACTION
					var transaction float32

					fmt.Print("Enter Transaction: $")
					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &transaction)

					if err != nil {
						fmt.Printf("Add Transaction Error: %v\n", err.Error())
					} else if transaction > 0 {
						fmt.Printf("Transaction '$%.2f' Added!\n", transaction)
						d.Transactions = append(d.Transactions, transaction)
					} else {
						fmt.Printf("Transaction '$%.2f' is Redundant, not added!\n", transaction)
					}

				case addDataChoices[2]: // GO BACK
					fmt.Println("Going back to Main Menu...")
					exitMenu = true

				default:
					fmt.Printf("Add Data Menu: Whoops, something went wrong!\n")
					return
				}

				// SHOW DATA BLOCK
				if exitMenu != true {
					fmt.Println("\n===== Block =====")
					d.Print()
				}
			}

		case mainChoices[1]: // VIEW DATA
			// LAUNCH MENU
			prompt = promptui.Select{
				Label: "View Menu",
				Items: viewChoices,
			}
			_, result, err = prompt.Run()
			handlePromptErr(err)

			// ACTION BASED ON SELECTION
			switch result {
			case viewChoices[0]: // DISPLAY LATEST MONTH
				ViewData(data, 1)

			case viewChoices[1]: // ASK FOR SPECIFIC MONTH + DISPLAY
				// Get Month from User
				fmt.Printf("Enter Specific Month: ")
				var month int
				text, _ := reader.ReadString('\n')
				_, err := fmt.Sscanf(text, "%d", &month)

				if err != nil {
					fmt.Printf("Specific Month Error: %v\n", err.Error())
				} else {
					// Find Month & Display
					d := FindDataMonth(data, month)
					if d != nil {
						d.Print()
					} else {
						fmt.Println("No Data Found for Month!")
					}
				}

			case viewChoices[2]: // VIEW LAST 3 MONTHS
				ViewData(data, 3)

			case viewChoices[3]: // RETURN TO PREVIOUS MENU
				fmt.Println("Returning to Main Menu...")
			default:
				fmt.Printf("View Menu: Whoops, something went wrong!\n")
				return

			}

		case mainChoices[2]: // SAVE DATA
			err1, err2 := SaveData(data, config, "data.json", "config.json")
			if err1 != nil {
				fmt.Printf("Data Save Failed! %v\n", err1)
			} else if err2 != nil {
				fmt.Printf("Config Save Failed! %v\n", err2)
			} else {
				fmt.Printf("Data and Config Saved!\n")
			}

		case mainChoices[3]: // SETTINGS
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
					fmt.Println("==== Configuration Settings ====")
					fmt.Printf("Data Length: %d\n", config.Properties.Data_Length)
					fmt.Printf("Savings: %.2f%%\n", config.Properties.Savings_Percentage)
					fmt.Printf("Liquid: %.2f%%\n\n", config.Properties.Liquid_Precentage)

				case settingChoices[1]: // CHANGE SAVINGS %
					// OBTAIN NEW SAVINGS PERCENTAGE
					var savingsP float32

					fmt.Print("New Savings %")
					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &savingsP)

					if err != nil {
						fmt.Printf("Settings Savings Error: %v\n", err)
					} else {
						fmt.Printf("Savings Set to '%.2f%%'\n\n", savingsP)
						config.Properties.Savings_Percentage = savingsP

						// UPDATE LIQUID TO MATCH 100%
						config.Properties.Liquid_Precentage = 100.00 - savingsP
					}

				case settingChoices[2]: // CHANGE LIQUID %
					// OBTAIN NEW LIQUID PERCENTAGE
					var liquidP float32

					fmt.Print("New Liquid %")
					text, _ := reader.ReadString('\n')
					_, err := fmt.Sscanf(text, "%f", &liquidP)

					if err != nil {
						fmt.Printf("Settings Liquid Error: %v\n", err)
					} else {
						fmt.Printf("Liquid Set to '%.2f%%'\n\n", liquidP)
						config.Properties.Liquid_Precentage = liquidP

						// UPDATE LIQUID TO MATCH 100%
						config.Properties.Savings_Percentage = 100.00 - liquidP
					}

				case settingChoices[3]: // GO BACK
					fmt.Println("Returning to Main Menu...")
					exitMenu = true
				default:
					fmt.Printf("Settings Menu: Whoops, something went wrong!\n")
					return
				}
			}

		case mainChoices[4]: // EXIT
			fmt.Print("Exiting...\n")
			return

		default:
			fmt.Printf("Main Menu: Whoops, something went wrong!\n")
			return
		}

		fmt.Println()
	}
}
