/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    dscv
 * @Date:    2021/7/9 3:12 下午
 * @package: dscv
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package selector

import (
	"context"
	"fmt"
	"github.com/jageros/hawos/log"
	protoc2 "github.com/jageros/hawos/protoc"
	registry2 "github.com/jageros/hawos/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

const Prefix = "/pbserver"

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	id        string
	namespace string
	name      string
	ttl       time.Duration
	ctx       context.Context
	cancel    context.CancelFunc
}

// Context with registry context.
func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// Namespace with registry namespance.
func Namespace(ns string) Option {
	return func(o *options) { o.namespace = ns }
}

func Name(nm string) Option {
	return func(o *options) {
		o.name = nm
	}
}

// RegisterTTL with register ttl.
func RegisterTTL(ttl time.Duration) Option {
	return func(o *options) { o.ttl = ttl }
}

// Registry is etcd registry.
type Registry struct {
	opts   *options
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease

	// watcher
	watchChan clientv3.WatchChan
	watcher   clientv3.Watcher
}

func (r *Registry) Init(id, name string) error {
	r.opts.id = id
	r.opts.name = name
	return nil
}

func (r *Registry) Serve() error {
	for {
		select {
		case <-r.opts.ctx.Done():
			return r.opts.ctx.Err()

		case <-r.watchChan:
			resp, err := r.kv.Get(r.opts.ctx, r.opts.namespace, clientv3.WithPrefix())
			if err != nil {
				return err
			}
			var items []*metaData
			for _, kv := range resp.Kvs {
				si, err := unmarshal(kv.Value)
				if err != nil {
					return err
				}
				items = append(items, si)
			}
			Update(items)
		}
	}
}

func (r *Registry) Stop() {
	r.opts.cancel()
	r.watcher.Close()
}

func (r *Registry) Info() string {
	return fmt.Sprintf("ServerName=%s ServerType=%s Namespace=%s", r.opts.name, "selector", r.opts.namespace)
}

// New creates etcd registry
func New(ctx context.Context, client *clientv3.Client, opts ...Option) (r *Registry) {
	ctx_, cancel := context.WithCancel(ctx)
	options := &options{
		ctx:       ctx_,
		cancel:    cancel,
		namespace: Prefix,
		ttl:       time.Second * 15,
	}
	for _, o := range opts {
		o(options)
	}
	r = &Registry{
		opts:    options,
		client:  client,
		kv:      clientv3.NewKV(client),
		watcher: clientv3.NewWatcher(client),
	}

	r.watchChan = r.watcher.Watch(r.opts.ctx, r.opts.namespace, clientv3.WithPrefix(), clientv3.WithRev(0))
	r.watcher.RequestProgress(context.Background())
	return
}

// Register the registration.
func (r *Registry) Register(_ registry2.Registrar) error {
	msgIds := protoc2.MsgIDs()
	if len(msgIds) <= 0 {
		return nil
	}
	md := &metaData{
		ID:     r.opts.id,
		Name:   r.opts.name,
		MsgIds: msgIds,
	}
	key := fmt.Sprintf("%s/%s/%s", r.opts.namespace, r.opts.name, md.ID)
	value, err := marshal(md)
	if err != nil {
		return err
	}
	if r.lease != nil {
		r.lease.Close()
	}
	r.lease = clientv3.NewLease(r.client)
	grant, err := r.lease.Grant(r.opts.ctx, int64(r.opts.ttl.Seconds()))
	if err != nil {
		return err
	}
	_, err = r.client.Put(r.opts.ctx, key, value, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}
	hb, err := r.client.KeepAlive(r.opts.ctx, grant.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case _, ok := <-hb:
				if !ok {
					return
				}
			case <-r.opts.ctx.Done():
				return
			}
		}
	}()
	log.Infof("Register Service Namespace=%s Name=%s ID=%s MsgIds=%+v", r.opts.namespace, md.Name, md.ID, md.MsgIds)
	return nil
}

// Deregister the registration.
func (r *Registry) Deregister() {
	defer func() {
		if r.lease != nil {
			r.lease.Close()
		}
	}()
	key := fmt.Sprintf("%s/%s/%s", r.opts.namespace, r.opts.name, r.opts.id)
	_, err := r.client.Delete(r.opts.ctx, key)
	if err != nil {
		log.Errorf("pdserver Deregister err: %v", err)
	} else {
		log.Infof("Deregister service Namespace=%s ServiceName=%s ID=%s", r.opts.namespace, r.opts.name, r.opts.id)
	}
}

// GetService return the service instances in memory according to the service name.
func (r *Registry) GetService() ([]*metaData, error) {
	key := r.opts.namespace
	resp, err := r.kv.Get(r.opts.ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var items []*metaData
	for _, kv := range resp.Kvs {
		si, err := unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		items = append(items, si)
	}
	return items, nil
}
