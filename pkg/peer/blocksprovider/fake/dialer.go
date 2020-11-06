// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"crypto/x509"
	"sync"

	"github.com/Hnampk/fabric-usable-internal/pkg/peer/blocksprovider"
	"google.golang.org/grpc"
)

type Dialer struct {
	DialStub        func(string, *x509.CertPool) (*grpc.ClientConn, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
		arg1 string
		arg2 *x509.CertPool
	}
	dialReturns struct {
		result1 *grpc.ClientConn
		result2 error
	}
	dialReturnsOnCall map[int]struct {
		result1 *grpc.ClientConn
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Dialer) Dial(arg1 string, arg2 *x509.CertPool) (*grpc.ClientConn, error) {
	fake.dialMutex.Lock()
	ret, specificReturn := fake.dialReturnsOnCall[len(fake.dialArgsForCall)]
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
		arg1 string
		arg2 *x509.CertPool
	}{arg1, arg2})
	fake.recordInvocation("Dial", []interface{}{arg1, arg2})
	fake.dialMutex.Unlock()
	if fake.DialStub != nil {
		return fake.DialStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.dialReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Dialer) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *Dialer) DialCalls(stub func(string, *x509.CertPool) (*grpc.ClientConn, error)) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = stub
}

func (fake *Dialer) DialArgsForCall(i int) (string, *x509.CertPool) {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	argsForCall := fake.dialArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Dialer) DialReturns(result1 *grpc.ClientConn, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 *grpc.ClientConn
		result2 error
	}{result1, result2}
}

func (fake *Dialer) DialReturnsOnCall(i int, result1 *grpc.ClientConn, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	if fake.dialReturnsOnCall == nil {
		fake.dialReturnsOnCall = make(map[int]struct {
			result1 *grpc.ClientConn
			result2 error
		})
	}
	fake.dialReturnsOnCall[i] = struct {
		result1 *grpc.ClientConn
		result2 error
	}{result1, result2}
}

func (fake *Dialer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Dialer) recordInvocation(key string, args []interface{}) {
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

var _ blocksprovider.Dialer = new(Dialer)
