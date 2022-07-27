# go-run-package-tests-in-custom-order

This is just a test project for me to learn Go as a new language. 

Here are some interesting things I want to share with others who are on a similar mission like me:

## Run tests of a package with custom selection and order

Please see  ./validator/validator_test.go with its test function TestPackageTestsInCustomOrder. It uses t.Run() to run a selection of tests in custom order.

You can run this single wrapper test function with:

    go test -v -run TestPackageTestsInCustomOrder ./validator
