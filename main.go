package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: <input.txt> <result.txt>")
	}

	words, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading file")
	}

	sentence := string(words)

	slices := strings.Split(sentence, "\n")
	
	  for i := range slices {
		  if slices[0] == "package   main"{
			  slices[i] = "package main"
		  } else if slices[2] == "import \"mt\""{
			slices[i]= " import \"fmt\""
		} else if slices[3] != " "{
			slices[i] = " "
		}else if slices[4] == "func ma(){"{
			slices[i]= "func main(){"
		}else if slices[5] == "fmt.Println(\"Hello)\""{
			slices[i] = "    fmt.Println(\"Hello World)\""
		}else if  slices[6] == ""{
			slices[i] = "}"
		} else if slices[7] != "\n"{
			slices[i] = "\n"
		}
	  }

}