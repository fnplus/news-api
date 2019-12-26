# news-api
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffnplus%2Fnews-api.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffnplus%2Fnews-api?ref=badge_shield)

A bot/worker/cron/process to fetch &amp; curate all the relevant news content from the web based on the user's subscription.

## How it works
This service is responsible for fetching the news related to the keywords that the users have subscribed to. The list of all keywords that users have subscribed to are present in the database.

  - The keywords are first fetched 
  - Multiple `workers` are created to to fetch news concurrently
  - The keyqords are queued for the workers to pickup 
  - Once a worker picks up a keyword, it will fetch the news from newsapi.org's REST api
  - Once the news is fetched, they are pushed to the database and the titles are cached locally. Because, the next time when the news is fetched, the same news should not be pushed to the dB again and again

This procedure is run multiple times in a day. Keep in mind, only the "new" news is sent to the database.

## Installation

 1. First, [fork and clone](https://github.com/fnplus/news-api) the repo. This will create a new folder named `news-api`.
 2. Navigate to `/news-api` and run:
   ```
   go get -d
   go build
   ./news-api
   ```

## How to contribute
Read [this](contributing.md)


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffnplus%2Fnews-api.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffnplus%2Fnews-api?ref=badge_large)