package widget

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)



func arabicToKanji(num int) string {

	// 漢字変換マップを作成します
	kansujiMap := map[int]string{
		0: "",
		1: "一",
		2: "二",
		3: "三",
		4: "四",
		5: "五",
		6: "六",
		7: "七",
		8: "八",
		9: "九",
	}

	// 日付の二桁を文字列に変換します
	dayStr := strconv.Itoa(num)

	// 漢字に変換した結果を格納する変数を初期化します
	var result string

	// 生成AIのロジックでうまく行かないので手作りする
	kanjiDigit := make([]string, len(dayStr))
	// 枠を作って、そこに対してセットするだけでいいはず。

	// 日付の各桁を漢字に変換します
	for i, digit := range dayStr {

		digitInt, _ := strconv.Atoi(string(digit))
		// 13なら、1→3の順にくる
		// 数値→漢字変換
		kanjiDigit[i] = kansujiMap[digitInt]

		// 1桁だけの数値の場合
		if len(dayStr) == 1 {
			result = kanjiDigit[i]
		} else {
			// 2桁ある数値の十の位の場合
			if i == 0 {
				// 13→十三/22→二十二/三十一→31
				switch digitInt {
				case 1:
					result = "十"
				case 2, 3:
					result = kanjiDigit[i] + "十"
				}

			} else {
				// 2桁ある数値の一の位の場合
				result += kanjiDigit[i]
			}

		}
	}

	return result
}

func getCalendar() (string, string) {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	timeNow := time.Now().In(loc)
	// formattedSS := timeNow.Format("05")

	// myTips = getTips(formattedSS)

	var weekdays = []string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}
	weekday := timeNow.Weekday()

	year := timeNow.Year() - 2018
	yearJP := arabicToKanji(year)
	monthJP := arabicToKanji(int(timeNow.Month()))
	dayJP := arabicToKanji(int(timeNow.Day()))
	dispItem := fmt.Sprintf("令和%v年 %v月 %v日 %v \n", yearJP, monthJP, dayJP, weekdays[weekday])

	// 現在の日付を取得する
	today := time.Now()

	currentYear := time.Now().Year()
	var yearALLDays float64
	if IsLeapYear(currentYear) {
		yearALLDays = 366
	} else {
		yearALLDays = 365
	}

	// 今年の最終日を取得する
	endOfYear := time.Date(today.Year(), time.December, 31, 0, 0, 0, 0, time.Local)

	// 残り日数を計算する
	remainingDays := endOfYear.Sub(today).Hours() / 24

	atDays := (remainingDays / yearALLDays) * 100

	yearHPper := fmt.Sprintf("%.2f%%", atDays)

	// 残り日数を表示する
	yearHP := fmt.Sprintf("今年の残り日数:%.1f 【%s】", remainingDays, yearHPper)

	return dispItem, yearHP
}
func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("start-CalendarHandler.")

	// これで表示できると思ったけど無理だったな、サーバ起動していないから？
	waitSec := 20
	dispItem, remainHP := getCalendar()

	optionItem := ""
	otherTZ := `<div class="clockTimerOther" ><div  id="clockHHMMOther">	</div></div>`
	// getパラメータの解析
	q := r.URL.Query()
	typeValue := q.Get("type")
	fileName := filepath.Base(typeValue)

	ext := filepath.Ext(fileName)
	runType := fileName[:len(fileName)-len(ext)]

	optionItem = "<br>" + remainHP

	otherTimeZone := "America/Edmonton"
	switch runType {
	case "simple":
		optionItem = ""
		otherTZ = ""
	case "remain":

		otherTZ = ""
	default:
		// TZ指定っぽい設定なら
		if strings.Contains(runType, "_") {
			tz := strings.Replace(runType, "_", "/", -1)
			otherTimeZone = tz
		}

	}

	fmt.Println(dispItem)
	viewPage := fmt.Sprintf(`	<!DOCTYPE html>
	<html lang="ja">
<head>
  <title>ストリーマー.widget[%v]</title>
  <link rel="stylesheet" href="https://scrapbox.io/api/code/cybernote/twitchbot-translatetext2nd/clock.css">
  
  <link rel="stylesheet" href="style.css">
  <link rel="stylesheet" href="https://scrapbox.io/api/code/cybernote/JPstreamer/style.css">

  </head>
  <body>
  <div class="container">
    <div class="black-board">%v</div>
    <div class="black-board-row" >%v</div>
	<div class="clockTimer" >
    <div  id="clockHHMM"</div>
	</div>
	%v

  </div>
  
  <script>
  function refreshPage(seconds) {
	  setTimeout(function () {
		  location.reload();
	  }, seconds * 1000);
  }
  refreshPage(%v);
  
</script>
<script>
function updateTime() {
    var now = new Date();
    var hours = now.getHours();
    var minutes = now.getMinutes();
    
    // ゼロパディング
    hours = ('0' + hours).slice(-2);
    minutes = ('0' + minutes).slice(-2);

    var timeString = hours + ':' + minutes;
    document.getElementById('clockHHMM').innerText = timeString;
}

function updateTimeOther() {
	var now = new Date();
  
	// JST以外の時間を表示
	var timeString = new Intl.DateTimeFormat('ja-JP', {
	  hour: 'numeric',
	  minute: 'numeric',
	  hour12: false,
	  timeZone: '%v'
	}).format(now);
	document.getElementById('clockHHMMOther').innerText = timeString;
  }

// 初回の更新
updateTime();
updateTimeOther();

// 1分ごとに更新
setInterval(updateTime, 60000);
</script>
<div id="dateclock"></div>
<div id="clock"></div>
<script src="https://scrapbox.io/api/code/cybernote/twitchbot-translatetext2nd/clock.js">
</script>

  </body>
</html>
	`, fileName, dispItem, optionItem, otherTZ, waitSec, otherTimeZone)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Vary", "Accept-Encoding")
	svgPage := "<h1>エラーが発生しました.</h1>"

	typeName := q.Get("type")
	if len(typeName) > 0 {

		// HTMLで終わっていること
		if strings.HasSuffix(typeName, ".html") {

		} else {
			return
		}
	} else {
		viewPage = svgPage
	}
	fmt.Println(optionItem)
	// HTMLの描画
	fmt.Fprint(w, viewPage)
}

// getTips is TIPSの読み込み
func getTips(keySS string) string {

	// 末尾の秒を取り出す
	msgrune := []rune(keySS)
	resultEnd := string(msgrune[len(msgrune)-1])

	// 現在の日付を取得する
	today := time.Now()

	currentYear := time.Now().Year()
	var yearALLDays float64
	if IsLeapYear(currentYear) {
		yearALLDays = 366
	} else {
		yearALLDays = 365
	}

	// 今年の最終日を取得する
	endOfYear := time.Date(today.Year(), time.December, 31, 0, 0, 0, 0, time.Local)

	// 残り日数を計算する
	remainingDays := endOfYear.Sub(today).Hours() / 24

	atDays := remainingDays / yearALLDays
	yearHPper := fmt.Sprintf("%.2f%%", atDays)

	// 残り日数を表示する
	yearHP := fmt.Sprintf("今年の残り日数:%.1f 【%v】", remainingDays, yearHPper)

	sep := " / "
	JPonly := "i'm sorry. I'm bad at Emglish. Japanese only."

	wishFood := "焼き鳥たべたい"
	likeFood := "鳥皮すき"
	liveQuiz := "記憶喪を消して観たい名作Best10の8番目は？"
	ret := "[hello]みたいに[]で囲うとJPに翻訳されるよ！"
	retEN := "If you enclose text in [ ], it will be translated to Japanese!"
	img := ` <img src="https://i.gyazo.com/c7ea69871ace51e44ab49ac027dd53ce.png"
	alt="おじさん" width="50" height="50"   />`
	eg := "e.g.[you sound like your cold is getting better]"
	// 基本a用の組み立て
	ret += sep
	ret += retEN
	ret += sep
	ret += eg
	switch resultEnd {
	case "0":
		ret = JPonly
	case "1":
		ret = yearHP
	case "2":
		ret = wishFood
	case "3":
		ret = likeFood
	case "4":
		ret = liveQuiz
	case "5":
		ret = img
	case "6":
		ret = yearHP
	case "7":
		ret = yearHP
	case "8":
		ret = yearHP
	case "9":
		ret = yearHP
	default:
	}

	return ret
}
func IsLeapYear(year int) bool {
	// 4で割り切れる年は閏年
	// ただし、100で割り切れる年は平年
	// ただし、400で割り切れる年はまた閏年
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
