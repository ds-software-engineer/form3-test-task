# Form3 Take Home Exercise

## Instructions

This exercise has been designed to be completed in 2-4 hours. The goal of this exercise is to write a client library 
in Go to access our fake [account API](http://api-docs.form3.tech/api.html#organisation-accounts) service. 

### Should
- Client library should be written in Go
- Implement only `Create`, `Fetch`, `List` only paging not filtering and `Delete` as [documented](http://api-docs.form3.tech/api.html#organisation-accounts)
- Focus on writing full-stack tests that cover the full range of expected and unexpected use-cases
 - Tests can be written in Go idomatic style or in BDD style. Form3 engineers tend to favour BDD. Make sure tests are easy to read
 - If you encounter any problems running the fake accountapi we would encourage you to do some debugging first, 
before reaching out for help

#### Docker-compose

 - Add your solution to the provided docker-compose file
 - We should be able to run `docker-compose up` and see your tests run against the provided account API service 

### Please don't
- Use a code generator to write the client library
- Implement an authentication scheme

## How to submit your exercise
- Create a private repository, copy the `docker-compose` from this repository
- Email our recruitment team to let them know you have finished the assignment, they will then ask you to add to reviewers