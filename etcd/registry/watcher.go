package registry

import (
	"context"
	registry2 "github.com/jageros/hawos/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	_ registry2.Watcher = &watcher{}
)

type watcher struct {
	key       string
	ctx       context.Context
	cancel    context.CancelFunc
	watchChan clientv3.WatchChan
	watcher   clientv3.Watcher
	kv        clientv3.KV
	ch        clientv3.WatchChan
}

func newWatcher(ctx context.Context, key string, client *clientv3.Client) *watcher {
	w := &watcher{
		key:     key,
		watcher: clientv3.NewWatcher(client),
		kv:      clientv3.NewKV(client),
	}
	w.ctx, w.cancel = context.WithCancel(ctx)
	w.watchChan = w.watcher.Watch(w.ctx, key, clientv3.WithPrefix(), clientv3.WithRev(0))
	w.watcher.RequestProgress(context.Background())
	return w
}

func (w *watcher) Next() ([]*registry2.ServiceInstance, error) {
	select {
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	case <-w.watchChan:
	}
	resp, err := w.kv.Get(w.ctx, w.key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var items []*registry2.ServiceInstance
	for _, kv := range resp.Kvs {
		si, err := unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		items = append(items, si)
	}
	return items, nil
}

func (w *watcher) Stop() error {
	w.cancel()
	return w.watcher.Close()
}
