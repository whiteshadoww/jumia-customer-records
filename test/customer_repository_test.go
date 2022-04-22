package test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.jumia.org/customers/pkg/utils"
	"testing"

	"go.jumia.org/customers/app/dbs"
	"go.jumia.org/customers/app/interfaces"
	"go.jumia.org/customers/app/models"
	"go.jumia.org/customers/app/repositories"
	"go.jumia.org/customers/app/schema"
)

var (
	customers = []*models.Customer{

		{
			ID:    0,
			Name:  "Walid Hammadi",
			Phone: "(212) 6007989253",
		},
		{
			ID:    1,
			Name:  "Yosaf Karrouch",
			Phone: "(212) 698054317",
		},
		{
			ID:    2,
			Name:  "Younes Boutikyad",
			Phone: "(212) 6546545369",
		},
		{
			ID:    3,
			Name:  "Houda Houda",
			Phone: "(212) 6617344445",
		},
		{
			ID:    4,
			Name:  "Chouf Malo",
			Phone: "(212) 691933626",
		},
		{
			ID:    5,
			Name:  "soufiane fritisse ",
			Phone: "(212) 633963130",
		},
		{
			ID:    6,
			Name:  "Nada Sofie",
			Phone: "(212) 654642448",
		},
		{
			ID:    7,
			Name:  "Edunildo Gomes Alberto ",
			Phone: "(258) 847651504",
		},
		{
			ID:    8,
			Name:  "Walla's Singz Junior",
			Phone: "(258) 846565883",
		},
		{
			ID:    9,
			Name:  "sevilton sylvestre",
			Phone: "(258) 849181828",
		},
		{
			ID:    10,
			Name:  "Tanvi Sachdeva",
			Phone: "(258) 84330678235",
		},
		{
			ID:    11,
			Name:  "Florencio Samuel",
			Phone: "(258) 847602609",
		},
		{
			ID:    12,
			Name:  "Solo Dolo",
			Phone: "(258) 042423566",
		},
		{
			ID:    13,
			Name:  "Pedro B 173",
			Phone: "(258) 823747618",
		},
		{
			ID:    14,
			Name:  "Ezequiel Fenias",
			Phone: "(258) 848826725",
		},
		{
			ID:    15,
			Name:  "JACKSON NELLY",
			Phone: "(256) 775069443",
		},
		{
			ID:    16,
			Name:  "Kiwanuka Budallah",
			Phone: "(256) 7503O6263",
		},
		{
			ID:    17,
			Name:  "VINEET SETH",
			Phone: "(256) 704244430",
		},
		{
			ID:    18,
			Name:  "Jokkene Richard",
			Phone: "(256) 7734127498",
		},
		{
			ID:    19,
			Name:  "Ogwal David",
			Phone: "(256) 7771031454",
		},
		{
			ID:    20,
			Name:  "pt shop 0901 Ultimo ",
			Phone: "(256) 3142345678",
		},
		{
			ID:    21,
			Name:  "Daniel Makori",
			Phone: "(256) 714660221",
		},
		{
			ID:    22,
			Name:  "shop23 sales",
			Phone: "(251) 9773199405",
		},
		{
			ID:    23,
			Name:  "Filimon Embaye",
			Phone: "(251) 914701723",
		},
		{
			ID:    24,
			Name:  "ABRAHAM NEGASH",
			Phone: "(251) 911203317",
		},
		{
			ID:    25,
			Name:  "ZEKARIAS KEBEDE",
			Phone: "(251) 9119454961",
		},
		{
			ID:    26,
			Name:  "EPHREM KINFE",
			Phone: "(251) 914148181",
		},
		{
			ID:    27,
			Name:  "Karim Niki",
			Phone: "(251) 966002259",
		},
		{
			ID:    28,
			Name:  "Frehiwot Teka",
			Phone: "(251) 988200000",
		},
		{
			ID:    29,
			Name:  "Fanetahune Abaia",
			Phone: "(251) 924418461",
		},
		{
			ID:    30,
			Name:  "Yonatan Tekelay",
			Phone: "(251) 911168450",
		},
		{
			ID:    31,
			Name:  "EMILE CHRISTIAN KOUKOU DIKANDA HONORE ",
			Phone: "(237) 697151594",
		},
		{
			ID:    32,
			Name:  "MICHAEL MICHAEL",
			Phone: "(237) 677046616",
		},
		{
			ID:    33,
			Name:  "ARREYMANYOR ROLAND TABOT",
			Phone: "(237) 6A0311634",
		},
		{
			ID:    34,
			Name:  "LOUIS PARFAIT OMBES NTSO",
			Phone: "(237) 673122155",
		},
		{
			ID:    35,
			Name:  "JOSEPH FELICIEN NOMO",
			Phone: "(237) 695539786",
		},
		{
			ID:    36,
			Name:  "SUGAR STARRK BARRAGAN",
			Phone: "(237) 6780009592",
		},
		{
			ID:    37,
			Name:  "WILLIAM KEMFANG",
			Phone: "(237) 6622284920",
		},
		{
			ID:    38,
			Name:  "THOMAS WILFRIED LOMO LOMO",
			Phone: "(237) 696443597",
		},
		{
			ID:    39,
			Name:  "Dominique mekontchou",
			Phone: "(237) 691816558",
		},
		{
			ID:    40,
			Name:  "Nelson Nelson",
			Phone: "(237) 699209115",
		},
	}

	customer = customers[0]
)

type CustomerRepositoryTestSuite struct {
	suite.Suite

	db   interfaces.IDatabase
	repo interfaces.ICustomerRepository
}

func (s *CustomerRepositoryTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()

	s.db = dbs.NewDatabase()
	s.repo = repositories.NewCustomerRepository(s.db)
}

func (s *CustomerRepositoryTestSuite) TestListFull() {
	custs, err := s.repo.List(&schema.CustomerQueryParam{
		Page:     1,
		PageSize: 100000,
	})
	s.Nil(err)
	s.NotNil(custs)
	s.Equal(len(customers), len(*custs))
}

func (s *CustomerRepositoryTestSuite) TestValidList() {
	custs, err := s.repo.List(&schema.CustomerQueryParam{
		State: utils.BoolPointer(true),
	})
	s.Nil(err)
	s.NotNil(custs)
	s.Equal(27, len(*custs))
}

func (s *CustomerRepositoryTestSuite) TestInValidList() {
	custs, err := s.repo.List(&schema.CustomerQueryParam{
		State: utils.BoolPointer(false),
	})
	s.Nil(err)
	s.NotNil(custs)
	s.Equal(14, len(*custs))
}

func (s *CustomerRepositoryTestSuite) TestCountryList() {
	custs, err := s.repo.List(&schema.CustomerQueryParam{
		Country: string(schema.Cameroon),
	})
	s.Nil(err)
	s.NotNil(custs)
	s.Equal(10, len(*custs))
}

func (s *CustomerRepositoryTestSuite) TestCountryStateList() {
	custs, err := s.repo.List(&schema.CustomerQueryParam{
		Country: string(schema.Cameroon),
		State:   utils.BoolPointer(true),
	})
	s.Nil(err)
	s.NotNil(custs)
	s.Equal(7, len(*custs))
}

func (s *CustomerRepositoryTestSuite) TestMaxLimit() {
	custs, err := s.repo.List(&schema.CustomerQueryParam{
		PageSize: 20,
	})
	s.Nil(err)
	s.NotNil(custs)
	s.Equal(20, len(*custs))
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}
