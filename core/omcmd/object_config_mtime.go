package omcmd

import (
	"context"
	"fmt"

	"github.com/opensvc/om3/core/commoncmd"
	"github.com/opensvc/om3/core/naming"
	"github.com/opensvc/om3/core/object"
	"github.com/opensvc/om3/core/objectaction"
	"github.com/opensvc/om3/util/timestamp"
)

type (
	CmdObjectConfigMtime struct {
		OptsGlobal
	}
)

func (t *CmdObjectConfigMtime) Run(selector, kind string) error {
	mergedSelector := commoncmd.MergeSelector(selector, t.ObjectSelector, kind, "")
	return objectaction.New(
		objectaction.LocalFirst(),
		objectaction.WithObjectSelector(mergedSelector),
		objectaction.WithLocal(t.Local),
		objectaction.WithOutput(t.Output),
		objectaction.WithColor(t.Color),
		objectaction.WithLocalFunc(func(ctx context.Context, p naming.Path) (interface{}, error) {
			o, err := object.New(p)
			if err != nil {
				return nil, err
			}
			c, ok := o.(object.Configurer)
			if !ok {
				return nil, fmt.Errorf("%s is not a configurer", o)
			}
			tm := c.Config().ModTime()
			return timestamp.New(tm).String(), nil
		}),
	).Do()
}
