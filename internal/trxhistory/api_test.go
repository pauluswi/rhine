package trxhistory

// func TestAPI(t *testing.T) {
// 	logger, _ := log.NewForTest()
// 	router := test.MockRouter(logger)
// 	repo := &mockRepository{items: []entity.TrxHistory{
// 		{uuid.NewV4().String(), "999999", "6281100099", "c", "0", 10000, time.Now().UTC(), time.Now().UTC()},
// 	}}
// 	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, logger)
// 	header := auth.MockAuthHeader()

// 	tests := []test.APITestCase{
// 		{"get all ok", "GET", "/getall", "", header, http.StatusOK, ``},
// 		//{"get ok", "GET", "/get/6281100099", "", header, http.StatusOK, ``},
// 		//{"get unknown", "GET", "/get/628110009911", "", nil, http.StatusNotFound, ""},
// 		// {"generate ok", "POST", "/generate", `{"customer_id":"6281100099"}`, header, http.StatusCreated, "*valid_until*"},
// 		// {"generate auth error", "POST", "/generate", `{"customer_id":"6281100099"}`, nil, http.StatusUnauthorized, ""},
// 		// {"generate input error", "POST", "/generate", `"customer_id":"6281100099"}`, header, http.StatusBadRequest, ""},
// 		// {"validate ok", "POST", "/validate", `{"token":"999999"}`, header, http.StatusCreated, "*valid_until*"},
// 		// {"validate auth error", "POST", "/validate", `{"CustomerID":"999999"}`, nil, http.StatusUnauthorized, ""},
// 		// {"validate input error", "POST", "/validate", `"CustomerID":"999999"}`, header, http.StatusBadRequest, ""},
// 	}
// 	for _, tc := range tests {
// 		test.Endpoint(t, router, tc)
// 	}
// }
