package controller

import (
	"a2goshell/misc"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMagnetInfo(c *gin.Context) {

	shellResult := ""

	//magnet:?xt=urn:btih:CHVZXE53FHTZLV34BBR5CFW7GSZX6A7G&dn=11111
	magnetLink := c.Param("link")

	if magnetLink == "" {
		misc.ReportErrorMessage(c, -1, "link is empty")
		return
	}

	torrentFile := ""
	shellResult, err := misc.ShellCommand("aria2c --bt-metadata-only=true --bt-save-metadata=true magnet:?xt=urn:btih:" + magnetLink)

	if err != nil && shellResult == "" {
		misc.ReportError(c, -2, err)
		return
	}

	lines := strings.Split(shellResult, "\n")

	for _, v := range lines {
		fmt.Println("Processing line:", v)
		if strings.Index(v, "Saving metadata as") > 0 {
			tmps := strings.Split(v, " ")
			torrentFile = tmps[6]
			println(tmps[6])
			break
		}
	}

	shellResult, err = misc.ShellCommand("aria2c " + torrentFile + " -S")

	if err != nil {
		misc.ReportError(c, -1, err)
		return
	}

	table := misc.ReadAria2BtTable(shellResult)

	c.JSON(0, table)
}
