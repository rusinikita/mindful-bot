package airtable

import (
	"errors"

	"github.com/rusinikita/discipline-bot/db"
)

func (b base) One(id db.ID, entity interface{}) error {
	record := record{}

	r, err := b.client.R().
		SetResult(&record).
		Get(db.TableName(entity) + "/" + string(id))
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	record.Fields["id"] = record.ID

	return decode(record.Fields, entity)
}
