How to setup:
0. Prerequisites
    > Postgres installed
    > Internet connection
    > Go installed
    > NPM installed

1. Prepare the database
    > Create new database & schema
    > Use the sql file "packform.sql" to populate it

    
2. Prepare the back-end
    > Clone from the given repo
    > Open the main.go, go to function "connect", then adjust the DSN connection string according to your username, password, database of your DB
    > In terminal, put the current dir into inside this folder project level, then run command "go run main.go"

3. Prepare the front-end
    > Clone from the given repo
    > In terminal, put the current dir into inside the frontend folder project level, then run command "npm install"
    > open the index.html in browser