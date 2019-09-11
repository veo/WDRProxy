# WDRProxy
WDRProxy in golang

WDRProxy is a simple forward proxy tool with WAF and CDN functions

## Use:

	./WDRProxy_[OS]_[ARCH] -h
	
	Usage of WDRProxy_[OS]_[ARCH]:
	  -l string
	        listen on ip:port (default "0.0.0.0:8080")
	  -r string
	        reverse proxy addr (default "http://www.xunblog.com")


	./WDRProxy_windows_amd64.exe -l "0.0.0.0:8081" -r "https://www.baidu.com"

	Listening on 0.0.0.0:8081, forwarding to https://www.baidu.com

