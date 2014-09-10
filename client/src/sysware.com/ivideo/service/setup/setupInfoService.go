package setup

type SetupInfoService interface {
	GetSetupInfo(setupInfo interface{}, fileName string) error
	SaveSetupInfo(setupInfo interface{}, fileName string, encoding string) error
}

func NewSetupInfoService() SetupInfoService {
	return &setupInfoServiceXmlImpl{}
}
