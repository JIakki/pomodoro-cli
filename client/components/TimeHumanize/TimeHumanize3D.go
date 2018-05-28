package humanize

import (
	"fmt"
	"strings"
)

const maxLetterHeight = 10

var linesOfNumbers [10][]string
var linesOfColon []string

const colon = `
 ___    
|\__\   
\|__|   
    ___ 
   |\__\
   \|__|
        
`

var numbersMap = []string{
	`
 ________     
|\   __  \    
\ \  \|\  \   
 \ \  \\\  \  
  \ \  \\\  \ 
   \ \_______\
    \|_______|
	`,
	`
  _____     
 / __  \    
|\/_|\  \   
\|/ \ \  \  
     \ \  \ 
      \ \__\
       \|__|
		`,
	`
  _______     
 /  ___  \    
/__/|_/  /|   
|__|//  / /   
    /  /_/__  
   |\________\
    \|_______|`,
	`
 ________     
|\_____  \    
\|____|\ /_   
      \|\  \  
     __\_\  \ 
    |\_______\
    \|_______|
	`,
	`
 ___   ___     
|\  \ |\  \    
\ \  \\_\  \   
 \ \______  \  
  \|_____|\  \ 
         \ \__\
          \|__|
	`,
	`
 ________      
|\   ____\     
\ \  \___|_    
 \ \_____  \   
  \|____|\  \  
    ____\_\  \ 
   |\_________\
   \|_________|
	`,
	`
 ________     
|\   ____\    
\ \  \___|    
 \ \  \____   
  \ \  ___  \ 
   \ \_______\
    \|_______|
	`,
	`
 ________  
|\_____  \ 
 \|___/  /|
     /  / /
    /  / / 
   /__/ /  
   |__|/   
  `,
	`
 ________     
|\   __  \    
\ \  \|\  \   
 \ \   __  \  
  \ \  \|\  \ 
   \ \_______\
    \|_______|
`,
	`
 ________     
|\  ___  \    
\ \____   \   
 \|____|\  \  
     __\_\  \ 
    |\_______\
    \|_______|
	`,
}

func init() {

	linesOfColon = strings.Split(colon, "\n")

	for i, number := range numbersMap {
		linesOfNumbers[i] = strings.Split(number, "\n")
	}
}

type TimeHumanize3D struct {
}

func NewTimeHumanize3D() *TimeHumanize3D {
	return &TimeHumanize3D{}
}

func (th *TimeHumanize3D) Convert(time int) string {
	var result string
	minutes := time / 60
	seconds := time % 60

	var symbolsArray = [][]string{
		linesOfNumbers[0],
		linesOfNumbers[minutes],
		linesOfColon,
		linesOfNumbers[0],
		linesOfNumbers[seconds],
	}

	for i := 0; i < maxLetterHeight; i++ {

		for _, lines := range symbolsArray {
			if i < len(lines) {

				result += lines[i]
			}
		}

		result += "\n"
	}

	fmt.Println(linesOfNumbers[minutes], linesOfColon)

	return result
}

func convertNumber(arr []string) {

}
