
# RpcError

REQUIRED on error. This member MUST NOT exist if there was no error triggered during invocation.

## Properties

Name | Type
------------ | -------------
`code` | number
`message` | string
`data` | object

## Example

```typescript
import type { RpcError } from ''

// TODO: Update the object below with actual values
const example = {
  "code": null,
  "message": null,
  "data": null,
} satisfies RpcError

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RpcError
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


