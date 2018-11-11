# web-crawler-in-go

[![Build Status](https://travis-ci.org/DaveLawes/web-crawler-in-go.svg?branch=master)](https://travis-ci.org/DaveLawes/web-crawler-in-go)

This is a simple web crawler in Go. Given a seed URL (domain name) the app searches through all href links and returns a list of all associated routes.

## How to Use

Ensure you have Go installed on your local machine then `git clone` this repo.

#### How to Run

To start the crawl navigate to `/cmd/webcrawler/` then run `go run main.go`.

#### How to Test

Each package has been created with its own test file. To test a package navigate to the relevant folder within `./pkg/` and run `go test`. To test all packages run `go test -v ./...` from the root of the project directory. 

HTTP requests are stubbed within the tests.

## Specification

* The crawler should be limited to one domain
* The crawler should not follow external links
* The crawler should print a simple site map


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

A user story to cover an identified edge case:

```
As a product owner
So I can visualise valid links only
I'd like the app to record valid relative urls only
```

## Domain Model

```                              
╔═══════════════════════╗     
║                         ║ specifies seed url
║        main.go          ║ starts crawl 
║                         ║ returns result
╚═══════════════════════╝  
            |
            |  
            | 
            |
            |      ╔════════════╗
            |      ║             ║  Creates go routine for each url found:   
            |------║   crawler   ║     - gets links and adds links to map of urls found
            |      ║             ║  Returns map of all urls   
            |      ╚════════════╝  
            |             |    
            |             |   
            |             |             ╔════════════╗  
            |             |             ║             ║  
            |             |-------------║   getBody   ║  Returns body of http response given a specific url 
            |             |             ║             ║  
            |             |             ╚════════════╝  
            |             |  
            |             |
            |             |             ╔════════════╗  
            |             |             ║             ║  
            |             |-------------║hrefExtractor║  Returns an array of all href links found on an html body 
            |                           ║             ║  
            |                           ╚════════════╝
            |          
            |
            |       ╔════════════╗
            |       ║             ║  
            |-------║ urlPrinter  ║  Pretty prints a map of urls found from the crawl
                    ║             ║
                    ╚════════════╝

```

`main.go` imports the required packages and starts the crawl. When all the links have been collected (returned in a map) they are transformed into a string and written to a .txt file.

The core package that `main.go` imports is the `crawler` package. The crawler package uses two channels to manage concurrency: 
1. urlQueue: whenever a valid URL is found it is added to this channel. As soon as a url is published to this channel it invokes a go routine that gets the links from the html body of that url.
2. urlCrawled: when all links have been retrieved from a html body a boolean true is published to this channel. This is used to increment an integer. This integer value is compared to the length of the map containing the urls in order to determine when the crawl is complete.

The go routine that extracts the links from the html body works as follow:
- `getBody` is a package imported by `crawler`. The `getBody` package makes an HTTP GET request to the provided address and returns the html response body. 
- When the response body is returned it is scanned for valid href links, this is carried out by the package `hrefExtractor`. `hrefExtractor` tokenizes the html and returns all valid href links.
- After all valid links are returned they are added to the map of URLs. This is a simple structure that contains keys for the URLs and values as an array of strings for the links found on that page. Because of the concurrent behaviour I had to use a lock to prevent multiple read/write attempts on the same resource (my URL map). The URLs added to the map are also added to the urlQueue channel - starting another go routine for each (initiating retrieval of their associated links).

## Design Process

- Read the requirements and ensure I have understood them
- Collect html test data
- Create user stories for core requirements and edge cases
- Create a domain model (having thought about what classes and architecture I want)
- Setup my project (github repo, TravisCI)
- Follow a red-green-refactor flow (starting with a feature test for my first user story)
- If edge cases are identified when working: add these as features/user stories to be completed after all core user stories have been satisfied
- When design is finished: update documentation

## Areas of Improvement

1. The structure I've used to record URLs and their associated links could be improved if I'd used a database. A database could contain more data and be manipulated (potentially) slightly easier through an ORM.
2. I could potentially utilise the net/http package to more simply manage absolute and relative URLs. Currently I use string manipulation to get the absolute URL, but maybe there are existing methods within the net/http package that could have made this easier for me.
3. I've limited my crawled URLs to 100 before existing the app. This was partly due to the time taken to crawl: which implies that I could improve the way I use concurrency in my app. Maybe a better way to do this would be to add URLs to the urlQueue channel as soon as they are discovered, as opposed to waiting until the entire HTML has been scanned. 
