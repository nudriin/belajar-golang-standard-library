package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("src/app/app.controller.ts"))        // Mengambil foldernya
	fmt.Println(path.Base("src/app/app.controller.ts"))       // Mengambil nama filenya
	fmt.Println(path.Ext("src/app/app.controller.ts"))        // Mengambil ekstensi
	fmt.Println(path.Join("src", "app", "app.controller.ts")) // Memembuat path dari beberapa folder dan file

	// Untuk path file system
	fmt.Println(filepath.Dir("src/app/app.controller.ts"))        // Mengambil foldernya
	fmt.Println(filepath.Base("src/app/app.controller.ts"))       // Mengambil nama filenya
	fmt.Println(filepath.Ext("src/app/app.controller.ts"))        // Mengambil ekstensi
	fmt.Println(filepath.IsAbs("src/app/app.controller.ts"))      // Apakah absolute = absolute itu darii C:/Windows/Users
	fmt.Println(filepath.IsLocal("src/app/app.controller.ts"))    // Apakah local = local ../src/app/main.ts
	fmt.Println(filepath.Join("src", "app", "app.controller.ts")) // Memembuat path dari beberapa folder dan file
}
