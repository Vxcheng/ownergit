package behavioral

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachine_GetStateName(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Reject()
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Approval()
}

func TestState(t *testing.T) {
	tests := []struct {
		name  string
		state state
		want  string
	}{
		{
			name:  start,
			state: new(startState),
			want:  start,
		},
		{
			name:  stop,
			state: new(stopState),
			want:  stop,
		},
	}

	c := new(stateContext)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.state.doAction(c)
			got := c.getState()
			if got != tt.want {
				t.Errorf("got is %s, want is %s", got, tt.want)
			}
			fmt.Printf("current state is '%s'\n", got)
		})
	}
}
