package common

const (
	USER_STATE_LOGOUT, USER_STATE_LOGIN = 0, 1
)

const (
	SERVER_STATE_OFFLINE, SERVER_STATE_ONLINE = 0, 1
)

const (
	SERVER_TYPE_CMS, SERVER_TYPE_SIP, SERVER_TYPE_MTS, SERVER_TYPE_APS = 0, 1, 2, 3
)

const (
	PAGE_TITLE_PREV, PAGE_TITLE_NEXT = "<<", ">>"
)

const (
	USER_NAME_ADMIN = "admin"
)

const (
	STATISTICS_LIST_CAPACITY = 50
)

const (
	STATISTICS_TIME_INTERVAL_SEC = 1
)

const (
	KB_SIZE, MB_SIZE, GB_SIZE, TB_SIZE, PB_SIZE = 1024, 1024 * 1024, 1024 * 1024 * 1024, 1024 * 1024 * 1024 * 1024, 1024 * 1024 * 1024 * 1024 * 1024
)

const (
	DEFAULT_CONFIG_FILE_NAME = "config.xml"
)

const (
	SERVER_COMMAND_QUERYSTATUS, SERVER_COMMAND_QUIT = "QueryStatus", "Quit"
)

const (
	SERVER_REQUEST_XML = `<?xml version="1.0" ?>
<Request>
  <SN>%d</SN>
  <CmdType>%s</CmdType>
</Request>`
)

const (
	RET_CODE_SUCCESS, RET_CODE_TIMEOUT = 0, -1
)

const (
	XML_ENCODE_UTF8, XML_ENCODE_GB2312 = "UTF-8", "gb2312"
)

const (
	SWD_RUN_ENV_DEBUG = "debug"
	SWD_RUN_ENV_PROD  = "prod"
)
