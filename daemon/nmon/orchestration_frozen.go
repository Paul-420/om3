package nmon

import "github.com/opensvc/om3/core/node"

func (t *Manager) orchestrateFrozen() {
	switch t.state.State {
	case node.MonitorStateIdle:
		t.frozenFromIdle()
	case node.MonitorStateFreezeSuccess:
		t.frozenFromFrozen()
	}
}

func (t *Manager) frozenFromIdle() {
	if t.frozenClearIfReached() {
		return
	}
	t.log.Infof("run action freeze")
	t.doTransitionAction(t.crmFreeze, node.MonitorStateFreezeProgress, node.MonitorStateFreezeSuccess, node.MonitorStateFreezeFailure)
	return
}

func (t *Manager) frozenFromFrozen() {
	if t.frozenClearIfReached() {
		t.state.State = node.MonitorStateIdle
		t.change = true
		return
	}
	return
}

func (t *Manager) frozenClearIfReached() bool {
	if nodeStatus := node.StatusData.GetByNode(t.localhost); nodeStatus != nil && !nodeStatus.FrozenAt.IsZero() {
		t.log.Infof("instance state is frozen, unset global expect")
		t.change = true
		t.state.GlobalExpect = node.MonitorGlobalExpectNone
		t.clearPending()
		return true
	}
	return false
}
