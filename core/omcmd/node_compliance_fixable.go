package omcmd

import (
	"github.com/opensvc/om3/core/nodeaction"
	"github.com/opensvc/om3/core/object"
)

type (
	CmdNodeComplianceFixable struct {
		OptsGlobal
		Moduleset    string
		Module       string
		NodeSelector string
		Force        bool
		Attach       bool
	}
)

func (t *CmdNodeComplianceFixable) Run() error {
	return nodeaction.New(
		nodeaction.WithFormat(t.Output),
		nodeaction.WithColor(t.Color),
		nodeaction.WithLocalFunc(func() (interface{}, error) {
			n, err := object.NewNode()
			if err != nil {
				return nil, err
			}
			comp, err := n.NewCompliance()
			if err != nil {
				return nil, err
			}
			run := comp.NewRun()
			run.SetModulesetsExpr(t.Moduleset)
			run.SetModulesExpr(t.Module)
			run.SetForce(t.Force)
			run.SetAttach(t.Attach)
			err = run.Fixable()
			return run, err
		}),
	).Do()
}
