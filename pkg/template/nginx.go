package template

import (
	"fmt"
	"strings"

	"github.com/kiraqjx/nginx-reload/pkg/entity"
	"github.com/kiraqjx/nginx-reload/pkg/vo"
)

type NginxTemplate struct {
	name   string
	header string
	footer string
}

func NewNginxTemplate(config vo.NginxTemplate) *NginxTemplate {
	return &NginxTemplate{
		name:   config.Name,
		header: config.Header,
		footer: config.Footer,
	}
}

func (n *NginxTemplate) Template(contents []string) string {
	var builder strings.Builder
	builder.WriteString(n.header)
	// contentsString := configsFillTemplate(contents)
	for _, content := range contents {
		builder.WriteString(content)
	}
	builder.WriteString(n.footer)
	return builder.String()
}

func (n *NginxTemplate) Start() *strings.Builder {
	var builder *strings.Builder
	builder.WriteString(n.header)
	return builder
}

func (n *NginxTemplate) AddContent(builder *strings.Builder, content entity.ProxyConfig) {
	contentString := configFillTemplate(content)
	builder.WriteString(contentString)
}

func (n *NginxTemplate) End(builder *strings.Builder) string {
	builder.WriteString(n.footer)
	return builder.String()
}

func (n *NginxTemplate) GetName() string {
	return n.name
}

const (
	template = "    upstream %s {\n        server %s;\n    }\n\n    server {\n        listen %s;\n        proxy_pass %s;\n    }\n"
)

func configFillTemplate(config entity.ProxyConfig) string {
	return fmt.Sprintf(template, config.Id, config.Datasource, config.Port, config.Id)
}

func ConfigsFillTemplate(configs []entity.ProxyConfig) []string {
	contents := make([]string, len(configs))
	for _, config := range configs {
		contents = append(contents, configFillTemplate(config))
	}
	return contents
}
