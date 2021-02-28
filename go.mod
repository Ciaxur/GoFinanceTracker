module GoFinanceTracker

go 1.16

replace (
	DataStructure => ./src/DataStructure
	Utils => ./src/Utils
)

require (
	DataStructure v0.0.0-00010101000000-000000000000
	Utils v0.0.0-00010101000000-000000000000
	github.com/manifoldco/promptui v0.8.0
)
