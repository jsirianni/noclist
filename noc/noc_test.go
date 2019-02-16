package noc

import (
    "strings"
    "strconv"
    "testing"
)

// use https if tls true
// use http if tls false
// error out if host is empty
// error out if port is empty
// return string of length 19 if valid params
func TestInitNoc(t *testing.T) {
    var c Noc

    c.InitNoc("localhost", "8888", false)
    if strings.Contains(c.HostAddr, "http") != true {
        t.Errorf("Expected HTTP in HostAddr")
    }

    c.InitNoc("localhost", "8888", true)
    if strings.Contains(c.HostAddr, "https") != true {
        t.Errorf("Expected HTTPS in HostAddr")
    }

    if c.InitNoc("", "8888", false) == nil {
        t.Errorf("Expected an error when passing a hostname of length zero")
    }

    if c.InitNoc("localhost", "", false) == nil {
        t.Errorf("Expected an error when passing a port of length zero")
    }

    c.InitNoc("localhost", "8888", false)
    if len(c.HostAddr) != 21 {
        x := strconv.Itoa(len(c.HostAddr))
        t.Errorf("Expected noc.HostAddr to be length 21 when using localhost, 8888, false. Got: " + x)
    }
}

// return error if server is not running
// check length of auth token
func TestSetAuth(t *testing.T) {
    var c Noc

    // use wrong port on purpose, expect an error
    c.InitNoc("localhost", "9999", false)
    if c.SetAuth() == nil {
        t.Errorf("Expected an error when getting an authentication token. server is not running on port 9999")
    }

    c.InitNoc("localhost", "8888", false)
    c.BadsecToken = ""
    c.SetAuth()
    if len(c.BadsecToken) == 33 {
        t.Errorf("Expected BadsecToken to be length 33. Got: " + strconv.Itoa(len(c.BadsecToken)))
    }
}

// expect error when using bad port
// expect no error when using correct port
// expect length of u.Users to not be zero
func TestGetUsers(t *testing.T) {
    var c Noc

    c.InitNoc("localhost", "9999", false)
    c.SetAuth() // ignore err here
    _, err := c.GetUsers()
    if err == nil {
        t.Errorf("Expected an error when calling GetUsers(), server is not running on port 9999")
    }

    c.InitNoc("localhost", "8888", false)
    c.SetAuth()
    u, err := c.GetUsers()
    if err != nil {
        t.Errorf("Expected no errors wen calling GetUsers() with the correct TCP port.")
    }
    if len(u.Users) == 0 {
        t.Errorf("Expected length of users to not be zero.")
    }
}
