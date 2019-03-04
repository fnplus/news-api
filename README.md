# community-news-bot
A bot/worker/cron/process to fetch &amp; curate all the relevant news content from the web based on the user's subscription.

## How it works
This service is responsible for fetching the news related to the keywords that the users have subscribed to. The list of all keywords that users have subscribed to are present in the database.

- The keywords are first fetched 
- Multiple `workers` are created to to fetch news concurrently
- The keyqords are queued for the workers to pickup 
- Once a worker picks up a keyword, it will fetch the news from newsapi.org's REST api
- Once the news is fetched, they are pushed to the database and the titles are cached locally. Because, the next time when the news is fetched, the same news should not be pushed to the dB again and again

This procedure is run multiple times in a day. Keep in mind, only the "new" news is sent to the database.

## How to contribute
Read [this](contributing.md)