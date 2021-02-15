package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/webgenie/go-in-action/chapter2/sample/search"
)

// RSS문서를 디코딩해 프로그램내에서 문서 데이터를 처리하기 위해 다음의 구조체 타입을 정의함.
type (
	// item 구조체는 RSS문서 내의 item태그에 정의된 필드들에 대응하는 필드들을 선언한다.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image 구조체는 RSS문서 내의 image태그에 정의된 필드들에 대응하는 필드를 선언
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel 구조체는 RSS문서 내의 channel태그에 정의된 필드에 대응하는 필드를 선언
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          string   `xml:"image"`
		Item           string   `xml:"item"`
	}

	// rssDocumnet 구조체는 RSS문서에 정의된 필드들에 대응하는 필드들을 정의한다.
	rssDocumnet struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// Matcher 인터페이스를 구현하는 rssMatcher 타입을 선언한다.
type rssMatcher struct{}

// init 함수를 통해 프로그램에 검색기를 등록한다.
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search 함수는 지정된 문서에서 검색어를 검색한다.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("피드종류[%s] 사이트[%s] 주소[%s] 에서 검색을 수행합니다.\n", feed.Type, feed.Name, feed.URI)

	// 검색할 데이터를 조회한다.
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// 제목에서 검색어를 검색한다.
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장한다.
		if matched {
			results = append(results, &search.Result{ // => 구조체표현식이므로 &를 사용해 슬라이스가 저장된 메모리 주소를 가져옴
					// append (덧붙일 슬라이스값, 슬라이스에 추가하고자 하는 값)
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 상세내용에서 검색어를 검색한다.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// 검색어가 발견되면 결과에 저장한다.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// HTTP GET 요청을 수행해서 RSS 피드를 요청한 후 결과를 디코딩한다.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("검색할 RSS 피드가 정의되지 않았습니다.")
	}

	// 웹에서 RSS문서를 조회한다.
	resp, err := http.Get(feed.URI)
	// resp : Response 타입에대한 포인터
	if err != nil {
		return nil, err
	}

	// 함수가 리턴될 때 응답 스트림을 닫는다.
	defer resp.Body.Close()

	// 상태 코드가 200인지를 검사해서 올바른 응답을 수신했는지를 확인한다.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP 응답 오류: %d\n", resp.StatusCode)
	}

	// RSS 피드문서를 구조체 타입으로 디코드한다.
	// 호출 ㅎ마수가 에러를 판단할 것이기 때문에 이함수에서는 에러를 처리하지 않는다.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
