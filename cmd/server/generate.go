package main

//go:generate echo "Generating templ files..."
//go:generate templ generate -path ../../templates
//go:generate echo "templ files generated"

//go:generate echo "Generating Tailwind CSS..."
//go:generate npx @tailwindcss/cli -i ../../public/static/css/input.css -o ../../public/static/css/output.css
//go:generate echo "Tailwind CSS generated"
