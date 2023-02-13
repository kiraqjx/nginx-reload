package dispenser

import (
	"bytes"
	"nginx-reload/pkg/ssh"
	"nginx-reload/pkg/store"
	"nginx-reload/pkg/template"
	"nginx-reload/pkg/vo"
	"sync"
)

type Dispenser struct {
	template *template.NginxTemplate
	targets  []*ssh.SshConnect
	store    store.Store
	lock     *sync.Mutex
}

// new dispenser
func NewDispenser(store store.Store, templateConfig vo.NginxTemplate, sshConfigs []vo.SshConfig) (*Dispenser, error) {
	sshList := make([]*ssh.SshConnect, len(sshConfigs))
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
	}, nil
}

// dispense
func (d *Dispenser) Do() error {
	d.lock.Lock()
	defer d.lock.Unlock()

	// get nginx conf
	contents := d.store.AllConfig()
	s := d.template.Template(contents)

	// dispenser by ssh
	for _, target := range d.targets {
		err := target.UploadFile(bytes.NewBuffer([]byte(s)), d.template.GetName())
		if err == nil {
			target.Reload()
		}
	}
	return nil
}
