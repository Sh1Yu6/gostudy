package meta

import (
	"sort"

	"github.com/gostudy/project4/db"
)

// FileMeta:
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, 0)
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:len(fMetaArray)]
}

func RemoveFileMeta(filehash string) {
	delete(fileMetas, filehash)
}

func UpdateFileMetaDB(fmeta FileMeta) bool {
	return db.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName,
		fmeta.Location, fmeta.FileSize)
}

func GetFileMetaDB(filehash string) (FileMeta, error) {
	tfile, err := db.GetFileMeta(filehash)
	if err != nil {
		return FileMeta{}, err
	}

	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil
}
