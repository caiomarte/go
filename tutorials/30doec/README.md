![series cover](../../images/30doec/cover-30doec.png)
# 30 Days of (Explained) Code

### Hello, refactored chimps!

This is a series of `[30]Post` destined to new Gophers and sympathizers.

In this series, we are going to explore fundamental Go concepts while solving basic algorithm problems. By the end, you should have a general view of Go and be able to understand and write some Go code on your own.

If you’re looking for a problem-based approach for beginners, this is the place for you!

---

### And how is it going to work?

We are going to solve 30 problems ranging from easy to medium level, one problem each post. These problems belong to [HackerRank’s Tutorial “30 Days of Code”](https://www.hackerrank.com/domains/tutorials/30-days-of-code), and they provide the basis we need to start exploring Go.

Each post will follow the same structure:

#### What is the problem?
A short explanation of what’s required to complete the challenge. It will describe what the problem is and what input and output we should expect.

#### What do we need to know about Go?
What are the Go concepts we need to know in order to solve the problem? This is our opportunity to discuss and learn the Go basics.

#### Step-by-step coding explained
Now that we have the theory, it’s time to code. As we solve the problem, I’ll explain each line or block of code based on the concepts we saw previously.

In this section, I will prioritize flow of thinking over flow of compilation. That means code may appear in different order in the explanation versus the algorithm. But don’t worry: I’ll place a number before each line indicating its position in the final structure.

#### Key Takeaways
A summary of what’s been seen for you to save for later.

#### Full commented solution
The full solution with all necessary comments for you to review.

---

### Table of contents

#### [Day 0: Hello, World](./d0/)

Outputting `Hello, World.` and an input retrieved from `Stdin` to the console.

<details>
 <summary>What to expect?</summary>
 
 ###### Theory:
 * Go packages and code structure
 * The `fmt`, `os`, and `bufio` packages from Go's standard library
 * 4 ways to declare variables in Go
  
 ###### Practice:
 * Reading input from `Stdin` using `bufio.NewScanner(os.Stdin).Scan()`
 * Saving input from `Stdin` using `bufio.NewScanner(os.Stdin).Text()`
 * Printing outpum to `Stdout` using `fmt.Print`, `fmt.Printf()`, and `fmt.Println()`
 
 ###### Tips:
 * Variable namimg convetion in Go
 
 ###### Resources:
 * [Key takeaways](./d0#key-takeaways)
 * [Full commented solution](./d0/main.go)
 * [Knowledge compination `.pdf`](../../resources/30doec/d0-c.pdf)

</details>

#### [Day 1: Data Types](./d1/)

Reading multiple inputs, executing basic operations, and outputing the results.

<details>
 <summary>What to expect?</summary>
 
 ###### Theory:
 * Numeric and text data types supported by Go
 * Go's `+` to add numbers and concatenate text
 * the only loop available in Go, and its different forms
 * all conditional structures, and how Go differs from other languages
 * how to convert a text into a number using Go's `strconv` package
  
 ###### Practice:
 * Reading multiple inputs using a `for` loop
 * Evaluating `switch case` conditions
 * Converting `string` to `uint` with `strconv.ParseUint()`
 * Converting `string` to `float` with `strconv.ParseFloat()`

 ###### Tips:
 * Using blank identifiers to handle unused values
 * Formatting `fmt.Printf()` output with verbs and `\n`
 * Outputting limited decimal places using `%f`
 
 ###### Resources:
 * [Key takeaways](./d1#key-takeaways)
 * [Full commented solution](./d1/main.go)
 * [Knowledge compination `.pdf`](../../resources/30doec/d1-c.pdf)

</details>
  
---

### What's next?

[Day 0: Hello, World](./d0/)

Outputting `Hello, World.` and an input retrieved from `Stdin` to the console.

<details>
 <summary>What to expect?</summary>
 
 ###### Theory:
 * Go packages and code structure
 * The `fmt`, `os`, and `bufio` packages from Go's standard library
 * 4 ways to declare variables in Go
  
 ###### Practice:
 * Reading input from `Stdin` using `bufio.NewScanner(os.Stdin).Scan()`
 * Saving input from `Stdin` using `bufio.NewScanner(os.Stdin).Text()`
 * Printing outpum to `Stdout` using `fmt.Print`, `fmt.Printf()`, and `fmt.Println()`
 
 ###### Tips:
 * Variable namimg convetion in Go
 
 ###### Resources:
 * [Key takeaways](./d0#key-takeaways)
 * [Full commented solution](./d0/main.go)
 * [Knowledge compination `.pdf`](../../resources/30doec/d0-c.pdf)

</details>
