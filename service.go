package main

func getItems() *ItemList {
	items := fetch()
	list := &ItemList{}
	list.Items = items
	return list
}

func addItems(itemList *ItemList) error {
	for _, item := range itemList.Items {
		error := add(item)
		if error != nil {
			return error
		}
	}
	return nil
}

func deleteItem(code string) error {
	return remove(code)
}
