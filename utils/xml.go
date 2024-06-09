package utils

import "fmt"

var urlXml = "<appmsg appid=\"\" sdkver=\"0\">\n\t\t<title>%s</title>\n\t\t<des>%s</des>\n\t\t<type>5</type>\n\t\t<url>%s</url>\n\t\t<appattach>\n\t\t\t<cdnthumbaeskey />\n\t\t\t<aeskey />\n\t\t</appattach>\n\t\t<thumburl>%s</thumburl>\n\t</appmsg>\n\t<scene>0</scene>\n\t<appinfo>\n\t\t<version>1</version>\n\t\t<appname></appname>\n\t</appinfo>\n\t<commenturl></commenturl>"

func BuildUrlXml(title, desc, url, thumbUrl string) string {
	return fmt.Sprintf(urlXml, title, desc, url, thumbUrl)
}
