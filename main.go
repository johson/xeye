package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Face struct {
	DeviceName string       `json:"device_name"`
	XeyeId     string       `json:"xeye_id"`
	Timestamp  uint64       `json:"timestamp"`
	BaseImg    string       `json:"base_img"`
	Faces      []FaceStruct `json:"faces"`
}

type FaceStruct struct {
	FaceId  int    `json:"face_id"`
	Top     int    `json:"top"`
	Left    int    `json:"left"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Conf    int    `json:"conf"`
	Quality int    `json:"quality"`
	PoseX   int    `json:"pose_x"`
	PoseY   int    `json:"pose_y"`
	PoseZ   int    `json:"pose_z"`
	FaceImg string `json:"face_img"`
	FaceEnd int    `json:"face_end"`
}

func main() {
	http.HandleFunc("/api/v1/postFaceEvent", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=================Beginning of a packet================")
		// fmt.Println(r.Header)
		var face Face
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&face)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Printf("%+v\n", face)
		fmt.Println("<<<<<<<<<<<<<<<<<Ending of a packet<<<<<<<<<<<<<<<<<<<")
	})
	log.Fatal(http.ListenAndServe(":8001", nil))
}
