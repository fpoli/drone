package database

import (
    . "github.com/drone/drone/pkg/model"
    "github.com/russross/meddler"
)

// Name of the Build table in the database
const containerTable = "containers"

// SQL Queries to retrieve a container by build_id.
const containerFindStmt = `
SELECT build_id, containers
FROM containers
WHERE build_id = ?
LIMIT 1
`

// SQL Queries to delete a container.
const containerDeleteStmt = `
DELETE FROM containers WHERE build_id = ?
`

// Returns the container with the given build_id.
func GetContainer(build_id int64) (*Container, error) {
    container := Container{}
    err := meddler.QueryRow(db, &container, containerFindStmt, build_id)
    return &container, err
}

// Creates a new Container.
func SaveContainer(container *Container) error {
    return meddler.Save(db, containerTable, container)
}

// Deletes an existing Container with the given build_id.
func DeleteContainer(build_id int64) error {
    _, err := db.Exec(containerDeleteStmt, build_id)
    return err
}
