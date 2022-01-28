package repository_test

import (
	"bitsports/ent"
	"bitsports/internal/product/adapter/repository"
	"bitsports/internal/product/domain"
	"bitsports/testutil"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductRepository_Query(t *testing.T) {
	t.Helper()

	_, client, teardown := testutil.NewUnit(t, container, "testproductquery")
	t.Cleanup(teardown)

	repo := repository.NewProduct(client)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (p []domain.Product, err error)
		assert  func(t *testing.T, p []domain.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get product's list",
			arrange: func(t *testing.T) {
				ctx := context.Background()

				c, err := client.Category.Create().
					SetCode("test").
					SetName("test").
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
			},
			act: func(ctx context.Context, t *testing.T) (p []domain.Product, err error) {
				return repo.Query(ctx)
			},
			assert: func(t *testing.T, got []domain.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 3, len(got))
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				teardown()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}

func TestProductRepository_Create(t *testing.T) {
	t.Helper()

	_, client, teardown := testutil.NewUnit(t, container, "testproductcreate")
	t.Cleanup(teardown)

	repo := repository.NewProduct(client)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (p *domain.Product, err error)
		assert  func(t *testing.T, p *domain.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name:    "It should create a product",
			arrange: func(t *testing.T) {},
			act: func(ctx context.Context, t *testing.T) (p *domain.Product, err error) {
				c, err := client.Category.Create().
					SetCode("test").
					SetName("test").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				product := domain.NewProduct{
					Name:       "test",
					Price:      100.25,
					Quantity:   10,
					CategoryID: c.ID,
				}
				return repo.Create(ctx, product)
			},
			assert: func(t *testing.T, got *domain.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "test", got.Name)
				assert.Equal(t, int64(10), got.Quantity)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				teardown()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}

func TestProductRepository_Update(t *testing.T) {
	t.Helper()

	_, client, teardown := testutil.NewUnit(t, container, "testproductupdate")
	t.Cleanup(teardown)

	repo := repository.NewProduct(client)

	type args struct {
		ctx       context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T) int64
		act     func(ctx context.Context, productID int64, t *testing.T) (p *domain.Product, err error)
		assert  func(t *testing.T, p *domain.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should update a product",
			arrange: func(t *testing.T) int64 {
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

				return p.ID

			},
			act: func(ctx context.Context, productID int64, t *testing.T) (p *domain.Product, err error) {
				var (
					name = "Zanahoria"
					price float64 = 2000
					quantity int64 = 20
				)

				cat, err := client.Category.Create().
					SetCode("123").
					SetName("Verduras").
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				prd := domain.UpdateProduct{
					ID:         productID,
					Name:      &name,
					Price:      &price,
					Quantity:   &quantity,
					CategoryID: &cat.ID,
				}
				return repo.Update(ctx, prd)
			},
			assert: func(t *testing.T, got *domain.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "Zanahoria", got.Name)
				assert.Equal(t, int64(20), got.Quantity)
				assert.Equal(t, float64(2000), got.Price)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				teardown()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.arrange(t)
			got, err := tt.act(tt.args.ctx, id, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}

func TestProductRepository_QueryByID(t *testing.T) {
	t.Helper()

	_, client, teardown := testutil.NewUnit(t, container, "testproductquerybyid")
	t.Cleanup(teardown)

	repo := repository.NewProduct(client)

	type args struct {
		ctx       context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T) int64
		act     func(ctx context.Context, productID int64, t *testing.T) (p *domain.Product, err error)
		assert  func(t *testing.T, p *domain.Product, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get product by id",
			arrange: func(t *testing.T) int64 {
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

				return p.ID

			},
			act: func(ctx context.Context, productID int64, t *testing.T) (p *domain.Product, err error) {
				return repo.QueryByID(ctx, productID)
			},
			assert: func(t *testing.T, got *domain.Product, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "Banano", got.Name)
				assert.Equal(t, int64(10), got.Quantity)
				assert.Equal(t, float64(1000), got.Price)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				teardown()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.arrange(t)
			got, err := tt.act(tt.args.ctx, id, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}

func TestProductRepository_Delete(t *testing.T) {
	t.Helper()

	_, client, teardown := testutil.NewUnit(t, container, "testproductdelete")
	t.Cleanup(teardown)

	repo := repository.NewProduct(client)

	type args struct {
		ctx       context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T) int64
		act     func(ctx context.Context, productID int64, t *testing.T) (bool, error)
		assert  func(t *testing.T, result bool, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should delete a product by id",
			arrange: func(t *testing.T) int64 {
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

				return p.ID

			},
			act: func(ctx context.Context, productID int64, t *testing.T) (bool, error) {
				return true, repo.Delete(ctx, productID)
			},
			assert: func(t *testing.T, got bool, err error) {
				assert.Nil(t, err)
				assert.True(t, got)
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				teardown()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.arrange(t)
			got, err := tt.act(tt.args.ctx, id, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}
