package db

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
)

func Compress(data []byte) []byte {
	return compressWithGZip(data)
}

func compressWithGZip(data []byte) []byte {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, 9)
	if err != nil {
		log.Println(err)
		return nil
	}
	if _, err := gz.Write(data); err != nil {
		log.Println(err)
		return nil
	}
	if err := gz.Close(); err != nil {
		log.Println(err)
		return nil
	}
	// fmt.Println(b.Bytes())
	return b.Bytes()
}

func UnCompress(data []byte) []byte {
	return unCompressWithGZip(data)
}

func unCompressWithGZip(cData []byte) []byte {
	reader, err := gzip.NewReader(bytes.NewReader(cData))
	if err != nil {
		log.Println(err)
		return nil
	}
	rData, err := io.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rData
}
