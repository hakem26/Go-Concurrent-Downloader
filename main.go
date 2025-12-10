package main

func DownloadFile(url, destDir string) error {
	//آدرس فایل

	//تبدیل به آدرس محلی

	//سخت فایل (حتما بسته شه)

	//درخواست HTTP (اینم حتما بسته شه)

	//ارور هندلینگ 404 یا 500 که فایل ناقص پاک کنه

	//ریختن فایل از نت به آدرس محلی

}

func SequentialDownloader(urls []string, destDir string) error {
	// حلقه ساده برای دانلود تک تک فایل ها

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
}

func main() {
	// ساخت urls جدید

	// اضافه کردن ConcurrentDownloader با ارور هندلینگ

	// لاگ نهایی
}
