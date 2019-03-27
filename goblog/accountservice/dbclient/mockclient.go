package dbclient

import (
	"github.com/callistanenterprise/goblog/accountservice/model"
	"github.com/stretchr/testify/mock"
)

type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
	args := m.Mock.Called(accountId)
	return args.Get(0).(model.Account), args.Error(1)
}
func (m *MockBoltClient) OpenBoltDb() {

}
func (m *MockBoltClient) Seed() {

}
