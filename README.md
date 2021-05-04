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
|```/create-client```|```POST```|create new client|```{"FirstNameC":"First","LastNameC":"Last","MiddleNameC":"Middle","PhoneNumberC":"+37525333333"}```|
|```/create-division```|```POST```|create new division|```{"DivisionName":"test"}```|
|```/create-role```|```POST```|create new role|```{"Name":"test","DivisionName":"testD"}```|
|```/create-service```|```POST```|create new service|```{"Name":"test","Cost":"22"}```|
|```/create-work```|```POST```|create new work|```{"WorkerId":"2","ServiceId":"3"}```|
|```/create-work-clients```|```POST```|create new work-clients|```{"FirstNameC":"First","LastNameC":"Last","MiddleNameC":"MiddleNameC","PhoneNumberC":"+37525333333"}```|
|```/create-worker```|```POST```|create new worker|```{"FirstName":"First","LastName":"Last","MiddleName":"MiddleNameC","PhoneNumber":"+37525333333","RoleId":"3"}```|
|```/delete-client```|```POST```|delete new client by id|```id```|
|```/delete-division```|```POST```|delete new division by id|```id```|
|```/delete-role```|```POST```|delete new role by id|```id```|
|```/delete-service```|```POST```|delete new service by id|```id```|
|```/delete-work```|```POST```|delete new work by id|```id```|
|```/delete-work-clients```|```POST```|delete new work-clients by id|```id```|
|```/delete-worker```|```POST```|delete new worker by id|```id```|
|```/update-client```|```POST```|update client|```{"id":"1",FirstNameC":"First","LastNameC":"Last","MiddleNameC":"Middle","PhoneNumberC":"+37525333333"}```|
|```/update-division```|```POST```|update division|```{"id":"1","DivisionName":"test"}```|
|```/update-service```|```POST```|update service|```{"id":"1","Name":"test","Cost":"22"}```|

## Usage

1. Run server on port ```8080```

> ```go run ./golang-training-enterprise/cmd/main.go```

2. Open URL ```http://localhost:8080```