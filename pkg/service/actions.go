package service

import (
	"fmt"
	"log"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/repository"
	"golang.org/x/crypto/ssh"
)

type ActionService struct {
	repo *repository.Repository
}

func NewActionService(repo *repository.Repository) *ActionService {
	return &ActionService{
		repo: repo,
	}
}

func (s *ActionService) KillBase(sshConfig *ent.SSH) (bool, error) {

	config := &ssh.ClientConfig{
		User: sshConfig.SSHUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshConfig.SSHPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to the SSH server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshConfig.Server, 22), config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	} else {
		fmt.Println("SUCCESS")
	}
	defer client.Close()

	// Run the first command
	command := "sudo service postgresql stop"
	output, err := s.RunCommand(client, command)
	if err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}
	fmt.Printf("Command 1 output:\n%s\n", string(output))

	return true, nil
}
func (s *ActionService) StopBase(tableName string) (bool, error) {
	return s.StopBase(tableName)
}
func (s *ActionService) RunBase(tableName string) (bool, error) {
	return s.RunBase(tableName)
}
func (s *ActionService) RestartBase(sshConfig *ent.SSH) error {

	err := s.repo.LastBaseCheckpoint()
	if err != nil {
		return err
	}

	// Create an SSH client configuration
	config := &ssh.ClientConfig{
		User: sshConfig.SSHUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshConfig.SSHPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshConfig.Server, 22), config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	} else {
		fmt.Println("SUCCESS")
	}
	defer client.Close()

	// Run the first command
	command := "sudo service postgresql stop"
	output, err := s.RunCommand(client, command)
	if err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}
	fmt.Printf("Command 1 output:\n%s\n", string(output))
	command = "sudo service postgresql start"
	output, err = s.RunCommand(client, command)
	if err != nil {
		log.Fatalf("Failed to restart server: %v", err)
	}
	fmt.Printf("Command 1 output:\n%s\n", string(output))

	return nil
}
func (s *ActionService) KillProcess(pid int64) (bool, error){
	return s.repo.KillProcess(pid)
}
func (s *ActionService) RunCommand(client *ssh.Client, command string) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.CombinedOutput(command)
}
