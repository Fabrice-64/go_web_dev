# Basic Use of Mongo DB
## Starting
```
mongod      // starts server
mongo       // starts CLI
```
## Check existing DB
```
show dbs        // it will show all databases available with sizes
use <dbname>    // will use the <dbname> database, will create if none exists
```
## Find Records
```
db.collectionName.find()
```

## Delete All Records
```
db.collectionName.deleteMany({})
```
