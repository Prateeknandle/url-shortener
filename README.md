# URL-SHORTENER

A basic URL shortener application which provides the shorten url for the actual url. It stores the shorteb url in `data.txt` file, corresponding to its actual/original url. 

# Setting - up

Follow the following steps: 

1. Clone the project - `git clone https://github.com/Prateeknandle/url-shortener.git`
2. Run `go mod tidy`
3. If you want to change the port, you can change it in the `main.go` file.
4. Run `go run main.go`
5. You can use `Postman` for requests.
6. Use `POST` request with URL - `http://localhost:3000/short`.
7. Provide the original url which we want to shorten under the body/raw section (add headers `Key : Content-Type` and `value : application/json`),
```
{
    "Long_url":"https://gobyexample.com"
}
```
8. After we send this request we'll get the shortened url in response.
9. Use the short url, and run it on your respective browser, it'll redirect to the original url.

Follow the following steps, if using docker image:
 
1. Link to Docker image - `https://hub.docker.com/r/prateeknandle/docker-gs-ping`
2. Use `docker pull prateeknandle/docker-gs-ping` to pull the image of the application.
3. Run `docker run -p 3000:3000 prateeknandle/docker-gs-ping` to start the server. PORT - 3000
4. Go to 5th point written above and follow the steps below it.

# Functionalities

Our server doesn't use any database to store the url, but rather we use `data.txt` file to store the values. If the file is not found in the working directory, server will automatically create the file and store the url. 

`apis` package handles the router and routes. Currently in keeping mind of the problem statement, only two routes are defined. One for shortening the url and other one for using the shortened url.

`handlers` packege contains the implementaion for shortening the url, handling data storage in the dynamic file, redirecting the user to the actual url when they use the shortened url and proving the same shortened url if they again ask short url for which the short url is already generated and stored in the dynamic file.

To shorten the url, we've used package `github.com/teris-io/shortid`, it enables the generation of short, fully unique, non-sequential and by default URL friendly Ids at a rate of hundreds of thousand per second. It guarantees uniqueness during the time period until 2050!

# Other details

We've used GO version - `go1.19 linux/amd64`