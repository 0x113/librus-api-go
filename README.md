<h1 align="center">
	<img src="https://imgur.com/XDq2eiP.png">
</h1>
<img src="https://travis-ci.com/0x113/librus-api-go.svg?branch=master">
<img src="https://img.shields.io/codefactor/grade/github/0x113/librus-api-go/master">
<img src="https://img.shields.io/github/repo-size/0x113/librus-api-go">
<img src="https://img.shields.io/github/go-mod/go-version/0x113/librus-api-go">

## Usage
1. Create new instance of `Librus` struct with student's username and password
2. Create new session via `Login` method
3. Call any other method ([List of available methods](#methods))

## Basic example
```go
package main

import (
	"fmt"
	"net/http"

	golibrus "github.com/0x113/librus-api-go"
)

func main() {
	httpClient := &http.Client{}
	client := &golibrus.Librus{
		Client: httpClient,
		Username: "student username", // only student, ends with "u", e.g. "1111111u"
		Password: "student password",
	}

	if err := client.Login(); err != nil {
		panic(err)
	}

	// here you can call any other method
	luckyNumber, err := client.GetLuckyNumber()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Date: %s, lucky number: %d", luckyNumber.LuckyNumberDay, luckyNumber.LuckyNumber)
	// output: Date: 2020-03-12, lucky number: 12
}
```

## Methods <a name="methods"></a>

#### `Login()` 
Required, saves access token for future usage

####  `GetLuckyNumber()`
Returns lucky number and date of it
<details>
<summary>Example</summary>
<p>

```go
// CREATE SESSION BEFORE [see basic example for details]
luckyNumber, err := client.GetLuckyNumber()
if err != nil {
	panic(err)
}
fmt.Printf("Date: %s, lucky number: %d", luckyNumber.LuckyNumberDay, luckyNumber.LuckyNumber)
// output: Date: 2020-03-12, lucky number: 12
```

</p>
</details>

#### `GetUserRealName()`
Returns user's first and last name
<details>
<summary>Example</summary>
<p>

```go
// CREATE SESSION BEFORE [see basic example for details]
fullName, err := client.GetUserRealName()
if err != nil {
	panic(err)
}
fmt.Printf("Full name: %s", fullName)
// output: Full name: Jan Kowalski
```

</p>
</details>

#### `GetUserClass()`
Returns new <a href="https://github.com/0x113/librus-api-go/blob/2698f602c640fa967d508a193488c89dd17f67c9/model.go#L57-L63" target="_blank">`ClassDetails`</a> object
<details>
<summary>Example</summary>
<p>

```go
// CREATE SESSION BEFORE [see basic example for details]
classDetails, err := client.GetUserClass()
if err != nil {
	panic(err)
}
fmt.Printf("Class: %s%s", classDetails.Number, classDetails.Symbol)
fmt.Printf("End of first semester: %s", classDetails.EndFirstSemester)
fmt.Printf("School year: %s -> %s", classDetails.BeginSchoolYear, classDetails.EndSchoolYear)
// output: Class: 3e
// output: End of first semester: 2020-01-03
// output: School year: 2019-09-02 -> 2020-06-23
```

</p>
</details>

#### `GetUserGrades()`
Returns info about every grade. Check <a href="https://github.com/0x113/librus-api-go/blob/2698f602c640fa967d508a193488c89dd17f67c9/model.go#L88-L93">`GradeDetails`</a> struct for more details.

<details>
<summary>Example</summary>
<p>

```go
// CREATE SESSION BEFORE [see basic example for details]
userGrades, err := client.GetUserGrades()
if err != nil {
	panic(err)
}
for _, grade := range userGrades {
	fmt.Printf("Grade %s, subject: %s, added by: %s %s", grade.Grade, grade.Subject, grade.AddedBy.FirstName, grade.AddedBy.LastName)
}
// output: ...
// output: Grade: 5, subject: Matematyka, added by: Jan Kowalski
// output: Grade: 4, subject: Fizyka, added by: Jan Nowak
// output: ...
```

</p>
</details>


#### `GetSubject(id)`
Returns subject name. 
<details>
<summary>Example</summary>
<p>

```go
// CREATE SESSION BEFORE [see basic example for details]
subject, err := client.GetSubject(12300212)
if err != nil {
	panic(err)
}

fmt.Printf("Subject name: %s", subject)
// output: Subject name: Matematyka
```

</p>
</details>

#### `GetGradeCategory(id)`
Returns grade category name
<details>
<summary>Example</summary>
<p>

```go
// CREATE SESSION BEFORE [see basic example for details]
category, err := client.GetCategory(12300212)
if err != nil {
	panic(err)
}

fmt.Printf("Category name: %s", category)
// output: Category name: Sprawdzian
```

</p>
</details>

#### `GetUser(id)`
Returns user's info, also teacher's info like first name and last name.
<details>
<summary>Example</summary>
<p>

```go
// CREATE SESSION BEFORE [see basic example for details]
userInfo, err := client.GetUser(12300212)
if err != nil {
	panic(err)
}

fmt.Printf("First name: %s, last name: %s", userInfo.FirstName, userInfo.LastName)
// output: First name: Jan, last name: Kowalski 
```

</p>
</details>

## TODO
- [x] Lucky number
- [x] Student info
- [x] Subject info
- [x] Student grades
- [x] School info
- [x] Attendances
- [ ] Messages
- [x] Timetable
	* [ ] Get classrooms
- [ ] Class free days
	* [x] Get all class free days
	* [ ] Filter data to get only for specific class
- [ ] Behaviour grades
- [ ] Parent teacher conferences
- [ ] What's new since last login
- [ ] Return in json format

## Projects using this API client
None :ok_hand:
