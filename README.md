nested
=====

The Nested Go package makes it easier to handle (too) big nested JSON structures.

How to use
=============
Lets say you have the next JSON:
```
j := `{
	"hello": {
		"this": {
			"is": {
				"an": {
					"example": "hi"
				}
			}
		}
	}
}`
```

Now if you unmarshal it to u (pseudocode)
```
u := unmarshal(j)
```

You will have a hard time accessing members in u, like
```
u.(map[string]interface{})["hello"].(map[string]interface{})["this"].(map[string]interface{})	//... etc etc
```

But not with package Nested!
```
magic, ok := nested.Get(u, "hello.this.is.an.example")
fmt.Println(magic, ok)
```

Will output something like:
```
"hi" true
```

It's magic, really!