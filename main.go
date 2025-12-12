package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadFile(url, destDir string) error {
	//آدرس فایل
	fileName := filepath.Base(url)
	filePath := filepath.Join(destDir, fileName)
	//سخت فایل (حتما بسته شه)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Downloading file ", url)
	start := time.Now()

	//درخواست HTTP (اینم حتما بسته شه)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//ارور هندلینگ 404 یا 500 که فایل ناقص پاک کنه
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response %s", res.Status)
	}

	//ریختن فایل از نت به آدرس محلی
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	
	fmt.Printf("Download %s took %s\n", fileName, time.Since(start))
	return nil
}

func SequentialDownloader(urls []string, destDir string) error {
	// حلقه ساده برای دانلود تک تک فایل ها
	return nil
}

// TODO Result struct

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	// ساخت پوشه به همراه کد دسترسی در صورت نبودن

	// برای هر دانلود یه نتیجه جدا

	//منتظر تموم شدن همه گو روتین ها

	// حداکثر چند همزمان در چنل

	// برای هر URL یه goroutine ساخته شود

	// وقتی همه تموم شدن، کانال results بسته شود

	// نتایج خوانده و نمایش داده شود
	return nil
}

func main() {
	// ساخت urls جدید
	url := "https://news-cdn.varzesh3.com/pictures/2025/12/12/D/ng3ktfp25.webp?w=350"

	// اضافه کردن ConcurrentDownloader با ارور هندلینگ
	err := DownloadFile(url, "./")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// لاگ نهایی
	fmt.Println("Done")
}
