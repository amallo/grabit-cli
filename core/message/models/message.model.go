package models

import "grabit-cli/core/identities/models"

type Message struct {
	Content string
	From    models.Identity
	To      models.Identity
	Id      string
}
