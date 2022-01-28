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

func TestProduct_Create(t *testing.T) {
	expect, client , teardown := e2e.Setup(t, container, "e2eproductcreate")
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
			name:    "It should create test product",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {

				c, err := client.Category.Create().
					SetCode("test").
					SetName("test").
					Save(context.Background())
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {
							createProduct( name: "Naranja", price: 11.52, quantity: 100,  categoryId:`+ strconv.FormatInt(c.ID, 10) + `) {
								price
								name
								id
								quantity
								createdAt
								updatedAt
						}
					}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got).Object()
				testUser := e2e.GetObject(data, "createProduct")
				testUser.Value("quantity").Number().Equal(100)
				testUser.Value("name").String().Equal("Naranja")
			},
			teardown: func(t *testing.T) { },
		},
		{
			name:    "It should NOT create test product when the length of the name is over",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {
				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {  
							createUser(name: "Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1", quantity: 20}) {   
								quantity
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

func TestProduct_Update(t *testing.T) {
	expect, client , teardown := e2e.Setup(t, container, "e2eproductupdate")
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
			name:    "It should update test product",
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
				p, err := client.Product.
					Create().
					SetName("Banano").
					SetPrice(1000).
					SetQuantity(10).
					SetCategoryID(c.ID).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {
							updateProduct( name: "Naranja", price: 11.52, quantity: 100,  id:`+ strconv.FormatInt(p.ID, 10) + `) {
								price
								name
								id
								quantity
								createdAt
								updatedAt
						}
					}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got).Object()
				testUser := e2e.GetObject(data, "updateProduct")
				testUser.Value("quantity").Number().Equal(100)
				testUser.Value("name").String().Equal("Naranja")
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

func TestProduct_Delete(t *testing.T) {
	expect, client , teardown := e2e.Setup(t, container, "e2eproductdelete")
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
			name:    "It should delete test product",
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
				p, err := client.Product.
					Create().
					SetName("Banano").
					SetPrice(1000).
					SetQuantity(10).
					SetCategoryID(c.ID).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {
							deleteProduct( id:`+ strconv.FormatInt(p.ID, 10) + `) {
								price
								name
								id
								quantity
								createdAt
								updatedAt
						}
					}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got).Object()
				testUser := e2e.GetObject(data, "deleteProduct")
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

