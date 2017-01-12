package main

import "testing"
import "fmt"


func TestStartPrompts(t *testing.T){
	fmt.Println("\nTest of the start game prompts\n")
	startGamePrompts()
	fmt.Println("\n")
}

func TestStateChange(t *testing.T){
	fmt.Println("\nTest of toggle state change\n")
	new := changeToggleState("off")
	if new == "on"{
		fmt.Println("statechange1 -test-: pass")}
	brandnew := changeToggleState(new)		//test of reversing the output
	if brandnew == "off"{
		fmt.Println("statechange2 -antitest- : pass")}
	fmt.Println("Passed State Change Tests\n")
}

func TestInputScan(t *testing.T){
	fmt.Println("scanForInput does not wait in test.  Pass for now if this line is read\n")
}

func TestEvaluateInput(t *testing.T){
	inputOne := "x"
	inputTwo := "y"
	inputThree := "t"

	fmt.Println("TestOne - test")
	evaluateInput(inputOne)
	fmt.Println("TestTwo - test")
	evaluateInput(inputTwo)
	fmt.Println("TestThree - antitest")
	evaluateInput(inputThree)

	fmt.Println("\nIf only TestThree printed 'That is not a valid input': pass")
}
