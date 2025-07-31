# TMDB-CLI-Tool

# 🎬 TMDB CLI App

A command-line tool built with Go to fetch movie data from [The Movie Database (TMDB)](https://www.themoviedb.org/). You can list **popular**, **top-rated**, or **upcoming** movies straight from your terminal.

---

## ✨ Features

- Fetches data from TMDB v3 API
- Supports listing:
  - ✅ Popular Movies
  - ✅ Top Rated Movies
  - ✅ Upcoming Movies
- Secure authentication using Bearer Token via `.env` file
- Clean CLI experience

---

## 📦 Requirements

- [Go](https://golang.org/) 1.18 or higher
- A [TMDB API Bearer Token](https://developer.themoviedb.org/docs/getting-started)

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/tmdb-cli-app.git
cd tmdb-cli-app
```

### 2. Create .env File
Create a .env file in the root directory:

```bash
TMDB_API_TOKEN=your_tmdb_bearer_token_here
```

Make sure the token does not contain quotes or extra whitespace. The program will automatically TrimSpace() and remove any accidental "

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the App
```bash
go run main.go --type [playing | popular | top | upcoming]
```
### Example:
```bash
go run main.go --type popular
```

How Authentication Works
Uses Bearer Token provided by TMDB
Loaded securely from .env using github.com/joho/godotenv
```bash
import (
    "os"
    "strings"
)

token := strings.TrimSpace(os.Getenv("TMDB_API_TOKEN"))
token = strings.Trim(token, "\"")
```



