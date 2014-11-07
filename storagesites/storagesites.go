package storagesites

import (
	"fmt"
	"github.com/jmcvetta/napping"
)

// User Object
type StorageSite struct {
	SITEID string,
	NAME string,
	SITETYPEID string,
	PROXYID string,
	SITETYPENAME string,
	EMAIL string,
	USERID string,
	FNAME string,
	LNAME string,
	ENABLED string,
	DATECREATED string,
	LOGGINGENABLED string,
	REVISIONCONTROL string,
	MAXREVISIONS string,
	CHECKOUTENABLED string,
	INDEXINGENABLED string,
	SHAREPERMISSIONS string,
	ARGUMENTSPACKET {}
}

type result struct{}

// api call: /storagesites/list.json
func (s StorageSite) Sites(sessionKey string) []{
	// Gets a list of all the storage sites

	return []

}

// api call: /storagesites/{siteId}/details.json
func (s StorageSite) Site(sessionKey, siteID string) string{
	// Gets details of a specific site

	return ""
}