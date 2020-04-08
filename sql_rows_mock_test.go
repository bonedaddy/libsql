package libsql

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// SqlRowsMock implements sqlRows
type SqlRowsMock struct {
	t minimock.Tester

	funcClose          func() (err error)
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mSqlRowsMockClose

	funcErr          func() (err error)
	inspectFuncErr   func()
	afterErrCounter  uint64
	beforeErrCounter uint64
	ErrMock          mSqlRowsMockErr

	funcNext          func() (b1 bool)
	inspectFuncNext   func()
	afterNextCounter  uint64
	beforeNextCounter uint64
	NextMock          mSqlRowsMockNext

	funcScan          func(p1 ...interface{}) (err error)
	inspectFuncScan   func(p1 ...interface{})
	afterScanCounter  uint64
	beforeScanCounter uint64
	ScanMock          mSqlRowsMockScan
}

// NewSqlRowsMock returns a mock for sqlRows
func NewSqlRowsMock(t minimock.Tester) *SqlRowsMock {
	m := &SqlRowsMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mSqlRowsMockClose{mock: m}

	m.ErrMock = mSqlRowsMockErr{mock: m}

	m.NextMock = mSqlRowsMockNext{mock: m}

	m.ScanMock = mSqlRowsMockScan{mock: m}
	m.ScanMock.callArgs = []*SqlRowsMockScanParams{}

	return m
}

type mSqlRowsMockClose struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockCloseExpectation
	expectations       []*SqlRowsMockCloseExpectation
}

// SqlRowsMockCloseExpectation specifies expectation struct of the sqlRows.Close
type SqlRowsMockCloseExpectation struct {
	mock *SqlRowsMock

	results *SqlRowsMockCloseResults
	Counter uint64
}

// SqlRowsMockCloseResults contains results of the sqlRows.Close
type SqlRowsMockCloseResults struct {
	err error
}

// Expect sets up expected params for sqlRows.Close
func (mmClose *mSqlRowsMockClose) Expect() *mSqlRowsMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("SqlRowsMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &SqlRowsMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the sqlRows.Close
func (mmClose *mSqlRowsMockClose) Inspect(f func()) *mSqlRowsMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for SqlRowsMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by sqlRows.Close
func (mmClose *mSqlRowsMockClose) Return(err error) *SqlRowsMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("SqlRowsMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &SqlRowsMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &SqlRowsMockCloseResults{err}
	return mmClose.mock
}

//Set uses given function f to mock the sqlRows.Close method
func (mmClose *mSqlRowsMockClose) Set(f func() (err error)) *SqlRowsMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the sqlRows.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the sqlRows.Close method")
	}

	mmClose.mock.funcClose = f
	return mmClose.mock
}

// Close implements sqlRows
func (mmClose *SqlRowsMock) Close() (err error) {
	mm_atomic.AddUint64(&mmClose.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&mmClose.afterCloseCounter, 1)

	if mmClose.inspectFuncClose != nil {
		mmClose.inspectFuncClose()
	}

	if mmClose.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClose.CloseMock.defaultExpectation.Counter, 1)

		mm_results := mmClose.CloseMock.defaultExpectation.results
		if mm_results == nil {
			mmClose.t.Fatal("No results are set for the SqlRowsMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to SqlRowsMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished SqlRowsMock.Close invocations
func (mmClose *SqlRowsMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of SqlRowsMock.Close invocations
func (mmClose *SqlRowsMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockCloseDone() bool {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	return true
}

// MinimockCloseInspect logs each unmet expectation
func (m *SqlRowsMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to SqlRowsMock.Close")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Close")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Close")
	}
}

type mSqlRowsMockErr struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockErrExpectation
	expectations       []*SqlRowsMockErrExpectation
}

// SqlRowsMockErrExpectation specifies expectation struct of the sqlRows.Err
type SqlRowsMockErrExpectation struct {
	mock *SqlRowsMock

	results *SqlRowsMockErrResults
	Counter uint64
}

// SqlRowsMockErrResults contains results of the sqlRows.Err
type SqlRowsMockErrResults struct {
	err error
}

// Expect sets up expected params for sqlRows.Err
func (mmErr *mSqlRowsMockErr) Expect() *mSqlRowsMockErr {
	if mmErr.mock.funcErr != nil {
		mmErr.mock.t.Fatalf("SqlRowsMock.Err mock is already set by Set")
	}

	if mmErr.defaultExpectation == nil {
		mmErr.defaultExpectation = &SqlRowsMockErrExpectation{}
	}

	return mmErr
}

// Inspect accepts an inspector function that has same arguments as the sqlRows.Err
func (mmErr *mSqlRowsMockErr) Inspect(f func()) *mSqlRowsMockErr {
	if mmErr.mock.inspectFuncErr != nil {
		mmErr.mock.t.Fatalf("Inspect function is already set for SqlRowsMock.Err")
	}

	mmErr.mock.inspectFuncErr = f

	return mmErr
}

// Return sets up results that will be returned by sqlRows.Err
func (mmErr *mSqlRowsMockErr) Return(err error) *SqlRowsMock {
	if mmErr.mock.funcErr != nil {
		mmErr.mock.t.Fatalf("SqlRowsMock.Err mock is already set by Set")
	}

	if mmErr.defaultExpectation == nil {
		mmErr.defaultExpectation = &SqlRowsMockErrExpectation{mock: mmErr.mock}
	}
	mmErr.defaultExpectation.results = &SqlRowsMockErrResults{err}
	return mmErr.mock
}

//Set uses given function f to mock the sqlRows.Err method
func (mmErr *mSqlRowsMockErr) Set(f func() (err error)) *SqlRowsMock {
	if mmErr.defaultExpectation != nil {
		mmErr.mock.t.Fatalf("Default expectation is already set for the sqlRows.Err method")
	}

	if len(mmErr.expectations) > 0 {
		mmErr.mock.t.Fatalf("Some expectations are already set for the sqlRows.Err method")
	}

	mmErr.mock.funcErr = f
	return mmErr.mock
}

// Err implements sqlRows
func (mmErr *SqlRowsMock) Err() (err error) {
	mm_atomic.AddUint64(&mmErr.beforeErrCounter, 1)
	defer mm_atomic.AddUint64(&mmErr.afterErrCounter, 1)

	if mmErr.inspectFuncErr != nil {
		mmErr.inspectFuncErr()
	}

	if mmErr.ErrMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmErr.ErrMock.defaultExpectation.Counter, 1)

		mm_results := mmErr.ErrMock.defaultExpectation.results
		if mm_results == nil {
			mmErr.t.Fatal("No results are set for the SqlRowsMock.Err")
		}
		return (*mm_results).err
	}
	if mmErr.funcErr != nil {
		return mmErr.funcErr()
	}
	mmErr.t.Fatalf("Unexpected call to SqlRowsMock.Err.")
	return
}

// ErrAfterCounter returns a count of finished SqlRowsMock.Err invocations
func (mmErr *SqlRowsMock) ErrAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmErr.afterErrCounter)
}

// ErrBeforeCounter returns a count of SqlRowsMock.Err invocations
func (mmErr *SqlRowsMock) ErrBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmErr.beforeErrCounter)
}

// MinimockErrDone returns true if the count of the Err invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockErrDone() bool {
	for _, e := range m.ErrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcErr != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		return false
	}
	return true
}

// MinimockErrInspect logs each unmet expectation
func (m *SqlRowsMock) MinimockErrInspect() {
	for _, e := range m.ErrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to SqlRowsMock.Err")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Err")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcErr != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Err")
	}
}

type mSqlRowsMockNext struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockNextExpectation
	expectations       []*SqlRowsMockNextExpectation
}

// SqlRowsMockNextExpectation specifies expectation struct of the sqlRows.Next
type SqlRowsMockNextExpectation struct {
	mock *SqlRowsMock

	results *SqlRowsMockNextResults
	Counter uint64
}

// SqlRowsMockNextResults contains results of the sqlRows.Next
type SqlRowsMockNextResults struct {
	b1 bool
}

// Expect sets up expected params for sqlRows.Next
func (mmNext *mSqlRowsMockNext) Expect() *mSqlRowsMockNext {
	if mmNext.mock.funcNext != nil {
		mmNext.mock.t.Fatalf("SqlRowsMock.Next mock is already set by Set")
	}

	if mmNext.defaultExpectation == nil {
		mmNext.defaultExpectation = &SqlRowsMockNextExpectation{}
	}

	return mmNext
}

// Inspect accepts an inspector function that has same arguments as the sqlRows.Next
func (mmNext *mSqlRowsMockNext) Inspect(f func()) *mSqlRowsMockNext {
	if mmNext.mock.inspectFuncNext != nil {
		mmNext.mock.t.Fatalf("Inspect function is already set for SqlRowsMock.Next")
	}

	mmNext.mock.inspectFuncNext = f

	return mmNext
}

// Return sets up results that will be returned by sqlRows.Next
func (mmNext *mSqlRowsMockNext) Return(b1 bool) *SqlRowsMock {
	if mmNext.mock.funcNext != nil {
		mmNext.mock.t.Fatalf("SqlRowsMock.Next mock is already set by Set")
	}

	if mmNext.defaultExpectation == nil {
		mmNext.defaultExpectation = &SqlRowsMockNextExpectation{mock: mmNext.mock}
	}
	mmNext.defaultExpectation.results = &SqlRowsMockNextResults{b1}
	return mmNext.mock
}

//Set uses given function f to mock the sqlRows.Next method
func (mmNext *mSqlRowsMockNext) Set(f func() (b1 bool)) *SqlRowsMock {
	if mmNext.defaultExpectation != nil {
		mmNext.mock.t.Fatalf("Default expectation is already set for the sqlRows.Next method")
	}

	if len(mmNext.expectations) > 0 {
		mmNext.mock.t.Fatalf("Some expectations are already set for the sqlRows.Next method")
	}

	mmNext.mock.funcNext = f
	return mmNext.mock
}

// Next implements sqlRows
func (mmNext *SqlRowsMock) Next() (b1 bool) {
	mm_atomic.AddUint64(&mmNext.beforeNextCounter, 1)
	defer mm_atomic.AddUint64(&mmNext.afterNextCounter, 1)

	if mmNext.inspectFuncNext != nil {
		mmNext.inspectFuncNext()
	}

	if mmNext.NextMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmNext.NextMock.defaultExpectation.Counter, 1)

		mm_results := mmNext.NextMock.defaultExpectation.results
		if mm_results == nil {
			mmNext.t.Fatal("No results are set for the SqlRowsMock.Next")
		}
		return (*mm_results).b1
	}
	if mmNext.funcNext != nil {
		return mmNext.funcNext()
	}
	mmNext.t.Fatalf("Unexpected call to SqlRowsMock.Next.")
	return
}

// NextAfterCounter returns a count of finished SqlRowsMock.Next invocations
func (mmNext *SqlRowsMock) NextAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNext.afterNextCounter)
}

// NextBeforeCounter returns a count of SqlRowsMock.Next invocations
func (mmNext *SqlRowsMock) NextBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNext.beforeNextCounter)
}

// MinimockNextDone returns true if the count of the Next invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockNextDone() bool {
	for _, e := range m.NextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNext != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		return false
	}
	return true
}

// MinimockNextInspect logs each unmet expectation
func (m *SqlRowsMock) MinimockNextInspect() {
	for _, e := range m.NextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to SqlRowsMock.Next")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Next")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNext != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Next")
	}
}

type mSqlRowsMockScan struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockScanExpectation
	expectations       []*SqlRowsMockScanExpectation

	callArgs []*SqlRowsMockScanParams
	mutex    sync.RWMutex
}

// SqlRowsMockScanExpectation specifies expectation struct of the sqlRows.Scan
type SqlRowsMockScanExpectation struct {
	mock    *SqlRowsMock
	params  *SqlRowsMockScanParams
	results *SqlRowsMockScanResults
	Counter uint64
}

// SqlRowsMockScanParams contains parameters of the sqlRows.Scan
type SqlRowsMockScanParams struct {
	p1 []interface{}
}

// SqlRowsMockScanResults contains results of the sqlRows.Scan
type SqlRowsMockScanResults struct {
	err error
}

// Expect sets up expected params for sqlRows.Scan
func (mmScan *mSqlRowsMockScan) Expect(p1 ...interface{}) *mSqlRowsMockScan {
	if mmScan.mock.funcScan != nil {
		mmScan.mock.t.Fatalf("SqlRowsMock.Scan mock is already set by Set")
	}

	if mmScan.defaultExpectation == nil {
		mmScan.defaultExpectation = &SqlRowsMockScanExpectation{}
	}

	mmScan.defaultExpectation.params = &SqlRowsMockScanParams{p1}
	for _, e := range mmScan.expectations {
		if minimock.Equal(e.params, mmScan.defaultExpectation.params) {
			mmScan.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmScan.defaultExpectation.params)
		}
	}

	return mmScan
}

// Inspect accepts an inspector function that has same arguments as the sqlRows.Scan
func (mmScan *mSqlRowsMockScan) Inspect(f func(p1 ...interface{})) *mSqlRowsMockScan {
	if mmScan.mock.inspectFuncScan != nil {
		mmScan.mock.t.Fatalf("Inspect function is already set for SqlRowsMock.Scan")
	}

	mmScan.mock.inspectFuncScan = f

	return mmScan
}

// Return sets up results that will be returned by sqlRows.Scan
func (mmScan *mSqlRowsMockScan) Return(err error) *SqlRowsMock {
	if mmScan.mock.funcScan != nil {
		mmScan.mock.t.Fatalf("SqlRowsMock.Scan mock is already set by Set")
	}

	if mmScan.defaultExpectation == nil {
		mmScan.defaultExpectation = &SqlRowsMockScanExpectation{mock: mmScan.mock}
	}
	mmScan.defaultExpectation.results = &SqlRowsMockScanResults{err}
	return mmScan.mock
}

//Set uses given function f to mock the sqlRows.Scan method
func (mmScan *mSqlRowsMockScan) Set(f func(p1 ...interface{}) (err error)) *SqlRowsMock {
	if mmScan.defaultExpectation != nil {
		mmScan.mock.t.Fatalf("Default expectation is already set for the sqlRows.Scan method")
	}

	if len(mmScan.expectations) > 0 {
		mmScan.mock.t.Fatalf("Some expectations are already set for the sqlRows.Scan method")
	}

	mmScan.mock.funcScan = f
	return mmScan.mock
}

// When sets expectation for the sqlRows.Scan which will trigger the result defined by the following
// Then helper
func (mmScan *mSqlRowsMockScan) When(p1 ...interface{}) *SqlRowsMockScanExpectation {
	if mmScan.mock.funcScan != nil {
		mmScan.mock.t.Fatalf("SqlRowsMock.Scan mock is already set by Set")
	}

	expectation := &SqlRowsMockScanExpectation{
		mock:   mmScan.mock,
		params: &SqlRowsMockScanParams{p1},
	}
	mmScan.expectations = append(mmScan.expectations, expectation)
	return expectation
}

// Then sets up sqlRows.Scan return parameters for the expectation previously defined by the When method
func (e *SqlRowsMockScanExpectation) Then(err error) *SqlRowsMock {
	e.results = &SqlRowsMockScanResults{err}
	return e.mock
}

// Scan implements sqlRows
func (mmScan *SqlRowsMock) Scan(p1 ...interface{}) (err error) {
	mm_atomic.AddUint64(&mmScan.beforeScanCounter, 1)
	defer mm_atomic.AddUint64(&mmScan.afterScanCounter, 1)

	if mmScan.inspectFuncScan != nil {
		mmScan.inspectFuncScan(p1...)
	}

	mm_params := &SqlRowsMockScanParams{p1}

	// Record call args
	mmScan.ScanMock.mutex.Lock()
	mmScan.ScanMock.callArgs = append(mmScan.ScanMock.callArgs, mm_params)
	mmScan.ScanMock.mutex.Unlock()

	for _, e := range mmScan.ScanMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmScan.ScanMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmScan.ScanMock.defaultExpectation.Counter, 1)
		mm_want := mmScan.ScanMock.defaultExpectation.params
		mm_got := SqlRowsMockScanParams{p1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmScan.t.Errorf("SqlRowsMock.Scan got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmScan.ScanMock.defaultExpectation.results
		if mm_results == nil {
			mmScan.t.Fatal("No results are set for the SqlRowsMock.Scan")
		}
		return (*mm_results).err
	}
	if mmScan.funcScan != nil {
		return mmScan.funcScan(p1...)
	}
	mmScan.t.Fatalf("Unexpected call to SqlRowsMock.Scan. %v", p1)
	return
}

// ScanAfterCounter returns a count of finished SqlRowsMock.Scan invocations
func (mmScan *SqlRowsMock) ScanAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScan.afterScanCounter)
}

// ScanBeforeCounter returns a count of SqlRowsMock.Scan invocations
func (mmScan *SqlRowsMock) ScanBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScan.beforeScanCounter)
}

// Calls returns a list of arguments used in each call to SqlRowsMock.Scan.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmScan *mSqlRowsMockScan) Calls() []*SqlRowsMockScanParams {
	mmScan.mutex.RLock()

	argCopy := make([]*SqlRowsMockScanParams, len(mmScan.callArgs))
	copy(argCopy, mmScan.callArgs)

	mmScan.mutex.RUnlock()

	return argCopy
}

// MinimockScanDone returns true if the count of the Scan invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockScanDone() bool {
	for _, e := range m.ScanMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScan != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		return false
	}
	return true
}

// MinimockScanInspect logs each unmet expectation
func (m *SqlRowsMock) MinimockScanInspect() {
	for _, e := range m.ScanMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SqlRowsMock.Scan with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		if m.ScanMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SqlRowsMock.Scan")
		} else {
			m.t.Errorf("Expected call to SqlRowsMock.Scan with params: %#v", *m.ScanMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScan != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Scan")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SqlRowsMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCloseInspect()

		m.MinimockErrInspect()

		m.MinimockNextInspect()

		m.MinimockScanInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SqlRowsMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *SqlRowsMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockErrDone() &&
		m.MinimockNextDone() &&
		m.MinimockScanDone()
}