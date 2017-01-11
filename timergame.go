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


//Notes template to be used on Golang projects
//(@func, @var, etc.)
//@desc -
//@params -
//@return -
//@type -
//@pkg -


//@func - userPrompts
//@desc - Used to print to terminal rules of the game
//@params - The function takes no parameters
//@return - The function has not return value
//@type - is not the method of a specific type
//@pkg - "fmt"	- This package requires "fmt" in order to function
func startGamePrompts(){
	fmt.Println("This program will record all on/off time.")
	fmt.Prinln("Rules of the game: 1. Start the game, 2. Toggle, 3. End the Game")
	fmt.Println("Game commands may be found below:")
	fmt.Println("Start Game = s")
	fmt.Println("Toggle On/Off = t")
	fmt.Println("End Game = x")
	fmt.Println("Type 's' into the counsel and hit Enter to get started!")
}

//Note: This will NOT be a goroutine
//@func - scanForInputDist
//@desc - Loop to retrieve user input
//@params 1. initStartTime float
		//initStartTime holds the time at start of game until changed via channel
//@return -
//@type -
//@pkg - "bufio", "os"
func scanForInput(initStartTime float64){
	reader1 := bufio.NewReader(os.Stdin)
	userInput, _ := reader1.ReadString('\n')
	return userInput
}

//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
func evaluateInput(theInput string, start string, end string){
	switch theInput {
	case start:
	case end:
	case true:
	case false:
	case default: fmt.Println("That is not a valid input.  If game not started, run again")
	}
}

//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
func recurStartInputScan(){
	scanForInput()
	evaluateStartInput()
}


//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
func evaluateStartInput(theInput string, startVar string){
	switch theInput{
	case startVar:
	default : recurStartInputScan();
	}
}

//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
func countOnOffTime(toggle bool, initStartTime float64){

}

//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
func alertUserStateChg(toggle bool){
	fmt.Println("The Game is now in ", toggle, " mode.")
}

//@func -changeToggleState
//@desc -Turns true to false and false to true
//@params 1.toggle Boolean
//@return 1. toggle
//toggle will be returned as opposite Boolean from when it was received.
//@type - none
//@pkg -none
func changeToggleState(toggle bool){
	if toggle == true{
		toggle = false
	}else if toggle == false{
		toggle = true
	}
	return toggle
}

func getCurrentTime(){

}

//@func - Container to start the game
//@desc -
//@params -
//@return -
//@type -
//@pkg -
func startGameContainer(){
	startGamePrompts()
	theInput := scanForInput()
	evaluateStartInput(theInput)

}


func mainGameContainer(initStartTime float64){

}

//@func -
//@desc -
//@params -
//@return -
//@type -
//@pkg -
func main(){

}
