# LibraryDatabase2.0

A simple program that takes an ISBN input from the user and requests information about the ISBN from GoogleAPI or OpenLibrary. If one fails then the information is gathered from the other API. There is also an option for the user to manually input a book. The program builds a database for use with the SQLite framework, saves the received information about the book to the database, and builds a table using jedib0t's go-pretty package. Besides SQLite and go-pretty, this program also contains a menu derived from the promptui package from manifoldco.

Give it a try, and let me know what you think!

![sample](https://github.com/jeffrey33704/LibraryDatabase2.0/blob/main/sample.gif)
