package e2e

import (
	"bitsports/ent"
	"bitsports/internal/product/adapter/graphql"
	"bitsports/internal/product/adapter/repository"
	"bitsports/internal/product/application/usecase"
	infraGraphql "bitsports/internal/product/infraestructure/graphql"
	"bitsports/internal/product/infraestructure/router"
	"bitsports/pkg/docker"

	"bitsports/testutil"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)


// Setup set up database and server for E2E test
func Setup(t *testing.T,  c *docker.Container, dbName string) (expect *httpexpect.Expect, client *ent.Client, teardown func()) {
	t.Helper()
	test := testutil.NewIntegration(t, c, dbName)

	repoProduct := repository.NewProduct(test.Client)
	ucProduct := usecase.NewProduct(repoProduct)
	repoCategory := repository.NewCategory(test.Client)
	ucCategory := usecase.NewCategory(repoCategory)

	schema :=  graphql.NewSchema(graphql.NewProductResolver(ucProduct),
		graphql.NewCategoryResolver(ucCategory))

	handler, err := infraGraphql.NewServer(schema)

	if err!= nil {
		test.Log.Fatal(err)
	}
	e := router.New(handler, test.Log, router.WithoutSecurity())

	srv := httptest.NewServer(e)

	return httpexpect.WithConfig(httpexpect.Config{
			BaseURL:  srv.URL,
			Reporter: httpexpect.NewAssertReporter(t),
			Printers: []httpexpect.Printer{
				httpexpect.NewDebugPrinter(t, true),
			},
		}), test.Client, func() {
			defer test.Teardown()
			defer srv.Close()
		}
}

// GetData gets data from graphql response.
func GetData(e *httpexpect.Response) *httpexpect.Value {
	return e.JSON().Path("$.data")
}

// GetObject return data from path.
// Path returns a new Value object for child object(s) matching given
// JSONPath expression.
// Example 1:
//  json := `{"users": [{"name": "john"}, {"name": "bob"}]}`
//  value := NewValue(t, json)
//
//  value.Path("$.users[0].name").String().Equal("john")
//  value.Path("$.users[1].name").String().Equal("bob")
func GetObject(obj *httpexpect.Object, path string) *httpexpect.Object {
	return obj.Path("$." + path).Object()
}

// GetErrors return errors from graphql response.
func GetErrors(e *httpexpect.Response) *httpexpect.Value {
	return e.JSON().Path("$.errors")
}
