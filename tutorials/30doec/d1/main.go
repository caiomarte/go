// Declaring the current source code as a package
package main	// 'main' is the executable package

// Importing external packages
import (
  "fmt"			// Formatted I/O functions
  "os"			// Universal interface to O.S. functionality
  "bufio"		// Buffered I/O functions
  "strconv"		// (0) String conversion functions
)

func main() {
  // Pre-declared variables
  var i uint64 = 4
  var d float64 = 4.0
  var s string = "HackerRank "
  // Pre-declared bufio.NewScanner() object to read input from STDIN
  scanner := bufio.NewScanner(os.Stdin)
  
  // Declaring variables to store input from STDIN
  var iV uint64		// (1) 'iV' stands for integer value
  var fV float64	// (2) 'fv' stands for float value
  var sV string		// (3) 'sV' stands for string value
  
  // (4) Iterating the necessary times to read the expected input
  // 3 inputs, therefore 3 loop iterations
  // Using default 'for' with initialization, condition, and post action
  for i := 0; i < 3; i++ {
    // Verifying if there's any input available
    scanner.Scan()
    // (5) Matching input order with input data type conversion
    switch i {
      case 0:	// (6) Converting 1st input from string to 10-base int64 
      iV, _ = strconv.ParseUint(scanner.Text(), 10, 64)
      case 1:	// (7) Converting 2nd input from string to float64
      fV, _ = strconv.ParseFloat(scanner.Text(), 64)
    case 2:		// Saving the third input as a string
      sV = scanner.Text()
    }
  }
  
  // Printing the sum of integers, floats, and strings
  // '%d' represents the integer 'i+iV'
  // '%.1f' represents the float 'd+fV' with a single decimal place
  // '%s' represents the concatenated string 's+sV'
  // '\n' breaks a line between outputs
  fmt.Printf("%d\n%.1f\n%s"), i+iV, d+fV, s+sV)
}

// + Tips:
// - You can only perform operations between operands of the same type
// - E.g. if 'i' is 'uint64', our variable 'iV' must be 'uint64'
// - Go's switch case doesn't need break statements
// - Go doesn’t compile code with unused variab
