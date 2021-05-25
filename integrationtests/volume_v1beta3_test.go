package integrationtests

import (
	"context"
	"fmt"
	"testing"

	diskv1beta3 "github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta3"
	"github.com/kubernetes-csi/csi-proxy/client/api/volume/v1beta3"
	diskv1beta3client "github.com/kubernetes-csi/csi-proxy/client/groups/disk/v1beta3"
	v1beta3client "github.com/kubernetes-csi/csi-proxy/client/groups/volume/v1beta3"
)

func v1beta3VolumeTests(t *testing.T) {
	var volumeClient *v1beta3client.Client
	var diskClient *diskv1beta3client.Client
	var err error

	if volumeClient, err = v1beta3client.NewClient(); err != nil {
		t.Fatalf("Client new error: %v", err)
	}
	defer volumeClient.Close()

	if diskClient, err = diskv1beta3client.NewClient(); err != nil {
		t.Fatalf("DiskClient new error: %v", err)
	}
	defer diskClient.Close()

	vhd, vhdCleanup := diskInit(t)
	defer vhdCleanup()

	listRequest := &v1beta3.ListVolumesOnDiskRequest{
		DiskNumber: vhd.DiskNumber,
	}
	listResponse, err := volumeClient.ListVolumesOnDisk(context.TODO(), listRequest)
	if err != nil {
		t.Fatalf("List response: %v", err)
	}

	volumeIDsLen := len(listResponse.VolumeIds)
	if volumeIDsLen != 1 {
		t.Fatalf("Number of volumes not equal to 1: %d", volumeIDsLen)
	}
	volumeID := listResponse.VolumeIds[0]

	isVolumeFormattedRequest := &v1beta3.IsVolumeFormattedRequest{
		VolumeId: volumeID,
	}
	isVolumeFormattedResponse, err := volumeClient.IsVolumeFormatted(context.TODO(), isVolumeFormattedRequest)
	if err != nil {
		t.Fatalf("Is volume formatted request error: %v", err)
	}
	if isVolumeFormattedResponse.Formatted {
		t.Fatal("Volume formatted. Unexpected !!")
	}

	formatVolumeRequest := &v1beta3.FormatVolumeRequest{
		VolumeId: volumeID,
	}
	_, err = volumeClient.FormatVolume(context.TODO(), formatVolumeRequest)
	if err != nil {
		t.Fatalf("Volume format failed. Error: %v", err)
	}

	isVolumeFormattedResponse, err = volumeClient.IsVolumeFormatted(context.TODO(), isVolumeFormattedRequest)
	if err != nil {
		t.Fatalf("Is volume formatted request error: %v", err)
	}
	if !isVolumeFormattedResponse.Formatted {
		t.Fatal("Volume should be formatted. Unexpected !!")
	}

	t.Logf("VolumeId %v", volumeID)
	volumeStatsRequest := &v1beta3.GetVolumeStatsRequest{
		VolumeId: volumeID,
	}

	volumeStatsResponse, err := volumeClient.GetVolumeStats(context.TODO(), volumeStatsRequest)
	if err != nil {
		t.Fatalf("VolumeStats request error: %v", err)
	}
	// For a volume formatted with 1GB it should be around 1GB, in practice it was 1056947712 bytes or 0.9844GB
	// let's compare with a range of +- 20MB
	if !sizeIsAround(t, volumeStatsResponse.TotalBytes, vhd.InitialSize) {
		t.Fatalf("volumeStatsResponse.TotalBytes reported is not valid, it is %v", volumeStatsResponse.TotalBytes)
	}

	// Resize the disk to twice its size (from 1GB to 2GB)
	// To resize a volume we need to resize the virtual hard disk first and then the partition
	cmd := fmt.Sprintf("Resize-VHD -Path %s -SizeBytes %d", vhd.Path, int64(volumeStatsResponse.TotalBytes*2))
	if out, err := runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %q. Out: %s.", err, cmd, out)
	}

	// Resize the volume to 1.5GB
	oldVolumeSize := volumeStatsResponse.TotalBytes
	newVolumeSize := int64(float32(oldVolumeSize) * 1.5)

	// This is the max partition size when doing a resize to 2GB
	//
	//    Get-PartitionSupportedSize -DiskNumber 7 -PartitionNumber 2 | ConvertTo-Json
	//    {
	//    	"SizeMin":  404725760,
	//    	"SizeMax":  2130689536
	//    }
	resizeVolumeRequest := &v1beta3.ResizeVolumeRequest{
		VolumeId: volumeID,
		// resize the partition to 1.5x times instead
		SizeBytes: newVolumeSize,
	}

	t.Logf("Attempt to resize volume from sizeBytes=%d to sizeBytes=%d", oldVolumeSize, newVolumeSize)

	_, err = volumeClient.ResizeVolume(context.TODO(), resizeVolumeRequest)
	if err != nil {
		t.Fatalf("Volume resize request failed. Error: %v", err)
	}

	volumeStatsResponse, err = volumeClient.GetVolumeStats(context.TODO(), volumeStatsRequest)
	if err != nil {
		t.Fatalf("VolumeStats request after resize error: %v", err)
	}
	// resizing from 1GB to approximately 1.5GB
	if !sizeIsAround(t, volumeStatsResponse.TotalBytes, newVolumeSize) {
		t.Fatalf("VolumeSize reported should be greater than the old size, it is %v", volumeStatsResponse.TotalBytes)
	}

	volumeDiskNumberRequest := &v1beta3.GetDiskNumberFromVolumeIDRequest{
		VolumeId: volumeID,
	}

	volumeDiskNumberResponse, err := volumeClient.GetDiskNumberFromVolumeID(context.TODO(), volumeDiskNumberRequest)
	if err != nil {
		t.Fatalf("GetDiskNumberFromVolumeID failed: %v", err)
	}

	diskNumberString := fmt.Sprintf("%d", volumeDiskNumberResponse.DiskNumber)

	diskStatsRequest := &diskv1beta3.DiskStatsRequest{
		DiskID: diskNumberString,
	}

	diskStatsResponse, err := diskClient.DiskStats(context.TODO(), diskStatsRequest)
	if err != nil {
		t.Fatalf("DiskStats request error: %v", err)
	}

	if diskStatsResponse.DiskSize < 0 {
		t.Fatalf("Invalid disk size was returned %v", diskStatsResponse.DiskSize)
	}

	// Mount the volume
	mountVolumeRequest := &v1beta3.MountVolumeRequest{
		VolumeId:   volumeID,
		TargetPath: vhd.Mount,
	}
	_, err = volumeClient.MountVolume(context.TODO(), mountVolumeRequest)
	if err != nil {
		t.Fatalf("Volume id %s mount to path %s failed. Error: %v", volumeID, vhd.Mount, err)
	}

	// Unmount the volume
	unmountVolumeRequest := &v1beta3.UnmountVolumeRequest{
		VolumeId:   volumeID,
		TargetPath: vhd.Mount,
	}
	_, err = volumeClient.UnmountVolume(context.TODO(), unmountVolumeRequest)
	if err != nil {
		t.Fatalf("Volume id %s mount to path %s failed. Error: %v", volumeID, vhd.Mount, err)
	}
}
