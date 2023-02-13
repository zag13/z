package timez

import (
	"testing"
	"time"
)

func TestGetDiffDays(t *testing.T) {
	t1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	t2 := time.Date(2023, 1, 1, 3, 0, 0, 0, time.Local)
	if GetDiffDays(t1, t2) != 0 {
		t.Errorf("GetDiffDays(%v, %v) != 0", t1, t2)
	}

	t1 = time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	t2 = time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)
	if GetDiffDays(t1, t2) != -1 {
		t.Errorf("GetDiffDays(%v, %v) != -1", t1, t2)
	}

	t1 = time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)
	t2 = time.Date(2023, 1, 1, 3, 0, 0, 0, time.Local)
	if GetDiffDays(t1, t2) != 1 {
		t.Errorf("GetDiffDays(%v, %v) != 1", t1, t2)
	}
}

func TestGetDiffDaysBySecond(t *testing.T) {
	t1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	t2 := time.Date(2023, 1, 1, 3, 0, 0, 0, time.Local)
	if GetDiffDaysBySecond(t1.Unix(), t2.Unix()) != 0 {
		t.Errorf("GetDiffDaysBySecond(%v, %v) != 0", t1, t2)
	}

	t1 = time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	t2 = time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)
	if GetDiffDaysBySecond(t1.Unix(), t2.Unix()) != -1 {
		t.Errorf("GetDiffDaysBySecond(%v, %v) != -1", t1, t2)
	}

	t1 = time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)
	t2 = time.Date(2023, 1, 1, 3, 0, 0, 0, time.Local)
	if GetDiffDaysBySecond(t1.Unix(), t2.Unix()) != 1 {
		t.Errorf("GetDiffDaysBySecond(%v, %v) != 1", t1, t2)
	}
}
