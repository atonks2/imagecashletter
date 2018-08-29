// Copyright 2018 The X9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

import (
	"testing"
	"time"
)

// mockCashLetterHeader creates a CashLetterHeader
func mockCashLetterHeader() *CashLetterHeader {
	clh := NewCashLetterHeader()
	clh.CollectionTypeIndicator = "01"
	clh.DestinationRoutingNumber = "231380104"
	clh.ECEInstitutionRoutingNumber = "121042882"
	clh.CashLetterBusinessDate = time.Now()
	clh.CashLetterCreationDate = time.Now()
	clh.CashLetterCreationTime = time.Now()
	clh.CashLetterRecordTypeIndicator = "N"
	clh.CashLetterDocumentationTypeIndicator = "A"
	clh.OriginatorContactName = "Contact Name"
	clh.OriginatorContactPhoneNumber = "5558675552"
	clh.CashLetterID = "A1"
	clh.FedWorkType = ""
	clh.ReturnsIndicator = ""
	clh.UserField = ""
	return clh
}

// testMockCashLetterHeader creates an ICL CashLetterHeader
func testMockCashLetterHeader(t testing.TB) {
	clh := mockCashLetterHeader()
	if err := clh.Validate(); err != nil {
		t.Error("mockCashLetterHeader does not validate and will break other tests: ", err)
	}
	if clh.recordType != "10" {
		t.Error("recordType does not validate and will break other tests")
	}
	if clh.CollectionTypeIndicator != "01" {
		t.Error("CollectionTypeIndicator does not validate and will break other tests")
	}
	if clh.DestinationRoutingNumber != "231380104" {
		t.Error("DestinationRoutingNumber does not validate and will break other tests")
	}
	if clh.ECEInstitutionRoutingNumber != "121042882" {
		t.Error("ECEInstitutionRoutingNumber does not validate and will break other tests")
	}
	if clh.CashLetterRecordTypeIndicator != "N" {
		t.Error("RecordTypeIndicator does not validate and will break other tests")
	}
	if clh.CashLetterDocumentationTypeIndicator != "A" {
		t.Error("DocumentationTypeIndicator does not validate and will break other tests")
	}
	if clh.OriginatorContactName != "Contact Name" {
		t.Error("OriginatorContactName does not validate and will break other tests")
	}
	if clh.OriginatorContactPhoneNumber != "5558675552" {
		t.Error("OriginatorContactPhoneNumber does not validate and will break other tests")
	}
	if clh.FedWorkType != "" {
		t.Error("FedWorkType does not validate and will break other tests")
	}
	if clh.ReturnsIndicator != "" {
		t.Error("ReturnsIndicator does not validate and will break other tests")
	}
	if clh.UserField != "" {
		t.Error("UserField does not validate and will break other tests")
	}
}

// TestMockCashLetterHeader tests creating an ICL CashLetterHeader
func TestMockCashLetterHeader(t *testing.T) {
	testMockCashLetterHeader(t)
}

// BenchmarkMockCashLetterHeader benchmarks creating an ICL CashLetterHeader
func BenchmarkMockCashLetterHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockCashLetterHeader(b)
	}
}
