package main

import (
	"fmt"
	cloud189 "panSDK/drivers/189"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	_189, err := cloud189.NewCloud189("15315496321", "mogen123")
	if err != nil {
		panic(err)
	}
	resp, err := _189.GetObjectFolderNodes("-13")
	fmt.Println(resp)
	fmt.Println(err)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
