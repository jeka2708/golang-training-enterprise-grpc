# golang-training-enterprise

## Description

##### Application for enterprise

|Path|Method|Description|
|:---:|:---:|:---:|
|```/api/v1/clients```|```GET```|get all clients|
|```/api/v1/divisions```|```GET```|get all divisions|
|```/api/v1/roles```|```GET```|get all roles|
|```/api/v1/services```|```GET```|get all services|
|```/api/v1/works```|```GET```|get all works|
|```/api/v1/works-clients```|```GET```|get all works-clients|
|```/api/v1/workers```|```GET```|get all workers|

|Path|Method|Description|Body example|
|:---:|:---:|:---:|:---|
|```/api/v1/clients```|```POST```|create new client|```{"FirstNameC":"First","LastNameC":"Last","MiddleNameC":"Middle","PhoneNumberC":"+37525333333"}```|
|```/api/v1/divisions```|```POST```|create new division|```{"DivisionName":"test"}```|
|```/api/v1/roles```|```POST```|create new role|```{"Name":"test","DivisionName":"testD"}```|
|```/api/v1/services```|```POST```|create new service|```{"Name":"test","Cost":"22"}```|
|```/api/v1/works```|```POST```|create new work|```{"WorkerId":"2","ServiceId":"3"}```|
|```/api/v1/work-client```|```POST```|create new work-clients|```{"FirstNameC":"First","LastNameC":"Last","MiddleNameC":"MiddleNameC","PhoneNumberC":"+37525333333"}```|
|```/api/v1/workers```|```POST```|create new worker|```{"FirstName":"First","LastName":"Last","MiddleName":"MiddleNameC","PhoneNumber":"+37525333333","RoleId":"3"}```|
|```/api/v1/clients```|```DELETE```|delete new client by id|```id```|
|```/api/v1/divisions```|```DELETE```|delete new division by id|```id```|
|```/api/v1/roles```|```DELETE```|delete new role by id|```id```|
|```/api/v1/services```|```DELETE```|delete new service by id|```id```|
|```/api/v1/works```|```DELETE```|delete new work by id|```id```|
|```/api/v1/work-client```|```DELETE```|delete new work-clients by id|```id```|
|```/api/v1/workers```|```DELETE```|delete new worker by id|```id```|
|```/api/v1/clients```|```PUT```|update client|```{"id":"1",FirstNameC":"First","LastNameC":"Last","MiddleNameC":"Middle","PhoneNumberC":"+37525333333"}```|
|```/api/v1/divisions```|```PUT```|update division|```{"id":"1","DivisionName":"test"}```|
|```/api/v1/services```|```PUT```|update service|```{"id":"1","Name":"test","Cost":"22"}```|

## Usage

1. Run server on port ```8080```

> ```docker-compose up```

2. Open URL ```http://localhost:8081```