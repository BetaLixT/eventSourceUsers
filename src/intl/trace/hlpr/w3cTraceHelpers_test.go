package hlpr

import "testing"

func TestParseTraceparentRaw(t *testing.T) {
  validtrcp := "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
  invldtrcp := []string {
    "00-0ax7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01",
    "003-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01",
    "00-0a7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01",
  }
      

  vers, tid, pid, flg, err := ParseTraceparentRaw(validtrcp)
  if err != nil {
    t.Errorf("failed %v", err)
  } else {
    if len(vers) != 1 {
      t.Errorf("invalid version length")
    }
    if len(tid) != 16 {
      t.Errorf("invalid traceid length")
    }
    if len(pid) != 8 {
      t.Errorf("invalid parentId length")
    }
    if len(flg) != 1 {
      t.Errorf("invalid flag length")
    } 
  }
  for idx, trcp := range invldtrcp {
    _, _, _, _, err = ParseTraceparentRaw(trcp)
    if err == nil {
      t.Errorf("Should fail to parse invalid traceparent %d", idx)
    }
  }
}

func TestParseTraceparent(t *testing.T) {
  validtrcp := "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
  invldtrcp := []string {
    "00-0ax7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01",
    "003-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01",
    "00-0a7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01",
  }
      

  vers, tid, pid, flg, err := ParseTraceparent(validtrcp)
  if err != nil {
    t.Errorf("failed %v", err)
  } else {
    if len(vers) != 1 {
      t.Errorf("invalid version length")
    }
    if len(tid) != 16 {
      t.Errorf("invalid traceid length")
    }
    if len(pid) != 8 {
      t.Errorf("invalid parentId length")
    }
    if len(flg) != 1 {
      t.Errorf("invalid flag length")
    } 
  }
  for idx, trcp := range invldtrcp {
    _, _, _, _, err = ParseTraceparent(trcp)
    if err == nil {
      t.Errorf("Should fail to parse invalid traceparent %d", idx)
    }
  }
}

func TestGenerateTrace(t *testing.T) {
  traceparent, err := GenerateNewTraceparent(true)
  if err != nil {
    t.Errorf("failed: %v", err)
    t.FailNow()
  }
  _, _, _, _, err = ParseTraceparentRaw(traceparent)
  if err != nil {
    t.Errorf("failed: %v", err)
    t.FailNow()
  }
}
