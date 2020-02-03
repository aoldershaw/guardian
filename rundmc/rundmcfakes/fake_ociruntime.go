// Code generated by counterfeiter. DO NOT EDIT.
package rundmcfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/guardian/rundmc"
	"code.cloudfoundry.org/guardian/rundmc/event"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"code.cloudfoundry.org/lager"
)

type FakeOCIRuntime struct {
	AttachStub        func(lager.Logger, string, string, garden.ProcessIO) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	attachReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	BundleInfoStub        func(lager.Logger, string) (string, goci.Bndl, error)
	bundleInfoMutex       sync.RWMutex
	bundleInfoArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	bundleInfoReturns struct {
		result1 string
		result2 goci.Bndl
		result3 error
	}
	bundleInfoReturnsOnCall map[int]struct {
		result1 string
		result2 goci.Bndl
		result3 error
	}
	ContainerHandlesStub        func() ([]string, error)
	containerHandlesMutex       sync.RWMutex
	containerHandlesArgsForCall []struct {
	}
	containerHandlesReturns struct {
		result1 []string
		result2 error
	}
	containerHandlesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ContainerPeaHandlesStub        func(lager.Logger, string) ([]string, error)
	containerPeaHandlesMutex       sync.RWMutex
	containerPeaHandlesArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	containerPeaHandlesReturns struct {
		result1 []string
		result2 error
	}
	containerPeaHandlesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	CreateStub        func(lager.Logger, string, goci.Bndl, garden.ProcessIO) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 goci.Bndl
		arg4 garden.ProcessIO
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	CreatePeaStub        func(lager.Logger, string, string, goci.Bndl, garden.ProcessIO) error
	createPeaMutex       sync.RWMutex
	createPeaArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 goci.Bndl
		arg5 garden.ProcessIO
	}
	createPeaReturns struct {
		result1 error
	}
	createPeaReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(lager.Logger, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	EventsStub        func(lager.Logger) (<-chan event.Event, error)
	eventsMutex       sync.RWMutex
	eventsArgsForCall []struct {
		arg1 lager.Logger
	}
	eventsReturns struct {
		result1 <-chan event.Event
		result2 error
	}
	eventsReturnsOnCall map[int]struct {
		result1 <-chan event.Event
		result2 error
	}
	ExecStub        func(lager.Logger, string, garden.ProcessSpec, garden.ProcessIO) (garden.Process, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 garden.ProcessSpec
		arg4 garden.ProcessIO
	}
	execReturns struct {
		result1 garden.Process
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	RemoveBundleStub        func(lager.Logger, string) error
	removeBundleMutex       sync.RWMutex
	removeBundleArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	removeBundleReturns struct {
		result1 error
	}
	removeBundleReturnsOnCall map[int]struct {
		result1 error
	}
	StateStub        func(lager.Logger, string) (rundmc.State, error)
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	stateReturns struct {
		result1 rundmc.State
		result2 error
	}
	stateReturnsOnCall map[int]struct {
		result1 rundmc.State
		result2 error
	}
	StatsStub        func(lager.Logger, string) (gardener.StatsContainerMetrics, error)
	statsMutex       sync.RWMutex
	statsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	statsReturns struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}
	statsReturnsOnCall map[int]struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOCIRuntime) Attach(arg1 lager.Logger, arg2 string, arg3 string, arg4 garden.ProcessIO) (garden.Process, error) {
	fake.attachMutex.Lock()
	ret, specificReturn := fake.attachReturnsOnCall[len(fake.attachArgsForCall)]
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessIO
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Attach", []interface{}{arg1, arg2, arg3, arg4})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.attachReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOCIRuntime) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeOCIRuntime) AttachCalls(stub func(lager.Logger, string, string, garden.ProcessIO) (garden.Process, error)) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = stub
}

func (fake *FakeOCIRuntime) AttachArgsForCall(i int) (lager.Logger, string, string, garden.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	argsForCall := fake.attachArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeOCIRuntime) AttachReturns(result1 garden.Process, result2 error) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) AttachReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = nil
	if fake.attachReturnsOnCall == nil {
		fake.attachReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.attachReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) BundleInfo(arg1 lager.Logger, arg2 string) (string, goci.Bndl, error) {
	fake.bundleInfoMutex.Lock()
	ret, specificReturn := fake.bundleInfoReturnsOnCall[len(fake.bundleInfoArgsForCall)]
	fake.bundleInfoArgsForCall = append(fake.bundleInfoArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("BundleInfo", []interface{}{arg1, arg2})
	fake.bundleInfoMutex.Unlock()
	if fake.BundleInfoStub != nil {
		return fake.BundleInfoStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.bundleInfoReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeOCIRuntime) BundleInfoCallCount() int {
	fake.bundleInfoMutex.RLock()
	defer fake.bundleInfoMutex.RUnlock()
	return len(fake.bundleInfoArgsForCall)
}

func (fake *FakeOCIRuntime) BundleInfoCalls(stub func(lager.Logger, string) (string, goci.Bndl, error)) {
	fake.bundleInfoMutex.Lock()
	defer fake.bundleInfoMutex.Unlock()
	fake.BundleInfoStub = stub
}

func (fake *FakeOCIRuntime) BundleInfoArgsForCall(i int) (lager.Logger, string) {
	fake.bundleInfoMutex.RLock()
	defer fake.bundleInfoMutex.RUnlock()
	argsForCall := fake.bundleInfoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOCIRuntime) BundleInfoReturns(result1 string, result2 goci.Bndl, result3 error) {
	fake.bundleInfoMutex.Lock()
	defer fake.bundleInfoMutex.Unlock()
	fake.BundleInfoStub = nil
	fake.bundleInfoReturns = struct {
		result1 string
		result2 goci.Bndl
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOCIRuntime) BundleInfoReturnsOnCall(i int, result1 string, result2 goci.Bndl, result3 error) {
	fake.bundleInfoMutex.Lock()
	defer fake.bundleInfoMutex.Unlock()
	fake.BundleInfoStub = nil
	if fake.bundleInfoReturnsOnCall == nil {
		fake.bundleInfoReturnsOnCall = make(map[int]struct {
			result1 string
			result2 goci.Bndl
			result3 error
		})
	}
	fake.bundleInfoReturnsOnCall[i] = struct {
		result1 string
		result2 goci.Bndl
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOCIRuntime) ContainerHandles() ([]string, error) {
	fake.containerHandlesMutex.Lock()
	ret, specificReturn := fake.containerHandlesReturnsOnCall[len(fake.containerHandlesArgsForCall)]
	fake.containerHandlesArgsForCall = append(fake.containerHandlesArgsForCall, struct {
	}{})
	fake.recordInvocation("ContainerHandles", []interface{}{})
	fake.containerHandlesMutex.Unlock()
	if fake.ContainerHandlesStub != nil {
		return fake.ContainerHandlesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.containerHandlesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOCIRuntime) ContainerHandlesCallCount() int {
	fake.containerHandlesMutex.RLock()
	defer fake.containerHandlesMutex.RUnlock()
	return len(fake.containerHandlesArgsForCall)
}

func (fake *FakeOCIRuntime) ContainerHandlesCalls(stub func() ([]string, error)) {
	fake.containerHandlesMutex.Lock()
	defer fake.containerHandlesMutex.Unlock()
	fake.ContainerHandlesStub = stub
}

func (fake *FakeOCIRuntime) ContainerHandlesReturns(result1 []string, result2 error) {
	fake.containerHandlesMutex.Lock()
	defer fake.containerHandlesMutex.Unlock()
	fake.ContainerHandlesStub = nil
	fake.containerHandlesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) ContainerHandlesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.containerHandlesMutex.Lock()
	defer fake.containerHandlesMutex.Unlock()
	fake.ContainerHandlesStub = nil
	if fake.containerHandlesReturnsOnCall == nil {
		fake.containerHandlesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.containerHandlesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) ContainerPeaHandles(arg1 lager.Logger, arg2 string) ([]string, error) {
	fake.containerPeaHandlesMutex.Lock()
	ret, specificReturn := fake.containerPeaHandlesReturnsOnCall[len(fake.containerPeaHandlesArgsForCall)]
	fake.containerPeaHandlesArgsForCall = append(fake.containerPeaHandlesArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ContainerPeaHandles", []interface{}{arg1, arg2})
	fake.containerPeaHandlesMutex.Unlock()
	if fake.ContainerPeaHandlesStub != nil {
		return fake.ContainerPeaHandlesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.containerPeaHandlesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOCIRuntime) ContainerPeaHandlesCallCount() int {
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	return len(fake.containerPeaHandlesArgsForCall)
}

func (fake *FakeOCIRuntime) ContainerPeaHandlesCalls(stub func(lager.Logger, string) ([]string, error)) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = stub
}

func (fake *FakeOCIRuntime) ContainerPeaHandlesArgsForCall(i int) (lager.Logger, string) {
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	argsForCall := fake.containerPeaHandlesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOCIRuntime) ContainerPeaHandlesReturns(result1 []string, result2 error) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = nil
	fake.containerPeaHandlesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) ContainerPeaHandlesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = nil
	if fake.containerPeaHandlesReturnsOnCall == nil {
		fake.containerPeaHandlesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.containerPeaHandlesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) Create(arg1 lager.Logger, arg2 string, arg3 goci.Bndl, arg4 garden.ProcessIO) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 goci.Bndl
		arg4 garden.ProcessIO
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeOCIRuntime) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeOCIRuntime) CreateCalls(stub func(lager.Logger, string, goci.Bndl, garden.ProcessIO) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeOCIRuntime) CreateArgsForCall(i int) (lager.Logger, string, goci.Bndl, garden.ProcessIO) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeOCIRuntime) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) CreatePea(arg1 lager.Logger, arg2 string, arg3 string, arg4 goci.Bndl, arg5 garden.ProcessIO) error {
	fake.createPeaMutex.Lock()
	ret, specificReturn := fake.createPeaReturnsOnCall[len(fake.createPeaArgsForCall)]
	fake.createPeaArgsForCall = append(fake.createPeaArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 goci.Bndl
		arg5 garden.ProcessIO
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("CreatePea", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createPeaMutex.Unlock()
	if fake.CreatePeaStub != nil {
		return fake.CreatePeaStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createPeaReturns
	return fakeReturns.result1
}

func (fake *FakeOCIRuntime) CreatePeaCallCount() int {
	fake.createPeaMutex.RLock()
	defer fake.createPeaMutex.RUnlock()
	return len(fake.createPeaArgsForCall)
}

func (fake *FakeOCIRuntime) CreatePeaCalls(stub func(lager.Logger, string, string, goci.Bndl, garden.ProcessIO) error) {
	fake.createPeaMutex.Lock()
	defer fake.createPeaMutex.Unlock()
	fake.CreatePeaStub = stub
}

func (fake *FakeOCIRuntime) CreatePeaArgsForCall(i int) (lager.Logger, string, string, goci.Bndl, garden.ProcessIO) {
	fake.createPeaMutex.RLock()
	defer fake.createPeaMutex.RUnlock()
	argsForCall := fake.createPeaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeOCIRuntime) CreatePeaReturns(result1 error) {
	fake.createPeaMutex.Lock()
	defer fake.createPeaMutex.Unlock()
	fake.CreatePeaStub = nil
	fake.createPeaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) CreatePeaReturnsOnCall(i int, result1 error) {
	fake.createPeaMutex.Lock()
	defer fake.createPeaMutex.Unlock()
	fake.CreatePeaStub = nil
	if fake.createPeaReturnsOnCall == nil {
		fake.createPeaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createPeaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) Delete(arg1 lager.Logger, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeOCIRuntime) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeOCIRuntime) DeleteCalls(stub func(lager.Logger, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeOCIRuntime) DeleteArgsForCall(i int) (lager.Logger, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOCIRuntime) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) Events(arg1 lager.Logger) (<-chan event.Event, error) {
	fake.eventsMutex.Lock()
	ret, specificReturn := fake.eventsReturnsOnCall[len(fake.eventsArgsForCall)]
	fake.eventsArgsForCall = append(fake.eventsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Events", []interface{}{arg1})
	fake.eventsMutex.Unlock()
	if fake.EventsStub != nil {
		return fake.EventsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.eventsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOCIRuntime) EventsCallCount() int {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	return len(fake.eventsArgsForCall)
}

func (fake *FakeOCIRuntime) EventsCalls(stub func(lager.Logger) (<-chan event.Event, error)) {
	fake.eventsMutex.Lock()
	defer fake.eventsMutex.Unlock()
	fake.EventsStub = stub
}

func (fake *FakeOCIRuntime) EventsArgsForCall(i int) lager.Logger {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	argsForCall := fake.eventsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOCIRuntime) EventsReturns(result1 <-chan event.Event, result2 error) {
	fake.eventsMutex.Lock()
	defer fake.eventsMutex.Unlock()
	fake.EventsStub = nil
	fake.eventsReturns = struct {
		result1 <-chan event.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) EventsReturnsOnCall(i int, result1 <-chan event.Event, result2 error) {
	fake.eventsMutex.Lock()
	defer fake.eventsMutex.Unlock()
	fake.EventsStub = nil
	if fake.eventsReturnsOnCall == nil {
		fake.eventsReturnsOnCall = make(map[int]struct {
			result1 <-chan event.Event
			result2 error
		})
	}
	fake.eventsReturnsOnCall[i] = struct {
		result1 <-chan event.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) Exec(arg1 lager.Logger, arg2 string, arg3 garden.ProcessSpec, arg4 garden.ProcessIO) (garden.Process, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 garden.ProcessSpec
		arg4 garden.ProcessIO
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2, arg3, arg4})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOCIRuntime) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeOCIRuntime) ExecCalls(stub func(lager.Logger, string, garden.ProcessSpec, garden.ProcessIO) (garden.Process, error)) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeOCIRuntime) ExecArgsForCall(i int) (lager.Logger, string, garden.ProcessSpec, garden.ProcessIO) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeOCIRuntime) ExecReturns(result1 garden.Process, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) ExecReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) RemoveBundle(arg1 lager.Logger, arg2 string) error {
	fake.removeBundleMutex.Lock()
	ret, specificReturn := fake.removeBundleReturnsOnCall[len(fake.removeBundleArgsForCall)]
	fake.removeBundleArgsForCall = append(fake.removeBundleArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("RemoveBundle", []interface{}{arg1, arg2})
	fake.removeBundleMutex.Unlock()
	if fake.RemoveBundleStub != nil {
		return fake.RemoveBundleStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeBundleReturns
	return fakeReturns.result1
}

func (fake *FakeOCIRuntime) RemoveBundleCallCount() int {
	fake.removeBundleMutex.RLock()
	defer fake.removeBundleMutex.RUnlock()
	return len(fake.removeBundleArgsForCall)
}

func (fake *FakeOCIRuntime) RemoveBundleCalls(stub func(lager.Logger, string) error) {
	fake.removeBundleMutex.Lock()
	defer fake.removeBundleMutex.Unlock()
	fake.RemoveBundleStub = stub
}

func (fake *FakeOCIRuntime) RemoveBundleArgsForCall(i int) (lager.Logger, string) {
	fake.removeBundleMutex.RLock()
	defer fake.removeBundleMutex.RUnlock()
	argsForCall := fake.removeBundleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOCIRuntime) RemoveBundleReturns(result1 error) {
	fake.removeBundleMutex.Lock()
	defer fake.removeBundleMutex.Unlock()
	fake.RemoveBundleStub = nil
	fake.removeBundleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) RemoveBundleReturnsOnCall(i int, result1 error) {
	fake.removeBundleMutex.Lock()
	defer fake.removeBundleMutex.Unlock()
	fake.RemoveBundleStub = nil
	if fake.removeBundleReturnsOnCall == nil {
		fake.removeBundleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeBundleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOCIRuntime) State(arg1 lager.Logger, arg2 string) (rundmc.State, error) {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("State", []interface{}{arg1, arg2})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.stateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOCIRuntime) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeOCIRuntime) StateCalls(stub func(lager.Logger, string) (rundmc.State, error)) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = stub
}

func (fake *FakeOCIRuntime) StateArgsForCall(i int) (lager.Logger, string) {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	argsForCall := fake.stateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOCIRuntime) StateReturns(result1 rundmc.State, result2 error) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 rundmc.State
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) StateReturnsOnCall(i int, result1 rundmc.State, result2 error) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 rundmc.State
			result2 error
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 rundmc.State
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) Stats(arg1 lager.Logger, arg2 string) (gardener.StatsContainerMetrics, error) {
	fake.statsMutex.Lock()
	ret, specificReturn := fake.statsReturnsOnCall[len(fake.statsArgsForCall)]
	fake.statsArgsForCall = append(fake.statsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Stats", []interface{}{arg1, arg2})
	fake.statsMutex.Unlock()
	if fake.StatsStub != nil {
		return fake.StatsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.statsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOCIRuntime) StatsCallCount() int {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return len(fake.statsArgsForCall)
}

func (fake *FakeOCIRuntime) StatsCalls(stub func(lager.Logger, string) (gardener.StatsContainerMetrics, error)) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = stub
}

func (fake *FakeOCIRuntime) StatsArgsForCall(i int) (lager.Logger, string) {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	argsForCall := fake.statsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOCIRuntime) StatsReturns(result1 gardener.StatsContainerMetrics, result2 error) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	fake.statsReturns = struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) StatsReturnsOnCall(i int, result1 gardener.StatsContainerMetrics, result2 error) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	if fake.statsReturnsOnCall == nil {
		fake.statsReturnsOnCall = make(map[int]struct {
			result1 gardener.StatsContainerMetrics
			result2 error
		})
	}
	fake.statsReturnsOnCall[i] = struct {
		result1 gardener.StatsContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeOCIRuntime) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	fake.bundleInfoMutex.RLock()
	defer fake.bundleInfoMutex.RUnlock()
	fake.containerHandlesMutex.RLock()
	defer fake.containerHandlesMutex.RUnlock()
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.createPeaMutex.RLock()
	defer fake.createPeaMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.removeBundleMutex.RLock()
	defer fake.removeBundleMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOCIRuntime) recordInvocation(key string, args []interface{}) {
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

var _ rundmc.OCIRuntime = new(FakeOCIRuntime)
