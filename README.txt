golang wanna be trello

docker-compose up -d
docker-compose down

go run main.go

localhost:4011/api
    /auth
        POST    /register
        POST    /login
    /users
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
    id          string(ObjectId)
    ownerId     string(ObjectId)
    owner       User
    name        string
    members     *[]User
    createdAt   time.Time
    updatedAt   time.Time

MembersTeam
    userId      string(ObjectId)
    teamId      string(ObjectId)

Note
    id          string(ObjectId)
    ownerId     string(ObjectId)
    owner       User
    topic       string
    description string
    createdAt   time.Time
    updatedAt   time.Time
    deadlineAt  time.Time




