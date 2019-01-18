// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	reconciler "github.com/operator-framework/operator-lifecycle-manager/pkg/controller/registry/reconciler"
)

type FakeRegistryReconciler struct {
	EnsureRegistryServerStub        func(*v1alpha1.CatalogSource) error
	ensureRegistryServerMutex       sync.RWMutex
	ensureRegistryServerArgsForCall []struct {
		arg1 *v1alpha1.CatalogSource
	}
	ensureRegistryServerReturns struct {
		result1 error
	}
	ensureRegistryServerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistryReconciler) EnsureRegistryServer(arg1 *v1alpha1.CatalogSource) error {
	fake.ensureRegistryServerMutex.Lock()
	ret, specificReturn := fake.ensureRegistryServerReturnsOnCall[len(fake.ensureRegistryServerArgsForCall)]
	fake.ensureRegistryServerArgsForCall = append(fake.ensureRegistryServerArgsForCall, struct {
		arg1 *v1alpha1.CatalogSource
	}{arg1})
	fake.recordInvocation("EnsureRegistryServer", []interface{}{arg1})
	fake.ensureRegistryServerMutex.Unlock()
	if fake.EnsureRegistryServerStub != nil {
		return fake.EnsureRegistryServerStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.ensureRegistryServerReturns
	return fakeReturns.result1
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerCallCount() int {
	fake.ensureRegistryServerMutex.RLock()
	defer fake.ensureRegistryServerMutex.RUnlock()
	return len(fake.ensureRegistryServerArgsForCall)
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerCalls(stub func(*v1alpha1.CatalogSource) error) {
	fake.ensureRegistryServerMutex.Lock()
	defer fake.ensureRegistryServerMutex.Unlock()
	fake.EnsureRegistryServerStub = stub
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerArgsForCall(i int) *v1alpha1.CatalogSource {
	fake.ensureRegistryServerMutex.RLock()
	defer fake.ensureRegistryServerMutex.RUnlock()
	argsForCall := fake.ensureRegistryServerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerReturns(result1 error) {
	fake.ensureRegistryServerMutex.Lock()
	defer fake.ensureRegistryServerMutex.Unlock()
	fake.EnsureRegistryServerStub = nil
	fake.ensureRegistryServerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistryReconciler) EnsureRegistryServerReturnsOnCall(i int, result1 error) {
	fake.ensureRegistryServerMutex.Lock()
	defer fake.ensureRegistryServerMutex.Unlock()
	fake.EnsureRegistryServerStub = nil
	if fake.ensureRegistryServerReturnsOnCall == nil {
		fake.ensureRegistryServerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.ensureRegistryServerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistryReconciler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ensureRegistryServerMutex.RLock()
	defer fake.ensureRegistryServerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistryReconciler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ reconciler.RegistryReconciler = new(FakeRegistryReconciler)
