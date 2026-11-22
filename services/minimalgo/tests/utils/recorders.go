package utils

// func HelperJoinTableConnection(t *testing.T, path string) (*websocket.Conn, *httptest.Server, *httptest.ResponseRecorder, error) {
// 	t.Helper()

// 	// app := HelperHttpServer(t)
// 	_, httpApp, _ := HelperFullServer(t)

// 	h := &http.BaseHandlers{}
// 	h.SetApp(httpApp)

// 	server := httptest.NewServer(httpnet.HandlerFunc(h.JoinWebsite))
// 	recorder := httptest.NewRecorder()

// 	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + path
// 	conn, response, err := websocket.DefaultDialer.Dial(wsURL, httpnet.Header{"Origin": []string{"GoTest"}})

// 	if err != nil {
// 		if response != nil {
// 			body, _ := io.ReadAll(response.Body)
// 			response.Body.Close()
// 			t.Fatalf("websocket dial failed (status %d): %v | body: %s", response.StatusCode, err, string(body))
// 		}
// 		t.Fatalf("websocket dial failed: %v", err)
// 	}

// 	return conn, server, recorder, err
// }

// func HelperRunJoinTableConnection(t *testing.T, path string) (*websocket.Conn, *httptest.Server, *httptest.ResponseRecorder, error) {
// 	t.Helper()

// 	s, httpApp, _ := HelperFullServer(t)

// 	go s.Start(t.Context())

// 	h := &http.BaseHandlers{}
// 	h.SetApp(httpApp)

// 	server := httptest.NewServer(httpnet.HandlerFunc(h.JoinWebsite))
// 	recorder := httptest.NewRecorder()

// 	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + path
// 	conn, response, err := websocket.DefaultDialer.Dial(wsURL, httpnet.Header{"Origin": []string{"GoTest"}})

// 	if err != nil {
// 		if response != nil {
// 			body, _ := io.ReadAll(response.Body)
// 			response.Body.Close()
// 			t.Fatalf("websocket dial failed (status %d): %v | body: %s", response.StatusCode, err, string(body))
// 		}
// 		t.Fatalf("websocket dial failed: %v", err)
// 	}

// 	return conn, server, recorder, err
// }
