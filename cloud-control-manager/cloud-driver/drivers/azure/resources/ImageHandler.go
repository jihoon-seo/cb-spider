package resources

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"

	idrv "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces"
	irs "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces/resources"
)

type AzureImageHandler struct {
	Region        idrv.RegionInfo
	Ctx           context.Context
	Client        *compute.ImagesClient
	VMImageClient *compute.VirtualMachineImagesClient
}

func (imageHandler *AzureImageHandler) setterImage(image compute.Image) *irs.ImageInfo {
	imageInfo := &irs.ImageInfo{
		IId: irs.IID{
			NameId:   *image.Name,
			SystemId: *image.Name,
		},
		GuestOS:      fmt.Sprint(image.ImageProperties.StorageProfile.OsDisk.OsType),
		Status:       *image.ProvisioningState,
		KeyValueList: []irs.KeyValue{{Key: "ResourceGroup", Value: imageHandler.Region.ResourceGroup}},
	}

	return imageInfo
}

func (imageHandler *AzureImageHandler) setterVMImage(image compute.VirtualMachineImage) *irs.ImageInfo {
	imageIdArr := strings.Split(*image.ID, "/")
	imageName := fmt.Sprintf("%s:%s:%s:%s", imageIdArr[8], imageIdArr[12], imageIdArr[14], imageIdArr[16])
	imageInfo := &irs.ImageInfo{
		IId: irs.IID{
			NameId:   imageName,
			SystemId: imageName,
		},
		GuestOS:      fmt.Sprint(image.OsDiskImage.OperatingSystem),
		KeyValueList: []irs.KeyValue{{Key: "ResourceGroup", Value: imageHandler.Region.ResourceGroup}},
	}

	return imageInfo
}

func (imageHandler *AzureImageHandler) CreateImage(imageReqInfo irs.ImageReqInfo) (irs.ImageInfo, error) {

	// @TODO: PublicIP 생성 요청 파라미터 정의 필요
	type ImageReqInfo struct {
		OSType string
		DiskId string
	}

	reqInfo := ImageReqInfo{
		//BlobUrl: "https://md-ds50xp550wh2.blob.core.windows.net/kt0lhznvgx2h/abcd?sv=2017-04-17&sr=b&si=b9674241-fb8e-4cb2-89c7-614d336dc3a7&sig=uvbqvAZQITSpxas%2BWosG%2FGOf6e%2BIBmWNxlUmvARnxiM%3D",
		OSType: "Linux",
		DiskId: "/subscriptions/cb592624-b77b-4a8f-bb13-0e5a48cae40f/resourceGroups/INNO-PLATFORM1-RSRC-GRUP/providers/Microsoft.Compute/disks/inno-test-vm_OsDisk_1_61bf675b990f4aa381d7ee3d766974aa",
		// edited by powerkim for test, 2019.08.13
		//DiskId: "/subscriptions/f1548292-2be3-4acd-84a4-6df079160846/resourceGroups/CB-RESOURCE-GROUP/providers/Microsoft.Compute/disks/vm_name_OsDisk_1_2d63d9cd754c4094b1b1fb6a98c36b71",
	}

	// Check Image Exists
	image, err := imageHandler.Client.Get(imageHandler.Ctx, imageHandler.Region.ResourceGroup, imageReqInfo.IId.NameId, "")
	if image.ID != nil {
		errMsg := fmt.Sprintf("Image with name %s already exist", imageReqInfo.IId.NameId)
		createErr := errors.New(errMsg)
		return irs.ImageInfo{}, createErr
	}

	createOpts := compute.Image{
		ImageProperties: &compute.ImageProperties{
			StorageProfile: &compute.ImageStorageProfile{
				OsDisk: &compute.ImageOSDisk{
					//BlobURI: to.StringPtr(reqInfo.BlobUrl),
					ManagedDisk: &compute.SubResource{
						ID: to.StringPtr(reqInfo.DiskId),
					},
					OsType: compute.OperatingSystemTypes(reqInfo.OSType),
				},
			},
		},
		Location: &imageHandler.Region.Region,
	}

	future, err := imageHandler.Client.CreateOrUpdate(imageHandler.Ctx, imageHandler.Region.ResourceGroup, imageReqInfo.IId.NameId, createOpts)
	if err != nil {
		return irs.ImageInfo{}, err
	}
	err = future.WaitForCompletionRef(imageHandler.Ctx, imageHandler.Client.Client)
	if err != nil {
		return irs.ImageInfo{}, err
	}

	// 생성된 Image 정보 리턴
	imageInfo, err := imageHandler.GetImage(imageReqInfo.IId)
	if err != nil {
		return irs.ImageInfo{}, err
	}
	return imageInfo, nil
}

func (imageHandler *AzureImageHandler) ListImage() ([]*irs.ImageInfo, error) {
	var imageList []*irs.ImageInfo

	publishers, err := imageHandler.VMImageClient.ListPublishers(imageHandler.Ctx, imageHandler.Region.Region)
	if err != nil {
		return nil, err
	}
	/*for _, p := range *publishers.Value {
		if strings.Contains(*p.Name, "azure") {
			spew.Dump(p)
		}
	}*/
	//publishers := []string{"OpenLogic", "CoreOS", "Debian", "SUSE", "RedHat", "Canonical", "MicrosoftWindowsServer"}
	for _, publisher := range *publishers.Value {
		offers, err := imageHandler.VMImageClient.ListOffers(imageHandler.Ctx, imageHandler.Region.Region, *publisher.Name)
		if err != nil {
			continue
		}
		for _, offer := range *offers.Value {
			skus, err := imageHandler.VMImageClient.ListSkus(imageHandler.Ctx, imageHandler.Region.Region, *publisher.Name, *offer.Name)
			if err != nil {
				continue
			}
			for _, sku := range *skus.Value {
				imageInfo, err := imageHandler.GetImage(irs.IID{NameId: fmt.Sprintf("%s:%s:%s", *publisher.Name, *offer.Name, *sku.Name)})
				if err != nil {
					continue
				}
				imageList = append(imageList, &imageInfo)
				if imageListLen := len(imageList); imageListLen%10 == 0 {
					fmt.Println(imageListLen)
				}
			}
		}
	}
	return imageList, nil
}

func (imageHandler *AzureImageHandler) GetImage(imageIID irs.IID) (irs.ImageInfo, error) {
	imageArr := strings.Split(imageIID.NameId, ":")

	// 해당 이미지 publisher, offer, skus 기준 version 목록 조회 (latest 기준 조회 불가)
	vmImageList, err := imageHandler.VMImageClient.List(imageHandler.Ctx, imageHandler.Region.Region, imageArr[0], imageArr[1], imageArr[2], "", to.Int32Ptr(1), "")
	if err != nil {
		return irs.ImageInfo{}, err
	}

	var imageVersion string
	if &vmImageList == nil {
		cblogger.Error(fmt.Sprintf("이미지 조회 실패, ImageIID=%s", imageIID.NameId))
		return irs.ImageInfo{}, errors.New(fmt.Sprintf("could not found image with imageId %s", imageIID.NameId))
	}
	if vmImageList.Value == nil {
		cblogger.Error(fmt.Sprintf("이미지 조회 실패, ImageIID=%s", imageIID.NameId))
		return irs.ImageInfo{}, errors.New(fmt.Sprintf("could not found image with imageId %s", imageIID.NameId))
	}
	if len(*vmImageList.Value) == 0 {
		return irs.ImageInfo{}, errors.New(fmt.Sprintf("could not found image with imageId %s", imageIID.NameId))
	} else {
		latestVmImage := (*vmImageList.Value)[0]
		imageIdArr := strings.Split(*latestVmImage.ID, "/")
		imageVersion = imageIdArr[len(imageIdArr)-1]
	}

	// 1개의 버전 정보를 기준으로 이미지 정보 조회
	vmImage, err := imageHandler.VMImageClient.Get(imageHandler.Ctx, imageHandler.Region.Region, imageArr[0], imageArr[1], imageArr[2], imageVersion)
	if err != nil {
		cblogger.Error(err)
		return irs.ImageInfo{}, err
	}

	imageInfo := imageHandler.setterVMImage(vmImage)
	return *imageInfo, nil
}

func (imageHandler *AzureImageHandler) DeleteImage(imageIID irs.IID) (bool, error) {
	future, err := imageHandler.Client.Delete(imageHandler.Ctx, imageHandler.Region.ResourceGroup, imageIID.NameId)
	if err != nil {
		return false, err
	}
	err = future.WaitForCompletionRef(imageHandler.Ctx, imageHandler.Client.Client)
	if err != nil {
		return false, err
	}
	return true, nil
}
