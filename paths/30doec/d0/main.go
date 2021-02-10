// Declaring the current source code as a package
package main    // 'main' is the executable package

// Importing external packages
import (
    "fmt"       // (1) Formatted I/O functions  
    "bufio"     // (2) Buffered I/O functions
    "os"        // (3) Universal interface to O.S. functionality
)

func main() {
    // (4) Creating a new bufio.NewScanner() object to read input from STDIN
    scanner := bufio.NewScanner(os.Stdin)
    
    // (5) Writing an output to STDOUT with fmt.Println()
    fmt.Println("Hello, World.")    // Println adds a line break \n at the end
    
    // Declaring a variable to hold our input from STDIN
    var is string                   // 'is' stands for inputString
    
    // Reading an input from STDIN with bufio.NewScanner.Scan()
    // It returns true or false to indicate if there's any input available
    // Since there's only 1 input, the for loop will run only once
    for scanner.Scan() {
        is = scanner.Text()         // (6) bufio.Scanner.Text reads Strings 
    }
    
    // Printing an output to STDOUT with fmt.Print()
    fmt.Print(is)                   // Print doesn't add a line break \n
}

// + Tips:
// - By convention, we use the first letters of the words to name variables
// - We use short declaration := for new variables with defined initial values
// - Variables declared with short declaration are block scoped
// - Block scopes are delimited by {}

// + References:
// (0) fmt package - https://golang.org/pkg/fmt/
// (1) bufio package - https://golang.org/pkg/bufio/
// (2) os package - https://golang.org/pkg/os/
// (3) bufio.Scanner - https://golang.org/pkg/bufio/#Scanner
// (4) fmt.Println - https://golang.org/pkg/fmt/#example_Println
// (5) bufio.Scanner.Text - https://golang.org/pkg/bufio/#Scanner.Text
