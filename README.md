# LibraryDatabase2.0

A simple program that takes an ISBN input from the user and requests information about the ISBN from GoogleAPI or OpenLibrary. If one fails then the information is gathered from the other API. There is also an option for the user to manually input a book. The program builds a database for use with the SQLite framework, saves the received information about the book to the database, and builds a table using jedib0t's go-pretty package. Besides SQLite and go-pretty, this program also contains a menu derived from the promptui package from manifoldco.

![sample](https://github.com/jeffrey33704/LibraryDatabase2.0/blob/main/sample.gif)

# Support for webcam scanning of barcode added
Now users may scan books using their webcams and barcodes.  However, the ease of use is dependent on the quality of one's webcam.  For example, my machine (Thinkpad t440p) has a low-quality webcam and so scanning is hit or miss; sometimes when scanning the barcode is picked up immediately, while other times it could take 20-30 seconds.  Furthermore, users must have OpenCV installed in order to utilize the GoCV package for webcam functionality.

Give it a try, and let me know what you think!

# Added the option to export the SQL database as JSON
The option to export the library database as a file containing the data in JSON format is now available.
