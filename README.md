# URL Shortener in Go

This is a barebones URL Shortener implementation in Go using the Gin web framework and MySQL. Also features a basic frontend.

## Local Setup
**Clone the repo**

`git clone https://github.com/sggts04/urlshortener-go`

**Set up environment variables**

Copy over the `.env.example` to `.env` and fill out the variables

`cp .env.example .env`

Your `.env` would look *something like this*
```
PORT=8000
MYSQL_USER=root
MYSQL_PASS=password
MYSQL_ADDR=localhost:3306
MYSQL_DB=urlshortener
```

The variables should be self-explanatory.

**Set up Database**

Run the migration in `migration.sql`

**Run Locally**

`go run .` and then navigate to `localhost:PORT` in your browser for the frontend or just use the API.
