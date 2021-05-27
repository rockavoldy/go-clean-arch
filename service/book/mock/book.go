// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_book is a generated GoMock package.
package mock_book

import (
	entity "go-clean-arch/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockReader) Get(id entity.ID) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReader)(nil).Get), id)
}

// List mocks base method.
func (m *MockReader) List() ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReaderMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReader)(nil).List))
}

// Search mocks base method.
func (m *MockReader) Search(query string) ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockReaderMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockReader)(nil).Search), query)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWriter) Create(e *entity.Book) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWriterMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriter)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockWriter) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWriterMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriter)(nil).Delete), id)
}

// Restore mocks base method.
func (m *MockWriter) Restore(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockWriterMockRecorder) Restore(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockWriter)(nil).Restore), id)
}

// Update mocks base method.
func (m *MockWriter) Update(e *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWriterMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWriter)(nil).Update), e)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(e *entity.Book) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockRepository) Get(id entity.ID) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// List mocks base method.
func (m *MockRepository) List() ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Restore mocks base method.
func (m *MockRepository) Restore(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockRepositoryMockRecorder) Restore(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockRepository)(nil).Restore), id)
}

// Search mocks base method.
func (m *MockRepository) Search(query string) ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// Update mocks base method.
func (m *MockRepository) Update(e *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), e)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockUseCase) CreateBook(title, author, isbn string, pages, qty int) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", title, author, isbn, pages, qty)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockUseCaseMockRecorder) CreateBook(title, author, isbn, pages, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockUseCase)(nil).CreateBook), title, author, isbn, pages, qty)
}

// DeleteBook mocks base method.
func (m *MockUseCase) DeleteBook(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockUseCaseMockRecorder) DeleteBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockUseCase)(nil).DeleteBook), id)
}

// GetBook mocks base method.
func (m *MockUseCase) GetBook(id entity.ID) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", id)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockUseCaseMockRecorder) GetBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockUseCase)(nil).GetBook), id)
}

// ListBooks mocks base method.
func (m *MockUseCase) ListBooks() ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBooks")
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBooks indicates an expected call of ListBooks.
func (mr *MockUseCaseMockRecorder) ListBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBooks", reflect.TypeOf((*MockUseCase)(nil).ListBooks))
}

// RestoreBook mocks base method.
func (m *MockUseCase) RestoreBook(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreBook", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreBook indicates an expected call of RestoreBook.
func (mr *MockUseCaseMockRecorder) RestoreBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreBook", reflect.TypeOf((*MockUseCase)(nil).RestoreBook), id)
}

// SearchBooks mocks base method.
func (m *MockUseCase) SearchBooks(query string) ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchBooks", query)
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchBooks indicates an expected call of SearchBooks.
func (mr *MockUseCaseMockRecorder) SearchBooks(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchBooks", reflect.TypeOf((*MockUseCase)(nil).SearchBooks), query)
}

// UpdateBook mocks base method.
func (m *MockUseCase) UpdateBook(e *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockUseCaseMockRecorder) UpdateBook(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockUseCase)(nil).UpdateBook), e)
}