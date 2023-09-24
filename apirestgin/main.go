package apirestgin

func Example() {
	err := routes().Run("localhost:8080")
	if err != nil {
		return
	}
}
