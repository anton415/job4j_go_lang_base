package tracker

import "github.com/google/uuid"

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.GetItems() {
		out.Out(item.toString())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id:")
	id := in.Get()
	out.Out("enter new name:")
	name := in.Get()
	tracker.UpdateItem(Item{ID: id, Name: name})
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id:")
	id := in.Get()
	tracker.DeleteItem(id)
}

type SearchUsecase struct{}

func (u SearchUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	for _, item := range tracker.SearchItems(name) {
		out.Out(item.toString())
	}
}
