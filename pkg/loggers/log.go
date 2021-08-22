/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-22 21:45:43
 */
package loggers

import "github/gphper/ginadmin/pkg/loggers/facade"

/*
* 通用info日志
 */
func LogInfo(path string, msg string, info map[string]string) {
	log := facade.NewZaplog(path)
	log.Info(msg, info)
}

/*
* 通用error日志
 */
func LogError(path string, msg string, info map[string]string) {
	log := facade.NewZaplog(path)
	log.Error(msg, info)
}
