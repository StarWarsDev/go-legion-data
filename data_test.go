package legiondata

import (
	"testing"
)

func TestGetData(t *testing.T) {
	data, err := GetData()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(data.Keywords) < 1 {
		t.Error("expected more than 0 keywords")
	}

	if len(data.Units) < 1 {
		t.Error("expected more than 0 units")
	}

	if len(data.Upgrades) < 1 {
		t.Error("expected more than 0 upgrades")
	}

	if len(data.CommandCards) < 1 {
		t.Error("expected more than 0 command cards")
	}

	if len(data.Sources) < 1 {
		t.Error("expected more than 0 sources")
	}
}
