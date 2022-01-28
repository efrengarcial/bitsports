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

func TestProduct_Query(t *testing.T) {
	expect, client, teardown := e2e.Setup(t, container, "e2eproductquery")
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
			name:    "It should query test products",
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
				products := []struct {
					name       string
					quantity   int64
					price      float64
					categoryID int64
				}{{name: "test", quantity: 10, price: 10.11, categoryID: c.ID},
					{name: "test2", quantity: 11, price: 15.2, categoryID: c.ID},
					{name: "test3", quantity: 12, price: 7.12, categoryID: c.ID}}
				bulk := make([]*ent.ProductCreate, len(products))
				for i, product := range products {
					bulk[i] = client.Product.Create().
						SetName(product.name).
						SetPrice(product.price).
						SetQuantity(product.quantity).
						SetCategoryID(product.categoryID)
				}

				_, err = client.Product.
					CreateBulk(bulk...).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						query {
							products { name,  price, createdAt, updatedAt, quantity, id, 
							  category { id, name , code , createdAt, updatedAt
							  }
							}  	
					}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got).Object()
				testUser := e2e.GetArray(data, "products")
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

func TestProduct_QueryByID(t *testing.T) {
	expect, client , teardown := e2e.Setup(t, container, "e2eproductquerybyid")
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
			name:    "It should query by id test product",
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
						query {
							product( id:`+ strconv.FormatInt(p.ID, 10) + `) {
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
				testUser := e2e.GetObject(data, "product")
				testUser.Value("quantity").Number().Equal(10)
				testUser.Value("name").String().Equal("Banano")
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
