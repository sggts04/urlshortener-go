# URL Shortener in Go

This is a barebones URL Shortener implementation in Go using the Gin web framework and MySQL. Also features a basic frontend.

## Local Setup
**Clone the repo**

`git clone https://github.com/sggts04/urlshortener-go`

**Set up environment variables**

Create a `.env` file in the root of the project and follow this format

```
PORT=
MYSQL_USER=
MYSQL_PASS=
MYSQL_ADDR=
MYSQL_DB=
```

The variables should be self-explanatory.

**Set up Database**

Run the migration in `migration.sql`

**Run Locally**

`go run .` and then navigate to `localhost:PORT` in your browser for the frontend or just use the API.
