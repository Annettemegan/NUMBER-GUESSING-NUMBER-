 Golang Number Guessing Game

 Overview
A simple web-based guessing game built in Go.

 Requirements
- Go 1.20+
- Browser

 Run
go run main.go

Open:
http://localhost:8080

Installation & Setup Instructions


Follow these steps to set up and run the Golang web application.


Step 1: Install Go
Download Go from the official website:
https://go.dev/dl/
Install it and verify installation using:
go version


Step 2: Create a Project Folder
Open your terminal and run:
mkdir guessing-game
cd guessing-game


Step 3: Create the Main File
Create a file named:
main.go

You can create it using:
notepad main.go

Or using VS Code:
code main.go


Step 4: Paste the Application Code
Paste the full Go web server code inside main.go.
Save the file.

Step 5: Run the Application
In the project folder, run:
go run main.go

You should see:
Server running at http://localhost:8080


Step 6: Open in Browser
Open your browser and go to:
http://localhost:8080

The number guessing game interface should load.
How to Use the Program

The homepage will display a number input field.

Enter any number between 1 and 100.

Click Submit Guess.

The system will respond with one of the following:

Too Low

Too High

Correct

If the guess is correct:

A new random number is automatically generated.

Click Play Again to restart.

The game continues until the user stops the server.







