// Author: Turing Zhu
// Date: 2021/8/23 8:03 PM
// File: json_test.go

package shamrock

import "testing"

type CustomType struct {
	StartTime int64  `json:"start_time"`
	Duration  int64  `json:"duration"`
	Type      string `json:"type"`
	Text      string `json:"text"`
	Content   struct {
		TemplateId string `json:"template_id"`
		Params     struct {
			Text      string `json:"text"`
			TestStyle struct {
				Font            string `json:"font"`
				FontSize        string `json:"font_size"`
				FontColor       string `json:"font_color"`
				FontAlpha       int    `json:"font_alpha"`
				FontBold        int    `json:"font_bold"`
				FontItalic      int    `json:"font_italic"`
				FontUline       int    `json:"font_uline"`
				FontAlign       string `json:"font_align"`
				ShadowColor     string `json:"shadow_color"`
				BottomColor     string `json:"bottom_color"`
				BottomAlpha     int    `json:"bottom_alpha"`
				BackgroundColor string `json:"background_color"`
				BackgroundAlpha int    `json:"background_alpha"`
				BorderWidth     int    `json:"border_width"`
				BorderColor     string `json:"border_color"`
			} `json:"test_style"`
		} `json:"params"`
	} `json:"content"`
}

func TestUnmarshalFile(t *testing.T) {
	filePath := "testData/custom_type.json"
	var variable CustomType
	resp, err := UnmarshalFile(filePath, &variable)
	if err != nil {
		t.Fatal(err)
	}
	variable = *(resp.(*CustomType))
}
