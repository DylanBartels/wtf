package help

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/maker"
	"github.com/wtfutil/wtf/utils"
)

// Display displays the output of the --help argument
func Display(moduleName string, config *config.Config) {
	if moduleName == "" {
		fmt.Println("\n  --module takes a module name as an argument, i.e: '--module=github'")
	} else {
		fmt.Printf("%s\n", helpFor(moduleName, config))
	}
}

func helpFor(moduleName string, config *config.Config) string {
	modConfig, _ := config.Get("wtf.mods." + moduleName)
	widget := maker.MakeWidget(nil, nil, moduleName, moduleName, modConfig, config)

	result := ""
	result += utils.StripColorTags(widget.HelpText())
	result += "\n"
	result += "Configuration Attributes"
	result += widget.ConfigText()
	return result
}
