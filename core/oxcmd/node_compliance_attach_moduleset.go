package oxcmd

import (
	"github.com/opensvc/om3/core/nodeaction"
)

type (
	CmdNodeComplianceAttachModuleset struct {
		OptsGlobal
		Moduleset    string
		NodeSelector string
	}
)

func (t *CmdNodeComplianceAttachModuleset) Run() error {
	return nodeaction.New(
		nodeaction.WithRemoteNodes(t.NodeSelector),
		nodeaction.WithFormat(t.Output),
		nodeaction.WithColor(t.Color),
	).Do()
}
