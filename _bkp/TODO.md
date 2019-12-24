## TODO

https://github.com/adigunhammedolalekan/go-contacts/blob/master/app/auth.go
https://gist.github.com/zmts/802dc9c3510d79fd40f9dc38a12bccfc
### Actors
- admin: user with admin access to system
- service company: service company employee
- guard: guarding company employee
- neighbor: authorized inhabitant

### Auth
- [x] POST /login - auth in app
- [ ] GET /logout - logout from app

### Requests
- [ ] POST /request - save request from neighbor
- [ ] GET /requests - get requests list (should support filtering and pagination)
- [ ] GET /request/{id} - get request data
- [ ] PUT /request/{id} - edit request
- [ ] DELETE /request/id - delete request

### Users
- [x] POST /user - create user
- [ ] GET /user/{id} - get user data
- [ ] GET /users - get users list (should support filtering and pagination)
- [ ] PUT /user/{id} - update user data
- [ ] DELETE /user/{id} - delete user

### Tasks
1. app should be able to authorize user  
    - at initial stage we'll have hardcoded guard, admin, and service company accounts
    - admin should be able to add neighbors
2. admin should be able to add guards, service company workers, neighbors
3. neighbor should be able to create request, view his requests, edit ongoing requests. in future - create request templates
4. guard should be able to see events list and change their statuses on complete/decline
5. there should be several statuses of a request: 
    - submitted: after user submits the request;
    - in_progress: when request in in progress of execution;
    - completed: when executor completes the request and it's done;
    - rejected: set by executor to notify that it's not possible to execute specific request;
    - cancelled: when user decides to cancel his request. 

