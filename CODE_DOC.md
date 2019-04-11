# Payment API - Code Documentation

This API was developed using tools below  
- Golang as programming language  
- MySQL as database  

The database was build in order to carry the information about client, buyer and payments.  
The system was build to serve as an REST API, with methods that supports include, consult and update infos.  
To organize codes, I've distributed functions especially on routes, handlers, functios and database. Focusing on keep code clean and clear as possible, trying to turn every future feature easy to implement.  
I've created an unit test to cases with static returns.  

## Preparing your machine
 - You should have Golang and MySQL installed on your machine.  

## Create and preparing database
- Create database, tables and columns using file "CREATE.SQL" (on folder ...\challenge\scripts).
- Run the initial inserts using file "INSERT.SQL" (on folder ...\challenge\scripts).

## How do I run the program?
1. Open CMD and open the folder where are the project (...\challenge\src).  
2. Execute the command "go run .", this will run all .Go files on folder.
3. Using an API test tool, you'll could now call methods that are listed on API doc. On folder of tests there are some created to Postman use (...\challenge\testes).

## How do I run the tests?
1. Open CMD and open the folder where is the project (...\challenge\src).  
2. Execute the command "go test -v", this will run the tests listing all with respective state.

