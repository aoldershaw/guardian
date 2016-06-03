// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/goci"
	"github.com/cloudfoundry-incubator/guardian/rundmc/depot"
)

type FakeBundleLoader struct {
	LoadStub        func(bundleDir string) (goci.Bndl, error)
	loadMutex       sync.RWMutex
	loadArgsForCall []struct {
		bundleDir string
	}
	loadReturns struct {
		result1 goci.Bndl
		result2 error
	}
}

func (fake *FakeBundleLoader) Load(bundleDir string) (goci.Bndl, error) {
	fake.loadMutex.Lock()
	fake.loadArgsForCall = append(fake.loadArgsForCall, struct {
		bundleDir string
	}{bundleDir})
	fake.loadMutex.Unlock()
	if fake.LoadStub != nil {
		return fake.LoadStub(bundleDir)
	} else {
		return fake.loadReturns.result1, fake.loadReturns.result2
	}
}

func (fake *FakeBundleLoader) LoadCallCount() int {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return len(fake.loadArgsForCall)
}

func (fake *FakeBundleLoader) LoadArgsForCall(i int) string {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return fake.loadArgsForCall[i].bundleDir
}

func (fake *FakeBundleLoader) LoadReturns(result1 goci.Bndl, result2 error) {
	fake.LoadStub = nil
	fake.loadReturns = struct {
		result1 goci.Bndl
		result2 error
	}{result1, result2}
}

var _ depot.BundleLoader = new(FakeBundleLoader)
