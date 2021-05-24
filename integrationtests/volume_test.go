package integrationtests

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	diskv1beta3 "github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta3"
	"github.com/kubernetes-csi/csi-proxy/client/api/volume/v1beta3"
	diskv1beta3client "github.com/kubernetes-csi/csi-proxy/client/groups/disk/v1beta3"
	v1beta3client "github.com/kubernetes-csi/csi-proxy/client/groups/volume/v1beta3"

	// pre v1beta3 requests have different mappings in some requests, checking with v1beta2 imports
	"github.com/kubernetes-csi/csi-proxy/client/api/volume/v1beta2"
	v1beta2client "github.com/kubernetes-csi/csi-proxy/client/groups/volume/v1beta2"
)

func runPowershellCmd(t *testing.T, command string) (string, error) {
	cmd := exec.Command("powershell", "/c", command)
	t.Logf("Executing command: %q", cmd.String())
	result, err := cmd.CombinedOutput()
	return string(result), err
}

func diskCleanup(t *testing.T, vhdxPath, mountPath, testPluginPath string) {
	if t.Failed() {
		t.Logf("Test failed. Skipping cleanup!")
		t.Logf("Mount path located at %s", mountPath)
		t.Logf("VHDx path located at %s", vhdxPath)
		return
	}
	var cmd, out string
	var err error

	cmd = fmt.Sprintf("Dismount-VHD -Path %s", vhdxPath)
	if out, err = runPowershellCmd(t, cmd); err != nil {
		t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}
	cmd = fmt.Sprintf("rm %s", vhdxPath)
	if out, err = runPowershellCmd(t, cmd); err != nil {
		t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}
	cmd = fmt.Sprintf("rmdir %s", mountPath)
	if out, err = runPowershellCmd(t, cmd); err != nil {
		t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}
	if testPluginPath != "" {
		cmd = fmt.Sprintf("rmdir %s", testPluginPath)
		if out, err = runPowershellCmd(t, cmd); err != nil {
			t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
		}
	}
}

func diskInit(t *testing.T, vhdxPath, mountPath, testPluginPath string) uint32 {
	var cmd, out string
	var err error
	const initialSize = 1 * 1024 * 1024 * 1024
	const partitionStyle = "GPT"

	cmd = fmt.Sprintf("mkdir %s", mountPath)
	if out, err = runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %q. Out: %s", err, cmd, out)
	}

	// Initialize the tests, using powershell directly.
	// Create the new vhdx
	cmd = fmt.Sprintf("New-VHD -Path %s -SizeBytes %d", vhdxPath, initialSize)
	if out, err = runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %q. Out: %s.", err, cmd, out)
	}

	// Mount the vhdx as a disk
	cmd = fmt.Sprintf("Mount-VHD -Path %s", vhdxPath)
	if out, err = runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %q. Out: %s", err, cmd, out)
	}

	var diskNum uint64
	var diskNumUnparsed string
	cmd = fmt.Sprintf("(Get-VHD -Path %s).DiskNumber", vhdxPath)

	if diskNumUnparsed, err = runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s", err, cmd)
	}
	if diskNum, err = strconv.ParseUint(strings.TrimRight(diskNumUnparsed, "\r\n"), 10, 32); err != nil {
		t.Fatalf("Error: %v", err)
	}

	cmd = fmt.Sprintf("Initialize-Disk -Number %d -PartitionStyle %s", diskNum, partitionStyle)
	if _, err = runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s", err, cmd)
	}

	cmd = fmt.Sprintf("New-Partition -DiskNumber %d -UseMaximumSize", diskNum)
	if _, err = runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s", err, cmd)
	}
	return uint32(diskNum)
}

func runNegativeListVolumeRequest(t *testing.T, client *v1beta3client.Client, diskNum uint32) {
	listRequest := &v1beta3.ListVolumesOnDiskRequest{
		DiskNumber: diskNum,
	}
	_, err := client.ListVolumesOnDisk(context.TODO(), listRequest)
	if err == nil {
		t.Fatalf("Empty error. Expected error for disknum:%d", diskNum)
	}
}

func runNegativeIsVolumeFormattedRequest(t *testing.T, client *v1beta3client.Client, volumeID string) {
	isVolumeFormattedRequest := &v1beta3.IsVolumeFormattedRequest{
		VolumeId: volumeID,
	}
	_, err := client.IsVolumeFormatted(context.TODO(), isVolumeFormattedRequest)
	if err == nil {
		t.Fatalf("Empty error. Expected error for volumeID: %s", volumeID)
	}
}

func runNegativeFormatVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID string) {
	formatVolumeRequest := &v1beta3.FormatVolumeRequest{
		VolumeId: volumeID,
	}
	_, err := client.FormatVolume(context.TODO(), formatVolumeRequest)
	if err == nil {
		t.Fatalf("Empty error. Expected error for volume id: %s", volumeID)
	}
}

func runNegativeResizeVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID string, size int64) {
	resizeVolumeRequest := &v1beta3.ResizeVolumeRequest{
		VolumeId:  volumeID,
		SizeBytes: size,
	}
	_, err := client.ResizeVolume(context.TODO(), resizeVolumeRequest)
	if err == nil {
		t.Fatalf("Error empty for volume resize. Volume: %s, Size: %d", volumeID, size)
	}
}

func runNegativeMountVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID, targetPath string) {
	// Mount the volume
	mountVolumeRequest := &v1beta3.MountVolumeRequest{
		VolumeId:   volumeID,
		TargetPath: targetPath,
	}

	_, err := client.MountVolume(context.TODO(), mountVolumeRequest)
	if err == nil {
		t.Fatalf("Error empty for volume(%s) mount to path %s.", volumeID, targetPath)
	}
}

func runNegativeUnmountVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID, targetPath string) {
	// Unmount the volume
	unmountVolumeRequest := &v1beta3.UnmountVolumeRequest{
		VolumeId:   volumeID,
		TargetPath: targetPath,
	}
	_, err := client.UnmountVolume(context.TODO(), unmountVolumeRequest)
	if err == nil {
		t.Fatalf("Empty error. Volume id %s dismount from path %s ", volumeID, targetPath)
	}
}

func runNegativeVolumeStatsRequest(t *testing.T, client *v1beta3client.Client, volumeID string) {
	// Get VolumeStats
	volumeStatsRequest := &v1beta3.GetVolumeStatsRequest{
		VolumeId: volumeID,
	}
	_, err := client.GetVolumeStats(context.TODO(), volumeStatsRequest)
	if err == nil {
		t.Errorf("Empty error. VolumeStats for id %s", volumeID)
	}
}

func negativeVolumeTests(t *testing.T) {
	var client *v1beta3client.Client
	var err error

	if client, err = v1beta3client.NewClient(); err != nil {
		t.Fatalf("Client new error: %v", err)
	}
	defer client.Close()

	// Empty volume test
	runNegativeIsVolumeFormattedRequest(t, client, "")
	// -ve volume id
	runNegativeIsVolumeFormattedRequest(t, client, "-1")

	// Format volume negative tests
	runNegativeFormatVolumeRequest(t, client, "")
	runNegativeFormatVolumeRequest(t, client, "-1")

	// Resize volume negative tests
	runNegativeResizeVolumeRequest(t, client, "", 2*1024*1024)
	runNegativeResizeVolumeRequest(t, client, "-1", 2*1024*1024)

	// Mount volume negative tests
	runNegativeMountVolumeRequest(t, client, "", "")
	runNegativeMountVolumeRequest(t, client, "-1", "")

	// Unmount volume negative tests
	runNegativeUnmountVolumeRequest(t, client, "", "")
	runNegativeUnmountVolumeRequest(t, client, "-1", "")

	runNegativeVolumeStatsRequest(t, client, "")
	runNegativeVolumeStatsRequest(t, client, "-1")
}

// clientCompatibilityTests tests that the API is compatible with versions that are before
// the latest, e.g. that a v1beta2 client can still use the server csi-proxy v1beta3
func clientCompatibilityTests(t *testing.T) {
	// it's intended for this client to be before v1beta3!
	// i.e. don't change it if there are upgrades
	var err error
	v1beta2Client, err := v1beta2client.NewClient()
	if err != nil {
		t.Fatalf("Failed to create new v1beta2 client, err=%+v", err)
	}

	var dismountVolumeRequest *v1beta2.DismountVolumeRequest
	dismountVolumeRequest = &v1beta2.DismountVolumeRequest{
		VolumeId: "",
		Path:     "",
	}

	var wanted string
	wanted = "volume id empty"
	_, err = v1beta2Client.DismountVolume(context.TODO(), dismountVolumeRequest)
	if err == nil || !strings.Contains(err.Error(), wanted) {
		t.Fatalf("Expected error to contain wanted=%s", wanted)
	}
}

func negativeDiskTests(t *testing.T) {
	var client *v1beta3client.Client
	var err error

	if client, err = v1beta3client.NewClient(); err != nil {
		t.Fatalf("Client new error: %v", err)
	}
	defer client.Close()
}

func simpleE2e(t *testing.T) {
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

	s1 := rand.NewSource(time.Now().UTC().UnixNano())
	r1 := rand.New(s1)

	testPluginPath := fmt.Sprintf("C:\\var\\lib\\kubelet\\plugins\\testplugin-%d.csi.io\\", r1.Intn(100))
	mountPath := fmt.Sprintf("%smount-%d", testPluginPath, r1.Intn(100))
	vhdxPath := fmt.Sprintf("%sdisk-%d.vhdx", testPluginPath, r1.Intn(100))

	defer diskCleanup(t, vhdxPath, mountPath, testPluginPath)
	diskNum := diskInit(t, vhdxPath, mountPath, testPluginPath)

	listRequest := &v1beta3.ListVolumesOnDiskRequest{
		DiskNumber: diskNum,
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

	if volumeStatsResponse.TotalBytes == -1 {
		t.Fatalf("volumeStatsResponse.TotalBytes reported is not valid, it is %v", volumeStatsResponse.TotalBytes)
	}

	// Resize the disk to twice its size (from 1GB to 2GB)
	oldSize := volumeStatsResponse.TotalBytes
	newSize := int64(oldSize * 2)

	// To resize a volume we need to resize the virtual hard disk first and then the partition
	cmd := fmt.Sprintf("Resize-VHD -Path %s -SizeBytes %d", vhdxPath, newSize)
	if out, err := runPowershellCmd(t, cmd); err != nil {
		t.Fatalf("Error: %v. Command: %q. Out: %s.", err, cmd, out)
	}

	// This is the max size a partition may be resized to (less than 2GB)
	//
	//    Get-PartitionSupportedSize -DiskNumber 7 -PartitionNumber 2 | ConvertTo-Json
	//    {
	//    	"SizeMin":  404725760,
	//    	"SizeMax":  2130689536
	//    }
	resizeVolumeRequest := &v1beta3.ResizeVolumeRequest{
		VolumeId: volumeID,
		// resize the partition to 1.5x times instead
		SizeBytes: int64(float32(oldSize) * 1.5),
	}

	t.Logf("Attempt to resize volume from sizeBytes=%d to sizeBytes=%d", oldSize, newSize)

	_, err = volumeClient.ResizeVolume(context.TODO(), resizeVolumeRequest)
	if err != nil {
		t.Fatalf("Volume resize request failed. Error: %v", err)
	}

	volumeStatsResponse, err = volumeClient.GetVolumeStats(context.TODO(), volumeStatsRequest)
	if err != nil {
		t.Fatalf("VolumeStats request after resize error: %v", err)
	}

	if volumeStatsResponse.TotalBytes <= oldSize {
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
		TargetPath: mountPath,
	}
	_, err = volumeClient.MountVolume(context.TODO(), mountVolumeRequest)
	if err != nil {
		t.Fatalf("Volume id %s mount to path %s failed. Error: %v", volumeID, mountPath, err)
	}

	// Unmount the volume
	unmountVolumeRequest := &v1beta3.UnmountVolumeRequest{
		VolumeId:   volumeID,
		TargetPath: mountPath,
	}
	_, err = volumeClient.UnmountVolume(context.TODO(), unmountVolumeRequest)
	if err != nil {
		t.Fatalf("Volume id %s mount to path %s failed. Error: %v", volumeID, mountPath, err)
	}
}

func TestVolumeAPIs(t *testing.T) {
	// TODO: This test will fail on Github Actions because Hyper-V is disabled
	// see https://github.com/actions/virtual-environments/pull/2525
	// Skip on GH actions till we find a better solution
	t.Run("SimpleE2E", func(t *testing.T) {
		skipTestOnCondition(t, isRunningOnGhActions())
		simpleE2e(t)
	})

	t.Run("NegativeDiskTests", func(t *testing.T) {
		negativeDiskTests(t)
	})
	t.Run("NegativeVolumeTests", func(t *testing.T) {
		negativeVolumeTests(t)
	})
	t.Run("API Compatibility Tests", func(t *testing.T) {
		clientCompatibilityTests(t)
	})
}
