package httpclient

import (
	`context`
	`log`
	`net/http`
	`testing`
	`time`

	`github.com/stretchr/testify/assert`
	`github.com/valyala/fasthttp`

	`github.com/sshawnta/TestTransport/models`
	`github.com/sshawnta/TestTransport/pkg/service`
	`github.com/sshawnta/TestTransport/pkg/service/httpserver`
)

const (
	serverAddr               = "localhost:8080"
	serverHost               = "localhost:8080"
	maxConns                 = 512
	maxRequestBodySize       = 15 * 1024 * 1024
	serverTimeout            = 1 * time.Millisecond
	serverLaunchingWaitSleep = 1 * time.Second

	// service name method
	getClientId = "GetClientId"
	postOrder   = "PostOrder"
	getCount    = "GetCount"
	getOrder    = "GetOrder"

	fail                     = "fail"
	emptyError               = ""
	errorGetBarcodesSupplier = "error ID"
	ID                       = 12
)

func TestClient_GetClientId(t *testing.T) {
	t.Run("TestGetIdSuccess", func(t *testing.T) {
		request := makeClientIdRequest()
		response := makeClientIdResponse()
		serviceMock := new(service.MockService)
		serviceMock.On(getClientId, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr,serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetClientId(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func TestClient_PostClientId(t *testing.T) {
	t.Run("TestPostSuccess", func(t *testing.T) {
		request := makePostClient()
		response := makeClientIdResponse()
		serviceMock := new(service.MockService)
		serviceMock.On(postOrder, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr,serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)
		resp, err := client.PostOrder(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}
func TestClient_GetCount(t *testing.T) {
	t.Run("TestGetCountSuccess", func(t *testing.T) {
		request := makeClientIdRequest()
		response := makeClientIdResponse()
		serviceMock := new(service.MockService)
		serviceMock.On(getCount, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr,serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetCount(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}


func TestClient_GetOrder(t *testing.T) {
	t.Run("TestGetOrderSuccess", func(t *testing.T) {
		response := makeClientIdResponse()
		serviceMock := new(service.MockService)
		serviceMock.On(getOrder, context.Background()).Return(response, nil)
		server, client := makeServerClient(serverAddr,serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetOrder(context.Background())
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}


func makeServerClient(serverAddr string, svc Service) (server *fasthttp.Server, client Service) {
	client = NewPreparedClient(
		serverAddr,
		serverHost,
		maxConns,
		nil,
		httpserver.NewErrorProcessor(http.StatusInternalServerError, "Iternal Error"),
		httpserver.NewError,
		httpserver.URIPathClientGetClientId,
		httpserver.URIPathClientPostOrder,
		httpserver.URIPathClientGetCountClient,
		httpserver.URIPathClientGetOrder,
		httpserver.HTTPMethodGetClientId,
		httpserver.HTTPMethodPostOrder,
		httpserver.HTTPMethodGetCount,
		httpserver.HTTPMethodGetOrder,
	)
	router := httpserver.NewPreparedServer(svc)
	server = &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: maxRequestBodySize,
		ReadTimeout:        serverTimeout,
	}

	go func() {
		err := server.ListenAndServe(serverAddr)
		if err != nil {
			log.Printf("server shut down err: %v", err)
		}
	}()
	return
}

func makeClientIdRequest() *models.GetClientId {
	return &models.GetClientId{
		Id: ID,
	}
}

func makePostClient()*models.PostClientId  {
	return &models.PostClientId{
		Id: ID,
	}
}

func makeClientIdResponse() models.Response{
	data := models.Data{Res:true}

	return models.Response{
		Data:&data,
	}

}
