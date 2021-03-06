// Code generated by counterfeiter. DO NOT EDIT.
package processorfakes

import (
	"sync"

	"github.com/remove-bg/go/processor"
)

type FakePromptInterface struct {
	ConfirmLargeBatchStub        func(int) bool
	confirmLargeBatchMutex       sync.RWMutex
	confirmLargeBatchArgsForCall []struct {
		arg1 int
	}
	confirmLargeBatchReturns struct {
		result1 bool
	}
	confirmLargeBatchReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePromptInterface) ConfirmLargeBatch(arg1 int) bool {
	fake.confirmLargeBatchMutex.Lock()
	ret, specificReturn := fake.confirmLargeBatchReturnsOnCall[len(fake.confirmLargeBatchArgsForCall)]
	fake.confirmLargeBatchArgsForCall = append(fake.confirmLargeBatchArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("ConfirmLargeBatch", []interface{}{arg1})
	fake.confirmLargeBatchMutex.Unlock()
	if fake.ConfirmLargeBatchStub != nil {
		return fake.ConfirmLargeBatchStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.confirmLargeBatchReturns
	return fakeReturns.result1
}

func (fake *FakePromptInterface) ConfirmLargeBatchCallCount() int {
	fake.confirmLargeBatchMutex.RLock()
	defer fake.confirmLargeBatchMutex.RUnlock()
	return len(fake.confirmLargeBatchArgsForCall)
}

func (fake *FakePromptInterface) ConfirmLargeBatchCalls(stub func(int) bool) {
	fake.confirmLargeBatchMutex.Lock()
	defer fake.confirmLargeBatchMutex.Unlock()
	fake.ConfirmLargeBatchStub = stub
}

func (fake *FakePromptInterface) ConfirmLargeBatchArgsForCall(i int) int {
	fake.confirmLargeBatchMutex.RLock()
	defer fake.confirmLargeBatchMutex.RUnlock()
	argsForCall := fake.confirmLargeBatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePromptInterface) ConfirmLargeBatchReturns(result1 bool) {
	fake.confirmLargeBatchMutex.Lock()
	defer fake.confirmLargeBatchMutex.Unlock()
	fake.ConfirmLargeBatchStub = nil
	fake.confirmLargeBatchReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePromptInterface) ConfirmLargeBatchReturnsOnCall(i int, result1 bool) {
	fake.confirmLargeBatchMutex.Lock()
	defer fake.confirmLargeBatchMutex.Unlock()
	fake.ConfirmLargeBatchStub = nil
	if fake.confirmLargeBatchReturnsOnCall == nil {
		fake.confirmLargeBatchReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.confirmLargeBatchReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakePromptInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.confirmLargeBatchMutex.RLock()
	defer fake.confirmLargeBatchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePromptInterface) recordInvocation(key string, args []interface{}) {
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

var _ processor.PromptInterface = new(FakePromptInterface)
