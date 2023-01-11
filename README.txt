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
    dashboards  *[]Dashboard
    createdAt   time.Time
    updatedAt   time.Time

MembersTeam
    userId      string
    teamId      string

Dashboard
    id          string
    teamId      string
    team        Team
    ownerUserId User
    ownerUser   User
    notes       *[]Note

Note
    id          string
    teamId      string
    ownerId     string
    owner       User
    topic       string
    description string
    status      bool
    comments     *[]Comment
    createdAt   time.Time
    updatedAt   time.Time
    deadlineAt  time.Time

Comment
    id          string
    teamId      string
    noteId      string
    userId      string
    description string
    createdAt   time.Time
    updatedAt   time.Time


