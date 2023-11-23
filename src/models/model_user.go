package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/kentlouisetonino/aggreflow/database/sqlc"
)

type User struct {
  ID        uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Name      string    `json:"name"`
}

func DatabaseUserToUser(dbUser sqlc.User) User {
  return User{
    ID:        dbUser.ID,
    CreatedAt: dbUser.CreatedAt,
    UpdatedAt: dbUser.UpdatedAt,
    Name:      dbUser.Name,
  }
}

