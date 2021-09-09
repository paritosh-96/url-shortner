# url-shortner

This is a go service that starts on 6080 port and opens up 3 apis
- /healthCheck -> to check the status of the service if up or not
- /api/shortenUrl -> takes a url and a key to generate short url and will be stored corresponding to that key
- /api/actualUrl -> takes the short url with the key and gives back short url.

The urls are stored for a key. Each key will have its mapping of URLs.


The capacity of storing url mapping per key is 10 for now. It can be configured (will be updated soon)
