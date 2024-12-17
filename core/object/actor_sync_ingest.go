package object

import (
	"context"

	"github.com/opensvc/om3/core/actioncontext"
	"github.com/opensvc/om3/core/resource"
)

func (t *actor) SyncIngest(ctx context.Context) error {
	ctx = actioncontext.WithProps(ctx, actioncontext.SyncIngest)
	if err := t.validateAction(); err != nil {
		return err
	}
	t.setenv("sync_ingest", false)
	unlock, err := t.lockAction(ctx)
	if err != nil {
		return err
	}
	defer unlock()
	return t.lockedSyncIngest(ctx)
}

func (t *actor) lockedSyncIngest(ctx context.Context) error {
	if err := t.masterSyncIngest(ctx); err != nil {
		return err
	}
	if err := t.slaveSyncIngest(ctx); err != nil {
		return err
	}
	return nil
}

func (t *actor) masterSyncIngest(ctx context.Context) error {
	return t.action(ctx, func(ctx context.Context, r resource.Driver) error {
		t.log.Attr("rid", r.RID()).Debugf("ingest resource data")
		return resource.Ingest(ctx, r)
	})
}

func (t *actor) slaveSyncIngest(ctx context.Context) error {
	return nil
}
