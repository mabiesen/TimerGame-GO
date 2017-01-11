package main

import "fmt"
import "time"
import "bufio"
import "os"

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


//Project Notes:  Had a hell of a time with locking down data type for user input
//Project Notes:  Opted to use int64 for time since converted to milliseconds


//Notes template to be used on Golang projects
//(@func, @var, etc.)
//@desc -
//@params - Line explaining each parameter
//@return - Line explaining each return
//@type -
//@pkg -
//@requiredFuncs if applicable
//@src if obtained in whole from another source


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

//Note: This will NOT be a goroutine
//@func - scanForInput
//@desc - Loop to retrieve user input, shave off end line
//@params - none
//@return 1. userInput String
//@type - none
//@pkg - "bufio", "os"
func scanForInput() string{
	reader1 := bufio.NewReader(os.Stdin)
	userInput, _ := reader1.ReadString('\n')
	return userInput
}

//@func - evaluateInput
//@desc - Evaluate user input during game for toggle or end game.
//@params 1. theInput String
//@theInput will be a single letter provided by user
//@return - none
//@type - none
//@pkg - none
func evaluateInput(theInput string){
	switch theInput {
	case "x":
	case "t":
	default: fmt.Println("That is not a valid input.")
	}
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
func compareInput(theInput string, comparator string){
	switch theInput{
		case comparator:
		default : recurCompareInput(comparator);
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

//This section holds the game containers

//preGameContainer waits for user to select start
func preGameContainer(){
	startGamePrompts()
	theInput := scanForInput()
	compareInput(theInput, "s")		//When types 's', we move on to mainGameContainer
	beginGameContainer()
}

//on user starting game, sets variables
func beginGameContainer(){
	gameStartTime := getCurrentTime()
	toggle := "off"
	inGameContainer(toggle, gameStartTime)
}

//this is the container where all the action will take place
func inGameContainer(toggle string, initStartTime int64){
	theInput := scanForInput()


func main(){

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
