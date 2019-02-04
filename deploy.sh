# !/usr/bin/sh

go build
scp ./community-news-bot fnplus_aws:/home/ubuntu/news-bot/.
scp ./firebase-creds.json fnplus_aws:/home/ubuntu/news-bot/.
scp ./config.yaml fnplus_aws:/home/ubuntu/news-bot/.