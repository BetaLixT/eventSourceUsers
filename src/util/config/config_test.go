package config

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
)

type OptionTest struct {
	SomeConfig  string
	OtherConfig int
	ThirdValue  string
	FourthValue string
	FifthValue  string
}

func TestInitializeConfig(t *testing.T) {
	os.Setenv("PORT", "43623")
	os.Setenv("BLOGGO_SIMPLE", "value00")
	os.Setenv("BLOGGO_OPTIONTEST__SOMECONFIG", "value01")
	os.Setenv("BLOGGO_OPTIONTEST__OTHERCONFIG", "435")
	os.Setenv("BLOGGO_OPTIONTEST__THIRDVALUE", "value02")
	os.Setenv("BLOGGO_OPTIONTEST__FOURTHVALUE", "value03")
	os.Setenv("BLOGGO_OPTIONTEST__FIFTHVALUE", "value04")
	lgr, err := zap.NewProduction()
	if err != nil {
		fmt.Println("failed to create logger.... why...?")
		t.FailNow()
	}
	cfg := NewConfig(lgr)
	if cfg.GetString("PORT") != "43623" {
		fmt.Println("port value is invalid")
		t.Fail()
	}
	if cfg.GetString("Simple") != "value00" {
		fmt.Println("value of simple is invalid")
		t.Fail()
	}
	opt := OptionTest{
		SomeConfig: cfg.GetString("OptionTest.SomeConfig"),
		OtherConfig: cfg.GetInt("OptionTest.OtherConfig"),
		ThirdValue: cfg.GetString("OptionTest.ThirdValue"),
		FourthValue: cfg.GetString("OptionTest.FourthValue"),
		FifthValue: cfg.GetString("OptionTest.FifthValue"),	
	}
	
	if opt.SomeConfig != "value01" {
		fmt.Printf("value SomeConfig is invalid %s\n", opt.SomeConfig)
		t.Fail()
	}
	if opt.OtherConfig != 435 {
		fmt.Printf("value OtherConfig is invalid %d\n", opt.OtherConfig)
		t.Fail()
	}
	if opt.ThirdValue != "value02" {
		fmt.Printf("value ThirdValue is invalid %s\n", opt.ThirdValue)
		t.Fail()
	}
	if opt.FourthValue != "value03" {
		fmt.Printf("value FourthValue is invalid %s\n", opt.FourthValue)
		t.Fail()
	}
	if opt.FifthValue!= "value04" {
		fmt.Printf("value FifthValue is invalid %s\n", opt.FifthValue)
		t.Fail()
	}
	for _, key := range cfg.AllKeys() {
		fmt.Println(key)
	}
}
