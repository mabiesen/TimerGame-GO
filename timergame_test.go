package main

import "testing"
import "fmt"

func TestInputScan(t *testing.T){
	fmt.Println("\nTest of input scan\n")
	scanForInput()
}

//Note: had issues running tests that look for user input.
//func TestRecurInput(t *testing.T){
//	recurCompareInput("x","x")
//}

func TestTimeMillis(t *testing.T){
	fmt.Println("\nTest of retrieving time since EPOCH\n")
	x := getCurrentTimeMillis()
	fmt.Println(x)
}

func TestArraySum(t *testing.T){
	fmt.Println("\nTest of array summation\n")
	sampleArray := []int64{22,12,10,6,1}
	x := sumOfArrayItems(sampleArray)
	testone := (x==51)
	fmt.Println(testone)
}

func TestArrayDisplay(t *testing.T){
	fmt.Println("\nTest of array display in people numbers\n")
	sampleArray := []int64{22,12,10,6,1}
	displayArrayAndIndexPlusOne(sampleArray, "Round")
}

func TestStartPrompts(t *testing.T){
	fmt.Println("\nTest of the start game prompts\n")
	startGamePrompts()
	fmt.Println("\n")
}

func TestStateChange(t *testing.T){
	fmt.Println("\nTest of toggle state change\n")
	x := changeToggleState("on","on","off")
	testone := (x=="off")
	fmt.Println(testone)

	y := changeToggleState("yes","no","yes")
	testtwo := (y=="no")
	fmt.Println(testtwo)
}

func TestAlertUserState(t *testing.T){
	fmt.Println("\nTest of user alert for game state change\n")
	alertUserStateChg()
}





//func TestEvaluateInput(t *testing.T){
	//inputOne := "x"
	//inputTwo := "y"
	//inputThree := "t"

//	fmt.Println("TestOne - test")
//	evaluateInput(inputOne)
//	fmt.Println("TestTwo - test")
//	evaluateInput(inputTwo)
//	fmt.Println("TestThree - antitest")
//	evaluateInput(inputThree)

//	fmt.Println("\nIf only TestThree printed 'That is not a valid input': pass")
//}
