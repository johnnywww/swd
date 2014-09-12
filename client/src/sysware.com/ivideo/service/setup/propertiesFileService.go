package setup

type PropertiesFileService interface {
	GetInfo(fileName string) (map[string]string, error)
	SaveInfo(infos map[string]string, fileName string) error
}

func NewPropertiesFileService() PropertiesFileService {
	return &propertiesFileServiceImpl{}
}
