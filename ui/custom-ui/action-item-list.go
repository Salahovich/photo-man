package customUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type VerticalActionItemList struct {
	VBox   *fyne.Container
	Items  []*ActionItemWidget
	radial bool
}

func NewVerticalActionItemList(radial bool, objects ...*ActionItemWidget) *VerticalActionItemList {
	var list *VerticalActionItemList

	list = &VerticalActionItemList{
		Items:  make([]*ActionItemWidget, 0),
		VBox:   container.NewVBox(),
		radial: radial,
	}

	list.Items = append(list.Items, objects...)
	list.VBox.Objects = make([]fyne.CanvasObject, len(list.Items))
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
		list.VBox.Objects[i] = actionItem
	}
	list.VBox.Refresh()
	return list
}
