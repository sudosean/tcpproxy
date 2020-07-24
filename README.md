# TCP Proxy project
In this project I am really just working with stdin/stdout in go using IO and bufio. I create a cmd that writes back whatever comes in from stdin (writer). I also create an echo server (echo) and a tcp proxy client that just returns 
whater comed from the website provided

# Blueprint/Boilerplate Used For This Golang Project
You can find more information about this project/repository and how to use it in following blog posts:

- [Ultimate Setup for Your Next Golang Project](https://towardsdatascience.com/ultimate-setup-for-your-next-golang-project-1cc989ad2a96)
- [Setting up GitHub Package Registry with Docker and Golang](https://towardsdatascience.com/setting-up-github-package-registry-with-docker-and-golang-7a75a2533139)
- [Building RESTful APIs in Golang](https://towardsdatascience.com/building-restful-apis-in-golang-e3fe6e3f8f95)
- [Setting Up Swagger Docs for Golang API](https://towardsdatascience.com/setting-up-swagger-docs-for-golang-api-8d0442263641)

### Adding New Libraries/Dependencies
```bash
go mod vendor
```
