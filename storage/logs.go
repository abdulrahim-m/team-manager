package storage

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"time"
)

func SaveGithubLog(data []byte) {
	var pretty bytes.Buffer

	err := json.Indent(&pretty, data, "", "  ")
	if err != nil {
		log.Println(err)
	}

	os.WriteFile("storage/logs/"+time.Now().Format("2006-01-02 15-04-05")+" - Payload", pretty.Bytes(), fs.ModePerm)
}
