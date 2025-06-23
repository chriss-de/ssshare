package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/chriss-de/ssshare/internal/backend"
	"github.com/chriss-de/ssshare/internal/helpers"

	"github.com/chriss-de/httpdirfs"
	dl "github.com/chriss-de/httpdirfs/dirlist"
)

func getFile(w http.ResponseWriter, r *http.Request) {
	var err error
	var groupID = r.PathValue("groupID")
	var shareID = r.PathValue("shareID")

	// get filepath from backend
	var filepath string

	if filepath, err = backend.GetFilePath(groupID, shareID); err != nil {
		if _, iErr := helpers.WriteJSON(w, http.StatusBadRequest, err); iErr != nil {
			slog.ErrorContext(r.Context(), "error writing response", "error", iErr, "original_error", err)
		}
		return
	}

	// get fd from OS
	var fd *os.File
	var fdStat os.FileInfo

	if fd, err = os.OpenFile(filepath, os.O_RDONLY, 0666); err != nil {
		if _, iErr := helpers.WriteJSON(w, http.StatusInternalServerError, err); iErr != nil {
			slog.ErrorContext(r.Context(), "error writing response", "error", iErr, "original_error", err)
		}
		return
	}
	if fdStat, err = fd.Stat(); err != nil {
		if _, iErr := helpers.WriteJSON(w, http.StatusInternalServerError, err); iErr != nil {
			slog.ErrorContext(r.Context(), "error writing response", "error", iErr, "original_error", err)
		}
		return
	}

	// switch for serving file or directory
	switch {
	case !fdStat.IsDir():
		w.Header().Set("Content-Disposition", "attachment; filename="+path.Base(filepath))
		http.ServeContent(w, r, path.Base(filepath), fdStat.ModTime(), fd)
	case fdStat.IsDir():
		if !strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, fmt.Sprintf("%s/", r.URL.Path), http.StatusTemporaryRedirect)
		}

		var hdfs *httpdirfs.HttpDirFs
		if hdfs, err = httpdirfs.NewHttpDirFs(filepath, httpdirfs.WithDirectoryListing(dl.NewHtmlDirectoryListing())); err != nil {
			if _, iErr := helpers.WriteJSON(w, http.StatusInternalServerError, err); iErr != nil {
				slog.ErrorContext(r.Context(), "error writing response", "error", iErr, "original_error", err)
			}
			return
		}

		http.StripPrefix(fmt.Sprintf("%s/%s/%s", urlPathSharePrefix, groupID, shareID), http.FileServer(hdfs)).ServeHTTP(w, r)
	}

}
