## Cron Expression Parser


### How to build
`go build main.go`


### How to run

`./main "*/15 0 1,15 * 1-5 /usr/bin/find"`


### Some context

The input should be passed inside double quotes ".

The program parses standard cron expression, so it should have 5 parameters: minute hour day(month) month day(week)





The output should be in the following format:

```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
