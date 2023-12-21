# ideas

Fields in sections will be defined by structs -> section
need to be able to activate/deactivate each field

Model of the form will be struct of sections -> fullform

Validations define methods on the specific fullform

when receiving patch for form, trigger validation

translations https://medium.com/@kyodo-tech/localization-for-dynamic-web-content-with-go-and-htmx-1f51fe81bf91

need a mechanism to replicate observer behavior on section post

maybe a type Control<T> {
    value T
    errors []formError
    enabled bool
}

# questions
what db? sql? mongo?
how to use translation keys in templates?
how to create a form?