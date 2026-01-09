package customUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ActionItemWidget struct {
	widget.Button
}

func NewActionItemWidget(icon fyne.Resource, onClick func()) *ActionItemWidget {
	var item *ActionItemWidget
	wrappedOnClickFunc := func() {
		switch item.Importance {
		case widget.LowImportance:
			item.Importance = widget.HighImportance
		case widget.HighImportance:
			item.Importance = widget.LowImportance
		}
		item.Button.Refresh()
		onClick()
	}

	item = &ActionItemWidget{}
	item.Icon = icon
	item.Importance = widget.LowImportance
	item.OnTapped = wrappedOnClickFunc
	item.ExtendBaseWidget(item)

	return item
}

func (item *ActionItemWidget) CreateRenderer() fyne.WidgetRenderer {
	return item.Button.CreateRenderer()
}
