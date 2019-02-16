package noc

import (
    "os"
    "fmt"
    "errors"
    "strconv"
    "strings"
    "net/http"
    "io/ioutil"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
)

// Noc struct represents the Noc client
type Noc struct {
    HostAddr string
    BadsecToken string
}

// Users struct represents the /users api response
type Users struct {
    Users []string
}

// InitNoc initilizes the Noc client
func (n *Noc) InitNoc(host string, port string, tls bool) error {
    if len(host) == 0 {
        return errors.New("hostname of length zero passed for 'host' in InitNoc()")
    }
    if len(port) == 0 {
        return errors.New("port of length zero passed for 'port' in InitNoc()")
    }

    if tls == true {
        n.HostAddr = "https://" + host + ":" + port
    } else {
        n.HostAddr = "http://" + host + ":" + port
    }

    return nil
}

// GetAuth sets the authentication token
func (n *Noc) SetAuth() error {
    var err error
    var r *http.Response

    // try up to three times
    for i := 0;  i<=2; i++ {
        r, err = http.Get(n.HostAddr + "/auth")
        if err != nil {
            //err = errors.New("Failed to get auth from BadSec server, is it running?")
            fmt.Fprintln(os.Stderr, err.Error())
        } else {
            defer r.Body.Close()
            if r.StatusCode == 200 {
                n.BadsecToken = r.Header.Get("Badsec-Authentication-token")
                if len(n.BadsecToken) != 36 {
                    msg := "Expected length of Badsec-Authentication-token was not correct."
                    msg = msg + "Expected 65, got " + strconv.Itoa(len(n.BadsecToken))
                    msg = msg + "\n" + n.BadsecToken
                    err = errors.New(msg)
                }

                // token is set. set err to nil, and return.
                err = nil
                return err
            }
            fmt.Fprintln(os.Stderr, i)
            fmt.Fprintln(os.Stderr, err.Error())
        }
    }

    // return the last error
    return err
}

// GetUsers returns a Users struct
func (n *Noc) GetUsers() (Users, error) {
    var u   Users
    var err error
    var req      *http.Request
    var resp     *http.Response
    var contents []byte

    if len(n.HostAddr) == 0 {
        return u, errors.New("HostAddr was not set properly. Was InitNoc() called?")
    }

    if len(n.BadsecToken) == 0 {
        return u, errors.New("BadsecToken was not set properly. Was InitNoc() called?")
    }

    // try up to three times
    for i := 0; i<=2; i++ {
        // build the client and request with headers
        client := &http.Client{}
        checksum := getRequestChecksum(n.BadsecToken, "/users")
        req, err = http.NewRequest("GET", n.HostAddr + "/users", nil)
        if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
        } else {
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set("Accept", "application/json")
            req.Header.Set("X-Request-Checksum", checksum)

            // perform the request and read the raw response body
            resp, err = client.Do(req)
            defer resp.Body.Close()
            contents, err = ioutil.ReadAll(resp.Body)

            // return on 200, else loop
            if resp.StatusCode == 200 {
                u.Users = strings.Split(string(contents), "\n")
                return u, nil
            } else {
                msg := "Server returned status: " + strconv.Itoa(resp.StatusCode) + "\n"
                msg = msg + string(contents)
                err = errors.New(msg)
            }
            fmt.Fprintln(os.Stderr, i)
            fmt.Fprintln(os.Stderr, err.Error())
        }

    }
    return u, err
}

// ToJson returns a Users struct as a json string
func (u *Users) ToJson() ([]byte, error) {
    j, err := json.Marshal(u.Users)
    if err != nil {
        return nil, err
    }
    return j, nil
}

// getRequestChecksum returns the Badsec-Authentication-Token checksum for
// a given api path
func getRequestChecksum(token string, path string) string {
    sum := sha256.New()
    sum.Write([]byte(token + path))

    // encode the sum to a string and return it
    return hex.EncodeToString(sum.Sum(nil))
}
