golang wanna be trello

docker-compose up -d
docker-compose down

go run main.go

localhost:4011/api
    /auth
        POST    /register
        POST    /login
    /user
        GET     /
        GET     /:id
        PUT     /:id
        DELETE  /:id
    /team
        POST    /
        GET     /
        GET     /:id

User

Team
    id          string
    ownerId     string
    owner       User
    name        string
    members     *[]User
    createdAt   time.Time
    updatedAt   time.Time

MembersTeam
    userId      string
    teamId      string

Note
    id          string
    ownerId     string
    owner       User
    topic       string
    description string
    createdAt   time.Time
    updatedAt   time.Time
    deadlineAt  time.Time




