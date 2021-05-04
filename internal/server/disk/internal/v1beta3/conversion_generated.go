// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta3

import (
	unsafe "unsafe"

	v1beta3 "github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta3"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/disk/internal"
)

func autoConvert_v1beta3_DiskIDs_To_internal_DiskIDs(in *v1beta3.DiskIDs, out *internal.DiskIDs) error {
	out.Identifiers = *(*map[string]string)(unsafe.Pointer(&in.Identifiers))
	return nil
}

// Convert_v1beta3_DiskIDs_To_internal_DiskIDs is an autogenerated conversion function.
func Convert_v1beta3_DiskIDs_To_internal_DiskIDs(in *v1beta3.DiskIDs, out *internal.DiskIDs) error {
	return autoConvert_v1beta3_DiskIDs_To_internal_DiskIDs(in, out)
}

func autoConvert_internal_DiskIDs_To_v1beta3_DiskIDs(in *internal.DiskIDs, out *v1beta3.DiskIDs) error {
	out.Identifiers = *(*map[string]string)(unsafe.Pointer(&in.Identifiers))
	return nil
}

// Convert_internal_DiskIDs_To_v1beta3_DiskIDs is an autogenerated conversion function.
func Convert_internal_DiskIDs_To_v1beta3_DiskIDs(in *internal.DiskIDs, out *v1beta3.DiskIDs) error {
	return autoConvert_internal_DiskIDs_To_v1beta3_DiskIDs(in, out)
}

func autoConvert_v1beta3_DiskLocation_To_internal_DiskLocation(in *v1beta3.DiskLocation, out *internal.DiskLocation) error {
	out.Adapter = in.Adapter
	out.Bus = in.Bus
	out.Target = in.Target
	out.LUNID = in.LUNID
	return nil
}

// Convert_v1beta3_DiskLocation_To_internal_DiskLocation is an autogenerated conversion function.
func Convert_v1beta3_DiskLocation_To_internal_DiskLocation(in *v1beta3.DiskLocation, out *internal.DiskLocation) error {
	return autoConvert_v1beta3_DiskLocation_To_internal_DiskLocation(in, out)
}

func autoConvert_internal_DiskLocation_To_v1beta3_DiskLocation(in *internal.DiskLocation, out *v1beta3.DiskLocation) error {
	out.Adapter = in.Adapter
	out.Bus = in.Bus
	out.Target = in.Target
	out.LUNID = in.LUNID
	return nil
}

// Convert_internal_DiskLocation_To_v1beta3_DiskLocation is an autogenerated conversion function.
func Convert_internal_DiskLocation_To_v1beta3_DiskLocation(in *internal.DiskLocation, out *v1beta3.DiskLocation) error {
	return autoConvert_internal_DiskLocation_To_v1beta3_DiskLocation(in, out)
}

func autoConvert_v1beta3_DiskStatsRequest_To_internal_DiskStatsRequest(in *v1beta3.DiskStatsRequest, out *internal.DiskStatsRequest) error {
	out.DiskID = in.DiskID
	return nil
}

func autoConvert_internal_DiskStatsRequest_To_v1beta3_DiskStatsRequest(in *internal.DiskStatsRequest, out *v1beta3.DiskStatsRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_internal_DiskStatsRequest_To_v1beta3_DiskStatsRequest is an autogenerated conversion function.
func Convert_internal_DiskStatsRequest_To_v1beta3_DiskStatsRequest(in *internal.DiskStatsRequest, out *v1beta3.DiskStatsRequest) error {
	return autoConvert_internal_DiskStatsRequest_To_v1beta3_DiskStatsRequest(in, out)
}

func autoConvert_v1beta3_DiskStatsResponse_To_internal_DiskStatsResponse(in *v1beta3.DiskStatsResponse, out *internal.DiskStatsResponse) error {
	out.DiskSize = in.DiskSize
	return nil
}

// Convert_v1beta3_DiskStatsResponse_To_internal_DiskStatsResponse is an autogenerated conversion function.
func Convert_v1beta3_DiskStatsResponse_To_internal_DiskStatsResponse(in *v1beta3.DiskStatsResponse, out *internal.DiskStatsResponse) error {
	return autoConvert_v1beta3_DiskStatsResponse_To_internal_DiskStatsResponse(in, out)
}

func autoConvert_internal_DiskStatsResponse_To_v1beta3_DiskStatsResponse(in *internal.DiskStatsResponse, out *v1beta3.DiskStatsResponse) error {
	out.DiskSize = in.DiskSize
	return nil
}

func autoConvert_v1beta3_GetAttachStateRequest_To_internal_GetAttachStateRequest(in *v1beta3.GetAttachStateRequest, out *internal.GetAttachStateRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_v1beta3_GetAttachStateRequest_To_internal_GetAttachStateRequest is an autogenerated conversion function.
func Convert_v1beta3_GetAttachStateRequest_To_internal_GetAttachStateRequest(in *v1beta3.GetAttachStateRequest, out *internal.GetAttachStateRequest) error {
	return autoConvert_v1beta3_GetAttachStateRequest_To_internal_GetAttachStateRequest(in, out)
}

func autoConvert_internal_GetAttachStateRequest_To_v1beta3_GetAttachStateRequest(in *internal.GetAttachStateRequest, out *v1beta3.GetAttachStateRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_internal_GetAttachStateRequest_To_v1beta3_GetAttachStateRequest is an autogenerated conversion function.
func Convert_internal_GetAttachStateRequest_To_v1beta3_GetAttachStateRequest(in *internal.GetAttachStateRequest, out *v1beta3.GetAttachStateRequest) error {
	return autoConvert_internal_GetAttachStateRequest_To_v1beta3_GetAttachStateRequest(in, out)
}

func autoConvert_v1beta3_GetAttachStateResponse_To_internal_GetAttachStateResponse(in *v1beta3.GetAttachStateResponse, out *internal.GetAttachStateResponse) error {
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_v1beta3_GetAttachStateResponse_To_internal_GetAttachStateResponse is an autogenerated conversion function.
func Convert_v1beta3_GetAttachStateResponse_To_internal_GetAttachStateResponse(in *v1beta3.GetAttachStateResponse, out *internal.GetAttachStateResponse) error {
	return autoConvert_v1beta3_GetAttachStateResponse_To_internal_GetAttachStateResponse(in, out)
}

func autoConvert_internal_GetAttachStateResponse_To_v1beta3_GetAttachStateResponse(in *internal.GetAttachStateResponse, out *v1beta3.GetAttachStateResponse) error {
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_internal_GetAttachStateResponse_To_v1beta3_GetAttachStateResponse is an autogenerated conversion function.
func Convert_internal_GetAttachStateResponse_To_v1beta3_GetAttachStateResponse(in *internal.GetAttachStateResponse, out *v1beta3.GetAttachStateResponse) error {
	return autoConvert_internal_GetAttachStateResponse_To_v1beta3_GetAttachStateResponse(in, out)
}

func autoConvert_v1beta3_ListDiskIDsRequest_To_internal_ListDiskIDsRequest(in *v1beta3.ListDiskIDsRequest, out *internal.ListDiskIDsRequest) error {
	return nil
}

// Convert_v1beta3_ListDiskIDsRequest_To_internal_ListDiskIDsRequest is an autogenerated conversion function.
func Convert_v1beta3_ListDiskIDsRequest_To_internal_ListDiskIDsRequest(in *v1beta3.ListDiskIDsRequest, out *internal.ListDiskIDsRequest) error {
	return autoConvert_v1beta3_ListDiskIDsRequest_To_internal_ListDiskIDsRequest(in, out)
}

func autoConvert_internal_ListDiskIDsRequest_To_v1beta3_ListDiskIDsRequest(in *internal.ListDiskIDsRequest, out *v1beta3.ListDiskIDsRequest) error {
	return nil
}

// Convert_internal_ListDiskIDsRequest_To_v1beta3_ListDiskIDsRequest is an autogenerated conversion function.
func Convert_internal_ListDiskIDsRequest_To_v1beta3_ListDiskIDsRequest(in *internal.ListDiskIDsRequest, out *v1beta3.ListDiskIDsRequest) error {
	return autoConvert_internal_ListDiskIDsRequest_To_v1beta3_ListDiskIDsRequest(in, out)
}

func autoConvert_v1beta3_ListDiskIDsResponse_To_internal_ListDiskIDsResponse(in *v1beta3.ListDiskIDsResponse, out *internal.ListDiskIDsResponse) error {
	if in.DiskIDs != nil {
		in, out := &in.DiskIDs, &out.DiskIDs
		*out = make(map[string]*internal.DiskIDs, len(*in))
		for key, val := range *in {
			newVal := new(*internal.DiskIDs)
			if err := Convert_v1beta3_DiskIDs_To_internal_DiskIDs(*&val, *newVal); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.DiskIDs = nil
	}
	return nil
}

// Convert_v1beta3_ListDiskIDsResponse_To_internal_ListDiskIDsResponse is an autogenerated conversion function.
func Convert_v1beta3_ListDiskIDsResponse_To_internal_ListDiskIDsResponse(in *v1beta3.ListDiskIDsResponse, out *internal.ListDiskIDsResponse) error {
	return autoConvert_v1beta3_ListDiskIDsResponse_To_internal_ListDiskIDsResponse(in, out)
}

func autoConvert_internal_ListDiskIDsResponse_To_v1beta3_ListDiskIDsResponse(in *internal.ListDiskIDsResponse, out *v1beta3.ListDiskIDsResponse) error {
	if in.DiskIDs != nil {
		in, out := &in.DiskIDs, &out.DiskIDs
		*out = make(map[string]*v1beta3.DiskIDs, len(*in))
		for key, val := range *in {
			newVal := new(*v1beta3.DiskIDs)
			if err := Convert_internal_DiskIDs_To_v1beta3_DiskIDs(*&val, *newVal); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.DiskIDs = nil
	}
	return nil
}

func autoConvert_v1beta3_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest(in *v1beta3.ListDiskLocationsRequest, out *internal.ListDiskLocationsRequest) error {
	return nil
}

// Convert_v1beta3_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest is an autogenerated conversion function.
func Convert_v1beta3_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest(in *v1beta3.ListDiskLocationsRequest, out *internal.ListDiskLocationsRequest) error {
	return autoConvert_v1beta3_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest(in, out)
}

func autoConvert_internal_ListDiskLocationsRequest_To_v1beta3_ListDiskLocationsRequest(in *internal.ListDiskLocationsRequest, out *v1beta3.ListDiskLocationsRequest) error {
	return nil
}

// Convert_internal_ListDiskLocationsRequest_To_v1beta3_ListDiskLocationsRequest is an autogenerated conversion function.
func Convert_internal_ListDiskLocationsRequest_To_v1beta3_ListDiskLocationsRequest(in *internal.ListDiskLocationsRequest, out *v1beta3.ListDiskLocationsRequest) error {
	return autoConvert_internal_ListDiskLocationsRequest_To_v1beta3_ListDiskLocationsRequest(in, out)
}

func autoConvert_v1beta3_ListDiskLocationsResponse_To_internal_ListDiskLocationsResponse(in *v1beta3.ListDiskLocationsResponse, out *internal.ListDiskLocationsResponse) error {
	if in.DiskLocations != nil {
		in, out := &in.DiskLocations, &out.DiskLocations
		*out = make(map[string]*internal.DiskLocation, len(*in))
		for key, val := range *in {
			newVal := new(*internal.DiskLocation)
			if err := Convert_v1beta3_DiskLocation_To_internal_DiskLocation(*&val, *newVal); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.DiskLocations = nil
	}
	return nil
}

// Convert_v1beta3_ListDiskLocationsResponse_To_internal_ListDiskLocationsResponse is an autogenerated conversion function.
func Convert_v1beta3_ListDiskLocationsResponse_To_internal_ListDiskLocationsResponse(in *v1beta3.ListDiskLocationsResponse, out *internal.ListDiskLocationsResponse) error {
	return autoConvert_v1beta3_ListDiskLocationsResponse_To_internal_ListDiskLocationsResponse(in, out)
}

func autoConvert_internal_ListDiskLocationsResponse_To_v1beta3_ListDiskLocationsResponse(in *internal.ListDiskLocationsResponse, out *v1beta3.ListDiskLocationsResponse) error {
	if in.DiskLocations != nil {
		in, out := &in.DiskLocations, &out.DiskLocations
		*out = make(map[string]*v1beta3.DiskLocation, len(*in))
		for key, val := range *in {
			newVal := new(*v1beta3.DiskLocation)
			if err := Convert_internal_DiskLocation_To_v1beta3_DiskLocation(*&val, *newVal); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.DiskLocations = nil
	}
	return nil
}

func autoConvert_v1beta3_PartitionDiskRequest_To_internal_PartitionDiskRequest(in *v1beta3.PartitionDiskRequest, out *internal.PartitionDiskRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_v1beta3_PartitionDiskRequest_To_internal_PartitionDiskRequest is an autogenerated conversion function.
func Convert_v1beta3_PartitionDiskRequest_To_internal_PartitionDiskRequest(in *v1beta3.PartitionDiskRequest, out *internal.PartitionDiskRequest) error {
	return autoConvert_v1beta3_PartitionDiskRequest_To_internal_PartitionDiskRequest(in, out)
}

func autoConvert_internal_PartitionDiskRequest_To_v1beta3_PartitionDiskRequest(in *internal.PartitionDiskRequest, out *v1beta3.PartitionDiskRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_internal_PartitionDiskRequest_To_v1beta3_PartitionDiskRequest is an autogenerated conversion function.
func Convert_internal_PartitionDiskRequest_To_v1beta3_PartitionDiskRequest(in *internal.PartitionDiskRequest, out *v1beta3.PartitionDiskRequest) error {
	return autoConvert_internal_PartitionDiskRequest_To_v1beta3_PartitionDiskRequest(in, out)
}

func autoConvert_v1beta3_PartitionDiskResponse_To_internal_PartitionDiskResponse(in *v1beta3.PartitionDiskResponse, out *internal.PartitionDiskResponse) error {
	return nil
}

// Convert_v1beta3_PartitionDiskResponse_To_internal_PartitionDiskResponse is an autogenerated conversion function.
func Convert_v1beta3_PartitionDiskResponse_To_internal_PartitionDiskResponse(in *v1beta3.PartitionDiskResponse, out *internal.PartitionDiskResponse) error {
	return autoConvert_v1beta3_PartitionDiskResponse_To_internal_PartitionDiskResponse(in, out)
}

func autoConvert_internal_PartitionDiskResponse_To_v1beta3_PartitionDiskResponse(in *internal.PartitionDiskResponse, out *v1beta3.PartitionDiskResponse) error {
	return nil
}

// Convert_internal_PartitionDiskResponse_To_v1beta3_PartitionDiskResponse is an autogenerated conversion function.
func Convert_internal_PartitionDiskResponse_To_v1beta3_PartitionDiskResponse(in *internal.PartitionDiskResponse, out *v1beta3.PartitionDiskResponse) error {
	return autoConvert_internal_PartitionDiskResponse_To_v1beta3_PartitionDiskResponse(in, out)
}

func autoConvert_v1beta3_RescanRequest_To_internal_RescanRequest(in *v1beta3.RescanRequest, out *internal.RescanRequest) error {
	return nil
}

// Convert_v1beta3_RescanRequest_To_internal_RescanRequest is an autogenerated conversion function.
func Convert_v1beta3_RescanRequest_To_internal_RescanRequest(in *v1beta3.RescanRequest, out *internal.RescanRequest) error {
	return autoConvert_v1beta3_RescanRequest_To_internal_RescanRequest(in, out)
}

func autoConvert_internal_RescanRequest_To_v1beta3_RescanRequest(in *internal.RescanRequest, out *v1beta3.RescanRequest) error {
	return nil
}

// Convert_internal_RescanRequest_To_v1beta3_RescanRequest is an autogenerated conversion function.
func Convert_internal_RescanRequest_To_v1beta3_RescanRequest(in *internal.RescanRequest, out *v1beta3.RescanRequest) error {
	return autoConvert_internal_RescanRequest_To_v1beta3_RescanRequest(in, out)
}

func autoConvert_v1beta3_RescanResponse_To_internal_RescanResponse(in *v1beta3.RescanResponse, out *internal.RescanResponse) error {
	return nil
}

// Convert_v1beta3_RescanResponse_To_internal_RescanResponse is an autogenerated conversion function.
func Convert_v1beta3_RescanResponse_To_internal_RescanResponse(in *v1beta3.RescanResponse, out *internal.RescanResponse) error {
	return autoConvert_v1beta3_RescanResponse_To_internal_RescanResponse(in, out)
}

func autoConvert_internal_RescanResponse_To_v1beta3_RescanResponse(in *internal.RescanResponse, out *v1beta3.RescanResponse) error {
	return nil
}

// Convert_internal_RescanResponse_To_v1beta3_RescanResponse is an autogenerated conversion function.
func Convert_internal_RescanResponse_To_v1beta3_RescanResponse(in *internal.RescanResponse, out *v1beta3.RescanResponse) error {
	return autoConvert_internal_RescanResponse_To_v1beta3_RescanResponse(in, out)
}

func autoConvert_v1beta3_SetAttachStateRequest_To_internal_SetAttachStateRequest(in *v1beta3.SetAttachStateRequest, out *internal.SetAttachStateRequest) error {
	out.DiskID = in.DiskID
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_v1beta3_SetAttachStateRequest_To_internal_SetAttachStateRequest is an autogenerated conversion function.
func Convert_v1beta3_SetAttachStateRequest_To_internal_SetAttachStateRequest(in *v1beta3.SetAttachStateRequest, out *internal.SetAttachStateRequest) error {
	return autoConvert_v1beta3_SetAttachStateRequest_To_internal_SetAttachStateRequest(in, out)
}

func autoConvert_internal_SetAttachStateRequest_To_v1beta3_SetAttachStateRequest(in *internal.SetAttachStateRequest, out *v1beta3.SetAttachStateRequest) error {
	out.DiskID = in.DiskID
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_internal_SetAttachStateRequest_To_v1beta3_SetAttachStateRequest is an autogenerated conversion function.
func Convert_internal_SetAttachStateRequest_To_v1beta3_SetAttachStateRequest(in *internal.SetAttachStateRequest, out *v1beta3.SetAttachStateRequest) error {
	return autoConvert_internal_SetAttachStateRequest_To_v1beta3_SetAttachStateRequest(in, out)
}

func autoConvert_v1beta3_SetAttachStateResponse_To_internal_SetAttachStateResponse(in *v1beta3.SetAttachStateResponse, out *internal.SetAttachStateResponse) error {
	return nil
}

// Convert_v1beta3_SetAttachStateResponse_To_internal_SetAttachStateResponse is an autogenerated conversion function.
func Convert_v1beta3_SetAttachStateResponse_To_internal_SetAttachStateResponse(in *v1beta3.SetAttachStateResponse, out *internal.SetAttachStateResponse) error {
	return autoConvert_v1beta3_SetAttachStateResponse_To_internal_SetAttachStateResponse(in, out)
}

func autoConvert_internal_SetAttachStateResponse_To_v1beta3_SetAttachStateResponse(in *internal.SetAttachStateResponse, out *v1beta3.SetAttachStateResponse) error {
	return nil
}

// Convert_internal_SetAttachStateResponse_To_v1beta3_SetAttachStateResponse is an autogenerated conversion function.
func Convert_internal_SetAttachStateResponse_To_v1beta3_SetAttachStateResponse(in *internal.SetAttachStateResponse, out *v1beta3.SetAttachStateResponse) error {
	return autoConvert_internal_SetAttachStateResponse_To_v1beta3_SetAttachStateResponse(in, out)
}
