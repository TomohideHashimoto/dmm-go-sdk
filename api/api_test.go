package api

import (
    "reflect"
    "testing"
    // "github.com/davecgh/go-spew/spew"
)

// func TestFoo(t *testing.T) {
//     testUrl := "https://api.dmm.com/affiliate/v3/itemList?affiliate_id=10278-999&api_id=UrwskPfkqQ0DuVry2gYL&site=DMM.R18"
//     actual, _ := RequestJson(testUrl)
//     spew.Dump(actual)
// }

func TestRequestJson(t *testing.T) {
    testUrl := "https://httpbin.org/get?foo=bar"
    actual, _ := RequestJson(testUrl)
    
    if actual == nil {
        t.Fatalf("response is not expected empty.")
    }

    if reflect.TypeOf(actual).String() != "map[string]interface {}" {
        t.Fatalf("response is expected inteface{}. but actual %s", reflect.TypeOf(actual).String())
    }
}

func TestValidateAffiliateId(t *testing.T) {
    val := "vcder56yuhnmkiuy-990";
    if !ValidateAffiliateId(val) {
        t.Fatalf("When value is %s, not expected false.", val)
    }

    val = "vcder56yuhnmkiuy-9";
    if ValidateAffiliateId(val) {
        t.Fatalf("When value is %s, not expected true.", val)
    }

    val = "vcder56yuhnmkiuy-";
    if ValidateAffiliateId(val) {
        t.Fatalf("When value is %s, not expected true.", val)
    }

    val = "-999";
    if ValidateAffiliateId(val) {
        t.Fatalf("When value is %s, not expected true.", val)
    }

    if ValidateAffiliateId("") {
        t.Fatalf("When value is empty, not expected true.")
    }
}

func TestValidateSite(t *testing.T) {
    if !ValidateSite(SITE_ALLAGES) {
        t.Fatalf("When value is %s, not expected false.", SITE_ALLAGES)
    }

    if !ValidateSite(SITE_ADULT) {
        t.Fatalf("When value is %s, not expected false.", SITE_ADULT)
    }

    if !ValidateSite("DMM.com") {
        t.Fatalf("When value is %s, not expected false.", "DMM.com")
    }

    if !ValidateSite("DMM.R18") {
        t.Fatalf("When value is %s, not expected false.", "DMM.R18")
    }

    if ValidateSite("DMM.co.jp") {
        t.Fatalf("When value is %s, not expected true.", "DMM.co.jp")
    }

    if ValidateSite("") {
        t.Fatalf("When value is empty, not expected true.")
    }
}

func TestValidateRange(t *testing.T) {
    var target int64
    var min    int64
    var max    int64
    
    target = 10
    min    = 1
    max    = 100
    if !ValidateRange(target, min, max) {
        t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
    }

    target = 1
    if !ValidateRange(target, min, max) {
        t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
    }

    target = 100
    if !ValidateRange(target, min, max) {
        t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
    }

    target = 0
    if ValidateRange(target, min, max) {
        t.Fatalf("When target value is %d, min is %d and max is %d, not expected true.", target, min, max)
    }

    target = 101
    if ValidateRange(target, min, max) {
        t.Fatalf("When target value is %d, min is %d and max is %d, not expected true.", target, min, max)
    }

    target = 10
    min    = 10
    max    = 10
    if !ValidateRange(target, min, max) {
        t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
    }    
}

func TestGetApiVersion(t *testing.T) {
    if GetApiVersion() != API_VERSION {
        t.Errorf("This value is expected to equal API_VERSION. actual:%s", GetApiVersion())
    }
}