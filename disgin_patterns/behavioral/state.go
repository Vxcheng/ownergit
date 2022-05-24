// 状态模式
// 这是一个工作流的例子，在企业内部或者是学校我们经常会看到很多审批流程
// 假设我们有一个报销的流程: 员工提交报销申请 -> 直属部门领导审批 -> 财务审批 -> 结束
// 在这个审批流中，处在不同的环节就是不同的状态
// 而流程的审批、驳回就是不同的事件
package behavioral

import "fmt"

// Machine 状态机
type Machine struct {
	state IState
}

// SetState 更新状态
func (m *Machine) SetState(state IState) {
	m.state = state
}

// GetStateName 获取当前状态
func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// IState 状态
type IState interface {
	// 审批通过
	Approval(m *Machine)
	// 驳回
	Reject(m *Machine)
	// 获取当前状态名称
	GetName() string
}

// leaderApproveState 直属领导审批
type leaderApproveState struct{}

// Approval 获取状态名字
func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}

// GetName 获取状态名字
func (leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

// Reject 获取状态名字
func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

// financeApproveState 财务审批
type financeApproveState struct{}

// Approval 审批通过
func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}

// 拒绝
func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

// GetName 获取名字
func (f financeApproveState) GetName() string {
	return "FinanceApproveState"
}

// GetFinanceApproveState GetFinanceApproveState
func GetFinanceApproveState() IState {
	return &financeApproveState{}
}

type state interface {
	doAction(c *stateContext)
	getValue() string
}

const (
	start = "start"
	stop  = "stop"
)

type startState struct{}

func (s *startState) doAction(c *stateContext) {
	fmt.Println("now start do something")
	c.setState(s)
}

func (s *startState) getValue() string {
	return start
}

type stopState struct{}

func (s *stopState) doAction(c *stateContext) {
	fmt.Println("now stop do something")
	c.setState(s)
}

func (s *stopState) getValue() string {
	return stop
}

type stateContext struct {
	state
}

func (c *stateContext) getState() string {
	return c.state.getValue()
}

func (c *stateContext) setState(state state) {
	c.state = state
}
