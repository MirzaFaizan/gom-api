# GOM-API

[![Build Status](https://travis-ci.org/MirzaFaizan/gom-api.svg?branch=master&style=for-the-badge)](https://travis-ci.org/MirzaFaizan/gom-api) [![iris](https://img.shields.io/badge/iris-powered-2196f3.svg?style=for-the-badge)](https://github.com/kataras/iris)

This is a base go lang API with CURD operations with MongoDB

  - Layered folder structure
  - Different Packages

### Tech

This API uses a number of open source projects to work properly:

* [Golang] - evented operations for the backend
* [Iris-go] - fast Golang network app framework [iris-go.com]
* [mongo-go] - mongoDB official drivers for Golang

And of course Dillinger itself is open source with a [public repository][dill]
 on GitHub.

Article link: <https://medium.com/@mirzafaizanejaz/go-lang-mongodb-iris-api-part-1-85024eb2d94d>.

### Installation

This API requires [Golang](https://golang.org/) v1.11+ to run.

Install the dependencies and devDependencies and start the server.

```sh
$ cd gom-api
$ go get -u
$ go app.go
```

### Development

Want to contribute? Great!

This API uses Rizla for fast developing.
Make a change in your file and instantaneously see your updates!

Open your favorite Terminal and run these commands.

```sh
cd gom-api
go build app.go
```
This will create the executeable GO file and pull in the necessary dependencies.

I am using Rizla Hot Reloading package (which hot reloads the server on code changes) map the port to whatever you wish on your host or set a PORT on your .env file.

```sh
rizla app.go
```

Verify the deployment by navigating to your server address in your preferred browser.

```sh
127.0.0.1:8080
```


### Todos

 - Write Tests
 - Add Bcrypt
 - Add JWT
 - Add Oauth

### Social

I am available on the following platforms if you have a question just reach out to me.

| Network | link |
| ------ | ------ |
| Twitter | [MirzaFaizanEjaz][PlTwitter] |
| Instagram | [le_programmer][PlIG] |
| Medium | [@mirzafaizanejaz][PlMe] |
| Stack Overflow | [faizan-ejaz][PlSO] |


License
----

MIT

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)


   [PlTwitter]: <https://twitter.com/mirzafaizanejaz>
   [PlIG]: <https://instagram.com/le_programmer>
   [PlMe]: <https://medium.com/@mirzafaizanejaz>
   [PlSO]: <https://stackoverflow.com/users/9268483/faizan-ejaz>
