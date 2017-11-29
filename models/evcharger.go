package models

/*
	xml: 공공데이터포털 api 데이터 xml 형식에 맞춤
	json: json 형태로 api 제공, 키 값 재정의
*/
type (
	Response struct {
		//Header    Header    `xml:"header" json:"header"`
		EvCharger EvCharger `xml:"body" json:"body`
	}

	Header struct {
		ResultCode string `xml:"resultCode" json:"resultCode"`
		ResultMsg  string `xml:"resultMsg" json:"resultMsg"`
		TotalCount string `xml:"totalCount" json:"totalCount"`
	}

	EvCharger struct {
		EvcItems EvcItems `xml:"items" json:"evcharger"`
	}

	EvcItems struct {
		EvcItem []EvcItem `xml:"item" json:"items"`
	}

	EvcItem struct {
		StatId    string `xml:"statId" json:"statId"`
		StatNm    string `xml:"statNm" json:"statNm"`
		ChgerId   string `xml:"chgerId" json:"chgerId"`
		ChgerType string `xml:"chgerType" json:"chgerType"`
		Stat      string `xml:"stat" json:"headstater"`
		AddrDoro  string `xml:"addrDoro" json:"addrDoro"`
		Lat       string `xml:"lat" json:"lat"`
		Lng       string `xml:"lng" json:"lng"`
		UseTime   string `xml:"useTime" json:"useTime"`
	}
)
