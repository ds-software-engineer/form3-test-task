# Form3 Take Home Exercise - Dmitriy Suhinin(suhinin.dmitriy@gmail.com)

### Prerequisite
- Go 1.18
- Makefile
- Docker

### Useful Commands
- to run integration tests over fake client and fake API.
```bash
docker-compose up 
```
- to get the help over existing `makefile` targets.
```bash
make help
```
- to run only unit tests without fake API.
```bash
make go_test_unit
```
- to run linters over the code.
```bash
make lint
```

### Possible Client Improvements
- more unit tests maybe for negative cases.
- extend `transport` layer to support more options.
- maybe re organize structure, re name some objects but not fully sure.
- re think about error handling inside the `transport` layer. right now not ideal and feels a bit ugly.
- maybe we even don't need to have `transport` layer. could be enough to configure overall client only with `base` url.
- logger will be useful in case when the client gets bigger.
- passing `context` to the clinet action also could be useful to control execution time.

### Possible API Improvements
- will be nice if API could return an `error` object using next format:
```json
{
  "code": 1233,
  "message": "super puper error"
}
```
in that case it is much easier to validate error only using `code` instead of comparing `message` value.
- in case of `error`, API returns in `error_message` the set of errors. It is a bit hard to handle them. Here could be different ways to improve:
  - return always only one `error` instead of set like right now.
  - return a bit more complicated `error` object where you can uniquely identify each error. right now everything just wrapper into `error_message` property.  

### Strange Things 
- official API documentation not really aligned with Fake API. I think Fake API a bit simplified.


