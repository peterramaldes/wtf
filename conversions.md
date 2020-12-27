# Formly 

```json

{
  "schema": {
    "title": "A registration form",
    "description": "A simple form example.",
    "type": "object",
    "required": [
      "firstName",
      "lastName"
    ],
    "properties": {
      "firstName": {
        "type": "string",
        "title": "First name",
        "default": "Chuck"
      },
      "lastName": {
        "type": "string",
        "title": "Last name"
      }
    }
  }
}

```

# Structure WSDL

```bash

├── Definitions (wsdl:definitions)
│   ├── Binding (wsdl:binding)
│   │   ├── PortType (wsdl:portType)
│   │   │   ├── Operation (wsdl:operation)
│   │   │   │   ├── Input (wsdl:input)
│   │   │   │   │   ├── Message (wsdl:message)
│   │   │   │   │   │   ├── Part (wsdl:part)
│   │   │   │   │   │   │   ├── Element (s:element)
│   │   │   │   │   │   │   │   ├── ComplexType (s:complexType)
│   │   │   │   │   │   │   │   │   ├── Sequence (s:sequence)
│   │   │   │   │   │   │   │   │   │   ├── Element (s:element)
│   │   │   │   ├── Output (wsdl:output)
│   │   │   │   │   ├── Message (wsdl:message)
│   │   │   │   │   │   ├── Part (wsdl:part)
│   │   │   │   │   │   │   ├── Element (s:element)
│   │   │   │   │   │   │   │   ├── ComplexType (s:complexType)
│   │   │   │   │   │   │   │   │   ├── Sequence (s:sequence)
│   │   │   │   │   │   │   │   │   │   ├── Element (s:element)

```

# Conversions

## Element (s:element)

name      -> properties[name]
type      -> properties[name].type
minOccurs -> 
maxOccurs -> 


### Data Types (properties[name].type)

s:int    -> number (note: remember to use `multipleOf`)
s:string -> string

