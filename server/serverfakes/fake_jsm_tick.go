// Code generated by counterfeiter. DO NOT EDIT.
package serverfakes

import (
	"sync"

	coolbeans_api_v1 "github.com/1xyz/coolbeans/api/v1"
	"github.com/1xyz/coolbeans/server"
)

type FakeJsmTick struct {
	TickStub        func() (*coolbeans_api_v1.TickResponse, error)
	tickMutex       sync.RWMutex
	tickArgsForCall []struct {
	}
	tickReturns struct {
		result1 *coolbeans_api_v1.TickResponse
		result2 error
	}
	tickReturnsOnCall map[int]struct {
		result1 *coolbeans_api_v1.TickResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJsmTick) Tick() (*coolbeans_api_v1.TickResponse, error) {
	fake.tickMutex.Lock()
	ret, specificReturn := fake.tickReturnsOnCall[len(fake.tickArgsForCall)]
	fake.tickArgsForCall = append(fake.tickArgsForCall, struct {
	}{})
	fake.recordInvocation("Tick", []interface{}{})
	fake.tickMutex.Unlock()
	if fake.TickStub != nil {
		return fake.TickStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.tickReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeJsmTick) TickCallCount() int {
	fake.tickMutex.RLock()
	defer fake.tickMutex.RUnlock()
	return len(fake.tickArgsForCall)
}

func (fake *FakeJsmTick) TickCalls(stub func() (*coolbeans_api_v1.TickResponse, error)) {
	fake.tickMutex.Lock()
	defer fake.tickMutex.Unlock()
	fake.TickStub = stub
}

func (fake *FakeJsmTick) TickReturns(result1 *coolbeans_api_v1.TickResponse, result2 error) {
	fake.tickMutex.Lock()
	defer fake.tickMutex.Unlock()
	fake.TickStub = nil
	fake.tickReturns = struct {
		result1 *coolbeans_api_v1.TickResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeJsmTick) TickReturnsOnCall(i int, result1 *coolbeans_api_v1.TickResponse, result2 error) {
	fake.tickMutex.Lock()
	defer fake.tickMutex.Unlock()
	fake.TickStub = nil
	if fake.tickReturnsOnCall == nil {
		fake.tickReturnsOnCall = make(map[int]struct {
			result1 *coolbeans_api_v1.TickResponse
			result2 error
		})
	}
	fake.tickReturnsOnCall[i] = struct {
		result1 *coolbeans_api_v1.TickResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeJsmTick) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tickMutex.RLock()
	defer fake.tickMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeJsmTick) recordInvocation(key string, args []interface{}) {
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

var _ server.JsmTick = new(FakeJsmTick)
