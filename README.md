# Todo list planner

## Description
This project is a web server with a Rest API under the hood for task management. It also includes a pre-compiled web interface.

## The main features
- Task Creation `POST /api/task`
- Editing a task `PUT /api/task`
- Getting information on the task `GET /api/task`
- Deleting a task `DELETE /api/task`
- Marking of task completion `POST /api/task/done`
- Getting the list of tasks with the ability to filter by a specific date or search by task name/comment `GET /api/tasks`
- Getting the next due date of a task `GET /api/nextdate`
- Authentication `POST /api/signin`

## Stack
- Go 1.24
- Sqlite3
- Used go packets: modernc.org/sqlite, github.com/golang-jwt/jwt/v5, github.com/go-chi/chi/v5

## Api endpoints

### Notes*
- To use Api endpoints with the ***authenticated** flag you need to authenticate via a request to `POST /api/signin`.
The **token** received from the response must be passed to **Cookie[“token”]** for each such request.
- The fields marked with * are required

### `POST /api/task` *authenticated
```json
    request:
    {
        "date": "20250613",
        "title": "Meeting at 10:00",
        "comment": "Discussion about plans for the future",
        "repeat": "d 2"
    }
```
**date**: next task start date

**title***: title of the task

**comment**: description of the task

**repeat**: rules for repeating the task, available formats with examples and description:
- “d 1” every day
- “d 7” every 7 days
- “y” every year on the same day.
- “w 7” reschedule the task for the next Sunday.
- “w 2,3” moving the task to the next Tuesday and Wednesday
- “m 4” every 4th day of every month
- “m -1” every last day of every month
- “m -2” every penultimate day of each month.
- “m 6 3,9,10” postponement to the 6th day of March, September, October

### `PUT /api/task` *authenticated
```json
    request:
    {
        "id": "1"
        "date": "20250615",
        "title": "Meeting at 10:00",
        "comment": "Discussion about plans for the future",
        "repeat": "w 7"
    }
```
**id***: identifier of the task

**date**: next task start date

**title***: title of the task

**comment**: description of the task

**repeat**: rules for repeating the task, available formats with examples and description:
- “d 1” every day
- “d 7” every 7 days
- “y” every year on the same day.
- “w 7” reschedule the task for the next Sunday.
- “w 2,3” moving the task to the next Tuesday and Wednesday
- “m 4” every 4th day of every month
- “m -1” every last day of every month
- “m -2” every penultimate day of each month.
- “m 6 3,9,10” postponement to the 6th day of March, September, October

### `GET /api/task?id=1` *authenticated
```json
    response:
    {
        "id": "1"
        "date": "20250615",
        "title": "Meeting at 10:00",
        "comment": "Discussion about plans for the future",
        "repeat": "w 7"
    }
```

### `DELETE /api/task?id=1` *authenticated

**id***: identifier of the task

### `POST /api/task/done?id=1` *authenticated

**id***: identifier of the task

### `GET /api/tasks?search=2025.06.01` *authenticated
```json
    response:
    [
        {
            "id": "5"
            "date": "20250601",
            "title": "Meeting at 20:00",
            "comment": "New strategy discussion",
            "repeat": "w 2"
        },
        {
            "id": "19"
            "date": "20250601",
            "title": "Meeting at 15:30",
            "comment": "Discussion about plans for the future",
            "repeat": "w 5"
        }
    ]
```
**search**: string for filtering and searching, valid formats:
- “2025.06.01” - for filtering by exact date match
- “about” - to search by task title or comment

### `GET /api/nextdate?now=20250615&date=20250613&repeat=w 7`
```
    response: 20250622
```

**now**: current time

**date***: next task start date

**repeat***: rules for repeating the task, available formats with examples and description:
- “d 1” every day
- “d 7” every 7 days
- “y” every year on the same day.
- “w 7” reschedule the task for the next Sunday.
- “w 2,3” moving the task to the next Tuesday and Wednesday
- “m 4” every 4th day of every month
- “m -1” every last day of every month
- “m -2” every penultimate day of each month.
- “m 6 3,9,10” postponement to the 6th day of March, September, October

### `POST /api/signin`
```json
    request:
    {
        "password": "qwerty"
    }

    response:
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxMTUxMjAsInVzZXJuYW1lIjoiIn0.1CQ7wGCFFijSvWQgZ3VZNH0rO_c_QkBc8l8rCZrc3QM"
    }
```
**password***: password to authenticate

------------------------------------------------------

## Install ##
- Clone a repository
```bash
   git clone https://github.com/TutunaruStanislav/todo-list-final.git
```
- Enter to project directory
```bash
    cd todo-list-final
```
- Create ENV file and edit it if needs
```bash
    cp .env.example .env
```

## Run locally ##
### Requirements ###
- Go 1.24 and higher
- git

### Run ###


- Install requirements
```bash
   go mod tidy
```
- Run server
```bash
   go run .
```

### Run API Tests ###

#### Notes* ####
- Ensure that the server is pre-started
- Update the **Token** variable value in the test settings file placed at `tests/settings.go` to the value received in response to the `POST /api/signin` request

- Run tests
```bash
   go test ./tests
```

## Run docker container ##
- Build docker image
```bash
    docker build --build-arg TODO_PORT=7540 -t todo-list-app .
```

- Run container
```bash
    docker run -p 7540:7540 todo-list-app
```

## Information
- The application will be available at [http://localhost:7540/](http://localhost:7540/)

- Default password: `qwerty`

## Author
Stanislav Tutunaru

## License
MIT