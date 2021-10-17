In order to check: 
1. Compile the program 
```go build -o solution```

2. Launch the program from one terminal window:
```./solution```

3. From another terminal window:
a. To Create a User
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```
b. To Get a User
```curl http://localhost:8080/user/<user uuid>```

c.To Delete this User:
```
curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/<user uuid>
```



