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

var dbSchemaCommands []string

func init() {
	var err error
	apiObj, err = NewAPI(*config.Values.MySQLConfig, nil)
	if err != nil {
		panic(err)
	}

	dbSchemaCommands = readDBSchema()

	resetDB()

	httpTester := httptest.NewServer(apiObj.getHandler())
	apiTester = baloo.New(httpTester.URL)
}

func resetDB() {
	for _, request := range dbSchemaCommands {
		if len(strings.TrimSpace(request)) == 0 {
			continue
		}
		_, err := apiObj.db.Exec(request)
		if err != nil {
			fmt.Println("-----Probably not pointing to an alive DB. Make sure all environment variables are set right-----")
			panic(err)
		}
	}
}

func readDBSchema() []string {
	dat, err := ioutil.ReadFile("../../../../db/schema.sql")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), ";")
}
