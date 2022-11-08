package usecase

import (
	"database/sql"
	"testing"

	"github.com/EduardoPPCaldas/fullcycle/internal/order/entity"
	"github.com/EduardoPPCaldas/fullcycle/internal/order/infra/database"
	"github.com/stretchr/testify/suite"
	_ "github.com/mattn/go-sqlite3"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("1", 10, 2)
	suite.NoError(err)
	err = order.CalculatePrice()
	suite.NoError(err)

	calculateFinalPriceInput := OrderInputDTO{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
	}

	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)

	suite.NoError(err)
	suite.Equal(order.ID, output.ID)
	suite.Equal(order.FinalPrice, output.FinalPrice)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
}