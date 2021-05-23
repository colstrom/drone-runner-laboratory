/*
Copyright 2021 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type UserHandler func(string, *v3.User) (*v3.User, error)

type UserController interface {
	generic.ControllerMeta
	UserClient

	OnChange(ctx context.Context, name string, sync UserHandler)
	OnRemove(ctx context.Context, name string, sync UserHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() UserCache
}

type UserClient interface {
	Create(*v3.User) (*v3.User, error)
	Update(*v3.User) (*v3.User, error)
	UpdateStatus(*v3.User) (*v3.User, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.User, error)
	List(opts metav1.ListOptions) (*v3.UserList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.User, err error)
}

type UserCache interface {
	Get(name string) (*v3.User, error)
	List(selector labels.Selector) ([]*v3.User, error)

	AddIndexer(indexName string, indexer UserIndexer)
	GetByIndex(indexName, key string) ([]*v3.User, error)
}

type UserIndexer func(obj *v3.User) ([]string, error)

type userController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewUserController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) UserController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &userController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromUserHandlerToHandler(sync UserHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.User
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.User))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *userController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.User))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateUserDeepCopyOnChange(client UserClient, obj *v3.User, handler func(obj *v3.User) (*v3.User, error)) (*v3.User, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *userController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *userController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *userController) OnChange(ctx context.Context, name string, sync UserHandler) {
	c.AddGenericHandler(ctx, name, FromUserHandlerToHandler(sync))
}

func (c *userController) OnRemove(ctx context.Context, name string, sync UserHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromUserHandlerToHandler(sync)))
}

func (c *userController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *userController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *userController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *userController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *userController) Cache() UserCache {
	return &userCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *userController) Create(obj *v3.User) (*v3.User, error) {
	result := &v3.User{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *userController) Update(obj *v3.User) (*v3.User, error) {
	result := &v3.User{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *userController) UpdateStatus(obj *v3.User) (*v3.User, error) {
	result := &v3.User{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *userController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *userController) Get(name string, options metav1.GetOptions) (*v3.User, error) {
	result := &v3.User{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *userController) List(opts metav1.ListOptions) (*v3.UserList, error) {
	result := &v3.UserList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *userController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *userController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.User, error) {
	result := &v3.User{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type userCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *userCache) Get(name string) (*v3.User, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.User), nil
}

func (c *userCache) List(selector labels.Selector) (ret []*v3.User, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.User))
	})

	return ret, err
}

func (c *userCache) AddIndexer(indexName string, indexer UserIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.User))
		},
	}))
}

func (c *userCache) GetByIndex(indexName, key string) (result []*v3.User, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.User, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.User))
	}
	return result, nil
}

type UserStatusHandler func(obj *v3.User, status v3.UserStatus) (v3.UserStatus, error)

type UserGeneratingHandler func(obj *v3.User, status v3.UserStatus) ([]runtime.Object, v3.UserStatus, error)

func RegisterUserStatusHandler(ctx context.Context, controller UserController, condition condition.Cond, name string, handler UserStatusHandler) {
	statusHandler := &userStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromUserHandlerToHandler(statusHandler.sync))
}

func RegisterUserGeneratingHandler(ctx context.Context, controller UserController, apply apply.Apply,
	condition condition.Cond, name string, handler UserGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &userGeneratingHandler{
		UserGeneratingHandler: handler,
		apply:                 apply,
		name:                  name,
		gvk:                   controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterUserStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type userStatusHandler struct {
	client    UserClient
	condition condition.Cond
	handler   UserStatusHandler
}

func (a *userStatusHandler) sync(key string, obj *v3.User) (*v3.User, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type userGeneratingHandler struct {
	UserGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *userGeneratingHandler) Remove(key string, obj *v3.User) (*v3.User, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.User{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *userGeneratingHandler) Handle(obj *v3.User, status v3.UserStatus) (v3.UserStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.UserGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
