package page

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"sync"
	"translator/tst/tt_ui/pack"
	"translator/util"
)

var (
	apiAboutUs  *AboutUs
	onceAboutUs sync.Once
)

func GetAboutUs() *AboutUs {
	onceAboutUs.Do(func() {
		apiAboutUs = new(AboutUs)
		apiAboutUs.id = util.Uid()
		apiAboutUs.name = "关于我们"
	})
	return apiAboutUs
}

type AboutUs struct {
	id         string
	name       string
	mainWindow *walk.MainWindow
	rootWidget *walk.Composite
}

func (customPage *AboutUs) GetId() string {
	return customPage.id
}

func (customPage *AboutUs) GetName() string {
	return customPage.name
}

func (customPage *AboutUs) BindWindow(win *walk.MainWindow) {
	customPage.mainWindow = win
}

func (customPage *AboutUs) SetVisible(isVisible bool) {
	if customPage.rootWidget != nil {
		customPage.rootWidget.SetVisible(isVisible)
	}
}

func (customPage *AboutUs) GetWidget() Widget {
	return StdRootWidget(&customPage.rootWidget,
		pack.TTGroupBox(
			pack.NewTTGroupBoxArgs(nil).
				SetTitle("关于我们").SetVisible(true).SetLayoutVBox(false).
				SetWidgets(
					pack.NewWidgetGroup().Append(
						pack.TTTextLabel(pack.NewTTTextLabelArgs(nil).
							SetCustomSize(Size{Width: 100, Height: 300}).
							SetText(
								`    你好，欢迎使用字幕翻译工具（以下简称：Anto），我是Anto的研发人员。。。叭叭叭，还是不写了
`)),
					).AppendZeroVSpacer().GetWidgets(),
				),
		),
	)
}

func (customPage *AboutUs) Reset() {

}