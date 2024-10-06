# Week 5 of golang assignment

In this assignment, you will build a RESTful API in Go (using the net/http built-in package) to manage a collection of "tasks". Each task will have a title, description, and a status (pending, completed). You will implement the following CRUD operations:

- Create a new task.
- Read a list of all tasks or a single task by ID.
- Update an existing task.
- Delete a task by ID.

Steps

- Create a golang module 
``$ go mod init task_list``

- Get the relevant repos from git
``$ go get .``

- Build
``$ go build main.go``

- Run
``$ ./main``

Usage

- Add tasks

``curl http://localhost:8080/tasks --include --header "Content-Type: application/json" --request "POST" --data '{\"id\":\"12\",\"title\":\"Grocery\",\"description\":\"Buy grocery\",\"status\":\"completed\"}'``

- Read Tasks

``curl http://localhost:8080/tasks`` 

- Read Task by ID

``curl http://localhost:8080/tasks/2`` (id = 2)

- Update Tasks

 ``curl http://localhost:8080/tasks/2 --include --header "Content-Type: application/json" --request "PUT" --data '{\"status\":\"completed\"}'``


- Delete Tasks

``curl http://localhost:8080/tasks/1 --include --request "DELETE"``

NOTE:
Window users have to run ``Remove-item alias:curl`` to avoid curl argument error

