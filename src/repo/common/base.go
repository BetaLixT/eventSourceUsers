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
  tx, err := b.dbctx.Beginx()
  if err != nil {
    return nil, fmt.Errorf(
      "error creating transaction: %w", err,
    )
  }
  chck := ExistsEntity{}

  // - Test if exists in stream
  err = tx.Get(
    &chck,
    CHECK_STREAM_VERSION_EXISTS_QUERY,
    stream,
    streamId,
    fromVer,
  )
  if err != nil {
    rberr := tx.Rollback()
    if rberr != nil {
      err = fmt.Errorf("%w: %w", rberr, err)
    }
    return nil, fmt.Errorf(
      "error checking for version validity: %w", err,
    )
  }
  if chck.Exists {
    _ = tx.Rollback() 
    return nil, &RepoError{code: ERROR_VERSION_INVALID}
  }

  var newEvent Event
  err = tx.Get(
    &newEvent,
    INSERT_EVENT_QUERY,
    stream,
    streamId,
    fromVer + 1,
    UPDATED_EVENT,
    update,
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

func (b *base) Delete(
  stream string,
  streamId string,
  fromVer int,
) (*Event, error) {
  tx, err := b.dbctx.Beginx()
  if err != nil {
    return nil, fmt.Errorf(
      "error creating transaction: %w", err,
    )
  }
  chck := ExistsEntity{}

  // - Test if exists in stream
  err = tx.Get(
    &chck,
    CHECK_STREAM_VERSION_EXISTS_QUERY,
    stream,
    streamId,
    fromVer,
  )
  if err != nil {
    rberr := tx.Rollback()
    if rberr != nil {
      err = fmt.Errorf("%w: %w", rberr, err)
    }
    return nil, fmt.Errorf(
      "error checking for version validity: %w", err,
    )
  }
  if chck.Exists {
    _ = tx.Rollback() 
    return nil, &RepoError{code: ERROR_VERSION_INVALID}
  }

  var newEvent Event
  err = tx.Get(
    &newEvent,
    INSERT_EVENT_QUERY,
    stream,
    streamId,
    fromVer + 1,
    DELETED_EVENT,
    nil,
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

const (
  CREATED_EVENT = "Created"
  UPDATED_EVENT = "Updated"
  DELETED_EVENT = "Deleted"
  
  CHECK_STREAM_EXISTS_QUERY = `
    SELECT EXISTS(
      SELECT * FROM events
      WHERE stream = ($1) && streamid = ($2)
    ) as exists`

  CHECK_STREAM_VERSION_EXISTS_QUERY = `
    SELECT EXISTS(
      SELECT * FROM events
      WHERE stream = ($1) && streamid = ($2) && streamversion = ($3)
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
