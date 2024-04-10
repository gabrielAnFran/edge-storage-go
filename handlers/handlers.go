package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	sdk "github.com/aziontech/azionapi-go-sdk/storage"
	"github.com/go-chi/chi"
)

type Bucket struct {
	Name       string `json:"name"`
	EdgeAccess string `json:"edge_access"`
}

type Response struct {
	Count   int      `json:"count"`
	Results []Bucket `json:"results"`
}

const (
	url = "https://api.azion.com/v4/"
)

type Client struct {
	apiClient sdk.APIClient
}

func ListBuckets(w http.ResponseWriter, r *http.Request) {

	var f *http.Client
	client := NewClient(f)

	res, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(resp))

}

func (c *Client) ListBuckets() (*Response, error) {
	ctx := context.Background()

	fmt.Println("Listing buckets....")

	req := c.apiClient.StorageAPI.StorageApiBucketsList(ctx)
	_, httpResp, err := req.Execute()
	if err != nil {
		fmt.Println(httpResp)
		return nil, err
	}
	bytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var response Response

	err = json.Unmarshal([]byte(bytes), &response)
	if err != nil {
		fmt.Println(err)
	}
	return &response, err
}

func CreateBucket(w http.ResponseWriter, r *http.Request) {

	var bucket Bucket
	err := json.NewDecoder(r.Body).Decode(&bucket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var f *http.Client
	client := NewClient(f)

	err = client.CreateBucket(bucket)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("Bucket created successfully"))

}

func (c *Client) CreateBucket(bucketRequest Bucket) error {
	ctx := context.Background()

	bucket := new(sdk.BucketCreate)

	bucket.Name = bucketRequest.Name
	bucket.EdgeAccess = sdk.EdgeAccessEnum(bucketRequest.EdgeAccess)

	fmt.Println("Creating bucket...")

	req := c.apiClient.StorageAPI.StorageApiBucketsCreate(ctx).BucketCreate(*bucket)
	_, httpResp, err := req.Execute()
	if err != nil {
		fmt.Println(httpResp)
		return err
	}
	_, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {

	var f *http.Client
	client := NewClient(f)

	bucketName := chi.URLParam(r, "name")
	if bucketName == "" {
		http.Error(w, "Bucket name is required", http.StatusBadRequest)
		return
	}

	err := client.DeleteBucket(bucketName)
	if err != nil {
		panic(err)
	}
}

func (c *Client) DeleteBucket(bucketName string) error {
	ctx := context.Background()

	fmt.Println("Deleting bucket...")

	_, httpResp, err := c.apiClient.StorageAPI.
		StorageApiBucketsDestroy(ctx, bucketName).Execute()
	if err != nil {
		fmt.Println(httpResp)
		return err
	}
	_, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	return nil
}

func NewClient(f *http.Client) *Client {

	conf := sdk.NewConfiguration()

	conf.HTTPClient = f
	conf.AddDefaultHeader("Authorization", "token <YOUR_API_TOKEN>") // https://www.azion.com/en/documentation/products/guides/personal-tokens/
	conf.AddDefaultHeader("Accept", "application/json;version=3")
	conf.Servers = sdk.ServerConfigurations{
		{URL: url},
	}

	return &Client{
		apiClient: *sdk.NewAPIClient(conf),
	}
}
