package ssh

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"regexp"
	"strings"
	"time"
	"unicode"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var regAnsi = regexp.MustCompile(ansi)

type NationalSSH struct {
	enable       bool
	newline      string
	client       *ssh.Client
	session      *ssh.Session
	writer       io.WriteCloser
	scanner      *bufio.Scanner
	cmdDelay     time.Duration
	regDelimiter *regexp.Regexp
	ignoreLineRe *regexp.Regexp
}

func NewNationalSSH(newlineDelimiter string) *NationalSSH {
	return &NationalSSH{
		enable:       false,
		newline:      newlineDelimiter,
		cmdDelay:     time.Millisecond * 10,
		regDelimiter: regexp.MustCompile(`(<.+>|\[.+\])$`),
		ignoreLineRe: regexp.MustCompile(`^\*`), // ^\*|Ctrl\+Z\.$
	}
}

func (m *NationalSSH) Close() error {
	m.closeSession()
	if m.client != nil {
		m.client.Close()
	}
	return nil
}

func (m *NationalSSH) closeSession() error {
	if m.session != nil {
		m.session.Close()
	}
	return nil
}

func (m *NationalSSH) CmdDelay(duration time.Duration) {
	m.cmdDelay = duration
}

func (m *NationalSSH) genConfig(user string, password string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 20 * time.Second,
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com",
				"arcfour256", "arcfour128", "aes128-cbc", "aes256-cbc", "3des-cbc", "des-cbc",
			},
		},
	}
	return config
}

func (m *NationalSSH) pipes(session *ssh.Session) (reader io.Reader, writer io.WriteCloser, err error) {
	writer, err = session.StdinPipe()
	if err != nil {
		return
	}

	reader, err = session.StdoutPipe()
	if err != nil {
		return
	}
	return
}

func (m *NationalSSH) initSession(client *ssh.Client) error {
	session, err := client.NewSession()
	if err != nil {
		return err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("vt100", 2048, 80, modes); err != nil {
		session.Close()
		return err
	}

	reader, writer, err := m.pipes(session)
	if err != nil {
		session.Close()
		return err
	}

	if err := session.Shell(); err != nil {
		session.Close()
		return err
	}

	m.session = session
	m.writer = writer
	m.scanner = bufio.NewScanner(reader)
	m.scanner.Split(bufio.ScanLines)

	return nil
}

func (m *NationalSSH) Dial(user string, password string, host string, port int64) error {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), m.genConfig(user, password))
	if err != nil {
		return err
	}

	if err := m.initSession(client); err != nil {
		client.Close()
		return err
	}
	//m.firstLogin = true

	if _, err := m.runCmd("screen-length disable"); err != nil {
		m.closeSession()
		client.Close()
		return err
	}

	m.client = client
	return nil
}

func (m *NationalSSH) Enable() error {

	return nil
}

func (m *NationalSSH) Disable() error {

	return nil
}

func (m *NationalSSH) runCmd(cmd string) ([]string, error) {
	if err := m.sendCmd(cmd); err != nil {
		return nil, err
	}
	return m.readDelimiter()
}

func (m *NationalSSH) RunCmd(cmd string, stripCmd ...bool) ([]string, error) {
	lines, err := m.runCmd(cmd)
	if err != nil {
		return nil, err
	}

	if len(stripCmd) > 0 && stripCmd[0] {
		for i, _ := range lines {
			if strings.Contains(lines[i], cmd) {
				lines = append(lines[:i], lines[i+1:]...)
				break
			}
		}
	}

	return lines, nil
}

func (m *NationalSSH) sendCmd(cmd string) error {
	cmd = strings.TrimSpace(cmd)
	cmd = fmt.Sprintf("%s\n\n", cmd)

	if _, err := m.writer.Write([]byte(cmd)); err != nil {
		return err
	}

	//<-time.After(m.cmdDelay)
	//if _, err := m.writer.Write([]byte(m.newline)); err != nil {
	//	return err
	//}
	return nil
}

func (m *NationalSSH) readDelimiter() ([]string, error) {
	var result []string
	for m.scanner.Scan() {
		//buf := stripAnsiEscapeCodes(m.scanner.Bytes())
		buf := m.scanner.Bytes()
		line := strings.Trim(string(buf), "\r")
		line = strings.TrimRightFunc(line, unicode.IsSpace)
		if m.ignoreLineRe.MatchString(line) {
			continue
		}

		if m.regDelimiter.MatchString(line) {
			result = append(result, line)
			break
		}
		result = append(result, line)
	}
	return result, nil
}

func stripAnsiEscapeCodes(buf []byte) []byte {
	return regAnsi.ReplaceAll(buf, nil)
}

func (m *NationalSSH) RunCmds(ignoreErr bool, cmds ...string) (out []string, err error) {
	for _, cmd := range cmds {
		values, err := m.RunCmd(cmd, false)

		if err != nil && !ignoreErr {
			return out, err
		}

		out = append(out, values...)
	}
	return
}
