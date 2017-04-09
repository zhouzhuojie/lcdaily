# lcdaily
Leetcode Daily - Sends yourself a random leetcode email every day.

# Preview

# Get started
```bash
# Here we use some env variables to configure the lcdaily

# your gmail account
# userEmail=username@gmail.com
# userPassword=123456

# your server host information for setting up adhoc url queries
# serverHost=10.10.10.10 or serverHost=sub.mydomain.com
# serverPort=8080

docker run -p 8080:8080 \
  -e "userEmail=username@gmail.com" -e "userPassword=123456" \ 
  -e "serverHost=localhost" -e "serverPort=8080" \
  -d zhouzhuojie/lcdaily
```
