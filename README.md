# Go API client for titan

The ultimate, language agnostic, container based job processing framework.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 0.4.6
- Package version: 0.4.6
- Build date: 2016-06-20T23:00:39.411Z
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./titan"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost:8080/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*GroupsApi* | [**GroupsGet**](docs/GroupsApi.md#groupsget) | **Get** /groups | Get all group names.
*GroupsApi* | [**GroupsNameGet**](docs/GroupsApi.md#groupsnameget) | **Get** /groups/{name} | Get information for a group.
*GroupsApi* | [**GroupsNamePut**](docs/GroupsApi.md#groupsnameput) | **Put** /groups/{name} | Create/update a job group.
*GroupsApi* | [**GroupsPost**](docs/GroupsApi.md#groupspost) | **Post** /groups | Post new group
*JobsApi* | [**GroupsNameJobsGet**](docs/JobsApi.md#groupsnamejobsget) | **Get** /groups/{name}/jobs | Get job list by group name.
*JobsApi* | [**GroupsNameJobsIdCancelPost**](docs/JobsApi.md#groupsnamejobsidcancelpost) | **Post** /groups/{name}/jobs/{id}/cancel | Cancel a job.
*JobsApi* | [**GroupsNameJobsIdDelete**](docs/JobsApi.md#groupsnamejobsiddelete) | **Delete** /groups/{name}/jobs/{id} | Delete the job.
*JobsApi* | [**GroupsNameJobsIdErrorPost**](docs/JobsApi.md#groupsnamejobsiderrorpost) | **Post** /groups/{name}/jobs/{id}/error | Mark job as failed.
*JobsApi* | [**GroupsNameJobsIdGet**](docs/JobsApi.md#groupsnamejobsidget) | **Get** /groups/{name}/jobs/{id} | Gets job by id
*JobsApi* | [**GroupsNameJobsIdLogGet**](docs/JobsApi.md#groupsnamejobsidlogget) | **Get** /groups/{name}/jobs/{id}/log | Get the log of a completed job.
*JobsApi* | [**GroupsNameJobsIdLogPost**](docs/JobsApi.md#groupsnamejobsidlogpost) | **Post** /groups/{name}/jobs/{id}/log | Send in a log for storage.
*JobsApi* | [**GroupsNameJobsIdRetryPost**](docs/JobsApi.md#groupsnamejobsidretrypost) | **Post** /groups/{name}/jobs/{id}/retry | Retry a job.
*JobsApi* | [**GroupsNameJobsIdStartPost**](docs/JobsApi.md#groupsnamejobsidstartpost) | **Post** /groups/{name}/jobs/{id}/start | Mark job as started, ie: status &#x3D; &#39;running&#39;
*JobsApi* | [**GroupsNameJobsIdSuccessPost**](docs/JobsApi.md#groupsnamejobsidsuccesspost) | **Post** /groups/{name}/jobs/{id}/success | Mark job as succeeded.
*JobsApi* | [**GroupsNameJobsIdTouchPost**](docs/JobsApi.md#groupsnamejobsidtouchpost) | **Post** /groups/{name}/jobs/{id}/touch | Extend job timeout.
*JobsApi* | [**GroupsNameJobsPost**](docs/JobsApi.md#groupsnamejobspost) | **Post** /groups/{name}/jobs | Enqueue Job
*JobsApi* | [**JobsGet**](docs/JobsApi.md#jobsget) | **Get** /jobs | Get next job.
*RunnerApi* | [**GroupsNameJobsIdErrorPost**](docs/RunnerApi.md#groupsnamejobsiderrorpost) | **Post** /groups/{name}/jobs/{id}/error | Mark job as failed.
*RunnerApi* | [**GroupsNameJobsIdStartPost**](docs/RunnerApi.md#groupsnamejobsidstartpost) | **Post** /groups/{name}/jobs/{id}/start | Mark job as started, ie: status &#x3D; &#39;running&#39;
*RunnerApi* | [**GroupsNameJobsIdSuccessPost**](docs/RunnerApi.md#groupsnamejobsidsuccesspost) | **Post** /groups/{name}/jobs/{id}/success | Mark job as succeeded.


## Documentation For Models

 - [Complete](docs/Complete.md)
 - [ErrorBody](docs/ErrorBody.md)
 - [Group](docs/Group.md)
 - [GroupWrapper](docs/GroupWrapper.md)
 - [GroupsWrapper](docs/GroupsWrapper.md)
 - [IdStatus](docs/IdStatus.md)
 - [Job](docs/Job.md)
 - [JobWrapper](docs/JobWrapper.md)
 - [JobsWrapper](docs/JobsWrapper.md)
 - [ModelError](docs/ModelError.md)
 - [NewJob](docs/NewJob.md)
 - [NewJobsWrapper](docs/NewJobsWrapper.md)
 - [Start](docs/Start.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author



