// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/robdimsdale/wundergo"
)

type FakeHTTPHelper struct {
	GetStub        func(url string) ([]byte, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		url string
	}
	getReturns struct {
		result1 []byte
		result2 error
	}
	PostStub        func(url string, body string) ([]byte, error)
	postMutex       sync.RWMutex
	postArgsForCall []struct {
		url  string
		body string
	}
	postReturns struct {
		result1 []byte
		result2 error
	}
	PutStub        func(url string, body string) ([]byte, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		url  string
		body string
	}
	putReturns struct {
		result1 []byte
		result2 error
	}
}

func (fake *FakeHTTPHelper) Get(url string) ([]byte, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		url string
	}{url})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(url)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeHTTPHelper) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeHTTPHelper) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].url
}

func (fake *FakeHTTPHelper) GetReturns(result1 []byte, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPHelper) Post(url string, body string) ([]byte, error) {
	fake.postMutex.Lock()
	fake.postArgsForCall = append(fake.postArgsForCall, struct {
		url  string
		body string
	}{url, body})
	fake.postMutex.Unlock()
	if fake.PostStub != nil {
		return fake.PostStub(url, body)
	} else {
		return fake.postReturns.result1, fake.postReturns.result2
	}
}

func (fake *FakeHTTPHelper) PostCallCount() int {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return len(fake.postArgsForCall)
}

func (fake *FakeHTTPHelper) PostArgsForCall(i int) (string, string) {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return fake.postArgsForCall[i].url, fake.postArgsForCall[i].body
}

func (fake *FakeHTTPHelper) PostReturns(result1 []byte, result2 error) {
	fake.PostStub = nil
	fake.postReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPHelper) Put(url string, body string) ([]byte, error) {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		url  string
		body string
	}{url, body})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(url, body)
	} else {
		return fake.putReturns.result1, fake.putReturns.result2
	}
}

func (fake *FakeHTTPHelper) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeHTTPHelper) PutArgsForCall(i int) (string, string) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].url, fake.putArgsForCall[i].body
}

func (fake *FakeHTTPHelper) PutReturns(result1 []byte, result2 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

var _ wundergo.HTTPHelper = new(FakeHTTPHelper)
