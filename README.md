# fastrbac - access control for golang

Access control system for Go inspired by GitHubs implementation - designed to be simple to use and fast.

## Interface types
- Owner - Someone who has permissions on stuff
- Object - An object to have permissions on
- Repository - Interface towards a database

## Data object types
- Trust - Defines permission on objects which you are not the owner
- Role - Defines permission on another owners objects

## Database implementations
- App Engine Datastore - `respositories/gae-datastore.go`

Currently there is only one concrete implementation of a database repository, the database logic has been refactored in such a manner that makes implementing a new database easy.