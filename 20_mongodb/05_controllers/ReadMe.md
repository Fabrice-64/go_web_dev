# Using Curl
Start your server

Enter this at the terminal

# Strings to implement in the terminal
curl http://localhost:8080/user/1
curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user
curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
IMPORTANT: Make sure you update your import statements to import packages from the correct location!