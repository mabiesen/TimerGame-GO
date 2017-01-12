Not all time in the game has been accounted for

The error lies in the definition of endTime.  endTime is defined in the evaluate input function:

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
