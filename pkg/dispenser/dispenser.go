package dispenser

import (
	"log"
	"sync"

	"github.com/kiraqjx/nginx-reload/pkg/ssh"
	"github.com/kiraqjx/nginx-reload/pkg/store"
	"github.com/kiraqjx/nginx-reload/pkg/template"
	"github.com/kiraqjx/nginx-reload/pkg/vo"
)

type Dispenser struct {
	template *template.NginxTemplate
	targets  []*ssh.SshConnect
	store    store.Store
	lock     *sync.Mutex
	debug    bool
}

// new dispenser
func NewDispenser(store store.Store, templateConfig vo.NginxTemplate, sshConfigs []vo.SshConfig, debug bool) (*Dispenser, error) {
	sshList := make([]*ssh.SshConnect, 0, len(sshConfigs))
	for _, sshConfig := range sshConfigs {
		sshConn, err := ssh.NewSshConnect(sshConfig)
		if err != nil {
			return nil, err
		}
		sshList = append(sshList, sshConn)
	}
	nginxTem := template.NewNginxTemplate(templateConfig)
	return &Dispenser{
		template: nginxTem,
		targets:  sshList,
		store:    store,
		debug:    debug,
		lock:     new(sync.Mutex),
	}, nil
}

// dispense
func (d *Dispenser) Do() error {
	d.lock.Lock()
	defer d.lock.Unlock()

	// get nginx conf
	contents := d.store.AllConfig()
	s := d.template.Template(contents)

	if d.debug {
		log.Fatalln(s)
		return nil
	}

	// dispenser by ssh
	for _, target := range d.targets {
		err := target.UploadFile([]byte(s), d.template.GetName())
		if err == nil {
			target.Reload()
		}
	}
	return nil
}
