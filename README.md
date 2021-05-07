# Youtube Fetcher

Use Google Youtube Data Api v3.

Use https://infinite-scroll.com/ for infinite scroll load videos.

> Infinite scrolling the right way:   
> https://medium.com/walmartglobaltech/infinite-scrolling-the-right-way-11b098a08815  
> https://jsfiddle.net/valkyris/43fmku20/693/  

# Quick Start

```
go run cmd/server/server.go
```

Look port on Mac:
```
lsof -i:8080
```

All your channelIds should list in `./internal/pkg/db/json/cids.txt`

# TODO

- [X] Start jobs with heartbeat
- [X] Infinite scroll loading complete on cid page.
- [X] Optimize view of index
- [ ] Search implement
- [ ] Scrollspy index

