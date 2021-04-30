# Youtube Fetcher

Use Google Youtube Data Api v3.

Use https://infinite-scroll.com/ for infinite scroll load videos.

> Infinite scrolling the right way: 
> https://medium.com/walmartglobaltech/infinite-scrolling-the-right-way-11b098a08815  
> https://jsfiddle.net/valkyris/43fmku20/693/  

# Quick Start

```
go run main.go
```

Look port on Mac:
```
lsof -i:8080
```

# TODO

- [X] Start jobs with heartbeat
- [ ] Load by Ajax on channels page.
  - [ ] nextCidHandler() for infinite request and return data
  - [ ] nextCidDerive() for render data from infinite request
  - [ ] cid.html need a link for next at bottom of page. it will be changed by js every scroll
- [X] Optimize view of index
- [ ] Scrollspy index

