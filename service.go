package main

func getItems() *ItemList {
	items := fetch()
	list := &ItemList{}
	list.Items = items
	return list
}

func addItems(itemList *ItemList) {
	for _, item := range itemList.Items {
		add(item)
	}
}

func deleteItem(code string) {
	remove(code)
}
