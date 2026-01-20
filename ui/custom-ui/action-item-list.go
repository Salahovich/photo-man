package customUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ActionItemList struct {
	Box    *fyne.Container
	Items  []*ActionItemWidget
	radial bool
}

func NewActionItemList(radial bool, vertical bool, objects ...*ActionItemWidget) *ActionItemList {
	var list *ActionItemList

	list = &ActionItemList{
		Items:  make([]*ActionItemWidget, 0),
		radial: radial,
	}
	if vertical {
		list.Box = container.NewVBox()
	} else {
		list.Box = container.NewHBox()
	}

	list.Items = append(list.Items, objects...)
	list.Box.Objects = make([]fyne.CanvasObject, len(list.Items))
	for i, actionItem := range list.Items {
		childTapped := actionItem.OnTapped
		actionItem.OnTapped = func() {
			if list.radial {
				for j, others := range list.Items {
					if j != i {
						others.Importance = widget.LowImportance
						others.Refresh()
					}
				}
			}
			childTapped()
		}
		list.Box.Objects[i] = actionItem
	}
	list.Box.Refresh()
	return list
}
