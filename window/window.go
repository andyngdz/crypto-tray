package window

import "github.com/wailsapp/wails/v3/pkg/application"

// NewSettings creates the settings window (hidden by default for tray app)
func NewSettings(app *application.App, icon []byte) *application.WebviewWindow {
	return app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:      "settings",
		Title:     "Crypto Tray Settings",
		Width:     640,
		Height:    800,
		MinWidth:  640,
		MinHeight: 600,
		Frameless: true,
		Hidden:    true,
		BackgroundColour: application.RGBA{
			Red: 0, Green: 0, Blue: 0, Alpha: 0,
		},
		Windows: application.WindowsWindow{
			HiddenOnTaskbar: true,
		},
		Linux: application.LinuxWindow{
			Icon: icon,
		},
		Mac: application.MacWindow{
			Backdrop: application.MacBackdropTranslucent,
		},
	})
}
