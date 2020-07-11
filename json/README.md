### json package

javaScript object notation (JSON) is an open standard file format, and data interchange format, that uses human-readable text to store and transmit data objects consisting of attribute–value pairs and array data types (or any other serializable value). 

json can have four peimitive types (strings, numbers, booleans, and null) and two structured types (objects and arrays)

json structure contains names\values:

```{"name1": "value1", "name2":{name2-1:"value2-1"}, "name3": 3, "name4": [1,2,3,4], "name5": null } ```

json representation of a person:

```
{
  "name": "John Mclain",
  "dyingHarder": true,
  "age": 32,
  "address": {
    "streetAddress": "Ben Yehuda 11",
    "city": "Tel Aviv",
    "state": "IL",
    "postalCode": "10021"
  },
  "phoneNumbers": [
    {
      "type": "home",
      "number": "03-678126"
    },
    {
      "type": "office",
      "number": "03-12396768"
    }
  ],
}
```
go package - https://godoc.org/encoding/json

go to json - https://mholt.github.io/json-to-go/

##### marshal

##### unmarshal

##### encoder



##### decoder  bbbbb.ihb,,,,,,m.,,.