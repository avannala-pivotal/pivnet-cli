// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeProductFileClient struct {
	ListStub        func(productSlug string, releaseVersion string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	listReturns struct {
		result1 error
	}
	GetStub        func(productSlug string, releaseVersion string, productFileID int) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		productSlug    string
		releaseVersion string
		productFileID  int
	}
	getReturns struct {
		result1 error
	}
	CreateStub        func(config go_pivnet.CreateProductFileConfig) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		config go_pivnet.CreateProductFileConfig
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(productFileID int, productSlug string, name *string, fileType *string, fileVersion *string, md5 *string, description *string) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		productFileID int
		productSlug   string
		name          *string
		fileType      *string
		fileVersion   *string
		md5           *string
		description   *string
	}
	updateReturns struct {
		result1 error
	}
	AddToReleaseStub        func(productSlug string, releaseVersion string, productFileID int) error
	addToReleaseMutex       sync.RWMutex
	addToReleaseArgsForCall []struct {
		productSlug    string
		releaseVersion string
		productFileID  int
	}
	addToReleaseReturns struct {
		result1 error
	}
	RemoveFromReleaseStub        func(productSlug string, releaseVersion string, productFileID int) error
	removeFromReleaseMutex       sync.RWMutex
	removeFromReleaseArgsForCall []struct {
		productSlug    string
		releaseVersion string
		productFileID  int
	}
	removeFromReleaseReturns struct {
		result1 error
	}
	AddToFileGroupStub        func(productSlug string, fileGroupID int, productFileID int) error
	addToFileGroupMutex       sync.RWMutex
	addToFileGroupArgsForCall []struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}
	addToFileGroupReturns struct {
		result1 error
	}
	RemoveFromFileGroupStub        func(productSlug string, fileGroupID int, productFileID int) error
	removeFromFileGroupMutex       sync.RWMutex
	removeFromFileGroupArgsForCall []struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}
	removeFromFileGroupReturns struct {
		result1 error
	}
	DeleteStub        func(productSlug string, productFileID int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		productSlug   string
		productFileID int
	}
	deleteReturns struct {
		result1 error
	}
	DownloadStub        func(productSlug string, releaseVersion string, productFileID int, filepath string, acceptEULA bool) error
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		productSlug    string
		releaseVersion string
		productFileID  int
		filepath       string
		acceptEULA     bool
	}
	downloadReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProductFileClient) List(productSlug string, releaseVersion string) error {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("List", []interface{}{productSlug, releaseVersion})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(productSlug, releaseVersion)
	} else {
		return fake.listReturns.result1
	}
}

func (fake *FakeProductFileClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeProductFileClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].productSlug, fake.listArgsForCall[i].releaseVersion
}

func (fake *FakeProductFileClient) ListReturns(result1 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) Get(productSlug string, releaseVersion string, productFileID int) error {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		productSlug    string
		releaseVersion string
		productFileID  int
	}{productSlug, releaseVersion, productFileID})
	fake.recordInvocation("Get", []interface{}{productSlug, releaseVersion, productFileID})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(productSlug, releaseVersion, productFileID)
	} else {
		return fake.getReturns.result1
	}
}

func (fake *FakeProductFileClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeProductFileClient) GetArgsForCall(i int) (string, string, int) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].productSlug, fake.getArgsForCall[i].releaseVersion, fake.getArgsForCall[i].productFileID
}

func (fake *FakeProductFileClient) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) Create(config go_pivnet.CreateProductFileConfig) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		config go_pivnet.CreateProductFileConfig
	}{config})
	fake.recordInvocation("Create", []interface{}{config})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(config)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeProductFileClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeProductFileClient) CreateArgsForCall(i int) go_pivnet.CreateProductFileConfig {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].config
}

func (fake *FakeProductFileClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) Update(productFileID int, productSlug string, name *string, fileType *string, fileVersion *string, md5 *string, description *string) error {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		productFileID int
		productSlug   string
		name          *string
		fileType      *string
		fileVersion   *string
		md5           *string
		description   *string
	}{productFileID, productSlug, name, fileType, fileVersion, md5, description})
	fake.recordInvocation("Update", []interface{}{productFileID, productSlug, name, fileType, fileVersion, md5, description})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(productFileID, productSlug, name, fileType, fileVersion, md5, description)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeProductFileClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeProductFileClient) UpdateArgsForCall(i int) (int, string, *string, *string, *string, *string, *string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].productFileID, fake.updateArgsForCall[i].productSlug, fake.updateArgsForCall[i].name, fake.updateArgsForCall[i].fileType, fake.updateArgsForCall[i].fileVersion, fake.updateArgsForCall[i].md5, fake.updateArgsForCall[i].description
}

func (fake *FakeProductFileClient) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) AddToRelease(productSlug string, releaseVersion string, productFileID int) error {
	fake.addToReleaseMutex.Lock()
	fake.addToReleaseArgsForCall = append(fake.addToReleaseArgsForCall, struct {
		productSlug    string
		releaseVersion string
		productFileID  int
	}{productSlug, releaseVersion, productFileID})
	fake.recordInvocation("AddToRelease", []interface{}{productSlug, releaseVersion, productFileID})
	fake.addToReleaseMutex.Unlock()
	if fake.AddToReleaseStub != nil {
		return fake.AddToReleaseStub(productSlug, releaseVersion, productFileID)
	} else {
		return fake.addToReleaseReturns.result1
	}
}

func (fake *FakeProductFileClient) AddToReleaseCallCount() int {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return len(fake.addToReleaseArgsForCall)
}

func (fake *FakeProductFileClient) AddToReleaseArgsForCall(i int) (string, string, int) {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return fake.addToReleaseArgsForCall[i].productSlug, fake.addToReleaseArgsForCall[i].releaseVersion, fake.addToReleaseArgsForCall[i].productFileID
}

func (fake *FakeProductFileClient) AddToReleaseReturns(result1 error) {
	fake.AddToReleaseStub = nil
	fake.addToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) RemoveFromRelease(productSlug string, releaseVersion string, productFileID int) error {
	fake.removeFromReleaseMutex.Lock()
	fake.removeFromReleaseArgsForCall = append(fake.removeFromReleaseArgsForCall, struct {
		productSlug    string
		releaseVersion string
		productFileID  int
	}{productSlug, releaseVersion, productFileID})
	fake.recordInvocation("RemoveFromRelease", []interface{}{productSlug, releaseVersion, productFileID})
	fake.removeFromReleaseMutex.Unlock()
	if fake.RemoveFromReleaseStub != nil {
		return fake.RemoveFromReleaseStub(productSlug, releaseVersion, productFileID)
	} else {
		return fake.removeFromReleaseReturns.result1
	}
}

func (fake *FakeProductFileClient) RemoveFromReleaseCallCount() int {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return len(fake.removeFromReleaseArgsForCall)
}

func (fake *FakeProductFileClient) RemoveFromReleaseArgsForCall(i int) (string, string, int) {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return fake.removeFromReleaseArgsForCall[i].productSlug, fake.removeFromReleaseArgsForCall[i].releaseVersion, fake.removeFromReleaseArgsForCall[i].productFileID
}

func (fake *FakeProductFileClient) RemoveFromReleaseReturns(result1 error) {
	fake.RemoveFromReleaseStub = nil
	fake.removeFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) AddToFileGroup(productSlug string, fileGroupID int, productFileID int) error {
	fake.addToFileGroupMutex.Lock()
	fake.addToFileGroupArgsForCall = append(fake.addToFileGroupArgsForCall, struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}{productSlug, fileGroupID, productFileID})
	fake.recordInvocation("AddToFileGroup", []interface{}{productSlug, fileGroupID, productFileID})
	fake.addToFileGroupMutex.Unlock()
	if fake.AddToFileGroupStub != nil {
		return fake.AddToFileGroupStub(productSlug, fileGroupID, productFileID)
	} else {
		return fake.addToFileGroupReturns.result1
	}
}

func (fake *FakeProductFileClient) AddToFileGroupCallCount() int {
	fake.addToFileGroupMutex.RLock()
	defer fake.addToFileGroupMutex.RUnlock()
	return len(fake.addToFileGroupArgsForCall)
}

func (fake *FakeProductFileClient) AddToFileGroupArgsForCall(i int) (string, int, int) {
	fake.addToFileGroupMutex.RLock()
	defer fake.addToFileGroupMutex.RUnlock()
	return fake.addToFileGroupArgsForCall[i].productSlug, fake.addToFileGroupArgsForCall[i].fileGroupID, fake.addToFileGroupArgsForCall[i].productFileID
}

func (fake *FakeProductFileClient) AddToFileGroupReturns(result1 error) {
	fake.AddToFileGroupStub = nil
	fake.addToFileGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) RemoveFromFileGroup(productSlug string, fileGroupID int, productFileID int) error {
	fake.removeFromFileGroupMutex.Lock()
	fake.removeFromFileGroupArgsForCall = append(fake.removeFromFileGroupArgsForCall, struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}{productSlug, fileGroupID, productFileID})
	fake.recordInvocation("RemoveFromFileGroup", []interface{}{productSlug, fileGroupID, productFileID})
	fake.removeFromFileGroupMutex.Unlock()
	if fake.RemoveFromFileGroupStub != nil {
		return fake.RemoveFromFileGroupStub(productSlug, fileGroupID, productFileID)
	} else {
		return fake.removeFromFileGroupReturns.result1
	}
}

func (fake *FakeProductFileClient) RemoveFromFileGroupCallCount() int {
	fake.removeFromFileGroupMutex.RLock()
	defer fake.removeFromFileGroupMutex.RUnlock()
	return len(fake.removeFromFileGroupArgsForCall)
}

func (fake *FakeProductFileClient) RemoveFromFileGroupArgsForCall(i int) (string, int, int) {
	fake.removeFromFileGroupMutex.RLock()
	defer fake.removeFromFileGroupMutex.RUnlock()
	return fake.removeFromFileGroupArgsForCall[i].productSlug, fake.removeFromFileGroupArgsForCall[i].fileGroupID, fake.removeFromFileGroupArgsForCall[i].productFileID
}

func (fake *FakeProductFileClient) RemoveFromFileGroupReturns(result1 error) {
	fake.RemoveFromFileGroupStub = nil
	fake.removeFromFileGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) Delete(productSlug string, productFileID int) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		productSlug   string
		productFileID int
	}{productSlug, productFileID})
	fake.recordInvocation("Delete", []interface{}{productSlug, productFileID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(productSlug, productFileID)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeProductFileClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeProductFileClient) DeleteArgsForCall(i int) (string, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].productSlug, fake.deleteArgsForCall[i].productFileID
}

func (fake *FakeProductFileClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) Download(productSlug string, releaseVersion string, productFileID int, filepath string, acceptEULA bool) error {
	fake.downloadMutex.Lock()
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		productSlug    string
		releaseVersion string
		productFileID  int
		filepath       string
		acceptEULA     bool
	}{productSlug, releaseVersion, productFileID, filepath, acceptEULA})
	fake.recordInvocation("Download", []interface{}{productSlug, releaseVersion, productFileID, filepath, acceptEULA})
	fake.downloadMutex.Unlock()
	if fake.DownloadStub != nil {
		return fake.DownloadStub(productSlug, releaseVersion, productFileID, filepath, acceptEULA)
	} else {
		return fake.downloadReturns.result1
	}
}

func (fake *FakeProductFileClient) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeProductFileClient) DownloadArgsForCall(i int) (string, string, int, string, bool) {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return fake.downloadArgsForCall[i].productSlug, fake.downloadArgsForCall[i].releaseVersion, fake.downloadArgsForCall[i].productFileID, fake.downloadArgsForCall[i].filepath, fake.downloadArgsForCall[i].acceptEULA
}

func (fake *FakeProductFileClient) DownloadReturns(result1 error) {
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProductFileClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	fake.addToFileGroupMutex.RLock()
	defer fake.addToFileGroupMutex.RUnlock()
	fake.removeFromFileGroupMutex.RLock()
	defer fake.removeFromFileGroupMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProductFileClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.ProductFileClient = new(FakeProductFileClient)
