# Build all go project present at src folder and save the generated binary at build folder
# The build process should optimise to build only the changed files and generate the smaller
# binary possible.

build:

# give file path as argument
mock:
	mockgen ./src/$(package) -destination=mocks/$(package).go

