// +build integration

package mysql

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/config"

	baloo "gopkg.in/h2non/baloo.v3"
)

var apiObj *API
var apiTester *baloo.Client

func init() {
	var err error
	apiObj, err = NewAPI(*config.Values.MySQLConfig, nil)
	if err != nil {
		panic(err)
	}

	resetDB(apiObj)

	httpTester := httptest.NewServer(apiObj.getHandler())
	apiTester = baloo.New(httpTester.URL)
}

func resetDB(api *API) {
	dat, err := ioutil.ReadFile("../../../../db/schema.sql")
	if err != nil {
		panic(err)
	}

	requests := strings.Split(string(dat), ";")

	for _, request := range requests {
		if len(strings.TrimSpace(request)) == 0 {
			continue
		}
		_, err = api.db.Exec(request)
		if err != nil {
			fmt.Println("-----Probably not pointing to an alive DB. Make sure all environment variables are set right-----")
			panic(err)
		}
	}
}
