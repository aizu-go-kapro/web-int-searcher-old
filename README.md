# web-int-searcher
web-int.u-aizu.ac.jp向け全検索エンジン

# start mongodb

```
sudo mongod --dbpath ./mongo/data  --logpath ./mongo/log/mongodb.log
```

## init
```
mkdir -p mongo/data/ mongo/log
touch mongo/log/mongodb.log
```

## find
```
db.data.find()
```

## count

```
db.data.find().count()
```


# mecab


## install 
```
export CGO_LDFLAGS="-L/usr/local/bin/mecab -lmecab -lstdc++"
export CGO_CFLAGS="-I/usr/local/bin/mecab/include"
go get github.com/shogo82148/go-mecab
```

## 

```
export CGO_LDFLAGS="`mecab-config --libs`"
export CGO_FLAGS="`mecab-config --inc-dir`"
go get github.com/shogo82148/go-mecab
```