# REST_I-O_Bound
- REST_IO_Bound — is a demo REST API that models I/O-bound tasks with lazy processing. Definition tasks automatically transition to the "Running" state and then to "Done" after a timer.
---
## Implemented capabilities
- CRUD API for tasks
- Asynchronous processing (with imitation of IO-bound operations)
- Generating and reading YAML configuration
---
## Installing  and Launching
```bash
    git clone https://github.com/IdzAnAG1/REST_I-O_Bound.git
    cd  /Rest_I-O_Bound
    make // make run - local launch
```
---
## Endpoints
| Method  | URL | Description   |
|---------|-----|---------------|
| GET     | /   | Get All Tasks | 
| POST    |   / | Create New Task | 
| GET     |   /:id| Get Task by ID|
| DELETE  | /:id | Delete Task by ID |
| --- | --- | --- |
---
## Example post structure 
```JSON
{
  "title": " Name of Task "
}
```
___

### Project structure
```text
.
├── cmd
│     └── REST_IO_Bound
│         └── main.go
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│     ├── config
│     │    ├── config_example.yaml
│     │    └── config.go
│     ├── handlers
│     │    └── handlers.go
│     ├── models
│     │    └── model.go
│     ├── server
│     │    └── server.go
│     └── variables
│         └── variables.go
├── Makefile
└── README.md


```