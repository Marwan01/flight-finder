Plane Seating Company

September 9th 2018

code developed on: Mac

steps to compile: ( assuming you have go installed globally on your machine)
1. Download and extract the given zip file.
2. In a terminal window, navigate to the project's folder.
3. Run the following command: go run db_1.go 
In case step 3. does not work, try copying the project folder to ~/go/src and running it from there.

bugs: none that I noticed so far

approach:

I created a reader to get input from the user on the command line and
then split it and put every single space-seperated string within a slice
of strings. Then I used the indexes of the slice to seperate commands. 
I validated for each command within a bunch of if statement blocks.
I also added validation for when a wrong command is entered. Every series of
commands calls a specific function to either add or load or find (or quit) 
data from/to a file. I put everything within a loop that will continue 
executing and asking for commands until the program is exited through entering
the key "q". 
on the file, I denoted city lines by a starting string of "c: " and a similar
approach was used for airlines and flights. This is what I use to distinguish
cities from flights from airports when trying to load a specific one.
