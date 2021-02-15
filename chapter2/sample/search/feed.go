package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json" // => 소문자: 비공개 상수로 외부 노출X

// 피드을 처리할 정보를 표현하는 구조체
type Feed struct { // 외부로 노출되는 구조체
	Name	string `json:"site"`
	URI		string `json:"link"`
	Type	string `json:"type"`
}

// RetrieveFeeds 함수는 피드 데이터 파일을 읽어 구조체로 변환한다.
func RetrieveFeeds()([]*Feed, error) {
	// 리턴 값1. Feed 타입 값들의 슬라이스에대한 포인터(주소값)
	// os 패키지를 사용해 파일을 연다
	file, err := os.Open(dataFile)
	// file : File타입 구조체에대한 포인터
	if err != nil {
		return nil, err
	}

	// defer 함수를 이용해 이 함수가 리턴될 때 앞서 열어둔 파일이 닫히도록 한다.
	/*
		defer 키워드
		함수가 리턴된 직후에 실행될 작업을 예약하는 키워드
	*/
	defer file.Close()

	// 파일을 읽어 Feed 구조체의 포인터의 슬라이스로 변환한다.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	/*
		json.NewDecoder
			file을 받아 디코딩할수 있는 Decoder타입의 포인터 값을 return.
			리턴값(포인터값)을 통해 Decode 메소드를 호출애 슬라이스 주소(&feeds)를 전달
		Decode
			func (dec *Decoder) Decode(v interface{}) error
			데이터 파일을 디코딩해 전달한 슬라이스(feeds)에 Feed타입 값들을 채운다.
	*/
	// 호출 함수가 오류를 처리할 수 있으므로 오류 처리는 하지 X
	return feeds, err
}