package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (s *Server) CreateRecord(c fuego.ContextWithBody[CreateRecordData]) (any, error) {
	ctx := c.Context()

	accountID := c.Value(TokenKey).(uuid.UUID)

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	err = s.Record.CreateRecord(ctx, accountID, body.Time)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Account with given email already existed",
		}
	}

	return nil, nil
}

func (s *Server) GetAllRecord(c fuego.ContextNoBody) ([]Record, error) {
	ctx := c.Context()

	raw, err := s.Record.GetAllRecord(ctx)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to get records",
		}
	}

	var records []Record
	copier.Copy(&records, &raw)

	return records, nil
}
