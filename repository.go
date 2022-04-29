package shopware6_admin_go_client

import (
	"bytes"
	"encoding/json"
	"github.com/develist/shopware6-admin-go-client/entity"
	"net/http"
	"reflect"
	"strings"
)

type BaseRepository[T any] interface {
	GetList() (*[]T, error)
	GetDetail(id int) (*T, error)
	Search(criteria entity.Criteria) (*[]T, error)
	SearchIds(criteria entity.Criteria) (*[]string, error)
}

type Repository[T any] struct {
	endpoints  Endpoints
	httpClient *http.Client
}

type Endpoints struct {
	entity    string
	search    string
	searchIds string
}

func CreateRepository[T any](baseUrl string, httpClient *http.Client) *Repository[T] {
	var a T
	endpoint := strings.ToLower(toKebabCase(reflect.TypeOf(a).Name()))
	client := httpClient
	entityEndPoint := strings.TrimRight(baseUrl, "/") + "/api/" + endpoint
	searchEndPoint := strings.TrimRight(baseUrl, "/") + "/api/search/" + endpoint
	searchIdsEndPoint := strings.TrimRight(baseUrl, "/") + "/api/search-ids/" + endpoint
	endpoints := Endpoints{entityEndPoint, searchEndPoint, searchIdsEndPoint}
	return &Repository[T]{endpoints, client}
}

func (g *Repository[T]) GetDetail(id string) (*T, error) {
	url := g.endpoints.entity + "/" + id
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	body, err := getBody(g.httpClient, req)
	if body == nil || err != nil {
		return nil, err
	}
	var result entity.Result[T]
	if err := json.Unmarshal(*body, &result); err != nil {
		return nil, err
	}
	return &result.Data, nil
}

func (g *Repository[T]) GetList() (*[]T, error) {
	req, err := http.NewRequest(http.MethodGet, g.endpoints.entity, nil)
	if err != nil {
		return nil, err
	}
	return queryArray[T](g.httpClient, req)
}

func (g *Repository[T]) Search(criteria entity.Criteria) (*[]T, error) {
	criteriaJSON, err := json.Marshal(criteria)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(http.MethodPost, g.endpoints.search, bytes.NewBuffer(criteriaJSON))
	return queryArray[T](g.httpClient, req)
}

func (g *Repository[T]) SearchIds(criteria entity.Criteria) (*[]string, error) {
	criteriaJSON, err := json.Marshal(criteria)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(http.MethodPost, g.endpoints.searchIds, bytes.NewBuffer(criteriaJSON))
	return queryArray[string](g.httpClient, req)
}
