# Ditto-Go-React

| Date | Status | Desc |
| --- | --- | --- |
| Oct 04, 2022 | ⚪️ | Tryna make a monorepo |
| Oct 04, 2022 | 🟢 | In progress |

## Memo

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

### Backup MongoDB

``` cmd
# AWS EC2: Docker → Ubuntu
sudo docker exec 210 mongodump --db ditto --out /mongodump_20221028
sudo docker cp 210:/mongodump_20221028 ~/mongodump_20221028
```


## References

- [Deploying Go + React to Heroku using Docker](https://levelup.gitconnected.com/deploying-go-react-to-heroku-using-docker-9844bf075228)
