package search

// 기본 검색기를 구현할 defaultMatcher 타입.
type defaultMatcher struct{} // 빈 구조체 : 타입은 필요하지만 타입의 상태관리가 필요하지 않는 경우 사용

// init 함수에서는 기본 검색기를 프로그램에 등록한다.
func init(){
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 함수는 기본 검색기의 동작을 구현한다.
// Matcher 인터페이스를 구현.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error){
	// 값 수신기(value receiver) : 해당 메서드는 지정된 수신기 타입에만 연결됨.
	// => Search 메소드는 defaultMatcher 타입의 값이나 포인터에 대해서만 호출 가능
	return nil, nil
}