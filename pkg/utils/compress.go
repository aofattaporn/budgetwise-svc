package utils

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"io"
	"os"

	"github.com/pkg/errors"
)

func Unzip(zippedFile *bytes.Buffer) (io.ReadCloser, error) {
	fileReader := bytes.NewReader(zippedFile.Bytes())
	zipReader, err := zip.NewReader(fileReader, fileReader.Size())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	file := zipReader.File[0]
	zippedfile, err := file.Open()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return zippedfile, nil
}

func UnzipReader(reader io.ReaderAt, size int64) (io.ReadCloser, error) {
	zipReader, err := zip.NewReader(reader, size)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	file := zipReader.File[0]
	zippedfile, err := file.Open()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return zippedfile, nil
}

func UnzipToFile(zippedFile *bytes.Buffer, filename string) error {
	fileReader := bytes.NewReader(zippedFile.Bytes())
	zipReader, err := zip.NewReader(fileReader, fileReader.Size())
	if err != nil {
		return errors.WithStack(err)
	}

	file := zipReader.File[0]

	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, rc)
	if err != nil {
		return err
	}

	return nil
}

func DecompressGzip(gzipFile *bytes.Buffer) (io.ReadCloser, error) {
	fileReader := bytes.NewReader(gzipFile.Bytes())
	gz, err := gzip.NewReader(fileReader)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gz, nil
}
