```markdown
# GORM Common Repository 

The gorm-common-repository is a Go package that simplifies common database operations
using the GORM ORM (Object-Relational Mapping) library. 
It provides a set of interfaces and functions to streamline CRUD (Create, Read, Update, Delete) operations and querying of database records.

## Features

- Create, read, update, and delete records in your database with ease.
- Designed to be adaptable to any GORM model.
- Handle errors gracefully and return meaningful error messages.
- Improve code re-usability by using common database patterns.
```
## Installation
To use this library in your Go project, you can install it using the `go get` command:
```shell
go get github.com/bondhansarker/gorm-common-repository
```

## Usage

```go
package main

import (
	"github.com/bondhansarker/gorm-common-repository"
	"gorm.io/gorm"
)

// User Model Definition
type User struct {
	ID    uint
	Name  string
	Email string
}

// UserRepository represents a repository for User models.
type UserRepository struct {
	dbClient *gorm.DB
	gorm_common_repository.CommonRepositoryInterface[User]
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(dbClient *gorm.DB) UserRepository {
	return UserRepository{
		dbClient:                  dbClient,
		CommonRepositoryInterface: gorm_common_repository.NewCommonRepository[User]("users", dbClient),
	}
}

// Post Model Definition
type Post struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
}

// PostRepository represents a repository for Post models.
type PostRepository struct {
	dbClient *gorm.DB
	gorm_common_repository.CommonRepositoryInterface[Post]
}

// NewPostRepository creates a new PostRepository instance.
func NewPostRepository(dbClient *gorm.DB) PostRepository {
	return PostRepository{
		dbClient:                  dbClient,
		CommonRepositoryInterface: gorm_common_repository.NewCommonRepository[Post]("posts", dbClient),
	}
}

func main() {
	// Please make sure that you have defined the gormDbClient variable which will connect with the database
	// Create a User repository instance
	userRepo := NewUserRepository(gormDbClient)

	// Use the repository to perform database operations
	// For example, fetching a record by field:
	user, err := userRepo.GetRecordByAttributes(map[string]interface{}{
		"name": "Bondhan",
	})
	if err != nil {
		// Handle the error
	}
	// 'user' now contains the fetched user record

	// Create a Post repository instance
	postRepo := NewPostRepository(gormDbClient)

	// Use the repository to perform database operations
	// For example, creating a new record:
	post := Post{UserID: user.ID, Title: "Hello world", Description: "This is an example"}
	createdPost, err := postRepo.CreateRecord(post)
	if err != nil {
		// Handle the error
	}
	// 'createdPost' now contains the newly created post record

	// Other operations like Get, Update, Delete, and Query are similarly straightforward
}

```

For detailed usage instructions and examples, refer to the [documentation](https://github.com/bondhansarker/gorm-common-repository).

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please [open an issue](https://github.com/bondhansarker/gorm-common-repository/issues) on GitHub.

## Author

- Bondhan Sarker
- GitHub: [github.com/bondhansarker](https://github.com/bondhansarker)
