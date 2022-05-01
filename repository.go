package shopware6_admin_go_client

import (
	"bytes"
	"encoding/json"
	"github.com/develist/shopware6-admin-go-client/entity"
	"net/http"
	"reflect"
	"strings"
)

type EntityRepository[T any] struct {
	endpoints  endpoints
	httpClient *http.Client
}

type RepositoryConfiguration struct {
	ApiBaseUrl string
	HttpClient *http.Client
}

type endpoints struct {
	entity    string
	search    string
	searchIds string
}

// CreateEntityRepository Creates a new generic repository. The Shopware 6 base API URL and a preconfigured
// http.Client is needed.
func CreateEntityRepository[T any](config *RepositoryConfiguration) *EntityRepository[T] {
	var a T
	endpoint := strings.ToLower(toKebabCase(reflect.TypeOf(a).Name()))
	client := config.HttpClient
	entityEndPoint := strings.TrimRight(config.ApiBaseUrl, "/") + "/" + endpoint
	searchEndPoint := strings.TrimRight(config.ApiBaseUrl, "/") + "/search/" + endpoint
	searchIdsEndPoint := strings.TrimRight(config.ApiBaseUrl, "/") + "/search-ids/" + endpoint
	endpoints := endpoints{entityEndPoint, searchEndPoint, searchIdsEndPoint}
	return &EntityRepository[T]{endpoints, client}
}

func (g *EntityRepository[T]) GetDetail(id string) (*T, error) {
	url := g.endpoints.entity + "/" + id
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
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

func (g *EntityRepository[T]) GetList() (*[]T, error) {
	req, err := http.NewRequest(http.MethodGet, g.endpoints.entity, nil)
	if err != nil {
		return nil, err
	}
	return queryArray[T](g.httpClient, req)
}

func (g *EntityRepository[T]) Search(criteria entity.Criteria) (*[]T, error) {
	criteriaJSON, err := json.Marshal(criteria)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, g.endpoints.search, bytes.NewBuffer(criteriaJSON))
	if err != nil {
		return nil, err
	}
	return queryArray[T](g.httpClient, req)
}

func (g *EntityRepository[T]) SearchIds(criteria entity.Criteria) (*[]string, error) {
	criteriaJSON, err := json.Marshal(criteria)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, g.endpoints.searchIds, bytes.NewBuffer(criteriaJSON))
	if err != nil {
		return nil, err
	}
	return queryArray[string](g.httpClient, req)
}

func (g *EntityRepository[T]) Create(item *T) error {
	entityJSON, err := json.Marshal(item)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, g.endpoints.entity, bytes.NewBuffer(entityJSON))
	if err != nil {
		return err
	}
	_, err = getBody(g.httpClient, req)
	if err != nil {
		return err
	}
	return nil
}

func (g *EntityRepository[T]) Update(id string, entity *T) error {
	url := g.endpoints.entity + "/" + id
	entityJSON, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(entityJSON))
	if err != nil {
		return err
	}
	_, err = getBody(g.httpClient, req)
	return err
}

func (g *EntityRepository[T]) Delete(id string) error {
	url := g.endpoints.entity + "/" + id
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	_, err = getBody(g.httpClient, req)
	if err != nil {
		return err
	}
	return nil
}
