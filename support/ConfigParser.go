package support

import (
	"bufio"
	"errors"
	"github.com/fatih/structs"
	"os"
	"os/user"
)

// CheckConfigFile checks for presence of the config file, then proceeds to
// its parsing and validation
func CheckConfigFile() (ConfigFile, error) {
	// configuration file is located into $HOME/.config/tsreddit.json
	userHome, _ := user.Current()
	configFilePath := userHome.HomeDir + "/.config/tsreddit.json"

	confFileObject, err := os.Open(configFilePath)
	defer confFileObject.Close()
	if err != nil {
		// probably no configuration file is present
		err = errors.New("error: configuration file not found in ~/.config/tsreddit.json")
	}
	cReader := bufio.NewReader(confFileObject)

	// our configuration file structure resides here!
	var configFile ConfigFile
	err = configFile.DecodeConfigFile(cReader)

	if err != nil {
		err = errors.New("error: your configuration file is malformed, check for any typos or missing properties")
	}

	// check for any nil Values
	for key, value := range structs.Map(configFile) {
		if value == "" {
			err = errors.New("error: '" + key + "' field not correctly defined")
		}
	}

	return configFile, err
}
