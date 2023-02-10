# MEDIUM-SCRAPER
## Web scraper using [chromedp](https://pkg.go.dev/github.com/chromedp/chromedp)
### Scraps [medium](https://medium.com/)
  - search articles
  - get article
## Telegram [Bot](https://core.telegram.org/bots)
### Bot
  - get user's search text and return a list of articles
  - let user choose a specific article from the list
  - return the full article
  - ability to cancel any running operation
### Bot's Database (redis)
  - add user
  - save user's last command
  - save user's last articles url
  - retrieve user

##  Running the application
### Define the environment variables
  - bot_token={your telegram bot api token}
  - redis_url={redis database url}
  - redis-password={redis database password}
### In the command line run
```
go run .
```