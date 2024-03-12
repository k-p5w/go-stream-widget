package widget

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 設定用定数
const (
	graphURLorg = "https://pixe.la/v1/users/%v/graphs/%v.svg"
	usrID       = "pppp-fan"
	usrPWD      = "pppp20240312"
	// https://pixe.la/v1/users/popocounter-sample/graphs/popo-tsk-sample.svg
	// popocounter-sample/pppp20221221/popo-tsk-sample
	// pppp-fan/pppp20240312/popo-tsk-counter
	graphName = "popo-tsk-counter"
	inputURL  = "https://go-stream-widget.vercel.app/widget/pppp_counter.html"
	// put
	pointAdd      = "https://pixe.la/v1/users/%v/graphs/%v/add"
	pointSet      = "https://pixe.la/v1/users/%v/graphs/%v"
	pointSubtract = "https://pixe.la/v1/users/%v/graphs/%v/subtract"
	// $ curl -X PUT https://pixe.la/v1/users/a-know/graphs/test-graph/subtract -H 'X-USER-TOKEN:thisissecret' -d '{"quantity":"5"}'
// {"message":"Success.","isSuccess":true}
// https://pixe.la/v1/users/popocounter-sample/graphs/popo-tsk-sample

)

// レスポンスボディをパースしてアクセストークンを取得
type GraphPutData struct {
	Date string `json:"date"`
	Qty  string `json:"quantity"`
}

func putPoint(usr string, pwd string, url string, pointDATE string, pointQTY string) {
	fmt.Println("start-putPoint.")

	apiURL := fmt.Sprintf(url, usr, graphName)

	requestData := map[string]string{
		"date":     pointDATE,
		"quantity": pointQTY,
	}
	jsonValue, _ := json.Marshal(requestData)

	// HTTPリクエストの作成
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		// エラー処理
		panic(err)
	}

	fmt.Println(apiURL)
	fmt.Println(requestData)
	req.Header.Set("X-USER-TOKEN", pwd)

	// HTTPリクエストの実行
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// エラー処理
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("ERR-CODE :%v \n ", resp)
		return
	}

}

func TskCounter(w http.ResponseWriter, r *http.Request) {

	// POSTされたパラメータを取得する
	regVal := r.FormValue("incrementValue")
	// // 文字列を float64 型に変換
	// incrementValue, err := strconv.ParseFloat(regVal, 64)
	// if err != nil {
	// 	// 変換に失敗した場合
	// 	http.Error(w, "Failed to parse incrementValue: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }
	regYmd := r.FormValue("selectedDate")
	putPoint(usrID, usrPWD, pointSet, regYmd, regVal)
	graphURL := fmt.Sprintf(graphURLorg, usrID, graphName)
	fmt.Printf("%v\n%v\n%v\n%v\n%v \n", usrID, usrPWD, pointSet, regYmd, regVal)
	/*
		                       https://pixe.la/v1/users/popocounter-sample/graphs/popo-tsk-sample
				$ curl -X POST https://pixe.la/v1/users/popocounter-sample/graphs/popo-tsk-sample -H 'X-USER-TOKEN:pppp20221221' -d '{"date":"20240312","quantity":"3"}'
	*/

	// if incrementValue > 0 {
	// 	pointSet

	// 	putPoint(usrID, usrPWD, pointAdd, regYmd, strconv.Itoa(int(math.Abs(incrementValue))))
	// } else {
	// 	putPoint(usrID, usrPWD, pointSubtract, regYmd, strconv.Itoa(int(math.Abs(incrementValue))))
	// }
	// // 取得したパラメータを出力する
	// fmt.Fprintf(w, "param1: %v\n", regVal)
	// fmt.Fprintf(w, "param2: %v\n", regYmd)

	// HTMLの描画
	viewPage := fmt.Sprintf(`	<!DOCTYPE html>
	<html lang="ja">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>舌打ちチェッカー</title>
  <style>
    body {
      font-family: monospace;
      background-color: #fff;
    }
    .dot-container {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
      align-items: center;
    }
    .dot {
      width: 10px;
      height: 10px;
      border: 1px solid #000;
      margin: 1px;
    }
    .dot-filled {
      background-color: #000;
    }
    .score {
      font-size: 24px;
      font-weight: bold;
      margin: 10px;
    }
  </style>
</head>
<body>
<img src="%v" alt="PPP SVG">
<a href="%v">舌打ち登録画面へ</a>
</body>
</html>	`, graphURL, inputURL)

	fmt.Fprint(w, viewPage)
}
