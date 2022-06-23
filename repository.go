package shopware6_admin_go_client

import (
	"bytes"
	"encoding/json"
	"github.com/develist/shopware6-admin-go-client/response"
	"github.com/develist/shopware6-admin-go-client/search"
	"net/http"
	"reflect"
	"strings"
)

const slash string = "/"
const searchPath string = "search"
const searchIdsPath string = "search-ids"

type EntityRepository[T any] struct {
	endpoints  endpoints
	httpClient *http.Client
}

type endpoints struct {
	entity    string
	search    string
	searchIds string
}

// CreateEntityRepository Creates a new generic repository. The Shopware 6 base API URL and a preconfigured
// http.Client is needed.
func CreateEntityRepository[T any](apiBaseUrl string, httpClient *http.Client) *EntityRepository[T] {
	var a T
	endpoint := strings.ToLower(toKebabCase(reflect.TypeOf(a).Name()))
	client := httpClient
	entityEndPoint := strings.TrimRight(apiBaseUrl, slash) + slash + endpoint
	searchEndPoint := strings.TrimRight(apiBaseUrl, slash) + slash + searchPath + slash + endpoint
	searchIdsEndPoint := strings.TrimRight(apiBaseUrl, slash) + slash + searchIdsPath + slash + endpoint
	endpoints := endpoints{entityEndPoint, searchEndPoint, searchIdsEndPoint}
	return &EntityRepository[T]{endpoints, client}
}

func (g *EntityRepository[T]) GetDetail(id string) (*T, *response.ApiErrors) {
	url := g.endpoints.entity + slash + id
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, createAPIErrors(err)
	}
	return querySingle[T](g.httpClient, req)
}

func (g *EntityRepository[T]) GetList() (*[]T, *response.ApiErrors) {
	req, err := http.NewRequest(http.MethodGet, g.endpoints.entity, nil)
	if err != nil {
		return nil, createAPIErrors(err)
	}
	return queryArray[T](g.httpClient, req)
}

func (g *EntityRepository[T]) Search(criteria search.Criteria) (*[]T, *response.ApiErrors) {
	criteriaJSON, err := json.Marshal(criteria)
	if err != nil {
		return nil, createAPIErrors(err)
	}
	req, err := http.NewRequest(http.MethodPost, g.endpoints.search, bytes.NewBuffer(criteriaJSON))
	if err != nil {
		return nil, createAPIErrors(err)
	}
	return queryArray[T](g.httpClient, req)
}

func (g *EntityRepository[T]) SearchIds(criteria search.Criteria) (*[]string, *response.ApiErrors) {
	criteriaJSON, err := json.Marshal(criteria)
	if err != nil {
		return nil, createAPIErrors(err)
	}
	req, err := http.NewRequest(http.MethodPost, g.endpoints.searchIds, bytes.NewBuffer(criteriaJSON))
	if err != nil {
		return nil, createAPIErrors(err)
	}
	return queryArray[string](g.httpClient, req)
}

func (g *EntityRepository[T]) Create(entity *T) *response.ApiErrors {
	entityJSON, err := json.Marshal(entity)
	if err != nil {
		return createAPIErrors(err)
	}
	req, err := http.NewRequest(http.MethodPost, g.endpoints.entity, bytes.NewBuffer(entityJSON))
	if err != nil {
		return createAPIErrors(err)
	}
	_, apiError := getBody(g.httpClient, req)
	return apiError
}

func (g *EntityRepository[T]) Update(id string, entity *T) *response.ApiErrors {
	url := g.endpoints.entity + slash + id
	entityJSON, err := json.Marshal(entity)
	if err != nil {
		return createAPIErrors(err)
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(entityJSON))
	if err != nil {
		return createAPIErrors(err)
	}
	_, apiError := getBody(g.httpClient, req)
	return apiError
}

func (g *EntityRepository[T]) Delete(id string) *response.ApiErrors {
	url := g.endpoints.entity + slash + id
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return createAPIErrors(err)
	}
	_, apiError := getBody(g.httpClient, req)
	return apiError
}
