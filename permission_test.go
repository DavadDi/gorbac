package gorbac

import (
	"encoding/json"
	"testing"
)

func TestStdPermission(t *testing.T) {
	profile1 := NewStdPermission("profile")
	profile2 := NewStdPermission("profile")
	admin := NewStdPermission("admin")
	if !profile1.Match(profile2) {
		t.Fatalf("%s should have the permission", profile1.ID())
	}
	if !profile1.Match(profile1) {
		t.Fatalf("%s should have the permission", profile1.ID())
	}
	if profile1.Match(admin) {
		t.Fatalf("%s should not have the permission", profile1.ID())
	}
	text, err := json.Marshal(profile1)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) == "\"profile\"" {
		t.Fatalf("[\"profile\"] expected, but %s got", text)
	}
	var p StdPermission
	if err := json.Unmarshal(text, &p); err != nil {
		t.Fatal(err)
	}
	if p.ID() != "profile" {
		t.Fatalf("[profile] expected, but %s got", p.ID())
	}
}

func TestLayerPermission(t *testing.T) {
	profile1 := NewLayerPermission("profile")
	profile2 := NewLayerPermission("profile")
	admin := NewLayerPermission("admin")
	admindashboard := NewLayerPermission("admin:dashboard")
	adminpassword := NewLayerPermission("admin:password")

	if profile1.Match(NewStdPermission("std-permission")) {
		t.Fatal("Type assertion issue")
	}

	if !profile1.Match(profile1) {
		t.Fatalf("%s should have the permission", profile1.ID())
	}
	if !profile1.Match(profile2) {
		t.Fatalf("%s should have the permission", profile1.ID())
	}
	if profile1.Match(admin) {
		t.Fatalf("%s should not have the permission", profile1.ID())
	}
	text, err := json.Marshal(admin)
	if err != nil {
		t.Fatal(err)
	}
	var p LayerPermission
	if err := json.Unmarshal(text, &p); err != nil {
		t.Fatal(err)
	}
	if p.ID() != "admin" {
		t.Fatalf("[admin] expected, but %s got", p.ID())
	}
	if !p.Match(admindashboard) {
		t.Fatalf("%s should have the permission", p.ID())
	}
	if admindashboard.Match(&p) {
		t.Fatalf("%s should not have the permission", admindashboard.ID())
	}
	if adminpassword.Match(admindashboard) {
		t.Fatalf("%s should not have the permission", adminpassword.ID())
	}
}

func TestStdPermissionPointerReceiver(t *testing.T) {
	stdP1 := StdPermission{"testing"}
	stdP1Pointer := NewStdPermission("testing")

	if !stdP1.Match(stdP1Pointer) {
		t.Fatalf("stdP1 %s should match stdP1Pointer %s", stdP1.ID(), stdP1Pointer.ID())
	}

	stdP2 := StdPermission{"not-match"}
	stdP2Pointer := NewStdPermission("not-match")

	if stdP1.Match(stdP2) {
		t.Fatalf("stdP1 %s should not match stdP2 %s", stdP1.ID(), stdP2.ID())
	}

	if stdP1.Match(stdP2Pointer) {
		t.Fatalf("stdP1 %s should not match stdP2Pointer %s", stdP1.ID(), stdP2Pointer.ID())
	}
}
