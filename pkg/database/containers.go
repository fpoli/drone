package database

import (
    . "github.com/drone/drone/pkg/model"
    "github.com/russross/meddler"
)

// Name of the Build table in the database
const containerTable = "containers"

// SQL Queries to retrieve a list of all Commits belonging to a Repo.
const containerStmt = `
SELECT id, containers
FROM containers
WHERE commit_id = ?
ORDER BY slug ASC
`

// SQL Queries to retrieve a container by id.
const containerFindStmt = `
SELECT id, containers
FROM containers
WHERE id = ?
LIMIT 1
`

// SQL Queries to delete a Commit.
const containerDeleteStmt = `
DELETE FROM containers WHERE build_id = ?
`

// Returns the container with the given ID.
func GetContainer(id int64) (*Container, error) {
    container := Container{}
    err := meddler.QueryRow(db, &container, containerFindStmt, id)
    return &container, err
}

// Creates a new Container.
func SaveContainer(container *Container) error {
    return meddler.Save(db, containerTable, container)
}

// Deletes an existing Container.
func DeleteContainer(id int64) error {
    _, err := db.Exec(containerDeleteStmt, id)
    return err
}

// Returns a list of all Containers associated
// with the specified Build ID.
func ListContainers(id int64) ([]*Container, error) {
    var containers []*Container
    err := meddler.QueryAll(db, &containers, containerStmt, id)
    return containers, err
}
