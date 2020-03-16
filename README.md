# Librus API client 

Simple and not completed librus api written in Go.

## Usage
1. Create new instance of `Librus` struct with student username and password
2. Call `Login` method
3. Call any other method

## Functions
#### `Login()` 
Required, saves access token for future usage

#### `GetData(url)` 
Retruns http.Response

####  `GetLuckyNumber()`
Returns lucky number and date of it

#### `GetUserRealName()`
Returns user's first and last name

#### `GetUserClass()`
Returns the class to which the user belongs

#### `GetUserGrades()`
Returns info about every grade. Check `GradeDetails` struct in `model.go` for more.

#### `GetSubject(id)`
Returns subject name. 

#### `GetGradeCategory(id)`
Returns grade category name

#### `GetUser(id)`
Returns user's info, also teacher's info like first name and last name.


## TODO
- [x] Lucky number
- [x] Student info
- [x] Subject info
- [x] Student grades
- [ ] School info
- [x] Attendances
- [ ] Messages
- [ ] Timetable
- [ ] Class free days
- [ ] Behaviour grades
- [ ] Parent teacher conferences
- [ ] What's new since last login
- [ ] Return in json format

## Projects using this API client
None :ok_hand:
