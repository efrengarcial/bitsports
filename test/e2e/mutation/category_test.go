package mutation_test

import (
	"bitsports/internal/product/infraestructure/router"
	"bitsports/testutil/e2e"
	"context"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"strconv"
	"testing"
)

func TestCategory_Create(t *testing.T) {
	expect, _ , teardown := e2e.Setup(t, container, "e2ecategorycreate")
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
			name:    "It should create test category",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {
							createCategory( name: "Frutas", code: "FRU") {
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
				testUser := e2e.GetObject(data, "createCategory")
				testUser.Value("code").String().Equal("FRU")
				testUser.Value("name").String().Equal("Frutas")
			},
			teardown: func(t *testing.T) { },
		},
		{
			name:    "It should NOT create test category when the length of the code is over",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {
				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {  
							createCategory(name: "FRU", code: "Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1T"}) {   
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
				data := e2e.GetData(got)
				data.Null()

				errors := e2e.GetErrors(got)
				errors.Array().Length().Equal(1)
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

func TestCategory_Update(t *testing.T) {
	expect, client , teardown := e2e.Setup(t, container, "e2ecategoryupdate")
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
			name:    "It should update test category",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {

				ctx := context.Background()

				c, err := client.Category.Create().
					SetCode("ABC").
					SetName("Frutal").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {
							updateCategory( name: "Frutas", code : "FRU" ,  id:`+ strconv.FormatInt(c.ID, 10) + `) {
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
				testUser := e2e.GetObject(data, "updateCategory")
				testUser.Value("code").String().Equal("FRU")
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

func TestCategory_Delete(t *testing.T) {
	expect, client , teardown := e2e.Setup(t, container, "e2ecategorydelete")
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
			name:    "It should delete test category",
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
						mutation {
							deleteCategory( id:`+ strconv.FormatInt(c.ID, 10) + `) {
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
				testUser := e2e.GetObject(data, "deleteCategory")
				testUser.Value("id").Null()
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
