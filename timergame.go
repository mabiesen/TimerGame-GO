package main

import "fmt"
import "time"
import "bufio"
import "os"
import "strings"

//Public variables
var onSlice []int64
var offSlice []int64
var totalTime int64 = 0

//Timer Game in Golang
//Game flow is as follows:  User is provided game directions and is prompted to
//start the game by pressing 's'.  During the game, user toggles on/off mode of
//timer.  the user can do this for as long as they like; when they have
//finished, the user types 'x'.  The game is then gracefully shut down and the
//user is shown:
//1.  total on and off time for the game
//2.  total on and off time by session
//3.  total game time
//4.  evaluation of whether all time was accounted for

//Project Method:
//1.  Used coding rules from Hongkiat, added addl specs for total func overview
//2.  Built personal library after code completion as a measure of refactoring
//3.  Created timergame_test.go file for function testing


//Project Notes:  Had a hell of a time with locking down data type for user input
//Project Notes:  Opted to use int64 for time since converted to milliseconds
//Project Notes:  Placement of endtime value in evaluateInput important
//Project Notes:  scanForInput was returning spaces. Why?  Added trim function to correct.
//Project Notes:  Whenver data type is in question, use go "reflect" pkg
//Project Notes:  Ideally, I would like to get rid of all public variables

//Additional Notes: Unused functions are found below the main function.
//retained for future projects

//Notes template to be used on Golang projects
//(@func, @var, etc.)
//@desc -
//@params - Line explaining each parameter
//@return - Line explaining each return
//@type -
//@pkg -
//@requiredFuncs if applicable
//@src if obtained in whole from another source




//Note: This will NOT be a goroutine
//@func - scanForInput
//@desc - Loop to retrieve user input, shave off end line
//@params - none
//@return 1. userInput String
//@type - none
//@pkg - "bufio", "os", "strings"
func scanForInput() string{
	fmt.Println("Waiting for input...")
	reader1 := bufio.NewReader(os.Stdin)
	userInput, _ := reader1.ReadString('\n')
	return strings.TrimSpace(userInput)
}


//@func - recurCompareInput
//@desc - Container function to wait for user to make the correct selection
//@params 1. comparator String
//@comparator in this case is 's'.  We need user to click 's' to start game
//@return - none
//@type - none
//@pkg - indirectly needed for requiredFuncs
//@requiredFuncs - scanForInput, compareInput
func recurCompareInput(comparator string){
	theInput := scanForInput()
	compareInput(theInput, comparator)
}

//@func - compareInput
//@desc -  Compares input to program criteria.  Calls recur function until expected output.
//@params 1. theInput string
//@theInput is the user input that is to be evaluated
//@params 2. comparator string
//@comparator is program specific to evaluate whether user input was expected.
//@return - none
//@type - none
//@pkg - none
//@requiredFuncs - recurCompareInput, scanForInput
func compareInput(theInput string, comparator string){
	switch theInput{
		case comparator: fmt.Println("Game starting in off mode.")
		default:
			fmt.Println("That is not a valid input.")
			recurCompareInput(comparator)
	}
}

//@func - alertUserStateChg
//@desc - Tells the user that they are now i on or off mode
//@params 1. toggle String
//@return - none
//@type - none
//@pkg - "fmt"
func alertUserStateChg(toggle string){
	fmt.Println("The Game is now in ", toggle, " mode.")
}

//@func - changeToggleState
//@desc - Change toggle from "on" to "off" etc.
//@params 1. toggle String
//@return 1. toggle String
//@type - none
//@pkg - none
func changeToggleState(toggle string) string{
	if toggle == "on"{
		toggle = "off"
	}else if toggle == "off"{
		toggle = "on"
	}
	return toggle
}

//@func - getCurrentTime
//@desc - Gets the current time since EPOCH in milliseconds AS INTEGER
//@params - none
//@return 1. currentTime int64
//@type - none
//@pkg - "time"
func getCurrentTime() int64{
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos/ 1000000
	currentTime := millis
	return currentTime
}

//@func - sumOfArrayItems
//@desc - Sums all items inside of array
//@params 1. array []int64
//@return 1. sum int64
//@type - none
//@pkg - none
func sumOfArrayItems(array []int64)int64{
	var sum int64 = 0
	for i := 0; i < len(array); i += 1 {
			sum = sum + array[i]
	}
	return sum
}

//@func - displayArrayAndIndex
//@desc - Displays array values with associated index numbers.
//@params 1. array []int64
//@return - none
//@type - none
//@pkg - "fmt"
func displayArrayAndIndex(array []int64){
	for i := 0; i < len(array); i += 1 {
			fmt.Println("Round ",i+1,": ",array[i])
	}
}

//This section holds the game containers
//ALL ITEMS BELOW THIS LINE ARE WRITTEN IN A PROGRAM SPECIFIC FASHION

//@func - startGamePrompts
//@desc - Used to print to terminal rules of the game
//@params - none
//@return - none
//@type - none
//@pkg - "fmt"
func startGamePrompts(){
	fmt.Println("This program will record all on/off time.")
	fmt.Println("Rules of the game: 1. Start the game, 2. Toggle, 3. End the Game")
	fmt.Println("Game commands may be found below:")
	fmt.Println("Start Game = s")
	fmt.Println("Toggle On/Off = t")
	fmt.Println("End Game = x")
	fmt.Println("Type 's' into the counsel and hit Enter to get started!")
}


//preGameContainer waits for user to select start
func preGameContainer(){
	startGamePrompts()
	theInput := scanForInput()
	compareInput(theInput, "s")		//When types 's', we move on to mainGameContainer
	beginGameContainer()
}

//on user starting game, sets variables
//This container will calculate total time for the game.
func beginGameContainer(){
	gameStartTime := getCurrentTime()
	toggle := "off"
	gameEndTime := recurGameContainer(toggle, gameStartTime)  //Will not come back until user quits game
	totalTime = gameEndTime - gameStartTime
	compareTime()
}

//this is the container where all the action will take place during game.
//evaluateInput does most of the heavy lifting
//container recurs until user indicates end of game
func recurGameContainer(toggle string, initStartTime int64) int64{
	theInput := scanForInput()
	endTime := evaluateInput(theInput, initStartTime, toggle)
	return endTime
}


//Note: This function feels really clunky
//Note: endTime here not acting as expected!
//@func - evaluateInput
//@desc - Evaluate user input during game for toggle or end game.
//@params 1. theInput String
//@theInput will be a single letter provided by user
//@return - none
//@type - none
//@pkg - none
func evaluateInput(theInput string, startTime int64, toggle string) int64{
	endTime := getCurrentTime()

	switch theInput {
	case "x":
		endTime = getCurrentTime()
		logTimeSpent(toggle, startTime, endTime)	//drops to end of switch statement
	case "t":
		logTimeSpent(toggle, startTime, endTime)
		toggle = changeToggleState(toggle)
		alertUserStateChg(toggle)
		recurGameContainer(toggle, endTime)
	default:
		fmt.Println("That is not a valid input.")
		recurGameContainer(toggle, startTime)		//returns start time as it was received
	}
	endTime = getCurrentTime()
	return endTime
}

//@func - logTimeSpent
//@desc - Used to log time spent on and off.  Store to public array.
//@params 1. toggle String
//@params 2. initStartTime int64
//@params 3. endTime int64
//@return - none - appends to public array
//@type - none
//@pkg - none
func logTimeSpent(toggle string, initStartTime int64, endTime int64){
	roundTime := endTime - initStartTime
	if toggle == "on"{
		//append to the on array
		onSlice = append(onSlice, roundTime)
		fmt.Println(roundTime)
		fmt.Println(onSlice)

	}else if toggle == "off"{
		//append to the off array
		offSlice = append(offSlice, roundTime)
		fmt.Println(roundTime)
		fmt.Println(offSlice)
	}
}

//@func - compareTime
//@desc - Compare totalTime to (on + off) time for quality control
//@params -
//@return -
//@type -
//@pkg -
func compareTime(){
	totalOnTime := sumOfArrayItems(onSlice)
	totalOffTime := sumOfArrayItems(offSlice)
	totalArrayTime := totalOnTime + totalOffTime
	variance := totalTime - totalArrayTime
	if variance < 5 && variance > (-5) {
		fmt.Println("Success! Variance within threshhold. Variance is ", variance, " milliseconds.")
		fmt.Println("Value of total game time: ", totalTime)
		fmt.Println("Value of total On time: ", totalOnTime)
		fmt.Println("Value of total Off time: ", totalOffTime)
		fmt.Println("Value of On + Off time: ", totalArrayTime)
	}else if totalTime != totalArrayTime{
		fmt.Println("Failure, the variance is: ", variance)
		fmt.Println("Value of total game time: ", totalTime)
		fmt.Println("Value of total On time: ", totalOnTime)
		fmt.Println("Value of total Off time: ", totalOffTime)
		fmt.Println("Value of On + Off time: ", totalArrayTime)
	}
}


func main(){
	preGameContainer()
}


//THIS FUNCTION DISREGARDED, TREATING ON/OFF AS STRING NOW
//@func -changeTruthState
//@desc -Turns true to false and false to true
//@params 1.toggle Boolean
//@return 1. toggleReturn Boolean
//toggle will be returned as opposite Boolean from when it was received.
//@type - none
//@pkg -none
//func changeTruthState(toggle bool) bool{
//	if toggle == true{
//		toggle = false
//	}else if toggle == false{
//		toggle = true
//	}
//	return toggle
//}


//THIS FUNCTION DISREGARDED, CONVERTING TIME TO INTEGER FOR WORK
//@func - roundNumber
//@desc - Rounds numbers to desired decimal place and on desired breaking point.
//@params 1. val float64
//@val is the number you want to round
//@params 2. roundOn float64
//@roundOn is an expression of the breaking point for rounding.  Normally .5
//@params 3. places integer
//@places is the number of decimal places you would like to round to
//@return 1. newVal float64
//@newVal is the rounded number
//@type - none
//@pkg - "math"
//@src - https://gist.github.com/DavidVaini/10308388
//func roundNumber(val float64, roundOn float64, places int ) float64 {
//	var round float64
//	pow := math.Pow(10, float64(places))
//	digit := pow * val
//	_, div := math.Modf(digit)
//	if div >= roundOn {
//		round = math.Ceil(digit)
//	} else {
//		round = math.Floor(digit)
//	}
//	newVal := round / pow
//	return newVal
//}


//@func - checkOnSliceSize
//@desc - Checks whether onSlice has reached capacity, if so doubles slice length
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func checkOnSliceSize(){
//	if cap(onSlice) == len(onSlice) {
//		doubleOnSliceSize()
//	}
//}

//@func - checkOffSliceSize
//@desc - Checks whether offSlice has reached capacity, if so doubles slice length
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func checkOffSliceSize(){
//	if cap(offSlice) == len(offSlice) {
//		doubleOffSliceSize()
//	}
//}

//@func - doubleOnSliceSize
//@desc - Double the size of onSlice when capacity is reached.
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func doubleOnSliceSize(){
//	newSlice := make([]int, len(onSlice), 2*cap(onSlice))
//	for i := range onSlice {
//			newSlice[i] = onSlice[i]
//	}
//	onSlice = newSlice
//}

//@func - doubleSliceSize
//@desc - Double the size of offSlice when capacity is reached.
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func doubleOffSliceSize(){
//	newSlice := make([]int, len(offSlice), 2*cap(offSlice))
//	for i := range offSlice {
//			newSlice[i] = offSlice[i]
//	}
//	offSlice = newSlice
//}

//@func - extendOnSlice
//@desc - Add new value to onSlice
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func extendOnSlice(element int64) []int64 {
//		n := len(onSlice)
//		onSlice[n] = element
//		return onSlice
//}

//@func - extendOffSlice
//@desc - Add new value to onSlice
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func extendOffSlice(element int64) []int64 {
//		n := len(offSlice)
//		offSlice = offSlice[0 : n+1]
//		offSlice[n] = element
//		return offSlice
//}

//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func appendOnSlice(items ...int){
//		for _, item := range items {
//				onSlice = extendOnSlice(item)
//		}
//}

//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
//@src - https://blog.golang.org/slices
//func appendOffSlice(items ...int){
//		for _, item := range items {
//				onSlice = extendOffSlice(item)
//		}
//}
