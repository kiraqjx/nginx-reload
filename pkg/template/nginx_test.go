package template

import (
	"fmt"
	"nginx-reload/pkg/entity"
	"nginx-reload/pkg/vo"
	"testing"
)

func TestTemplate(t *testing.T) {
	templateConfig := &vo.NginxTemplate{
		Name:   "nginx.conf",
		Header: "worker processes 8;\n\nevents{\n    worker_connections 1024;\n}\n\nstream{\n",
		Footer: "}",
	}
	template := NewNginxTemplate(*templateConfig)
	contents := []entity.ProxyConfig{{
		Id:         "id",
		Datasource: "127.0.0.1:3306",
		Port:       "3306",
	}}
	result := template.Template(contents)
	fmt.Println(result)
}
