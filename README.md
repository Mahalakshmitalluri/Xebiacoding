# Xebiacoding 
The Coding is done using two different methods with two go files.
# main.go file  in the solution folder, it is used for the Capitalization  of every third alphanumeric character in a string.
We have printed here the "aspiration.com" with the capitalization for every third element. If we want to check some other string or give some other input, we can replace with the aspiration.com and go for it.
To run this code, do the "go run " command with the file name along with it. Further ,i have used if else condition, whenever if the index is at the 3rd position in the aspiration string ,it will give us an upper case ,else we are passing it through the lower case. Then we are adding this temp string to "result" variable name or else we are just simply adding string to "result". In main function, we are printing and testing the string "Aspiration.com".
# main.go file in the mapper folder
Capitalization of every third element is done here with a different function signature using interface. We have done the string conversion here.
Created an interface, which represents two methods as Transform Rune(pos int) and GetValueAsRuneSlice() []rune by using transform rune, converting 
third character string into uppercase, whereas with getvalueasruneslice,we are storing the converted string value into rune. After that, we are printing the given string "Aspiration.com".

