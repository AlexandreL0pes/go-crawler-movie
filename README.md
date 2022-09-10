# go-crawler-movie
a movie crawler built with go for study purposes.

The crawler performance can be improved. Right now, the crawler processes an average of **1319 movies por minute**.


### db
To init the database, run this command:
```
docker-compose up
```

### helpers
Create sqlite database
```
make create db
```

Run the crawler with:
```
make run
```

### TODO

- [ ] During the crawling process some movies aren't persisted in the database
- [ ] Improve the logs, add debug and better visualization
- [ ] create a concurrent func to write in database

