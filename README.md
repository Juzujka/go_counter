
# go_counter

Go test

## Task requirements

Read from stdin Urls.  
Send Http request with every Url.  
Count "Go" in the body of every response.  
Print the number of "Go" in every response and total number of "Go".  


Process Urls simultaneously but no more then 5 routines.  

## Running the tests

### Go test

go to the ./go_counter_proc
```
cd ./go_counter_proc
```
and run 
```
go test
```

if golang.org and google.com does not change too much then you will see something like
```
test with url https://golang.org returns 20 err <nil>
test with url https://google.org returns 30 err <nil>
Count for https://x.x: 0
Count for https://golang.org: 20
Count for https://google.com: 6
PASS
ok  	go_counter/go_counter_proc	... s
```

### Manual test

```
echo -e 'https://golang.org\nhttps://golang.org' | go run ./main.go
```

returns

```
Count for https://golang.org: 20
Count for https://golang.org: 20
Total: 40

```

