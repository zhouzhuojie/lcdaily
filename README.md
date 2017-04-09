# lcdaily
Leetcode Daily - Send random leetcode email every day

# Get started
```bash
$ # pick your server port to serve, for example 8080
$ # then -p 8080:8080 and -e "serverPort=8080"
$ docker run -p 8080:8080 -e "userEmail=username@gmail.com" -e "userPassword=123456" -e "serverHost=localhost" -e "serverPort=8080" -d zhouzhuojie/lcdaily
```
