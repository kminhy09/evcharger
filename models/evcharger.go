package models

import (
	"database/sql"
	"fmt"
)

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
		EvCharger EvcItems `xml:"items" json:"evcharger"`
	}

	EvcItems struct {
		EvcItems []EvcItem `xml:"item" json:"items"`
	}

	EvcItem struct {
		EvchargerID int     `xml:"evchargerId" json:"evchargerId"`
		StatID      string  `xml:"statId" json:"statId"`
		StatNm      string  `xml:"statNm" json:"statNm"`
		ChgerID     string  `xml:"chgerId" json:"chgerId"`
		ChgerType   string  `xml:"chgerType" json:"chgerType"`
		HeadStater  int     `xml:"headstater" json:"headstater"`
		AddrDoro    string  `xml:"addrDoro" json:"addrDoro"`
		Lat         float32 `xml:"lat" json:"lat"`
		Lng         float32 `xml:"lng" json:"lng"`
		UseTime     string  `xml:"useTime" json:"useTime"`
	}
)

func GetEvChargers(db *sql.DB) EvCharger {
	sqlStatement := "SELECT * FROM ev_charger"
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := EvCharger{}
	for rows.Next() {
		evcItem := EvcItem{}
		err2 := rows.Scan(
			&evcItem.EvchargerID,
			&evcItem.StatID,
			&evcItem.StatNm,
			&evcItem.ChgerID,
			&evcItem.ChgerType,
			&evcItem.HeadStater,
			&evcItem.AddrDoro,
			&evcItem.Lat,
			&evcItem.Lng,
			&evcItem.UseTime)
		if err2 != nil {
			panic(err2)
		}
		result.EvCharger.EvcItems = append(result.EvCharger.EvcItems, evcItem)
	}
	return result
}

// PutTask into DB
func PutEvChargers(db *sql.DB, evcItems []EvcItem) error {
	var err error

	fmt.Println(len(evcItems))
	for index, evcItem := range evcItems {

		//fmt.Printf("##### evcItems %d >>> %s\n", index, string(evcItem.StatId))

		sqlStatement := `
						INSERT INTO ev_charger (
								statid, 
								statnm, 
								chgerid, 
								chgertype, 
								headstater, 
								addrdoro, 
								lat, 
								lng, 
								usetime
						) 
						VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
						RETURNING evchargerid;
						`
		result := 0
		err = db.QueryRow(sqlStatement, evcItem.StatID, evcItem.StatNm, evcItem.ChgerID, evcItem.ChgerType, evcItem.HeadStater, evcItem.AddrDoro, evcItem.Lat, evcItem.Lng, evcItem.UseTime).Scan(&result)
		if err != nil {
			panic(err)
		}
		fmt.Println(index, evcItem)

	}

	return err
}

// DeleteTask from DB
func DeleteEvChargers(db *sql.DB) (int64, error) {
	sqlStatement := `DELETE FROM ev_charger;`

	res, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return count, err
}
