"package main\n\nimport (\n \t\"fmt\"\n \t\"log\"\n \t\"os\"\n \t\"net/http\"\n)\n\nfunc main() {\n \thttp.HandleFunc(\"/\", func(w http.ResponseWriter, r *http.Request) {\n \t\tfmt.Fprintf(w, \"<h1>Futuristic Go Web Application</h1>\")\n \t})\n \tfmt.Println(\"Listening on port :8080\");\n \terr := http.ListenAndServe(\":8080\", nil)\n \tif err != nil {\n \t\tlog.Fatal(\"ListenAndServe: \", err)\n \t}\n}"