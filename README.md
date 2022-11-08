# Ditto-Go-React

| Date | Status | Desc |
| --- | --- | --- |
| Oct 04, 2022 | ⚪️ | Tryna make a monorepo |
| Oct 04, 2022 | 🟢 | In progress |

## Memo

### Ubuntu port 80

``` cmd
sudo kill -9 `sudo lsof -t -i:80`
```

### Ubuntu install React

``` cmd
sudo apt install npm

npm --version
node --version
```

### Ubuntu update Go version

``` cmd
sudo apt-get remove golang-go
sudo apt-get remove --auto-remove golang-go


sudo rm -rvf /usr/local/go

wget https://golang.org/dl/go1.19.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.19.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### Clear MongoDB trash record

``` cmd
db.act.deleteOne({"_id": ObjectId('000000000000000000000000')})
```

### Backup MongoDB

``` cmd
# AWS EC2: Docker → Ubuntu
sudo docker exec 25b mongodump --db ditto --out /mongodump_20221106
sudo docker cp 25b:/mongodump_20221106 ~/mongodump_20221106
```

### Restore MongoDB

``` cmd
# Windows
docker cp mongodump_20221106 4e7:/mongodump_20221106
docker exec -i 4e7 /usr/bin/mongorestore --db ditto /mongodump_20221106/ditto

# AWS EC2: Ubuntu → Docker
sudo docker cp mongodump_20221106 25b:/mongodump_20221106
sudo docker exec -i 25b /usr/bin/mongorestore --db ditto /mongodump_20221106/ditto
```

### Docker deployment

``` cmd
docker build -t ditto-go-react .
docker run -p 80:8080 -d ditto-go-react
docker-compose up -d --build web
```

### Delete Redis data

``` redis-cli
flushall
```

## References

- [Deploying Go + React to Heroku using Docker](https://levelup.gitconnected.com/deploying-go-react-to-heroku-using-docker-9844bf075228)
- [mgo - MongoDB driver for Go](https://github.com/go-mgo/mgo)