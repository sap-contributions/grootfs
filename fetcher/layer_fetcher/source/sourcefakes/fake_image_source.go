// Code generated by counterfeiter. DO NOT EDIT.
package sourcefakes

import (
	"context"
	"io"
	"sync"

	"code.cloudfoundry.org/grootfs/fetcher/layer_fetcher/source"
	"github.com/containers/image/types"
	digest "github.com/opencontainers/go-digest"
)

type FakeImageSource struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	GetBlobStub        func(context.Context, types.BlobInfo, types.BlobInfoCache) (io.ReadCloser, int64, error)
	getBlobMutex       sync.RWMutex
	getBlobArgsForCall []struct {
		arg1 context.Context
		arg2 types.BlobInfo
		arg3 types.BlobInfoCache
	}
	getBlobReturns struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}
	getBlobReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}
	GetManifestStub        func(context.Context, *digest.Digest) ([]byte, string, error)
	getManifestMutex       sync.RWMutex
	getManifestArgsForCall []struct {
		arg1 context.Context
		arg2 *digest.Digest
	}
	getManifestReturns struct {
		result1 []byte
		result2 string
		result3 error
	}
	getManifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 string
		result3 error
	}
	GetSignaturesStub        func(context.Context, *digest.Digest) ([][]byte, error)
	getSignaturesMutex       sync.RWMutex
	getSignaturesArgsForCall []struct {
		arg1 context.Context
		arg2 *digest.Digest
	}
	getSignaturesReturns struct {
		result1 [][]byte
		result2 error
	}
	getSignaturesReturnsOnCall map[int]struct {
		result1 [][]byte
		result2 error
	}
	HasThreadSafeGetBlobStub        func() bool
	hasThreadSafeGetBlobMutex       sync.RWMutex
	hasThreadSafeGetBlobArgsForCall []struct {
	}
	hasThreadSafeGetBlobReturns struct {
		result1 bool
	}
	hasThreadSafeGetBlobReturnsOnCall map[int]struct {
		result1 bool
	}
	LayerInfosForCopyStub        func(context.Context) ([]types.BlobInfo, error)
	layerInfosForCopyMutex       sync.RWMutex
	layerInfosForCopyArgsForCall []struct {
		arg1 context.Context
	}
	layerInfosForCopyReturns struct {
		result1 []types.BlobInfo
		result2 error
	}
	layerInfosForCopyReturnsOnCall map[int]struct {
		result1 []types.BlobInfo
		result2 error
	}
	ReferenceStub        func() types.ImageReference
	referenceMutex       sync.RWMutex
	referenceArgsForCall []struct {
	}
	referenceReturns struct {
		result1 types.ImageReference
	}
	referenceReturnsOnCall map[int]struct {
		result1 types.ImageReference
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageSource) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeImageSource) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeImageSource) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeImageSource) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageSource) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageSource) GetBlob(arg1 context.Context, arg2 types.BlobInfo, arg3 types.BlobInfoCache) (io.ReadCloser, int64, error) {
	fake.getBlobMutex.Lock()
	ret, specificReturn := fake.getBlobReturnsOnCall[len(fake.getBlobArgsForCall)]
	fake.getBlobArgsForCall = append(fake.getBlobArgsForCall, struct {
		arg1 context.Context
		arg2 types.BlobInfo
		arg3 types.BlobInfoCache
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetBlob", []interface{}{arg1, arg2, arg3})
	fake.getBlobMutex.Unlock()
	if fake.GetBlobStub != nil {
		return fake.GetBlobStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getBlobReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeImageSource) GetBlobCallCount() int {
	fake.getBlobMutex.RLock()
	defer fake.getBlobMutex.RUnlock()
	return len(fake.getBlobArgsForCall)
}

func (fake *FakeImageSource) GetBlobCalls(stub func(context.Context, types.BlobInfo, types.BlobInfoCache) (io.ReadCloser, int64, error)) {
	fake.getBlobMutex.Lock()
	defer fake.getBlobMutex.Unlock()
	fake.GetBlobStub = stub
}

func (fake *FakeImageSource) GetBlobArgsForCall(i int) (context.Context, types.BlobInfo, types.BlobInfoCache) {
	fake.getBlobMutex.RLock()
	defer fake.getBlobMutex.RUnlock()
	argsForCall := fake.getBlobArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImageSource) GetBlobReturns(result1 io.ReadCloser, result2 int64, result3 error) {
	fake.getBlobMutex.Lock()
	defer fake.getBlobMutex.Unlock()
	fake.GetBlobStub = nil
	fake.getBlobReturns = struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImageSource) GetBlobReturnsOnCall(i int, result1 io.ReadCloser, result2 int64, result3 error) {
	fake.getBlobMutex.Lock()
	defer fake.getBlobMutex.Unlock()
	fake.GetBlobStub = nil
	if fake.getBlobReturnsOnCall == nil {
		fake.getBlobReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 int64
			result3 error
		})
	}
	fake.getBlobReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImageSource) GetManifest(arg1 context.Context, arg2 *digest.Digest) ([]byte, string, error) {
	fake.getManifestMutex.Lock()
	ret, specificReturn := fake.getManifestReturnsOnCall[len(fake.getManifestArgsForCall)]
	fake.getManifestArgsForCall = append(fake.getManifestArgsForCall, struct {
		arg1 context.Context
		arg2 *digest.Digest
	}{arg1, arg2})
	fake.recordInvocation("GetManifest", []interface{}{arg1, arg2})
	fake.getManifestMutex.Unlock()
	if fake.GetManifestStub != nil {
		return fake.GetManifestStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getManifestReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeImageSource) GetManifestCallCount() int {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return len(fake.getManifestArgsForCall)
}

func (fake *FakeImageSource) GetManifestCalls(stub func(context.Context, *digest.Digest) ([]byte, string, error)) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = stub
}

func (fake *FakeImageSource) GetManifestArgsForCall(i int) (context.Context, *digest.Digest) {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	argsForCall := fake.getManifestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageSource) GetManifestReturns(result1 []byte, result2 string, result3 error) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = nil
	fake.getManifestReturns = struct {
		result1 []byte
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImageSource) GetManifestReturnsOnCall(i int, result1 []byte, result2 string, result3 error) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = nil
	if fake.getManifestReturnsOnCall == nil {
		fake.getManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 string
			result3 error
		})
	}
	fake.getManifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImageSource) GetSignatures(arg1 context.Context, arg2 *digest.Digest) ([][]byte, error) {
	fake.getSignaturesMutex.Lock()
	ret, specificReturn := fake.getSignaturesReturnsOnCall[len(fake.getSignaturesArgsForCall)]
	fake.getSignaturesArgsForCall = append(fake.getSignaturesArgsForCall, struct {
		arg1 context.Context
		arg2 *digest.Digest
	}{arg1, arg2})
	fake.recordInvocation("GetSignatures", []interface{}{arg1, arg2})
	fake.getSignaturesMutex.Unlock()
	if fake.GetSignaturesStub != nil {
		return fake.GetSignaturesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSignaturesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageSource) GetSignaturesCallCount() int {
	fake.getSignaturesMutex.RLock()
	defer fake.getSignaturesMutex.RUnlock()
	return len(fake.getSignaturesArgsForCall)
}

func (fake *FakeImageSource) GetSignaturesCalls(stub func(context.Context, *digest.Digest) ([][]byte, error)) {
	fake.getSignaturesMutex.Lock()
	defer fake.getSignaturesMutex.Unlock()
	fake.GetSignaturesStub = stub
}

func (fake *FakeImageSource) GetSignaturesArgsForCall(i int) (context.Context, *digest.Digest) {
	fake.getSignaturesMutex.RLock()
	defer fake.getSignaturesMutex.RUnlock()
	argsForCall := fake.getSignaturesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageSource) GetSignaturesReturns(result1 [][]byte, result2 error) {
	fake.getSignaturesMutex.Lock()
	defer fake.getSignaturesMutex.Unlock()
	fake.GetSignaturesStub = nil
	fake.getSignaturesReturns = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImageSource) GetSignaturesReturnsOnCall(i int, result1 [][]byte, result2 error) {
	fake.getSignaturesMutex.Lock()
	defer fake.getSignaturesMutex.Unlock()
	fake.GetSignaturesStub = nil
	if fake.getSignaturesReturnsOnCall == nil {
		fake.getSignaturesReturnsOnCall = make(map[int]struct {
			result1 [][]byte
			result2 error
		})
	}
	fake.getSignaturesReturnsOnCall[i] = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImageSource) HasThreadSafeGetBlob() bool {
	fake.hasThreadSafeGetBlobMutex.Lock()
	ret, specificReturn := fake.hasThreadSafeGetBlobReturnsOnCall[len(fake.hasThreadSafeGetBlobArgsForCall)]
	fake.hasThreadSafeGetBlobArgsForCall = append(fake.hasThreadSafeGetBlobArgsForCall, struct {
	}{})
	fake.recordInvocation("HasThreadSafeGetBlob", []interface{}{})
	fake.hasThreadSafeGetBlobMutex.Unlock()
	if fake.HasThreadSafeGetBlobStub != nil {
		return fake.HasThreadSafeGetBlobStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hasThreadSafeGetBlobReturns
	return fakeReturns.result1
}

func (fake *FakeImageSource) HasThreadSafeGetBlobCallCount() int {
	fake.hasThreadSafeGetBlobMutex.RLock()
	defer fake.hasThreadSafeGetBlobMutex.RUnlock()
	return len(fake.hasThreadSafeGetBlobArgsForCall)
}

func (fake *FakeImageSource) HasThreadSafeGetBlobCalls(stub func() bool) {
	fake.hasThreadSafeGetBlobMutex.Lock()
	defer fake.hasThreadSafeGetBlobMutex.Unlock()
	fake.HasThreadSafeGetBlobStub = stub
}

func (fake *FakeImageSource) HasThreadSafeGetBlobReturns(result1 bool) {
	fake.hasThreadSafeGetBlobMutex.Lock()
	defer fake.hasThreadSafeGetBlobMutex.Unlock()
	fake.HasThreadSafeGetBlobStub = nil
	fake.hasThreadSafeGetBlobReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImageSource) HasThreadSafeGetBlobReturnsOnCall(i int, result1 bool) {
	fake.hasThreadSafeGetBlobMutex.Lock()
	defer fake.hasThreadSafeGetBlobMutex.Unlock()
	fake.HasThreadSafeGetBlobStub = nil
	if fake.hasThreadSafeGetBlobReturnsOnCall == nil {
		fake.hasThreadSafeGetBlobReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasThreadSafeGetBlobReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImageSource) LayerInfosForCopy(arg1 context.Context) ([]types.BlobInfo, error) {
	fake.layerInfosForCopyMutex.Lock()
	ret, specificReturn := fake.layerInfosForCopyReturnsOnCall[len(fake.layerInfosForCopyArgsForCall)]
	fake.layerInfosForCopyArgsForCall = append(fake.layerInfosForCopyArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("LayerInfosForCopy", []interface{}{arg1})
	fake.layerInfosForCopyMutex.Unlock()
	if fake.LayerInfosForCopyStub != nil {
		return fake.LayerInfosForCopyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.layerInfosForCopyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageSource) LayerInfosForCopyCallCount() int {
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	return len(fake.layerInfosForCopyArgsForCall)
}

func (fake *FakeImageSource) LayerInfosForCopyCalls(stub func(context.Context) ([]types.BlobInfo, error)) {
	fake.layerInfosForCopyMutex.Lock()
	defer fake.layerInfosForCopyMutex.Unlock()
	fake.LayerInfosForCopyStub = stub
}

func (fake *FakeImageSource) LayerInfosForCopyArgsForCall(i int) context.Context {
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	argsForCall := fake.layerInfosForCopyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImageSource) LayerInfosForCopyReturns(result1 []types.BlobInfo, result2 error) {
	fake.layerInfosForCopyMutex.Lock()
	defer fake.layerInfosForCopyMutex.Unlock()
	fake.LayerInfosForCopyStub = nil
	fake.layerInfosForCopyReturns = struct {
		result1 []types.BlobInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeImageSource) LayerInfosForCopyReturnsOnCall(i int, result1 []types.BlobInfo, result2 error) {
	fake.layerInfosForCopyMutex.Lock()
	defer fake.layerInfosForCopyMutex.Unlock()
	fake.LayerInfosForCopyStub = nil
	if fake.layerInfosForCopyReturnsOnCall == nil {
		fake.layerInfosForCopyReturnsOnCall = make(map[int]struct {
			result1 []types.BlobInfo
			result2 error
		})
	}
	fake.layerInfosForCopyReturnsOnCall[i] = struct {
		result1 []types.BlobInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeImageSource) Reference() types.ImageReference {
	fake.referenceMutex.Lock()
	ret, specificReturn := fake.referenceReturnsOnCall[len(fake.referenceArgsForCall)]
	fake.referenceArgsForCall = append(fake.referenceArgsForCall, struct {
	}{})
	fake.recordInvocation("Reference", []interface{}{})
	fake.referenceMutex.Unlock()
	if fake.ReferenceStub != nil {
		return fake.ReferenceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.referenceReturns
	return fakeReturns.result1
}

func (fake *FakeImageSource) ReferenceCallCount() int {
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	return len(fake.referenceArgsForCall)
}

func (fake *FakeImageSource) ReferenceCalls(stub func() types.ImageReference) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = stub
}

func (fake *FakeImageSource) ReferenceReturns(result1 types.ImageReference) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = nil
	fake.referenceReturns = struct {
		result1 types.ImageReference
	}{result1}
}

func (fake *FakeImageSource) ReferenceReturnsOnCall(i int, result1 types.ImageReference) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = nil
	if fake.referenceReturnsOnCall == nil {
		fake.referenceReturnsOnCall = make(map[int]struct {
			result1 types.ImageReference
		})
	}
	fake.referenceReturnsOnCall[i] = struct {
		result1 types.ImageReference
	}{result1}
}

func (fake *FakeImageSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.getBlobMutex.RLock()
	defer fake.getBlobMutex.RUnlock()
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	fake.getSignaturesMutex.RLock()
	defer fake.getSignaturesMutex.RUnlock()
	fake.hasThreadSafeGetBlobMutex.RLock()
	defer fake.hasThreadSafeGetBlobMutex.RUnlock()
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageSource) recordInvocation(key string, args []interface{}) {
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

var _ source.ImageSource = new(FakeImageSource)
