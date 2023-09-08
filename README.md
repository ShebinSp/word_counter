WORD COUNTER v 1.0
------------

This is a simple word counter written in go.
-----------  -----------   ------- ---------
* Ensure go is installed on your system. Then-
* Create a new folder named word_counter and open vs code in that folder(or any other text editor/ even bash is ok).
TIP: in linux open the folder word_counter and press 'f4' key to open terminal in current folder, then-
type 'code .' to open vs code as the current folder as working directory.

* Creating a go.mod
Type go mod init to create a module for your project.

* INTRODUCTION
For this tool, you import the bufio package to read text, the fmt package to print formatted output,
the io package which provides the io.Reader interface, and the os package so you can use operating-
system resources.

The word counter have two functions: main and count.
* The main function is the starting point of the program. All Go programs that will be
compiled into executable files requre this function.
* The count function , which will perform the actual counting of the words. This function receives
a single input argument: an io.Reader interface(for now think of an io.Reader as any Go type from which 
you can read data). In this case, the function will receive the contents of the STDIN to process.
-------------------------------------------------------------------------------------------------

Operations
----------

main function:-
The main function will call the count function and print out that function's return value using fmt.Println function.

count function:-
The count function uses the NewScanner function from the bufio package to create a new scanner. A scanner is a
convenient way of reading data delimited by spaces or new lines. By default, a scanner reads lines of data,
so we instruct the scanner to read words instead by setting the split function of the scanner to bufio.ScanWords.
We then define a variable , wCtr, to hold the word count and increment it by looping through each token using the
scanner.Scan function and adding 1 to the counter each time. We then return the word count
----------------------------------------------------------------------------------------------------
After compliting the word counter program we have to check the program is working as we expected.
So to write a test case, create a new file 'main_test.go' in the same location where the 'main.go' is.
----------------------------------------------------------------------------------------------------

TEST
----
main_test.go
------------
This test file containse a single test called TestConuntWords. In this test, we create a new buffer of bytes from
a string containing four words and pass hte buffer into count function. If this function returns anything other than 4,
the test doesn't pass and we raise an error that shows what we expected and what we actually got insted.

To execute the test, use the go test tool like this:

* To ensure you are in the directory where the main.go and main_test.go files are, run this command -
$ ls
go.mod   main.go   main_test.go  <- if you see files like this then run the next command.

* To test the file run this command
$ go test -v
 === RUN   TestCountWords
 --- PASS: TestCountWords (0.00s)
PASS
ok  --here you will see your file directory--
--------------------------------------------------------------------------------------------------------

BUILD
-----
The test passes, so you can compile the program with fo build.
$ go build

This creates the word_counter executable in the current directory:
$ ls
go.mod  main.go  main_test.go  word_counter

Test the program out by passing it an input string:
$ echo "My first command line tool with Go" | ./word_counter
7

Check the words in main.go file using word_counter
**********************************************************************************************************
*                                     cat main.go | ./word_counter                                       
**********************************************************************************************************
THIS PROGRAM WORKS AS EXPECTED. LET'S ADD THE ABILITY TO COUNT LINES TO THIS TOOL

* Adding Command-Line Flags
----------------------------
Here we created a command line flag "-l".
if the user input "-l" when running the program, then number of lines will appear insted of words.
**********************************************************************************************************
cat main.go | ./word_counter -l
**********************************************************************************************************

* check the project on git hub with version.
  command line and count lines added in version 2.0.0

------------------------------------------------------------------------------------------------------------
* New feature added countBytes to count the bytes in version 2.1.0
add command line flag -b to count the number of bytes
**********************************************************************************************************
cat main.go | ./word_counter -b
**********************************************************************************************************

