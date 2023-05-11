package badgeify

import "testing"

func TestLookupCoverage(t *testing.T) {
	cov, err := LookupCoverage()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if cov != 98.59 {
		t.Errorf("Expected coverage of 98.59, got %v", cov)
	}
}
