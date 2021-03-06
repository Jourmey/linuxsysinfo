package linuxsysinfo

import (
	"encoding/json"
	"testing"
)

func TestCreatCPUInfo(t *testing.T) {
	i, err := CreatCPUInfo()
	if err != nil {
		t.Fatal(err)
	}
	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatMemInfo(t *testing.T) {
	i, err := CreatMemInfo()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatIfConfigInfos(t *testing.T) {
	i, err := CreatIfConfigInfos()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatVersionInfo(t *testing.T) {
	i, err := CreatVersionInfo()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatDfInfo(t *testing.T) {
	i, err := CreatDfInfo()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestCreatNetInfo(t *testing.T) {
	i, err := CreatNetInfo()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ii))
}

func TestAll(t *testing.T) {
	type info struct {
		CPUInfo *CPUInfo
		MemInfo *MemInfo
		//IfConfigInfos []IfConfigInfo
		VersionInfo *VersionInfo
		DfInfos     []DfInfo
		NetInfos    []NetInfo
	}

	i := new(info)
	var err error
	i.CPUInfo, err = CreatCPUInfo()
	if err != nil {
		t.Fatal(err)
	}
	i.MemInfo, err = CreatMemInfo()
	if err != nil {
		t.Fatal(err)
	}
	//i.IfConfigInfos, err = CreatIfConfigInfos()
	//if err != nil {
	//	t.Fatal(err)
	//}
	i.VersionInfo, err = CreatVersionInfo()
	if err != nil {
		t.Fatal(err)
	}
	i.DfInfos, err = CreatDfInfo()
	if err != nil {
		t.Fatal(err)
	}
	i.NetInfos, err = CreatNetInfo()
	if err != nil {
		t.Fatal(err)
	}

	ii, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(ii))
}
