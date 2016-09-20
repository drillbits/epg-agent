package agent

import "testing"

func TestReadConfigPath(t *testing.T) {
	conf, err := ReadConfigPath("./epg-agent.toml.sample")
	if err != nil {
		t.Fatal(err)
	}

	for _, src := range conf.Sources {
		t.Logf("src: %s path: %s", src.Name, src.Path)
	}

	for _, dst := range conf.Dests {
		t.Logf("dst: %s type: %s", dst.Name, dst.Type)
		t.Logf(" db: %#v", dst.Database)
	}
}
