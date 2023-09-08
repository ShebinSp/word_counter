WORD COUNTER
------------

This is a simple word counter written in go.

For this tool, you import the bufio package to read text, the fmt package to print formatted output,
the io package which provides the io.Reader interface, and the os package so you can use operating-
system resources.

The word counter have two functions: main and count.
*) The main function is the starting point of the program. All Go programs that will be
compiled into executable files requre this function.
*) The count function , which will perform the actual counting of the words. This function receives
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

