If a variable/function name is capitalized then it is exported from the package.

`Pizza` and `Pi` are exported names while `pineapple` and `cake` are not.

When importing a package, you can only refer to the publically available exported names. An unexported name is not accessible from outside the package and hence is private.