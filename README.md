# URL Shortener

A simple and lightweight URL shortener service written in Go (Golang). This project allows users to shorten long URLs into a shorter format and also provides redirection from the short URL to the original URL.

---

## Features

- Generate a shortened URL for any valid long URL.
- Automatically redirect to the original URL when accessing the shortened URL.
- In-memory storage for mapping short URLs to long URLs.
- Simple, lightweight, and fast implementation.

---

## How It Works

1. **Shorten URL**:
   - The `/shorten` endpoint accepts a JSON payload with the original URL.
   - It generates a random 6-character string as the short URL key.
   - Returns the shortened URL as a response.

2. **Redirect**:
   - Accessing the shortened URL redirects the user to the original URL stored in the system.

---

## Requirements

- **Go**: Version 1.16 or later.
- Any HTTP client (e.g., Postman, cURL, or a web browser).

---

## Setup and Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. **Run the application**:
   ```bash
   go run main.go
   ```

3. **The server will start at http://localhost:8080.**
