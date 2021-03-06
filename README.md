# Cookie-Finder API

An API to find the best cookies around! Honestly this is just a play API to practice integrating configurations, logging, and mongo into a database.

In order to run test against the handler users will need to make sure there is a instance of MongoDB running locally.

---

### Things to do 

- [x] Set up database configurations (mongo)
- [x] Set up configs (viper)
- [x] Set up formatted logging (Zap)
- [-] Write test for http request (httptest)
- [ ] Set-up JWTs for secured API
- [x] Set-up linter

--- 

### Thoughts

- How do I want to organize how the app connects to the database?
- *Give a package everything it needs to operate on it's own.*
- I need to test the database behavior, as well as, the endpoints too
- Moving forward it makes sense for the endpoints to return a value for testing purposes
  - Changed mind. With the httptest library there's no need for this. The mux docs has some good info on how to handle this [here.](https://github.com/gorilla/mux#testing-handlers)
- How do I separate the database calls from the handler logic?
- It doesn't make sense to have database configuration in the helper package
- The Update function shouldn't be able to take any other parameters or queries to avoid security concerns.
- How do I set-up timeouts on my endpoints?