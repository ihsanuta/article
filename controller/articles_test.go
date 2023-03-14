package controller

import (
	"article/app/models"
	usecase "article/app/usecase"
	"article/app/usecase/articles/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func performRequest(c *gin.Context, ctrl *controller) *httptest.ResponseRecorder {
	router := gin.Default()
	router.POST("/api/v1/articles", func(c *gin.Context) {
		ctrl.CreateArticles(c)
	})

	req, _ := http.NewRequest(c.Request.Method, c.Request.URL.Path, c.Request.Body)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func performGetRequest(c *gin.Context, ctrl *controller) *httptest.ResponseRecorder {
	router := gin.Default()
	router.GET("/api/v1/articles", func(c *gin.Context) {
		ctrl.GetArticles(c)
	})

	req, _ := http.NewRequest(c.Request.Method, c.Request.URL.Path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func Test_controller_GetArticles(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mocksUsecase := mocks.NewArticles(t)
	uc := &usecase.Usecase{
		Articles: mocksUsecase,
	}
	type fields struct {
		usecase *usecase.Usecase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		response   []models.Article
		wantErr    error
		wantStatus int
	}{
		// TODO: Add test cases.
		{
			name: "SHOULD Success with result",
			fields: fields{
				usecase: uc,
			},
			args: args{
				c: &gin.Context{
					Request: &http.Request{
						Method: "GET",
						URL: &url.URL{
							Path: "/api/v1/articles",
						},
						Header: http.Header{},
					},
				},
			},
			response: []models.Article{
				{ID: 1,
					Author:  "",
					Title:   "dia",
					Body:    "test",
					Created: time.Now(),
				},
			},
			wantStatus: http.StatusOK,
			wantErr:    nil,
		},
		{
			name: "SHOULD Success without result",
			fields: fields{
				usecase: uc,
			},
			args: args{
				c: &gin.Context{
					Request: &http.Request{
						Method: "GET",
						URL: &url.URL{
							Path: "/api/v1/articles",
						},
						Header: http.Header{},
					},
				},
			},
			response:   []models.Article{},
			wantErr:    nil,
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksUsecase.On("GetArticles", mock.Anything).Return(tt.response, tt.wantErr)
			ctrl := &controller{
				usecase: tt.fields.usecase,
			}

			tt.args.c.Request.Header.Set("Content-Type", "application/json")

			// Perform a GET request with that handler.
			w := performGetRequest(tt.args.c, ctrl)

			// Assert we encoded correctly,
			// the request gives a 200
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func Test_controller_CreateArticles(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mocksUsecase := mocks.NewArticles(t)
	uc := &usecase.Usecase{
		Articles: mocksUsecase,
	}
	type fields struct {
		usecase *usecase.Usecase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		payload    models.Article
		wantErr    error
		wantStatus int
	}{
		// TODO: Add test cases.
		{
			name: "SHOULD Bad Request",
			fields: fields{
				usecase: uc,
			},
			args: args{
				c: &gin.Context{
					Request: &http.Request{
						Method: "POST",
						URL: &url.URL{
							Path: "/api/v1/articles",
						},
						Header: http.Header{},
					},
				},
			},
			payload: models.Article{
				ID:      1,
				Author:  "",
				Title:   "dia",
				Body:    "test",
				Created: time.Now(),
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    nil,
		},
		{
			name: "SHOULD Success",
			fields: fields{
				usecase: uc,
			},
			args: args{
				c: &gin.Context{
					Request: &http.Request{
						Method: "POST",
						URL: &url.URL{
							Path: "/api/v1/articles",
						},
						Header: http.Header{},
					},
				},
			},
			payload: models.Article{
				ID:      1,
				Author:  "saya",
				Title:   "dia",
				Body:    "test",
				Created: time.Now(),
			},
			wantErr:    fmt.Errorf("Bad Request"),
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksUsecase.On("CreateArticles", mock.Anything, mock.Anything).Return(tt.payload, tt.wantErr)
			ctrl := &controller{
				usecase: tt.fields.usecase,
			}

			jsonbytes, err := json.Marshal(tt.payload)
			if err != nil {
				panic(err)
			}

			tt.args.c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
			tt.args.c.Request.Header.Set("Content-Type", "application/json")

			// Perform a GET request with that handler.
			w := performRequest(tt.args.c, ctrl)

			// Assert we encoded correctly,
			// the request gives a 200
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
