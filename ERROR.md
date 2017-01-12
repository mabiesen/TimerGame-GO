Not all time in the game has been accounted for. 

The game is showing, on average, a variance of two milliseconds.  Occasionally I see variances as high as 6 milliseconds.

endTime is defined in the evaluate input function.  The error lies in the value of endTime when case "x" occurs.  

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
      

Variants to the above equation:
* If I remove endtime statement from line 23, the variance increases.
* If I remove endtime statement from line 12, the variance increases.
* If I remove both, the variance increases.

This is very odd.  One would assume that defining end time above the switch statement would be sufficient.
