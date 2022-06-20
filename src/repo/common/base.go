package common

import "fmt"

type base struct {
  dbctx IDatabaseContext
  notifier INotifier
}

func (b *base) Create(
  stream string,
  streamId string,
  newObj interface{},
) (*Event, error) {
  tx, err := b.dbctx.Beginx()
  if err != nil {
    return nil, fmt.Errorf(
      "error creating transaction: %w", err,
    )
  }
  chck := ExistsEntity{}

  // - Test if exists in stream
  err = tx.Get(&chck, CHECK_STREAM_EXISTS_QUERY, stream, streamId)
  if err != nil {
    rberr := tx.Rollback()
    if rberr != nil {
      err = fmt.Errorf("%w: %w", rberr, err)
    }
    return nil, fmt.Errorf(
      "error checking for stream conflicts : %w", err,
    )
  }
  if chck.Exists {
    _ = tx.Rollback() 
    return nil, &RepoError{code: ERROR_STREAM_CONFLICT}
  }

  var newEvent Event
  err = tx.Get(
    &newEvent,
    INSERT_EVENT_QUERY,
    stream,
    streamId,
    0,
    CREATED_EVENT,
    newObj,
  )
  if err != nil {
    rberr := tx.Rollback()
    if rberr != nil {
      err = fmt.Errorf("%w: %w", rberr, err)
    }
    return nil, fmt.Errorf(
      "error inserting new event : %w", err,
    )
  }
  
  err = tx.Commit();
  if err != nil {
    rberr := tx.Rollback()
    if rberr != nil {
      err = fmt.Errorf("%w: %w", rberr, err)
    }
    return nil, fmt.Errorf(
      "error commiting transaction: %w", err,
    ) 
  }
  b.notifier.EnqueueEventNotification(&newEvent)
  return &newEvent, nil
}

func (b *base) Update(
  stream string,
  streamId string,
  fromVer int,
  update interface{},
) (*Event, error) {

}

func (b *base) Delete(
  stream string,
  streamId string,
  fromVer int,
) (*Event, error) {

}

const (
  CREATED_EVENT = "Created"
  
  CHECK_STREAM_EXISTS_QUERY = `
    SELECT EXISTS(
      SELECT * FROM events
      WHERE stream = ($1) && streamid = ($2)
    ) as exists`

  INSERT_EVENT_QUERY = `
    INSERT INTO events (
      stream,
      streamid,
      streamversion,
      event,
      data
    ) VALUES (
      $1, $2, $3, $4, $5
    )
  `
)
