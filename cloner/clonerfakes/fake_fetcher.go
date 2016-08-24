// This file was generated by counterfeiter
package clonerfakes

import (
	"net/url"
	"sync"

	"code.cloudfoundry.org/grootfs/cloner"
	"code.cloudfoundry.org/lager"
)

type FakeFetcher struct {
	ImageInfoStub        func(logger lager.Logger, imageURL *url.URL) (cloner.ImageInfo, error)
	imageInfoMutex       sync.RWMutex
	imageInfoArgsForCall []struct {
		logger   lager.Logger
		imageURL *url.URL
	}
	imageInfoReturns struct {
		result1 cloner.ImageInfo
		result2 error
	}
	StreamerStub        func(logger lager.Logger, imageURL *url.URL) (cloner.Streamer, error)
	streamerMutex       sync.RWMutex
	streamerArgsForCall []struct {
		logger   lager.Logger
		imageURL *url.URL
	}
	streamerReturns struct {
		result1 cloner.Streamer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetcher) ImageInfo(logger lager.Logger, imageURL *url.URL) (cloner.ImageInfo, error) {
	fake.imageInfoMutex.Lock()
	fake.imageInfoArgsForCall = append(fake.imageInfoArgsForCall, struct {
		logger   lager.Logger
		imageURL *url.URL
	}{logger, imageURL})
	fake.recordInvocation("ImageInfo", []interface{}{logger, imageURL})
	fake.imageInfoMutex.Unlock()
	if fake.ImageInfoStub != nil {
		return fake.ImageInfoStub(logger, imageURL)
	} else {
		return fake.imageInfoReturns.result1, fake.imageInfoReturns.result2
	}
}

func (fake *FakeFetcher) ImageInfoCallCount() int {
	fake.imageInfoMutex.RLock()
	defer fake.imageInfoMutex.RUnlock()
	return len(fake.imageInfoArgsForCall)
}

func (fake *FakeFetcher) ImageInfoArgsForCall(i int) (lager.Logger, *url.URL) {
	fake.imageInfoMutex.RLock()
	defer fake.imageInfoMutex.RUnlock()
	return fake.imageInfoArgsForCall[i].logger, fake.imageInfoArgsForCall[i].imageURL
}

func (fake *FakeFetcher) ImageInfoReturns(result1 cloner.ImageInfo, result2 error) {
	fake.ImageInfoStub = nil
	fake.imageInfoReturns = struct {
		result1 cloner.ImageInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) Streamer(logger lager.Logger, imageURL *url.URL) (cloner.Streamer, error) {
	fake.streamerMutex.Lock()
	fake.streamerArgsForCall = append(fake.streamerArgsForCall, struct {
		logger   lager.Logger
		imageURL *url.URL
	}{logger, imageURL})
	fake.recordInvocation("Streamer", []interface{}{logger, imageURL})
	fake.streamerMutex.Unlock()
	if fake.StreamerStub != nil {
		return fake.StreamerStub(logger, imageURL)
	} else {
		return fake.streamerReturns.result1, fake.streamerReturns.result2
	}
}

func (fake *FakeFetcher) StreamerCallCount() int {
	fake.streamerMutex.RLock()
	defer fake.streamerMutex.RUnlock()
	return len(fake.streamerArgsForCall)
}

func (fake *FakeFetcher) StreamerArgsForCall(i int) (lager.Logger, *url.URL) {
	fake.streamerMutex.RLock()
	defer fake.streamerMutex.RUnlock()
	return fake.streamerArgsForCall[i].logger, fake.streamerArgsForCall[i].imageURL
}

func (fake *FakeFetcher) StreamerReturns(result1 cloner.Streamer, result2 error) {
	fake.StreamerStub = nil
	fake.streamerReturns = struct {
		result1 cloner.Streamer
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.imageInfoMutex.RLock()
	defer fake.imageInfoMutex.RUnlock()
	fake.streamerMutex.RLock()
	defer fake.streamerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFetcher) recordInvocation(key string, args []interface{}) {
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

var _ cloner.Fetcher = new(FakeFetcher)
