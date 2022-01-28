package query_test

import (
	"bitsports/ent"
	"bitsports/internal/product/infraestructure/router"
	"bitsports/testutil/e2e"
	"context"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"strconv"
	"testing"
)

func TestCategory_Query(t *testing.T) {
	expect, client, teardown := e2e.Setup(t, container, "e2ecategoryquery")
	defer teardown()

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(t *testing.T) *httpexpect.Response
		assert  func(t *testing.T, got *httpexpect.Response)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name:    "It should query test categories",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {

				ctx := context.Background()

				categories := []struct {
					name string
					code string
				}{{name: "test", code: "A1"},
					{name: "test2", code: "A2"},
					{name: "test3", code: "A3"}}
				bulk := make([]*ent.CategoryCreate, len(categories))
				for i, category := range categories {
					bulk[i] = client.Category.Create().
						SetName(category.name).
						SetCode(category.code)
				}

				_, err := client.Category.
					CreateBulk(bulk...).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						query {
							categories { name,  code, createdAt, updatedAt, id }
						}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got).Object()
				testUser := e2e.GetArray(data, "categories")
				testUser.Length().Equal(3)
			},
			teardown: func(t *testing.T) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got := tt.act(t)
			tt.assert(t, got)
			tt.teardown(t)
		})
	}
}

func TestCategory_QueryByID(t *testing.T) {
	expect, client , teardown := e2e.Setup(t, container, "e2ecategoryquerybyid")
	defer teardown()

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(t *testing.T) *httpexpect.Response
		assert  func(t *testing.T, got *httpexpect.Response)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name:    "It should query by id test category",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {

				ctx := context.Background()

				c, err := client.Category.Create().
					SetCode("ABC").
					SetName("Frutas").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						query {
							category( id:`+ strconv.FormatInt(c.ID, 10) + `) {
								code
								name
								id
								createdAt
								updatedAt
						}
					}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got).Object()
				testUser := e2e.GetObject(data, "category")
				testUser.Value("code").String().Equal("ABC")
				testUser.Value("name").String().Equal("Frutas")
			},
			teardown: func(t *testing.T) { },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got := tt.act(t)
			tt.assert(t, got)
			tt.teardown(t)
		})
	}
}
