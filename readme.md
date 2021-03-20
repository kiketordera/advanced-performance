# Template project for Golang using gin

In this project, you can find an easy, ready, and documented implementation for a simple landing page written in Golang and plain SCSS & HTML with best practices.

The project is equipped with all functionality you need to create your own project for rendering a simple website from here. This project uses Clean Architecture

<br>

# Clean architecture

The Clean Architecture was proposed by Robert C. Martin in his book, and two of the most recommended book to read about this topic are (both written by himself):<br>

> Clean Architecture: A Craftsman's Guide to Software Structure and Design

> Clean Code: A Handbook of Agile Software Craftsmanship

<br>
Clean Architecture push us to separate stable business rules (higher-level abstractions, domain layer) from volatile technical details (lower-level details), defining clear boundaries. The main building block is the Dependency Rule: source code dependencies must point only inward, toward higher-level policies. <br><br>
  <img  src="media/documentation/inward-dependency.png" alt="Sublime's custom image"/><br><br>

For OOP, the concept insists on polymorphism (interfaces, or abstract classes) which brings us to dependency inversion. DI allows the source code dependency (the inheritance relationship) to points in the invert direction. <br><br>

The architecture pattern used in this project is Domain-driven design (DDD), based on clean architecture principles. The main idea is to isolate the domain layer, where all the business logic is happening. <br><br>

To make that isolation possible, we use interfaces in the domain layer to communicate with other layers. For example, we have repositories for the Database, the mail service, and the Authentification. <br> <br>

## SCSS rem pattern

In the `SCSS` code, we refer to `rem` units to measure the elements. That means that all these values will be pointing to the root element `font-size`, in our case the `html{}` SCSS class. <br><br>

So, since all margins, height, font-size, and measurements, in general, are in rem units, which means are pointing all to one element, we can implement our media queries in only that element, and then we will have our project being resizable in every measurement, whiteout needed to put media queries all around our code. <br><br>

The html class is:

```
html {
  font-size: 62.5%;
  @media (max-width: 1250px) {
    font-size: 55%;
  }
  @media (max-width: 650px) {
    font-size: 47%;
  }
  @media (max-width: 500px) {
    font-size: 40%;
  }
  @media (max-width: 428px) {
    font-size: 35%;
  }
  @media (max-width: 375px) {
    font-size: 32%;
  }
}
```

We only need to implement our media queries here, maintaining our code much more clean and maintainable. And the best of all is that once we implement the media queries in the HTML element, anything we code with the rem unit will be mobile friendly without needed to make anything.

The initial value of `62,5%` is to make calculations easy, since in that value `1rem=10px`, so if we want to make `width: 35px;`, its equivalent in `rem` is `width: 3.5rem;` <br>
You can find full explanation why the value is `62,5%` [here](https://www.sitepoint.com/understanding-and-using-rem-units-in-css/)<br><br>

### Shortcuts SCSS

- \# = id
- \- = siblins
- $ = index variable for when we use \* (multiply)
- { } = text inside the HTML
- ( ) = gouping. Inside the brakets will be instructions, and after we can put + for siblins or whatever we want to put<br><br>

## Database

This project is working with any kind of database. Since we use interfaces to isolate layers, the connection between the Database and the application is just an implementation of the interface defined.

All the logic restrictions are in the business logic, you will only need to implement the interface. The implementation available in this project is for a local NoSQL database called BoltHold, [here full documentation](https://github.com/timshannon/bolthold) <br><br>

## Login

We have implemented a custom login, making use of [jwt-go tokens](github.com/dgrijalva/jwt-go) and cookies encryption. The implementation is in the `repositories` path and we use an interface in the domain layer to separate layers, so we could change this logging for another logging by just implementing the same interface from the domain with another service.

To make effective the methods contained in the interface, we need to use the corresponding middleware, in the case of [jwt-go tokens](github.com/dgrijalva/jwt-go):

```
r.Use(h.Login.CheckToken())
```

## i18n

Internationalization, often abbreviated as i18n, is the process through which products can be prepared to be taken to other countries. In our case, just mean being able to change the language of the website and show the default language of our user's browser.

The package we use [gin-i18n](github.com/suisrc/gin-i18n) is very simple and effective. You will be using `.toml` files, stored in the `/media/text/` directory.

The allowed names of the files must match the i18n international codes:

```
af, ar, az, be, bg, bn, bs, ca, cs, cy, da, de, de-AT, de-CH, de-DE, el,
el-CY, en, en-AU, en-CA, en-GB en-IE, en-IN, en-NZ, en-US, en-ZA, en-CY,
en-TT, eo, es, es-419, es-AR, es-CL, es-CO, es-CR, es-EC, es-ES es-MX, es-NI,
 es-PA, es-PE, es-US, es-VE, et, eu, fa, fi, fr, fr-CA, fr-CH, fr-FR, gl, he,
 hi, hi-IN, hr hu, id, is, it, it-CH, ja, ka, km, kn, ko, lb, lo, lt, lv, mk,
 ml, mn, mr-IN, ms, nb, ne, nl, nn, oc, or pa, pl, pt, pt-BR, rm, ro, ru, sk,
sl, sq, sr, st, sw, ta, te, th, tl, tr, tt, ug, ur, uz, vi, wo, zh-CN,
zh-HK, zh-TW, zh-YUE

```

We need to implement the package as a middleware

```
r.Use(i18n.Serve(bundle))
```

<br>

## Security

We use a middleware called [gin-secure](https://github.com/gin-contrib/secure) to secure our app with strict security settings. Default parameters configure are:

- SSLRedirect
- IsDevelopment
- STSSeconds
- STSIncludeSubdomains
- FrameDeny
- ContentTypeNosniff
- BrowserXssFilter
- ContentSecurityPolicy
- SSLProxyHeaders

<br>

# Run the project

You have available different ways of making your code run

<br>

## With Golang

Move to the directory of the main.go file

```
cd app
```

Get all the dependencies

```
go get ./...
```

Run the main.go file

```
go run main.go
```

### Thourbleshooting the code

- go mod init
- go mod tidy

<br>

## With Docker

Go to the root of the folder, and create the image.

```
docker build -t nameImage .
```

Once the image has been created, run the container

```
docker rum name
```

## With Air

With Air, you can reload automatically your Golang code every time you save a file.

<br>

### For making Hotreloading for the first time

Run these commands in this order

```
export GOPATH=$HOME/go

go get -u github.com/cosmtrek/air

curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

alias air='$(go env GOPATH)/bin/air'

air
```

<br>

### For making Hotreloading **AFTER** the first time

```
export GOPATH=$HOME/go
alias air='$(go env GOPATH)/bin/air'
air
```

### Troubleshooting Air

If you have a permission denied error, run the following

```
chmod u+x air
```

If we want to use the hot reload with the make file, write

```
make watch
```

# Run the tests

The `main_test.go` file is in the `app/` location.

<br>

# Configure Firebase for this project

https://medium.com/wesionary-team/authenticate-rest-api-in-go-with-firebase-authentication-36cdf7c254c

<br>

# Documentation

To have consistency across projects, we rely on the OpenAPI Specification (formerly Swagger Specification). Is an API description format for REST APIs. An OpenAPI file allows you to describe your entire API, including:

- Available endpoints (/users) and operations on each endpoint (GET /users, POST /users)
- Operation parameters Input and output for each operation
- Authentication methods
- Contact information, license, terms of use, and other information.

To see the generated documentation, you can run the project and visit http://localhost:8080/swagger/index.html

To initiate the documentation for the first time:
From `app/` directory

```
swag init
```

A cool article about Swagger in Golang [here](https://martinheinz.dev/blog/9)
<br><br>

## Troubleshooting documentation

If you get ERROR: `swag: command not found` then run

```
export PATH=$(go env GOPATH)/bin:$PATH
```

<br>

# Tutorials

## GO templates

> https://blog.gopheracademy.com/advent-2017/using-go-templates/

> https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0

<br>

## JSON and validation

> http://brandonokert.com/articles/json-management-patterns-in-go/

> http://brandonokert.com/articles/json-management-patterns-in-go/#easy-validation

> https://tutorialedge.net/golang/secure-coding-in-go-input-validation/

<br>

## Security

> https://www.veracode.com/blog/secure-development/use-golang-these-mistakes-could-compromise-your-apps-security

<br>

## Apps for testing your security

> https://www.zaproxy.org/

<br>
