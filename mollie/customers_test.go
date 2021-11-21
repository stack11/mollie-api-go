package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type CustomersTestSuite struct{ suite.Suite }

func (cs *CustomersTestSuite) SetupSuite() { setEnv() }

func (cs *CustomersTestSuite) TearDownSuite() { unsetEnv() }

func (cs *CustomersTestSuite) TestCustomerService_Get() {
	type args struct {
		ctx      context.Context
		customer string
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get mollie customers works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetCustomerResponse))
			},
		},
		{
			"get mollie customers, an error is returned from the server",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"get mollie customers, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get mollie customers, invalid url when building request",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		cs.T().Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s", c.args.customer), c.handler)
			c.pre()
			cc, err := tClient.Customers.Get(c.args.ctx, c.args.customer)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.Equal(c.args.customer, cc.ID)
			}
		})
	}
}

func (cs *CustomersTestSuite) TestCustomersService_Create() {
	type args struct {
		ctx      context.Context
		customer Customer
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"create mollie customers works as expected",
			http.StatusAccepted,
			args{
				context.Background(),
				Customer{Locale: German},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "POST")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.CreateCustomerResponse))
			},
		},
		{
			"create mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
			args{
				context.Background(),
				Customer{},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"create mollie customers, an error occurs when parsing json",
			http.StatusInternalServerError,
			args{
				context.Background(),
				Customer{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				Customer{},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		cs.T().Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc("/v2/customers", c.handler)
			c.pre()
			cc, err := tClient.Customers.Create(c.args.ctx, c.args.customer)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.IsType(&Customer{}, cc)
			}
		})
	}
}

func (cs *CustomersTestSuite) TestCustomersService_Update() {
	type args struct {
		ctx        context.Context
		customerID string
		customer   Customer
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"update mollie customers works as expected",
			http.StatusAccepted,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Customer{Locale: French},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "PATCH")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.UpdateCustomerResponse))
			},
		},
		{
			"update mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Customer{},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"update mollie customers, an error occurs when parsing json",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Customer{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Customer{},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		cs.T().Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s", c.args.customerID), c.handler)
			c.pre()
			cc, err := tClient.Customers.Update(c.args.ctx, c.args.customerID, c.args.customer)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.IsType(&Customer{}, cc)
			}
		})
	}
}

func (cs *CustomersTestSuite) TestCustomersService_List() {
	type args struct {
		ctx     context.Context
		options *ListCustomersOptions
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"list mollie customers with options works as expected",
			http.StatusAccepted,
			args{
				context.Background(),
				&ListCustomersOptions{
					SequenceType: OneOffSequence,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListCustomersResponse))
			},
		},
		{
			"list mollie customers works as expected",
			http.StatusAccepted,
			args{
				context.Background(),
				nil,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListCustomersResponse))
			},
		},
		{
			"list mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"list mollie customers, an error occurs when parsing json",
			http.StatusInternalServerError,
			args{
				context.Background(),
				&ListCustomersOptions{
					SequenceType: OneOffSequence,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				&ListCustomersOptions{
					SequenceType: OneOffSequence,
				},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		cs.T().Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc("/v2/customers", c.handler)
			c.pre()
			cc, err := tClient.Customers.List(c.args.ctx, c.args.options)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.IsType(&CustomersList{}, cc)
			}
		})
	}
}

func (cs *CustomersTestSuite) TestCustomersService_Delete() {
	type args struct {
		ctx      context.Context
		customer string
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"delete mollie customers with options works as expected",
			http.StatusNoContent,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "DELETE")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"delete mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"delete mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		cs.T().Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s", c.args.customer), c.handler)
			c.pre()
			err := tClient.Customers.Delete(c.args.ctx, c.args.customer)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
			}
		})
	}
}

func (cs *CustomersTestSuite) TestCustomerService_GetPayments() {
	type args struct {
		ctx      context.Context
		customer string
		options  *ListCustomersOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get mollie customers payments works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				nil,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"get mollie customers payments with options works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				&ListCustomersOptions{Limit: 100},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"get mollie customers payments, an error is returned from the server",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				&ListCustomersOptions{SequenceType: RecurringSequence},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"get mollie customers payments, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get mollie customers payments, invalid url when building request",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				nil,
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		cs.T().Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/payments", c.args.customer), c.handler)
			c.pre()
			cc, err := tClient.Customers.GetPayments(c.args.ctx, c.args.customer, c.args.options)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.NotZero(cc.Count)
			}
		})
	}
}

func (cs *CustomersTestSuite) TestCustomerService_CreatePayment() {
	type args struct {
		ctx      context.Context
		customer string
		payment  Payment
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"create mollie customers payments works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Payment{TestMode: true},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(cs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(cs.T(), r, "POST")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"create mollie customers payments, an error is returned from the server",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Payment{TestMode: true},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"create mollie customers payments, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Payment{TestMode: true},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create mollie customers payments, invalid url when building request",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				Payment{TestMode: true},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		cs.T().Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/payments", c.args.customer), c.handler)
			c.pre()
			cc, err := tClient.Customers.CreatePayment(c.args.ctx, c.args.customer, c.args.payment)
			if c.wantErr {
				cs.NotNil(err)
				cs.EqualError(err, c.err.Error())
			} else {
				cs.Nil(err)
				cs.IsType(&Payment{}, cc)
			}
		})
	}
}

func TestCustomersService(t *testing.T) {
	suite.Run(t, new(CustomersTestSuite))
}
