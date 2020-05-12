
# go_counter

Go test

## Task requirements

Read from stdin Urls.  
Send Http request with every Url.  
Count "Go" in the body of every response.  
Print the number of "Go" in every response and total number of "Go".  


Process Urls simultaneously but no more then 5 routines.  

## Running the tests

```
echo -e 'https://golang.org\nhttps://golang.org' | go run ./main.go
```

returns

```
Count for https://golang.org: 20
Count for https://golang.org: 20
Total: 40

```
