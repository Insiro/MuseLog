package out

import (
	"fmt"
	"reflect"
	"strings"
)

// Table 구조체를 테이블 형식으로 출력.
func Table(data interface{}, title ...string) {
	v := reflect.ValueOf(data).Elem()
	t := v.Type()

	// 최대 너비 계산을 위한 변수 초기화
	maxFieldWidth := 0
	maxValueWidth := 0

	// 필드 이름과 값의 최대 길이 계산
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := fmt.Sprintf("%v", v.Field(i).Interface())

		if len(fieldName) > maxFieldWidth {
			maxFieldWidth = len(fieldName)
		}
		if len(fieldValue) > maxValueWidth {
			maxValueWidth = len(fieldValue)
		}
	}

	if len(title) > 0 && title[0] != "" {
		tableTitle := title[0]

		// 타이틀 길이를 고려해 구분선 길이 설정
		totalWidth := maxFieldWidth + maxValueWidth + 3
		titleLen := len(tableTitle)
		if titleLen > totalWidth {
			totalWidth = titleLen
			maxValueWidth += titleLen - totalWidth
		}

		// 타이틀 출력
		fmt.Println("+", strings.Repeat("-", totalWidth), "+")
		fmt.Printf("| %-*s |\n", totalWidth, tableTitle)
	}
	// 필드 구분선 및 헤더 출력
	fmt.Println("+", strings.Repeat("-", maxFieldWidth), "+", strings.Repeat("-", maxValueWidth), "+")
	fmt.Printf("| %-*s | %-*s |\n", maxFieldWidth, "Field", maxValueWidth, "Value")
	fmt.Println("+", strings.Repeat("-", maxFieldWidth), "+", strings.Repeat("-", maxValueWidth), "+")

	// 테이블 본문 출력
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fmt.Printf("| %-*s | %-*v |\n", maxFieldWidth, fieldName, maxValueWidth, fieldValue)
	}

	// 마지막 구분선
	fmt.Println("+", strings.Repeat("-", maxFieldWidth), "+", strings.Repeat("-", maxValueWidth), "+")
}
