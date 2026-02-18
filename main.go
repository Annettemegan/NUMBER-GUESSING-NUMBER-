
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var number int

func main() {
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(100) + 1

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/guess", guessHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Number Guessing Game</title>
		<style>
			body {
				margin: 0;
				height: 100vh;
				display: flex;
				justify-content: center;
				align-items: center;
				background: linear-gradient(135deg, #667eea, #764ba2);
				font-family: Arial, sans-serif;
			}
			.card {
				background: white;
				padding: 40px;
				border-radius: 15px;
				box-shadow: 0 10px 25px rgba(0,0,0,0.2);
				text-align: center;
				width: 320px;
			}
			h1 {
				margin-bottom: 10px;
			}
			p {
				color: #555;
			}
			input {
				width: 100%;
				padding: 10px;
				margin: 15px 0;
				border-radius: 8px;
				border: 1px solid #ccc;
				font-size: 16px;
			}
			button {
				width: 100%;
				padding: 10px;
				border: none;
				border-radius: 8px;
				background: #667eea;
				color: white;
				font-size: 16px;
				cursor: pointer;
				transition: 0.3s;
			}
			button:hover {
				background: #5a67d8;
			}
		</style>
	</head>
	<body>
		<div class="card">
			<h1>🎯 Guess the Number</h1>
			<p>Choose a number between 1 and 100</p>
			<form action="/guess" method="POST">
				<input type="number" name="guess" required>
				<button type="submit">Submit Guess</button>
			</form>
		</div>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	guessStr := r.FormValue("guess")
	guess, _ := strconv.Atoi(guessStr)

	message := ""

	if guess < number {
		message = "Too Low 🔻"
	} else if guess > number {
		message = "Too High 🔺"
	} else {
		message = "Correct 🎉 Refresh to play again!"
		number = rand.Intn(100) + 1
	}

	response := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Result</title>
		<style>
			body {
				margin: 0;
				height: 100vh;
				display: flex;
				justify-content: center;
				align-items: center;
				background: linear-gradient(135deg, #ff9a9e, #fad0c4);
				font-family: Arial, sans-serif;
			}
			.card {
				background: white;
				padding: 40px;
				border-radius: 15px;
				box-shadow: 0 10px 25px rgba(0,0,0,0.2);
				text-align: center;
				width: 320px;
			}
			a {
				display: inline-block;
				margin-top: 20px;
				text-decoration: none;
				color: white;
				background: #ff6b6b;
				padding: 10px 20px;
				border-radius: 8px;
				transition: 0.3s;
			}
			a:hover {
				background: #ee5253;
			}
		</style>
	</head>
	<body>
		<div class="card">
			<h2>` + message + `</h2>
			<a href="/">Play Again</a>
		</div>
	</body>
	</html>
	`

	fmt.Fprint(w, response)
}
