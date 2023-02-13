package ssh

import (
	"fmt"
	"io"
	"log"
	"net"
	"nginx-reload/pkg/vo"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SshConnect struct {
	target      string
	reloadShell string
	config      vo.SshConfig
	client      *ssh.Client
}

// new ssh connect
func NewSshConnect(config vo.SshConfig) (*SshConnect, error) {
	sshConfig := ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{ssh.Password(config.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	sshClient, err := ssh.Dial("tcp", addr, &sshConfig)
	if err != nil {
		log.Fatalln("err connect ssh: ", err)
		return nil, err
	}
	return &SshConnect{
		target:      config.TargetPath,
		reloadShell: config.NginxPath + " -s reload",
		config:      config,
		client:      sshClient,
	}, nil
}

// exec shell script
func (s *SshConnect) Reload() (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		log.Fatalln("create new session error: ", err)
		return "", err
	}
	defer session.Close()
	output, err := session.CombinedOutput(s.reloadShell)
	if err != nil {
		log.Fatalln("exec shell script error: ", err)
		return "", err
	}
	return string(output), nil
}

// upload file from local path to target path
func (s *SshConnect) UploadFile(local io.Writer, fileName string) error {
	sftpClient, err := sftp.NewClient(s.client)
	if err != nil {
		log.Fatalln("create sftp client error: ", err)
		return err
	}
	dstFile, err := sftpClient.Create(path.Join(s.target, fileName))
	if err != nil {
		log.Fatalln("create file transfer channel error: ", err)
		return err
	}
	defer dstFile.Close()
	_, err = dstFile.WriteTo(local)
	if err != nil {
		log.Fatalln("upload file error: ", err)
		return err
	}
	return nil
}
