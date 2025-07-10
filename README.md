# Simple Webcrawler

Concurrency using crawler to count internal links of a domain

## Usage

```bash
go run . <url> <thread count> <max pages>
```

**url** - domain to be crawled  
**thread count** - maximum number of threads spawned for scanning  
**max pages** - sensibility limit, stops scan if unique page count reaches it
