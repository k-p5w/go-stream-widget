function updateClock() {
    // 現在の時刻を取得する
    const now = new Date();
  
    // 時、分、秒を取得する
    const hour = now.getHours();
    const minute = now.getMinutes();
    const second = now.getSeconds();
  
    // 時刻を表示する
    document.getElementById("clock").textContent = `${hour}:${minute}:${second}`;
  
    // 1秒後に再び関数を呼び出す
    setTimeout(updateClock, 1000);
  }
  
  // ページを読み込んだときに関数を呼び出す
  window.onload = updateClock;

  // 日付を取得する
const date = now.getFullYear() + "/" + (now.getMonth() + 1) + "/" + now.getDate();

// 日付を表示する
document.getElementById("dateclock").textContent += ` ${date}`;
