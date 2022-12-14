package storage

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/ThorstenHans/azb/pkg/config"
)

func getClient() (*azblob.Client, error) {
	cfg := config.Load()
	if !cfg.IsConfigured() {
		return nil, fmt.Errorf("azb not configured. Please invoke azb init")
	}
	url := fmt.Sprintf("https://%s.blob.core.windows.net/", cfg.StorageAccountName)
	if cfg.UseCliAuth {
		cred, err := azidentity.NewAzureCLICredential(nil)
		if err != nil {
			return nil, err
		}
		return azblob.NewClient(url, cred, nil)
	}
	cred, err := azblob.NewSharedKeyCredential(cfg.StorageAccountKey, cfg.StorageAccountKey)
	if err != nil {
		return nil, err
	}
	return azblob.NewClientWithSharedKeyCredential(url, cred, nil)
}

func UploadFile(f string) error {

	client, err := getClient()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(f, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = client.UploadFile(context.TODO(), "uploads", f, file, nil)
	if err != nil {
		return err
	}
	return nil
}

func DownloadFile(f, t string) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	dl, err := os.Create(path.Join(t, f))
	if err != nil {
		return err
	}
	defer dl.Close()
	_, err = client.DownloadFile(context.Background(), "uploads", f, dl, nil)
	if err != nil {
		return err
	}
	return nil
}

func ListFiles() error {
	client, err := getClient()
	if err != nil {
		return err
	}

	pager := client.NewListBlobsFlatPager("uploads", nil)
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			return err
		}
		for _, blob := range page.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
	return nil
}

func DeleteBlob(blobName string) error {
	client, err := getClient()
	if err != nil {
		return err
	}

	_, err = client.DeleteBlob(context.Background(), "uploads", blobName, nil)
	return err
}
