package apollo

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func writeFile(t *testing.T, filename string, content string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
}

func TestNewAgent(t *testing.T) {
	a, err := NewAgent("")
	if err != nil {
		t.Fatal(err)
	}
	defer a.Close()
}

func TestAgent_Get(t *testing.T) {
	// write the configuration file before new agent
	content := `# 降级
[degrade]
mysql=0
redis=1
tusd=0
	
[degrade.fileSharing]
push_server=2
key1=1
`
	writeFile(t, DefaultApolloIni, content)

	// new agent
	a, err := NewAgent(DefaultApolloIni)
	if err != nil {
		t.Fatal(err)
	}
	defer a.Close()

	// test init agent: first load
	assert.EqualValues(t, 0, a.Get("degrade.fileSharing", "mysql").MustInt(0))
	assert.EqualValues(t, 1, a.Get("degrade.fileSharing", "redis").MustInt(0))
	assert.EqualValues(t, 2, a.Get("degrade.fileSharing", "push_server").MustInt(0))
	assert.EqualValues(t, 1, a.Get("degrade.fileSharing", "key1").MustInt(0))

	// modify the configuration file
	content = `# 降级
[degrade]
mysql=0
redis=0
tusd=0
	
[degrade.fileSharing]
push_server=5
key1=10
key2=2
`
	writeFile(t, DefaultApolloIni, content)
	time.Sleep(1 * time.Second) // wait for reloading

	/// test agent: second load after modify
	assert.EqualValues(t, 0, a.Get("degrade.fileSharing", "mysql").MustInt(0))
	assert.EqualValues(t, 0, a.Get("degrade.fileSharing", "redis").MustInt(0))
	assert.EqualValues(t, 5, a.Get("degrade.fileSharing", "push_server").MustInt(0))
	assert.EqualValues(t, 10, a.Get("degrade.fileSharing", "key1").MustInt(0))
	assert.EqualValues(t, 2, a.Get("degrade.fileSharing", "key2").MustInt(0))
}
