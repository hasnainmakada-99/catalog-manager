/**
 * Copyright 2020 Napptive
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package catalog_manager

import (
	"context"
	"fmt"
	"github.com/napptive/catalog-manager/internal/pkg/entities"
	grpc_catalog_common_go "github.com/napptive/grpc-catalog-common-go"
	grpc_catalog_go "github.com/napptive/grpc-catalog-go"
	"github.com/napptive/nerrors/pkg/nerrors"
	"io"
)

const appAddedMsg = "%s added to catalog"
const appRemovedMsg = "%s removed from catalog"

type Handler struct {
	manager Manager
}

// TODO: Check update/get concurrency

func NewHandler(manager Manager) *Handler {
	return &Handler{manager: manager}
}

// Add a new application in the catalog
func (h *Handler) Add(server grpc_catalog_go.Catalog_AddServer) error {

	// TODO: create a map to load the files and avoid send a file twice
	applicationName := ""
	var applicationFiles []*entities.FileInfo

	for {
		// From https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc-1
		request, err := server.Recv()
		if err == io.EOF {
			if err := h.manager.Add(applicationName, applicationFiles); err != nil {
				return nerrors.FromError(err).ToGRPC()
			} else {
				return server.SendAndClose(&grpc_catalog_common_go.OpResponse{
					Status:     grpc_catalog_common_go.OpStatus_SUCCESS,
					StatusName: grpc_catalog_common_go.OpStatus_SUCCESS.String(),
					UserInfo:   fmt.Sprintf(appAddedMsg, applicationName),
				})
			}
		}
		if err != nil {
			return nerrors.FromError(err).ToGRPC()
		}

		// the first time save the application name
		if applicationName == "" {
			applicationName = request.ApplicationName
		}

		// if the name is other than the saved one -> ERROR
		// it is not allowed sending different applications in the same stream
		if request.ApplicationName != applicationName {
			sErr := nerrors.NewFailedPreconditionError("not allowed sending different applications in the same stream")
			return nerrors.FromError(sErr).ToGRPC()
		}
		// Append the files
		applicationFiles = append(applicationFiles, entities.NewFileInfo(request.File))

	}
}

// Download an application from catalog
func (h *Handler) Download(request *grpc_catalog_go.DownloadApplicationRequest, server grpc_catalog_go.Catalog_DownloadServer) error {
	if err := request.Validate(); err != nil {
		return nerrors.FromError(err).ToGRPC()
	}

	files, err := h.manager.Download(request.ApplicationName)
	if err != nil {
		return nerrors.FromError(err).ToGRPC()
	}

	for _, file := range files {
		if err := server.Send(file.ToGRPC()); err != nil {
			return nerrors.NewInternalErrorFrom(err, "unable to send the file").ToGRPC()
		}
	}

	return nil
}

//Remove an application from the catalog
func (h *Handler) Remove(ctx context.Context, request *grpc_catalog_go.RemoveApplicationRequest) (*grpc_catalog_common_go.OpResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.FromError(err).ToGRPC()
	}

	if err := h.manager.Remove(request.ApplicationName); err != nil {
		return nil, nerrors.FromError(err).ToGRPC()
	}

	return &grpc_catalog_common_go.OpResponse{
		Status:     grpc_catalog_common_go.OpStatus_SUCCESS,
		StatusName: grpc_catalog_common_go.OpStatus_SUCCESS.String(),
		UserInfo:   fmt.Sprintf(appRemovedMsg, request.ApplicationName),
	}, nil
}

// List returns a list with all the applications
func (h *Handler) List(ctx context.Context, request *grpc_catalog_common_go.EmptyRequest) (*grpc_catalog_go.ApplicationList, error) {
	returned, err := h.manager.List()
	if err != nil {
		return nil, nerrors.FromError(err).ToGRPC()
	}

	summaryList := make([]*grpc_catalog_go.ApplicationSummary, 0)
	for _, app := range returned {
		summaryList = append(summaryList, app.ToApplicationSummary())
	}
	return &grpc_catalog_go.ApplicationList{Applications: summaryList}, nil

}

// Info returns the detail of a given application
func (h *Handler) Info(ctx context.Context, request *grpc_catalog_go.InfoApplicationRequest) (*grpc_catalog_go.InfoApplicationResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.FromError(err).ToGRPC()
	}

	retrieved, err := h.manager.Get(request.ApplicationName)
	if err != nil {
		return nil, nerrors.FromError(err).ToGRPC()
	}
	return &grpc_catalog_go.InfoApplicationResponse{
		RepositoryName:  retrieved.Repository,
		ApplicationName: retrieved.ApplicationName,
		Tag:             retrieved.Tag,
		MetadataFile:    []byte(retrieved.Metadata),
		ReadmeFile:      []byte(retrieved.Readme),
		Metadata:        retrieved.MetadataObj.ToGRPC(),
	}, nil
}
