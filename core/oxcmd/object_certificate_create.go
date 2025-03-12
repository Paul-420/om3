package oxcmd

import (
	"github.com/opensvc/om3/core/objectaction"
)

type (
	CmdSecGenCert struct {
		OptsGlobal
	}
)

func (t *CmdSecGenCert) Run(selector, kind string) error {
	mergedSelector := mergeSelector(selector, t.ObjectSelector, kind, "")
	return objectaction.New(
		objectaction.WithColor(t.Color),
		objectaction.WithOutput(t.Output),
		objectaction.WithObjectSelector(mergedSelector),
	).Do()
}
