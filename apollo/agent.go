package apollo

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"

	"github.com/fsnotify/fsnotify"
)

const DefaultApolloIni = "./apollo/apollo.ini"

type Agent struct {
	config  *ini.File
	file    string
	watcher *fsnotify.Watcher
}

func NewAgent(file string) (*Agent, error) {
	if file == "" {
		file = DefaultApolloIni
	}
	dir := filepath.Dir(file)
	_, err := os.Lstat(file)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil && !os.IsExist(err) {
			return nil, err
		}
		if f, err := os.Create(file); err != nil {
			return nil, err
		} else {
			f.Close()
		}
	}
	cfg, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	// Watch the directory, not the file itself.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	if err = watcher.Add(dir); err != nil {
		return nil, err
	}
	a := &Agent{
		config:  cfg,
		file:    file,
		watcher: watcher,
	}
	go a.watchLoop()
	return a, nil
}

func (a *Agent) watchLoop() {
	// Start listening for events.
	for {
		select {
		case event, ok := <-a.watcher.Events:
			if !ok {
				return
			}
			//log.Println("event:", event)

			if filepath.Dir(event.Name) != filepath.Dir(a.file) || filepath.Base(event.Name) != filepath.Base(a.file) {
				log.Println("continue modified file:", event.Name, a.file)
				continue
			}

			if event.Has(fsnotify.Write) {
				log.Println("modified file:", event)
				cfg, err := ini.Load(a.file)
				if err != nil {
					log.Println("load ini file err:", err)
					continue
				}
				a.config = cfg
			}
		case err, ok := <-a.watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func (a *Agent) Close() {
	a.watcher.Close()
}

func (a *Agent) Get(section, key string) *ini.Key {
	return a.config.Section(section).Key(key)
}

// ----------------------------package functions-------------------------------

var agent *Agent // default agent

// InitAgent initializes the default apollo agent with the given file
func InitAgent(file string) error {
	return nil
}

// GetKey get apollo key with section
func GetKey(section, key string) *ini.Key {
	return agent.Get(section, key)
}
