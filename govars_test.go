package govars

import "testing"

func TestLoading(t *testing.T) {

	configuration := GoVarsConfig{}
	configuration.Load("sample.ini");

	if configuration.GetString("stringVariable") != `This is a "string" variable` {
		t.Error(`Expected "This is a "string" variable", got:`, configuration.GetString("stringVariable"))
	}

	if configuration.GetString("emptyStringVariable") != `` {
		t.Error(`Expected "", got:`, configuration.GetString("emptyStringVariable"))
	}

	if configuration.GetInt("intVariable") != 42 {
		t.Error(`Expected "42", got:`, configuration.GetInt("intVariable"))
	}

	if configuration.GetFloat32("floatVariable") != 3.14 {
		t.Error(`Expected "3.14", got:`, configuration.GetFloat32("floatVariable"))
	}

	if configuration.GetFloat64("floatVariable") != 3.14 {
		t.Error(`Expected "3.14", got:`, configuration.GetFloat64("floatVariable"))
	}

}