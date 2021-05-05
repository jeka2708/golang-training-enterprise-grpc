# golang-training-enterprise

## Description

##### Application for enterprise

|Path|Method|Description|
|:---:|:---:|:---:|
|```/clients```|```GET```|get all clients|
|```/divisions```|```GET```|get all divisions|
|```/roles```|```GET```|get all roles|
|```/services```|```GET```|get all services|
|```/works```|```GET```|get all works|
|```/works-clients```|```GET```|get all works-clients|
|```/workers```|```GET```|get all workers|

|Path|Method|Description|Body example|
|:---:|:---:|:---:|:---|
|```/clients```|```POST```|create new client|```{"FirstNameC":"First","LastNameC":"Last","MiddleNameC":"Middle","PhoneNumberC":"+37525333333"}```|
|```/divisions```|```POST```|create new division|```{"DivisionName":"test"}```|
|```/roles```|```POST```|create new role|```{"Name":"test","DivisionName":"testD"}```|
|```/services```|```POST```|create new service|```{"Name":"test","Cost":"22"}```|
|```/works```|```POST```|create new work|```{"WorkerId":"2","ServiceId":"3"}```|
|```/work-client```|```POST```|create new work-clients|```{"FirstNameC":"First","LastNameC":"Last","MiddleNameC":"MiddleNameC","PhoneNumberC":"+37525333333"}```|
|```/workers```|```POST```|create new worker|```{"FirstName":"First","LastName":"Last","MiddleName":"MiddleNameC","PhoneNumber":"+37525333333","RoleId":"3"}```|

|```/clients```|```DELETE```|delete new client by id|```id```|
|```/divisions```|```DELETE```|delete new division by id|```id```|
|```/roles```|```DELETE```|delete new role by id|```id```|
|```/services```|```DELETE```|delete new service by id|```id```|
|```/works```|```DELETE```|delete new work by id|```id```|
|```/work-client```|```DELETE```|delete new work-clients by id|```id```|
|```/workers```|```DELETE```|delete new worker by id|```id```|
|```/clients```|```PUT```|update client|```{"id":"1",FirstNameC":"First","LastNameC":"Last","MiddleNameC":"Middle","PhoneNumberC":"+37525333333"}```|
|```/divisions```|```PUT```|update division|```{"id":"1","DivisionName":"test"}```|
|```/services```|```PUT```|update service|```{"id":"1","Name":"test","Cost":"22"}```|

## Usage

1. Run server on port ```8080```

> ```go run ./golang-training-enterprise/cmd/main.go```

2. Open URL ```http://localhost:8080```