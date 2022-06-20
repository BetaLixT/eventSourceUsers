package common

import "fmt"

type base struct {
  dbctx *IDatabaseContext
}

func (b *base) Create(
  stream string,
  streamId string,
  newObj interface{},
) (Event, interface{}, error) {
  tx, err := (*b.dbctx).Beginx()
  if err != nil {
    return Event{}, nil, fmt.Errorf(
      "error creating transaction: %w", err,
    )
  }
  chck := ExistsEntity{}

  // - Test if exists in stream
  err = (*tx).Get(&chck, CHECK_STREAM_EXISTS, stream, streamId)
  if err != nil {
    rberr := (*tx).Rollback()
    if rberr != nil {
      err = fmt.Errorf("%w: %w", rberr, err)
    }
    return Event{}, nil, fmt.Errorf(
      "error checking for stream conflicts : %w", err,
    )
  }
  if chck.Exists {
    _ = (*tx).Rollback() 
    return Event{}, nil, &RepoError{code: ERROR_STREAM_CONFLICT}
  }

  err = (*tx).Commit();
  if err != nil {
    rberr := (*tx).Rollback()
    if rberr != nil {
      err = fmt.Errorf("%w: %w", rberr, err)
    }
    return Event{}, nil, fmt.Errorf(
      "error commiting transaction: %w", err,
    ) 
  }
}

func (b *base) Update(int, interface{}) (Event, interface{}) {

}
func (b *base) Delete(int, interface{}) (Event, interface{}) {

}

const (
  CHECK_STREAM_EXISTS = `
    SELECT EXISTS(
      SELECT * FROM events
      WHERE stream = ($1) && streamid = ($2)
    ) as exists`

  INSERT_EVENT = `

  `
)
