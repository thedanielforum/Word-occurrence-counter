#Programming practice
#####Small serverlet that accepts as input a body of text, such as that from a book, and returns the top ten most-used words along with how many times they occur in the text.

##Endpoints
POST: /
- HEADER "limit" (default 10) - limits the results returned
- BODY - text to be parsed

GET: /
- N/A - returns nil

##How to use
This app is made for google [app-engine](https://cloud.google.com/appengine/), so clone the repo and run
```cmd
go get
```
to download all the required packages, after this you can test locally
using the [App Engine SDK for Go](https://cloud.google.com/appengine/docs/go/download) from google or [deploy](https://cloud.google.com/appengine/docs/go/tools/uploadinganapp) to GCP.