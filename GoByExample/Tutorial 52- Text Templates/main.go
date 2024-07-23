package main

import (
	"os"
	"text/template"
)

func main() {

	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n") // {{.}} is where it will insert the value
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n")) // .Must panics if Parse fails

	t1.Execute(os.Stdout, "some text") // Run the template at os.Stdout with the second parameter as a value
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template { // Helper function that creates a template with the name "Name" and text "t"
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n") // Use helper function. {{.Name}} is the field

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"}) // Set value to be a field

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse", // maps can also be used
	})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n") // Check if value is empty
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n") // Create template that goes through range
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
