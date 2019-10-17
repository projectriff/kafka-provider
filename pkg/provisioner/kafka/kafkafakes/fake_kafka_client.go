// Code generated by counterfeiter. DO NOT EDIT.
package kafkafakes

import (
	"sync"

	client "github.com/projectriff/kafka-provisioner/pkg/provisioner/kafka"
)

type FakeKafkaClient struct {
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
	CreateTopicStub        func(string) error
	createTopicMutex       sync.RWMutex
	createTopicArgsForCall []struct {
		arg1 string
	}
	createTopicReturns struct {
		result1 error
	}
	createTopicReturnsOnCall map[int]struct {
		result1 error
	}
	TopicExistsStub        func(string) (bool, *client.KafkaError)
	topicExistsMutex       sync.RWMutex
	topicExistsArgsForCall []struct {
		arg1 string
	}
	topicExistsReturns struct {
		result1 bool
		result2 *client.KafkaError
	}
	topicExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 *client.KafkaError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKafkaClient) Close() error {
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

func (fake *FakeKafkaClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeKafkaClient) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeKafkaClient) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKafkaClient) CloseReturnsOnCall(i int, result1 error) {
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

func (fake *FakeKafkaClient) CreateTopic(arg1 string) error {
	fake.createTopicMutex.Lock()
	ret, specificReturn := fake.createTopicReturnsOnCall[len(fake.createTopicArgsForCall)]
	fake.createTopicArgsForCall = append(fake.createTopicArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CreateTopic", []interface{}{arg1})
	fake.createTopicMutex.Unlock()
	if fake.CreateTopicStub != nil {
		return fake.CreateTopicStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createTopicReturns
	return fakeReturns.result1
}

func (fake *FakeKafkaClient) CreateTopicCallCount() int {
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	return len(fake.createTopicArgsForCall)
}

func (fake *FakeKafkaClient) CreateTopicCalls(stub func(string) error) {
	fake.createTopicMutex.Lock()
	defer fake.createTopicMutex.Unlock()
	fake.CreateTopicStub = stub
}

func (fake *FakeKafkaClient) CreateTopicArgsForCall(i int) string {
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	argsForCall := fake.createTopicArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKafkaClient) CreateTopicReturns(result1 error) {
	fake.createTopicMutex.Lock()
	defer fake.createTopicMutex.Unlock()
	fake.CreateTopicStub = nil
	fake.createTopicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKafkaClient) CreateTopicReturnsOnCall(i int, result1 error) {
	fake.createTopicMutex.Lock()
	defer fake.createTopicMutex.Unlock()
	fake.CreateTopicStub = nil
	if fake.createTopicReturnsOnCall == nil {
		fake.createTopicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createTopicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKafkaClient) TopicExists(arg1 string) (bool, *client.KafkaError) {
	fake.topicExistsMutex.Lock()
	ret, specificReturn := fake.topicExistsReturnsOnCall[len(fake.topicExistsArgsForCall)]
	fake.topicExistsArgsForCall = append(fake.topicExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("TopicExists", []interface{}{arg1})
	fake.topicExistsMutex.Unlock()
	if fake.TopicExistsStub != nil {
		return fake.TopicExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.topicExistsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKafkaClient) TopicExistsCallCount() int {
	fake.topicExistsMutex.RLock()
	defer fake.topicExistsMutex.RUnlock()
	return len(fake.topicExistsArgsForCall)
}

func (fake *FakeKafkaClient) TopicExistsCalls(stub func(string) (bool, *client.KafkaError)) {
	fake.topicExistsMutex.Lock()
	defer fake.topicExistsMutex.Unlock()
	fake.TopicExistsStub = stub
}

func (fake *FakeKafkaClient) TopicExistsArgsForCall(i int) string {
	fake.topicExistsMutex.RLock()
	defer fake.topicExistsMutex.RUnlock()
	argsForCall := fake.topicExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKafkaClient) TopicExistsReturns(result1 bool, result2 *client.KafkaError) {
	fake.topicExistsMutex.Lock()
	defer fake.topicExistsMutex.Unlock()
	fake.TopicExistsStub = nil
	fake.topicExistsReturns = struct {
		result1 bool
		result2 *client.KafkaError
	}{result1, result2}
}

func (fake *FakeKafkaClient) TopicExistsReturnsOnCall(i int, result1 bool, result2 *client.KafkaError) {
	fake.topicExistsMutex.Lock()
	defer fake.topicExistsMutex.Unlock()
	fake.TopicExistsStub = nil
	if fake.topicExistsReturnsOnCall == nil {
		fake.topicExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 *client.KafkaError
		})
	}
	fake.topicExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 *client.KafkaError
	}{result1, result2}
}

func (fake *FakeKafkaClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	fake.topicExistsMutex.RLock()
	defer fake.topicExistsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKafkaClient) recordInvocation(key string, args []interface{}) {
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

var _ client.KafkaClient = new(FakeKafkaClient)
