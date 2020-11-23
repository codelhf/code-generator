# MacOS uses app bundles for GUI apps
mkdir -p code-generator.app/Contents/MacOS
go build -o code-generator.app/Contents/MacOS/code-generator
# Or click on the app in Finder
open code-generator.app
