# web-crawler-in-go

This is a simple web crawler in Go. Given a seed URL (domain name) the app searches through all href links and returns a list of all associated routes.

## Specification

* The crawler should be limited to one domain
* The crawler should not follow external links
* The crawler should print a simple site amp


## User Stories

Translating the specification above into some simple user stories:

```
As a product owner
So I can keep track of our website
I'd like an app to crawl through all links included within our domain

As a product owner
So I can keep track of only my links
I'd like the web crawler limited to the domain I specify

As a product owner
So I can visualise the links we maintain
I'd like the app to print all the links in our domain

```

## Domain Model


