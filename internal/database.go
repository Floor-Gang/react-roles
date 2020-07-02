package internal

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "os"
)

type Controller struct {
    db *sql.DB
}

type Roles struct {
    channelID, messageID, reaction, role string
}

// This will get a new database controller
func GetController(location string) Controller {
    if _, err := os.Stat(location); err != nil {
        _, err := os.Create(location)

        if err != nil {
            panic(err)
        }
    }

    db, err := sql.Open("sqlite3", location)

    if err != nil {
        panic(err)
    } else {
        controller := Controller{db: db}
        if err = controller.init(); err != nil {
            panic(err)
        }
        return controller
    }
}

// This initializes the database table if it doesn't exist
func (c Controller) init() error {
    _, err := c.db.Exec(
        "CREATE TABLE IF NOT EXISTS roles (id INTEGER  PRIMARY  KEY , channel_id TEXT, message_id TEXT, reaction TEXT, role TEXT)",
    )

    return err
}

// Create a new role reaction row
func (c Controller) createRoleReaction(channelID string, messageID string, reaction string, role string) error {
    statement, err := c.db.Prepare(
        "INSERT INTO roles (channel_id, message_id, reaction, role) VALUES (?, ?, ?, ?)",
    )

    if err != nil {
        return err
    }

    _, err  = statement.Exec(channelID, messageID, reaction, role)

    return err
}

// Get all
func (c Controller) getAll() ([]Roles, error) {
    var result []Roles

    request, err := c.db.Query("SELECT * FROM roles")

    if err != nil {
        return result, err
    }

    for request.Next() {
        roles := Roles{
          channelID: "",
          messageID: "",
          reaction: "",
          role: "",
        }

        err := request.Scan(&roles.channelID, &roles.messageID, &roles.reaction, &roles.role)

        if err != nil {
            return result, err
        } else {
            result = append(result, roles)
        }
    }

    return result, nil
}