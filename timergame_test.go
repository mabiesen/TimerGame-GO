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

func TestCompareInput(t *testing.T){
	compareInput("x","x")
	compareInput("x","y")
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
