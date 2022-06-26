# Tour Checker
This was a test project to try out deploying AWS Lambda with cron tasks. This project checks the Stanford website to see if there are any available tours and notifies me if so.

## Building

To deploy this to AWS Lambda, add the Discord Webhook URL environmental variable `WEBHOOK_URL` and build this program:

```
GOARCH=amd64 GOOS=linux go build main.go
```
Lastly, create a zip file of the executable to upload to AWS:

```
zip tour.zip main
```
