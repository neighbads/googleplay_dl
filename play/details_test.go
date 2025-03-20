package play

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDetails(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	home = filepath.ToSlash(home) + "/google/play"
	data, err := os.ReadFile(home + "/Token")
	if err != nil {
		t.Fatal(err)
	}
	var token1 Token
	err = token1.Unmarshal(data)
	if err != nil {
		t.Fatal(err)
	}
	auth1, err := token1.Auth()
	if err != nil {
		t.Fatal(err)
	}
	for _, app1 := range apps {
		data, err = os.ReadFile(home + "/" + app1.abi)
		if err != nil {
			t.Fatal(err)
		}
		var check Checkin
		err = check.Unmarshal(data)
		if err != nil {
			t.Fatal(err)
		}
		detail, err := auth1.Details(check, app1.id, false)
		if err != nil {
			t.Fatal(err)
		}
		if detail.Downloads() == 0 {
			t.Fatal("downloads")
		}
		if detail.Name() == "" {
			t.Fatal("name")
		}
		var ok bool
		for data := range detail[0].Get(8) {
			for range data.GetVarint(1) {
				ok = true
			}
		}
		if !ok {
			t.Fatal("field 8 1")
		}
		if detail.field_8_2() == "" {
			t.Fatal("field 8 2")
		}
		if detail.field_13_1_4() == "" {
			t.Fatal("field 13 1 4")
		}
		app1.date = func() string {
			time1, err := time.Parse("Jan 2, 2006", detail.field_13_1_16())
			if err != nil {
				t.Fatal(err)
			}
			return time1.Format("2006-01-02")
		}()
		ok = false
		for range detail.field_13_1_17() {
			ok = true
		}
		if !ok {
			t.Fatal("field 13 1 17")
		}
		if detail.field_13_1_82_1_1() == "" {
			t.Fatal("field 13 1 82 1 1")
		}
		if detail.field_15_18() == "" {
			t.Fatal("field_15_18")
		}
		if detail.size() == 0 {
			t.Fatal("size")
		}
		if detail.version_code() == 0 {
			t.Fatal("version code")
		}
		fmt.Printf("%#v,\n", app1)
		time.Sleep(99 * time.Millisecond)
	}
}

type app_test struct {
	date string
	abi  string
	id   string
}

var apps = []app_test{
	{date: "2024-08-02", abi: "x86", id: "com.wakdev.wdnfc"},
	{date: "2024-12-06", abi: "armeabi-v7a", id: "com.sygic.aura"},
	{date: "2024-12-10", abi: "x86", id: "com.clearchannel.iheartradio.controller"},
	{date: "2024-12-16", abi: "x86", id: "com.amctve.amcfullepisodes"},
	{date: "2024-12-24", abi: "armeabi-v7a", id: "com.xiaomi.smarthome"},
	{date: "2025-01-10", abi: "x86", id: "app.source.getcontact"},
	{date: "2025-01-13", abi: "armeabi-v7a", id: "com.binance.dev"},
	{date: "2025-01-15", abi: "armeabi-v7a-leanback", id: "com.netflix.ninja"},
	{date: "2025-01-16", abi: "arm64-v8a", id: "com.kakaogames.twodin"},
	{date: "2025-01-16", abi: "armeabi-v7a", id: "com.axis.drawingdesk.v3"},
	{date: "2025-01-17", abi: "x86", id: "com.google.android.apps.walletnfcrel"},
	{date: "2025-01-17", abi: "x86", id: "com.pinterest"},
	{date: "2025-01-17", abi: "x86-leanback", id: "com.roku.web.trc"},
	{date: "2025-01-18", abi: "x86", id: "kr.sira.metal"},
	{date: "2025-01-20", abi: "x86", id: "com.instagram.android"},
	{date: "2025-01-21", abi: "arm64-v8a", id: "com.app.xt"},
	{date: "2025-01-21", abi: "x86", id: "com.busuu.android.enc"},
	{date: "2025-01-21", abi: "x86", id: "org.thoughtcrime.securesms"},
	{date: "2025-01-22", abi: "armeabi-v7a", id: "com.madhead.tos.zh"},
	{date: "2025-01-22", abi: "x86", id: "br.com.rodrigokolb.realdrum"},
	{date: "2025-01-22", abi: "x86", id: "com.cabify.rider"},
	{date: "2025-01-22", abi: "x86", id: "com.google.android.youtube"},
}
