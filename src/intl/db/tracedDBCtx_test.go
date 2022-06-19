package db

import (
	"testing"

	"github.com/betalixt/eventSourceUsers/intl/trace"
	"github.com/betalixt/eventSourceUsers/optn"
	"github.com/betalixt/eventSourceUsers/util/config"
	"github.com/betalixt/eventSourceUsers/util/logger"
	"github.com/spf13/viper"
)

func TestNewTracedDBCtx(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("failed with %v", err)
			t.FailNow()
		}
	}()
  lgr := logger.NewLogger()
	berr := config.InitializeConfig(lgr, "../../cfg")
	if berr != nil {
		t.Errorf("failed to create config")
		t.FailNow()
	}
	cfg := viper.GetViper()
	db := NewDatabase(optn.NewDatabaseOptions(cfg))
	tracer := trace.NewZapTracer(lgr)
  tdb := NewTracedDBContext(db, tracer, "test")
  
  chck := ExistsEntity{}
  err := tdb.Get(&chck, CheckTimestampProceduresExist)
  if err != nil {
    t.Errorf("failed with %v", err)
    t.FailNow()
  }
}
