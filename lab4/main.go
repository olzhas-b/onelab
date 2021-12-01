package main

import "encoding/json"
type Data struct {
	FileAttributes             *[]DownloadFileAttribute `json:"fileAttributes" validator:"fileAttributes"`
}
type DownloadFileAttribute struct {
	FileId   *int64  `json:"fileId"`
	FileName *string `json:"fileName"`
	FileSize *int    `json:"fileSize"`
}
func main() {
	downloadFileAttribute := make([]DownloadFileAttribute, 0)
	var data Data = Data{FileAttributes: &downloadFileAttribute}
	var id int64 = 12
	filename := "text1"
	fileSize := 1000024
	downloadFileAttribute = append(downloadFileAttribute, DownloadFileAttribute{FileId: &id, FileName: &filename, FileSize: &fileSize})
	downloadFileAttribute = append(downloadFileAttribute, DownloadFileAttribute{FileId: &id, FileName: &filename, FileSize: &fileSize})
	downloadFileAttribute = append(downloadFileAttribute, DownloadFileAttribute{FileId: &id, FileName: &filename, FileSize: &fileSize})
	text, _ := json.MarshalIndent(&data, "", "\t")
	println(string(text))
}
