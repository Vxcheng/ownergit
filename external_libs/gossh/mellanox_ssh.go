package gossh

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"regexp"
	"strings"
	"time"
	"unicode"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var regAnsi = regexp.MustCompile(ansi)

type MellanoxSSH struct {
	enable       bool
	newline      string
	client       *ssh.Client
	session      *ssh.Session
	writer       io.Writer
	scanner      *bufio.Scanner
	cmdDelay     time.Duration
	regDelimiter *regexp.Regexp
}

func NewMellanoxSSH(newlineDelimiter string) *MellanoxSSH {
	return &MellanoxSSH{
		enable:       false,
		newline:      newlineDelimiter,
		cmdDelay:     time.Millisecond * 10,
		regDelimiter: regexp.MustCompile(`([>#]\s*\r*)$`),
	}
}

func (m *MellanoxSSH) Close() error {
	m.closeSession()
	if m.client != nil {
		m.client.Close()
	}
	return nil
}

func (m *MellanoxSSH) closeSession() error {
	if m.session != nil {
		m.session.Close()
	}
	return nil
}

func (m *MellanoxSSH) CmdDelay(duration time.Duration) {
	m.cmdDelay = duration
}

func (m *MellanoxSSH) genConfig(user string, password string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(_, _ string, questions []string, _ []bool) (answers []string, err error) {
				answers = make([]string, len(questions))
				for n, _ := range questions {
					answers[n] = password
				}
				return answers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

func (m *MellanoxSSH) pipes(session *ssh.Session) (reader io.Reader, writer io.Writer, err error) {
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

func (m *MellanoxSSH) initSession(client *ssh.Client) error {
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
func (m *MellanoxSSH) Dial(user string, password string, host string, port int64) error {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), m.genConfig(user, password))
	if err != nil {
		return err
	}

	if err := m.initSession(client); err != nil {
		client.Close()
		return err
	}

	if _, err := m.runCmd("", false); err != nil {
		m.closeSession()
		client.Close()
		return err
	}

	m.client = client
	return nil
}

func (m *MellanoxSSH) Enable() error {
	if m.enable {
		return nil
	}

	if _, err := m.runCmd("enable", false); err != nil {
		return err
	}
	m.enable = true
	return nil
}

func (m *MellanoxSSH) Disable() error {
	if !m.enable {
		return nil
	}

	if _, err := m.runCmd("disable", false); err != nil {
		return err
	}
	m.enable = false
	return nil
}

func (m *MellanoxSSH) runCmd(cmd string, ignorePrefixPromote bool) ([]string, error) {
	if err := m.sendCmd(cmd); err != nil {
		return nil, err
	}
	return m.readDelimiter(ignorePrefixPromote)
}

func (m *MellanoxSSH) RunCmd(cmd string, stripCmd ...bool) ([]string, error) {
	lines, err := m.runCmd(cmd, true)
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

func (m *MellanoxSSH) sendCmd(cmd string) error {
	cmd = strings.TrimSpace(cmd)
	if _, err := m.writer.Write([]byte(fmt.Sprintf("%s\n", cmd))); err != nil {
		return err
	}

	<-time.After(m.cmdDelay)
	if _, err := m.writer.Write([]byte(m.newline)); err != nil {
		return err
	}
	return nil
}

func (m *MellanoxSSH) readDelimiter(ignorePrefixPromote bool) ([]string, error) {
	var result []string
	for m.scanner.Scan() {
		buf := stripAnsiEscapeCodes(m.scanner.Bytes())
		line := strings.TrimLeft(string(buf), "\r")
		line = strings.TrimRightFunc(line, unicode.IsSpace)
		if m.regDelimiter.MatchString(line) {
			if !ignorePrefixPromote || len(result) != 0 {
				break
			} else {
				continue
			}
		}
		result = append(result, line)
	}
	return result, nil
}

func stripAnsiEscapeCodes(buf []byte) []byte {
	return regAnsi.ReplaceAll(buf, nil)
}
